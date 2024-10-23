package exloggo

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	errorJSON      = "json error"
	RequestBodyKey = "body_string"
)

var (
	logFilePrefix = "log-"
	systemLogging = "system logging"
)

type LogData struct {
	Level           string       `json:"level"`
	Message         string       `json:"message"`
	Point           string       `json:"point"`
	FilePoint       string       `json:"file_point"`
	RequestId       string       `json:"request_id"`
	ClientRequestId string       `json:"client_request_id"`
	Timestamp       string       `json:"timestamp"`
	ServerVersion   string       `json:"server_version"`
	Request         *RequestData `json:"request"`
	Extends         interface{}  `json:"extends"`
}

type RequestData struct {
	RequestIP      string            `json:"ip"`
	RequestBody    interface{}       `json:"body"`
	RequestURI     string            `json:"uri"`
	RequestHeaders map[string]string `json:"headers"`
}

func (l *LogData) appendRequestData(c *gin.Context) error {
	body := getRequestBody(c)
	headers, err := getRequestHeaders(c)
	if err != nil {
		Error(err.Error())
		return err
	}

	l.Request = &RequestData{
		RequestIP:      c.ClientIP(),
		RequestBody:    body,
		RequestURI:     c.Request.RequestURI,
		RequestHeaders: headers,
	}

	return nil
}

func init() {
	initRelative()
	loggerMode = DevelopmentMode
}

func initRelative() {
	_, fileName, _, _ := runtime.Caller(0)
	prefixPath = filepath.ToSlash(filepath.Dir(filepath.Dir(fileName))) + "/"
}

func getRelativePath(path string) string {
	return strings.TrimPrefix(filepath.ToSlash(path), prefixPath)
}

func saveLogWithRequestData(message string, level string, extend interface{}, c *gin.Context) {
	foundation := getFoundationLogData(message, level)
	if err := foundation.appendRequestData(c); err != nil {
		Error(err.Error())
	}
	outputLogData(foundation, extend)
}

func saveLog(message string, level string, extend interface{}) {
	foundation := getFoundationLogData(message, level)
	outputLogData(foundation, extend)
}

func outputLogData(foundation *LogData, extend interface{}) {
	foundation.Extends = extend

	dataJson, err := json.Marshal(foundation)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if loggerMode == DevelopmentMode {
		fmt.Println(string(dataJson))
	}
	outputFile(string(dataJson))
}

func getFoundationLogData(message string, level string) *LogData {
	var file string
	var line int
	var filePoint string

	if level == levelRequestResult {
		_, file, line, _ = runtime.Caller(4)
	} else {
		_, file, line, _ = runtime.Caller(3)
	}

	switch runtime.GOOS {
	case "linux", "darwin":
		filePoint = fmt.Sprintf("file://%s:%d", file, line)
	case "windows":
		filePoint = fmt.Sprintf("file:///%s:%d", file, line)
	default:
		filePoint = fmt.Sprintf("%s:%d", file, line)
	}

	point := fmt.Sprintf(`%s:%d`, getRelativePath(file), line)

	timestamp := time.Now().UTC().Format(time.RFC3339)
	contextBody := GetContextBody()
	if contextBody == nil {
		return &LogData{
			Level:           level,
			Message:         message,
			Point:           point,
			FilePoint:       filePoint,
			RequestId:       systemLogging,
			ClientRequestId: systemLogging,
			Timestamp:       timestamp,
			ServerVersion:   serverVersion,
		}
	}

	return &LogData{
		Level:           level,
		Message:         message,
		Point:           point,
		FilePoint:       filePoint,
		RequestId:       contextBody.ResponseHeaders.RequestId,
		ClientRequestId: contextBody.ResponseHeaders.ClientRequestId,
		Timestamp:       timestamp,
		ServerVersion:   serverVersion,
	}
}

func outputFile(logData string) {
	var file *os.File

	timeNow := time.Now().UTC()

	date := fmt.Sprintf(`%d-%.2d-%.2d`, timeNow.Year(), timeNow.Month(), timeNow.Day())

	monthDate := fmt.Sprintf(`%d-%.2d`, timeNow.Year(), timeNow.Month())
	monthPath := fmt.Sprintf(`%s/%s`, logsDirectoryPath, monthDate)

	filePath := fmt.Sprintf(`%s/%s/%s%s.txt`, logsDirectoryPath, monthDate, logFilePrefix, date)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(monthPath, 0777); err != nil {
			log.Println(fmt.Sprintf(`%s: %s: %s`, "error mkdir all", err.Error(), monthPath))
			return
		}
		_, err = os.Create(filePath)
		if err != nil {
			log.Println(fmt.Sprintf(`%s: %s: %s`, "error create log file", err.Error(), filePath))
			return
		}
	}

	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Println("error open log file: " + filePath)
		return
	}

	_, err = file.WriteString(logData + "\n")
	if err != nil {
		log.Println("error write log file: " + filePath)
		return
	}

	if err = file.Close(); err != nil {
		log.Println("error close log file: " + filePath)
		return
	}
}

func getRequestBody(c *gin.Context) interface{} {
	var body interface{}
	stringBody := c.GetString(RequestBodyKey)
	if stringBody == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(stringBody), &body); err != nil {
		Error(err.Error())
	}
	return body
}

func getRequestHeaders(c *gin.Context) (map[string]string, error) {
	mapHeaders := make(map[string]string)
	for key, values := range c.Request.Header {
		mapHeaders[key] = strings.Join(values, ",")
	}
	return mapHeaders, nil
}
