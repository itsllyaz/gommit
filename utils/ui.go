package utils

import "github.com/fatih/color"

// Logo is the ASCII banner shown in command headers.
var Logo string = `#     _______  _______  _______  _______ __________________
#    (  ____ \(  ___  )(       )(       )\__   __/\__   __/
#    | (    \/| (   ) || () () || () () |   ) (      ) (
#    | |      | |   | || || || || || || |   | |      | |
#    | | ____ | |   | || |(_)| || |(_)| |   | |      | |
#    | | \_  )| |   | || |   | || |   | |   | |      | |
#    | (___) || (___) || )   ( || )   ( |___) (___   | |
#    (_______)(_______)|/     \||/     \|\_______/   )_(
#                                                          `

// Tagline is shown alongside branding output.
var Tagline = "Your personal Git assistant"

// Red prints a formatted line in red.
func Red(format string, args ...any) {
	color.Red(format, args...)
}

// Green prints a formatted line in green.
func Green(format string, args ...any) {
	color.Green(format, args...)
}

// Yellow prints a formatted line in yellow.
func Yellow(format string, args ...any) {
	color.Yellow(format, args...)
}
