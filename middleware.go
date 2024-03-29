package exloggo

import (
	"context"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

const (
	headerExceptionStatus = 432
	headerError           = "an error occurred while working with headers"
	headerEmpty           = "empty field"
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

type ContextBody struct {
	RequestHeaders  *RequestHeaders
	ResponseHeaders *ResponseHeaders
}

var contextBodyStore sync.Map

func MiddlewareWithRequiredHeaders(c *gin.Context) {
	var requestHeaders RequestHeaders
	var requestTime = time.Now().UTC().Round(time.Second)

	if err := c.ShouldBindHeader(&requestHeaders); err != nil {
		c.AbortWithStatusJSON(headerExceptionStatus, headerError)
		return
	}

	responseHeaders := ResponseHeaders{
		RequestTimeString: requestTime.Format(time.RFC1123),
		ClientRequestId:   requestHeaders.ClientRequestId,
		RequestTime:       requestTime,
		RequestId:         getUUID(),
	}

	body := ContextBody{
		RequestHeaders:  &requestHeaders,
		ResponseHeaders: &responseHeaders,
	}

	ctx := context.Background()
	contextBodyStore.Store(ctx, &body)

	c.Next()

	contextBodyStore.Delete(ctx)
}

func Middleware(c *gin.Context) {
	var requestTime = time.Now().UTC().Round(time.Second)

	requestHeaders := RequestHeaders{
		ClientRequestId: headerEmpty,
	}

	responseHeaders := ResponseHeaders{
		RequestTimeString: requestTime.Format(time.RFC1123),
		ClientRequestId:   requestHeaders.ClientRequestId,
		RequestTime:       requestTime,
		RequestId:         getUUID(),
	}

	body := ContextBody{
		RequestHeaders:  &requestHeaders,
		ResponseHeaders: &responseHeaders,
	}

	ctx := context.Background()
	contextBodyStore.Store(ctx, &body)

	c.Next()

	contextBodyStore.Delete(ctx)
}
