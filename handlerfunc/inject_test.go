package handlerfunc

import (
	"context"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInjectHandler(t *testing.T) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/index", nil)
	if err != nil {
		t.Error(err)
		return
	}
	request.Response = &http.Response{Header: map[string][]string{}}

	handlerCtx := &gin.Context{Request: request}
	InjectTraceHandler(handlerCtx)
	t.Log(handlerCtx.Request.Header)
	t.Log(handlerCtx.Request.Response.Header)
}

func TestInjectTimeOutHandler(t *testing.T) {
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/index", nil)
	if err != nil {
		t.Error(err)
		return
	}
	request.Response = &http.Response{Header: map[string][]string{}}

	handlerCtx := &gin.Context{Request: request}
	InjectTraceHandler(handlerCtx)
	InjectTimeOutHandler(100)(handlerCtx)
	t.Log(handlerCtx.Request.Header)

	handlerCtx.Request = handlerCtx.Request.Clone(context.TODO())
	handlerCtx.Request.Header.Set(xTimeout, "9000xx")
	InjectTimeOutHandler(30)(handlerCtx)
	t.Log(handlerCtx.Request.Header)
	t.Log(handlerCtx.Request.Response.Header)
}
