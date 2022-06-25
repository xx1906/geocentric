package handlerfunc

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dijkvy/geocentric/tag"
)

const (
	xTrace   = "X-Trace"
	xTimeout = "X-Age"
)

// InjectTraceHandler 注入 value handler func
func InjectTraceHandler(ctx *gin.Context) {
	c := ctx.Request.Context()
	if tag.Extract(c) != nil {
		ctx.Next()
		return
	}

	value := strings.ToLower(base64.StdEncoding.EncodeToString([]byte(uuid.NewString())))
	c = tag.Inject(c, value)
	ctx.Request = ctx.Request.WithContext(c)
	ctx.Request.Header.Set(xTrace, value)
	ctx.Header(xTrace, value)
	ctx.Next()
}

// InjectTimeOutHandler 向请求上下文中注入过期时间
func InjectTimeOutHandler(timeoutMillSec time.Duration) gin.HandlerFunc {
	// fixed 修复共享变量 (timeoutMillSec) 的 bug
	timeoutMillSec = timeoutMillSec * time.Millisecond
	return func(ctx *gin.Context) {
		var timeoutMill = timeoutMillSec
		c := ctx.Request.Context()
		deadline, ok := c.Deadline()
		if ok && deadline.Before(deadline) {
			if err := ctx.AbortWithError(500, http.ErrHandlerTimeout); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "InjectTimeOutHandler abort but error %s", err)
			}
			return
		}
		if ok {
			// todo 这里需要处理 header 的 trace 和 time 信息被删除的情况
			ctx.Next()
			return
		}
		// 从 header 中提取过期时间, 如果没有或者格式不正确不做处理
		if v := ctx.Request.Header.Get(xTimeout); v != "" {
			if parseInt, err := strconv.ParseInt(v, 10, 64); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "InjectTimeOutHandler parse timeoutMillSec format error %s", err)
			} else {
				timeoutMill = time.Duration(parseInt)
			}
		}
		// 注入过期时间
		timeoutValue := fmt.Sprintf("%d", timeoutMill)
		var cancel func()
		c, cancel = context.WithTimeout(c, timeoutMill)
		go func() {
			select {
			case <-c.Done():
				cancel()
			}
		}()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Request.Header.Set(xTimeout, timeoutValue)
		ctx.Header(xTimeout, timeoutValue)
		ctx.Next()
	}
}
