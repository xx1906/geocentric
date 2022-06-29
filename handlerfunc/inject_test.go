package handlerfunc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestInjectHandler(t *testing.T) {
	engine := gin.Default()
	// 随机端口
	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		t.Error(err)
		return
	}
	// 使用 插件
	engine.Use(InjectTraceHandler)
	engine.GET("/inject", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// 异步开启服务
	go func() {
		if err := engine.RunListener(listener); err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(listener.Addr().String())
	t.Log(listener.Addr().String())
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "http://"+listener.Addr().String()+"/inject", nil)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second * 3)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Header)
	defer resp.Body.Close()

	// 关闭监听
	listener.Close()
}

func TestInjectTimeOutHandler(t *testing.T) {
	engine := gin.Default()
	// 随机端口
	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		t.Error(err)
		return
	}
	// 使用 插件
	engine.Use(InjectTraceHandler)
	engine.Use(InjectTimeOutHandler(time.Second * 5))
	engine.GET("/inject", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// 异步开启服务
	go func() {
		if err := engine.RunListener(listener); err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Second * 3)

	fmt.Println(listener.Addr().String())
	t.Log(listener.Addr().String())
	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, "http://"+listener.Addr().String()+"/inject", nil)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.Header)
	defer resp.Body.Close()

	// 关闭监听
	listener.Close()
}
