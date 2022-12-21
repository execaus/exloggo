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

func Middleware(c *gin.Context) {
	var requestHeaders RequestHeaders
	var requestTime = time.Now().UTC()

	if err := c.ShouldBindHeader(&requestHeaders); err != nil {
		c.AbortWithStatusJSON(headerExceptionStatus, headerError)
		return
	}

	responseHeaders := ResponseHeaders{
		RequestTimeString: requestTime.Format(time.RFC1123),
		ClientRequestId:   requestHeaders.ClientRequestId,
		RequestTime:       requestTime,
		RequestId:         GetUUID(),
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
