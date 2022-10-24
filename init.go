package exloggo

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	levelInfo          = "INFO"
	levelWarning       = "WARNING"
	levelError         = "ERROR"
	levelDebug         = "DEBUG"
	levelFatal         = "FATAL"
	levelRequestResult = "REQUEST_RESULT"
)

const (
	ReleaseMode     = "RELEASE"
	DevelopmentMode = "DEVELOPMENT"
)

const (
	defaultServerVersion   = "dev"
	defaultOutputDirectory = "logs/"
	defaultMode            = DevelopmentMode
)

type Parameters struct {
	Mode          *string
	ServerVersion *string
	Directory     *string
}

var loggerMode string
var prefixPath string
var serverVersion string
var logsDirectoryPath string

func Inject(router *gin.Engine) {
	router.Use(middleware)
}

func SetParameters(params *Parameters) error {
	if err := setMode(params.Mode); err != nil {
		Error(errorInvalidMode, nil)
		return err
	}
	setServerVersion(params.ServerVersion)
	setOutputDirectory(params.Directory)

	return nil
}

func setMode(mode *string) error {
	if mode != nil {
		if *mode != DevelopmentMode && *mode != ReleaseMode {
			return errors.New(errorInvalidMode)
		}
	} else {
		loggerMode = DevelopmentMode
		return nil
	}
	loggerMode = *mode
	return nil
}

func setServerVersion(version *string) {
	if version == nil {
		serverVersion = defaultServerVersion
		return
	}
	serverVersion = *version
}

func setOutputDirectory(path *string) {
	if path == nil {
		logsDirectoryPath = defaultOutputDirectory
		return
	}
	logsDirectoryPath = *path
}
