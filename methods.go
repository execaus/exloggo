package exloggo

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func SetParameters(params *Parameters) error {
	if err := setMode(params.Mode); err != nil {
		Error(errorInvalidMode)
		return err
	}
	setServerVersion(params.ServerVersion)
	setOutputDirectory(params.Directory)

	return nil
}

func GetContextBody() *ContextBody {
	ctx := context.Background()
	body, ok := contextBodyStore.Load(ctx)
	if !ok {
		return nil
	}
	return body.(*ContextBody)
}

func Info(message string) {
	saveLog(message, levelInfo, nil)
}

func InfoWithExtension(message string, extend interface{}) {
	saveLog(message, levelInfo, extend)
}

func Infof(template string, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelInfo, nil)
}

func InfofWithExtension(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelInfo, extend)
}

func Warning(message string) {
	saveLog(message, levelWarning, nil)
}
func WarningWithExtension(message string, extend interface{}) {
	saveLog(message, levelWarning, extend)
}

func Warningf(template string, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelWarning, nil)
}

func WarningfWithExtension(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelWarning, extend)
}

func Error(message string) {
	saveLog(message, levelError, nil)
}

func ErrorWithExtension(message string, extend interface{}) {
	saveLog(message, levelError, extend)
}

func Errorf(template string, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelError, nil)
}

func ErrorfWithExtension(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelError, extend)
}

func Debug(message string) {
	if loggerMode == DevelopmentMode {
		saveLog(message, levelDebug, nil)
	}
}

func DebugWithExtension(message string, extend interface{}) {
	if loggerMode == DevelopmentMode {
		saveLog(message, levelDebug, extend)
	}
}

func Debugf(template string, a ...any) {
	message := fmt.Sprintf(template, a)
	if loggerMode == DevelopmentMode {
		saveLog(message, levelDebug, nil)
	}
}

func DebugfWithExtension(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	if loggerMode == DevelopmentMode {
		saveLog(message, levelDebug, extend)
	}
}

func Fatal(message string) {
	saveLog(message, levelFatal, nil)
	os.Exit(1)
}

func FatalWithExtension(message string, extend interface{}) {
	saveLog(message, levelFatal, extend)
	os.Exit(1)
}

func Fatalf(template string, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelFatal, nil)
	os.Exit(1)
}

func FatalfWithExtension(template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLog(message, levelFatal, extend)
	os.Exit(1)
}

func RequestResult(c *gin.Context, message string) {
	saveLogWithRequestData(message, levelRequestResult, nil, c)
}

func RequestResultWithExtension(c *gin.Context, message string, extend interface{}) {
	saveLogWithRequestData(message, levelRequestResult, extend, c)
}

func RequestResultf(c *gin.Context, template string, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLogWithRequestData(message, levelRequestResult, nil, c)
}

func RequestResultfWithExtension(c *gin.Context, template string, extend interface{}, a ...any) {
	message := fmt.Sprintf(template, a)
	saveLogWithRequestData(message, levelRequestResult, extend, c)
}
