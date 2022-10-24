package exloggo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	accessDenied          = "accessDenied"
	generalException      = "generalException"
	itemNotFound          = "itemNotFound"
	invalidRequest        = "invalidRequest"
	resourceModified      = "resourceModified"
	unAuthenticated       = "unAuthenticated"
	sendEmail             = "sendEmail"
	sendEmailError        = "smtp server error"
	conditionNotMet       = "conditionNotMet"
	resourceDeleted       = "resourceDeleted"
	notConfirmed          = "notConfirmed"
	timeoutExpired        = "timeoutExpired"
	headerException       = "headerException"
	sendEmailErrorStatus  = 506
	headerExceptionStatus = 432
)

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SendOK sends the response to the client without the body with code http.StatusNoContent.
func SendOK(c *gin.Context) {
	sendResponse(c, http.StatusNoContent, gin.H{})
}

// SendOKWithBody sends the response to the client with the code http.StatusOK.
func SendOKWithBody(c *gin.Context, data interface{}) {
	sendResponse(c, http.StatusOK, data)
}

// SendAccessDenied sends the response to the client without the body with code http.StatusForbidden.
func SendAccessDenied(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    accessDenied,
		Message: message,
	}

	saveRequestLog(c, message)
	sendResponse(c, http.StatusForbidden, data)
}

// SendGeneralException sends the response to the client without the body with code http.StatusInternalServerError.
func SendGeneralException(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    generalException,
		Message: message,
	}
	sendResponse(c, http.StatusInternalServerError, data)
}

// SendItemNotFound sends the response to the client without the body with code http.StatusNotFound.
func SendItemNotFound(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    itemNotFound,
		Message: message,
	}
	RequestResult(message, nil, c)
	sendResponse(c, http.StatusNotFound, data)
}

// SendInvalidRequest sends the response to the client without the body with code http.StatusBadRequest.
func SendInvalidRequest(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    invalidRequest,
		Message: message,
	}
	RequestResult(message, nil, c)
	sendResponse(c, http.StatusBadRequest, data)
}

// SendResourceModified sends the response to the client without the body with code http.StatusInternalServerError.
func SendResourceModified(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    resourceModified,
		Message: message,
	}
	saveRequestLog(c, message)
	sendResponse(c, http.StatusInternalServerError, data)
}

// SendUnAuthenticated sends the response to the client without the body with code http.StatusUnauthorized.
func SendUnAuthenticated(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    unAuthenticated,
		Message: message,
	}
	saveRequestLog(c, message)
	sendResponse(c, http.StatusUnauthorized, data)
}

// SendSendEmailError sends the response to the client without the body with code 506.
func SendSendEmailError(c *gin.Context) {
	data := ErrorBody{
		Code:    sendEmail,
		Message: sendEmailError,
	}
	sendResponse(c, sendEmailErrorStatus, data)
}

// SendConditionNotMet sends the response to the client without the body with code http.StatusPreconditionFailed.
func SendConditionNotMet(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    conditionNotMet,
		Message: message,
	}
	RequestResult(message, nil, c)
	sendResponse(c, http.StatusPreconditionFailed, data)
}

// SendResourceDeleted sends the response to the client without the body with code http.StatusGone.
func SendResourceDeleted(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    resourceDeleted,
		Message: message,
	}
	RequestResult(message, nil, c)
	sendResponse(c, http.StatusGone, data)
}

// SendNotConfirmed sends the response to the client without the body with code http.StatusPreconditionFailed.
func SendNotConfirmed(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    notConfirmed,
		Message: message,
	}
	RequestResult(message, nil, c)
	sendResponse(c, http.StatusPreconditionFailed, data)
}

// SendTimeoutExpired sends the response to the client without the body with code http.StatusPreconditionFailed.
func SendTimeoutExpired(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    timeoutExpired,
		Message: message,
	}
	sendResponse(c, http.StatusPreconditionFailed, data)
}

// SendHeaderException sends the response to the client without the body with code 432.
func SendHeaderException(c *gin.Context, message string) {
	data := ErrorBody{
		Code:    headerException,
		Message: message,
	}
	RequestResult(message, nil, c)
	sendResponse(c, headerExceptionStatus, data)
}

func saveRequestLog(c *gin.Context, message string) {
	// TODO extend - account info
	RequestResult(message, nil, c)
}

func sendResponse(c *gin.Context, status int, data interface{}) {
	c.AbortWithStatusJSON(status, data)
	UntieGoroutineRequestId()
}
