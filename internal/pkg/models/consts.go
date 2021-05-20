package models

import "os"

// "DEBUG" : будет слушать запросы с localhost:3000
// "DEPLOY" : будет слушать запросы с lepick.ru
var Mode string = os.Getenv("MODE")

func GetDomain() string {
	if Mode == "DEBUG" {
		return "localhost:3000"
	}

	return "lepick.ru"
}

func GetSecure() bool {
	if Mode == "DEBUG" {
		return false
	}

	return true
}
