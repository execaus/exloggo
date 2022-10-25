package exloggo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
	"time"
)

const (
	errorHeaders       = "an error occurred while working with headers"
	RequestHeadersKey  = "ExGoLogRequestHeadersKey"
	ResponseHeadersKey = "ExGoLogResponseHeadersKey"
)

type RequestHeaders struct {
	ClientRequestId string `header:"Client-Request-ID" binding:"required,uuid"`
}

type ResponseHeaders struct {
	RequestTime       time.Time
	RequestTimeString string `header:"Timestamp" binding:"required,datetime=Mon, 02 Jan 2006 15:04:05 MST"`
	ClientRequestId   string `header:"Client-Request-ID" binding:"required,uuid"`
	RequestId         string `header:"Request-ID" binding:"required,uuid"`
}

func middleware(c *gin.Context) {
	var requestHeaders RequestHeaders
	var requestTime = time.Now().UTC()

	if err := c.ShouldBindHeader(&requestHeaders); err != nil {
		SendHeaderException(c, err.Error())
		return
	}

	responseHeaders := ResponseHeaders{
		RequestTimeString: requestTime.Format(time.RFC1123),
		ClientRequestId:   requestHeaders.ClientRequestId,
		RequestTime:       requestTime,
		RequestId:         GetUUIDv7(),
	}

	c.Set(RequestHeadersKey, &requestHeaders)
	c.Set(ResponseHeadersKey, &responseHeaders)

	if err := setResponseHeaders(c, &responseHeaders); err != nil {
		SendGeneralException(c, errorHeaders)
		return
	}

	c.Next()
}

func setResponseHeaders(c *gin.Context, headers interface{}) error {
	a := reflect.ValueOf(headers)
	fieldsCount := reflect.ValueOf(headers).Elem().NumField()
	if a.Kind() != reflect.Ptr {
		Error("wrong type struct", nil)
		return errors.New("wrong type struct")
	}
	for x := 0; x < fieldsCount; x++ {
		headerName := reflect.TypeOf(headers).Elem().Field(x).Tag.Get("header")
		value := reflect.ValueOf(headers).Elem().Field(x).String()
		c.Header(headerName, value)
	}

	return nil
}
