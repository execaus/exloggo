package exloggo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	errorHeaders       = "an error occurred while working with headers"
	ResponseHeadersKey = "ExGoLogResponseHeadersKey"
)

type requestHeaders struct {
	ClientRequestId string `header:"Client-Request-ID" binding:"required,uuid"`
}

type responseHeaders struct {
	RequestTime       time.Time
	RequestTimeString string `header:"Timestamp" binding:"required,datetime=Mon, 02 Jan 2006 15:04:05 MST"`
	ClientRequestId   string `header:"Client-Request-ID" binding:"required,uuid"`
	RequestId         string `header:"Request-ID" binding:"required,uuid"`
}

func middleware(c *gin.Context) {
	headers, err := getResponseHeaders(c)
	if err != nil {
		SendHeaderException(c, err.Error())
		return
	}
	BindGoroutineRequestId(headers.RequestId, headers.ClientRequestId)
	c.Next()
}

func getResponseHeaders(c *gin.Context) (*responseHeaders, error) {
	headers, exists := c.Get(ResponseHeadersKey)
	if !exists {
		Error(errorHeaders, nil)
		return nil, errors.New(errorHeaders)
	}
	return headers.(*responseHeaders), nil
}
