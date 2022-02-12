package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

//目标 1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
//用errgroup 实现任务并行处理并返回结果
func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{Addr: addr, Handler: handler}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

type Handler struct {
}

func (h *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	stop := make(chan struct{})
	g.Go(func() error {
		return server("0.0.0.0:8080", &Handler{}, stop)
	})

	g.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			close(stop)
			return errors.New("用户关闭")
		case <-ctx.Done():
			return errors.New("http服务关闭")
		}

	})

	err := g.Wait()
	fmt.Printf("关闭服务来源：%v", err)

}
