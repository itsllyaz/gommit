// Package sdk provides GitHub and GitLab API clients for gommit.
package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is a generic REST API client for forge platforms.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	UserAgent  string
}

// NewClient creates an SDK client for a forge base URL.
func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: strings.TrimRight(baseURL, "/"),
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		UserAgent: "gommit-sdk/1.0",
	}
}

// GitHub returns a client configured for GitHub API v3.
func GitHub(token string) *Client {
	return NewClient("https://api.github.com", token)
}

// GitLab returns a client configured for GitLab API v4.
func GitLab(baseURL, token string) *Client {
	if baseURL == "" {
		baseURL = "https://gitlab.com/api/v4"
	}
	return NewClient(baseURL, token)
}

// Request performs an authenticated HTTP request.
func (c *Client) Request(ctx context.Context, method, path string, body io.Reader) (*http.Response, error) {
	u, err := url.JoinPath(c.BaseURL, path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, u, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return c.HTTPClient.Do(req)
}

// GetJSON performs GET and decodes JSON response.
func (c *Client) GetJSON(ctx context.Context, path string, out any) error {
	resp, err := c.Request(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("api %s: status %d: %s", path, resp.StatusCode, string(b))
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// Repository represents basic repository metadata.
type Repository struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	URL         string `json:"html_url"`
}

// PullRequest represents a PR/MR summary.
type PullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
	User   string `json:"user.login"`
	URL    string `json:"html_url"`
}

// Issue represents an issue summary.
type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
	Labels []string
	URL    string `json:"html_url"`
}

// RepoStats aggregates repository statistics.
type RepoStats struct {
	OpenPRs    int
	OpenIssues int
	Commits    int
	Contributors int
}

// GitHubService wraps GitHub-specific endpoints.
type GitHubService struct {
	client *Client
}

func NewGitHubService(token string) *GitHubService {
	return &GitHubService{client: GitHub(token)}
}

func (g *GitHubService) GetRepo(ctx context.Context, owner, repo string) (Repository, error) {
	var r Repository
	path := fmt.Sprintf("/repos/%s/%s", owner, repo)
	err := g.client.GetJSON(ctx, path, &r)
	return r, err
}

func (g *GitHubService) ListPullRequests(ctx context.Context, owner, repo, state string) ([]PullRequest, error) {
	var prs []PullRequest
	if state == "" {
		state = "open"
	}
	path := fmt.Sprintf("/repos/%s/%s/pulls?state=%s", owner, repo, state)
	err := g.client.GetJSON(ctx, path, &prs)
	return prs, err
}

func (g *GitHubService) ListIssues(ctx context.Context, owner, repo, state string) ([]Issue, error) {
	var issues []Issue
	if state == "" {
		state = "open"
	}
	path := fmt.Sprintf("/repos/%s/%s/issues?state=%s", owner, repo, state)
	err := g.client.GetJSON(ctx, path, &issues)
	return issues, err
}

// GitLabService wraps GitLab-specific endpoints.
type GitLabService struct {
	client *Client
}

func NewGitLabService(baseURL, token string) *GitLabService {
	return &GitLabService{client: GitLab(baseURL, token)}
}

func (gl *GitLabService) GetProject(ctx context.Context, id string) (Repository, error) {
	var r Repository
	path := fmt.Sprintf("/projects/%s", url.PathEscape(id))
	err := gl.client.GetJSON(ctx, path, &r)
	return r, err
}

func (gl *GitLabService) ListMergeRequests(ctx context.Context, projectID, state string) ([]PullRequest, error) {
	var mrs []PullRequest
	if state == "" {
		state = "opened"
	}
	path := fmt.Sprintf("/projects/%s/merge_requests?state=%s", url.PathEscape(projectID), state)
	err := gl.client.GetJSON(ctx, path, &mrs)
	return mrs, err
}

// ParseRemote extracts forge info from a git remote URL.
func ParseRemote(raw string) (forge, owner, repo string, err error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimSuffix(raw, ".git")
	switch {
	case strings.Contains(raw, "github.com"):
		forge = "github"
		parts := strings.Split(raw, "github.com/")
		if len(parts) < 2 {
			return "", "", "", fmt.Errorf("invalid github remote")
		}
		seg := strings.Split(parts[1], "/")
		if len(seg) < 2 {
			return "", "", "", fmt.Errorf("invalid github remote path")
		}
		return forge, seg[0], seg[1], nil
	case strings.Contains(raw, "gitlab"):
		forge = "gitlab"
		return forge, "", raw, nil
	default:
		return "unknown", "", raw, nil
	}
}
