package exloggo

import (
	"errors"
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
	Mode          string
	ServerVersion string
	Directory     string
}

var loggerMode string
var prefixPath string
var serverVersion string
var logsDirectoryPath string

func setMode(mode string) error {
	if mode != "" {
		if mode != DevelopmentMode && mode != ReleaseMode {
			return errors.New(errorInvalidMode)
		}
	} else {
		loggerMode = DevelopmentMode
		return nil
	}
	loggerMode = mode
	return nil
}

func setServerVersion(version string) {
	if version == "" {
		serverVersion = defaultServerVersion
		return
	}
	serverVersion = version
}

func setOutputDirectory(path string) {
	if path == "" {
		logsDirectoryPath = defaultOutputDirectory
		return
	}
	logsDirectoryPath = path
}
