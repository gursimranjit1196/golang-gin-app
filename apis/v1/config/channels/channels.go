package channels

import (
	"gin-app/apis/v1/constants"
	"gin-app/apis/v1/utils/loggers"
)

var CriticalErrorChannel chan interface{}

func InitChannels() {
	CriticalErrorChannel = make(chan interface{})
	go handleCritialErrors()
}

func handleCritialErrors() {
	for {
		loggers.Log(constants.CriticalErrorMsg, <-CriticalErrorChannel)
		// Send emails or take necessary actions on critical errors in app.
	}
}
