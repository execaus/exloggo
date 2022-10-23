package exloggo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Info(message string, extend interface{}) {
	saveLog(message, levelInfo, extend)
}

func Infof(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelInfo, extend)
}

func Warning(message string, extend interface{}) {
	saveLog(message, levelWarning, extend)
}

func Warningf(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelWarning, extend)
}

func Error(message string, extend interface{}) {
	saveLog(message, levelError, extend)
}

func Errorf(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelError, extend)
}

func Debug(message string, extend interface{}) {
	if loggerMode == DevelopmentMode {
		saveLog(message, levelDebug, extend)
	}
}

func Debugf(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	if loggerMode == DevelopmentMode {
		saveLog(message, levelDebug, extend)
	}
}

func Fatal(message string, extend interface{}) {
	saveLog(message, levelFatal, extend)
	os.Exit(1)
}

func Fatalf(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelFatal, extend)
	os.Exit(1)
}

func RequestResult(message string, extend interface{}, c *gin.Context) {
	saveLogWithRequestData(message, levelRequestResult, extend, c)
}

func RequestResultf(template string, extend interface{}, c *gin.Context, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLogWithRequestData(message, levelRequestResult, extend, c)
}
