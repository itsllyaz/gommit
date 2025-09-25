package utils

import "time"


func GetTimeInTimezone(timezone string)(string, error){
	location, err := time.LoadLocation(timezone)
	if err != nil{
		return "", err
	}

	currentTime := time.Now().In(location)
	return currentTime.Format(time.RFC1123), nil
}
