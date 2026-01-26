// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"likegozero/example/internal/config"
	"likegozero/example/internal/handler"
	"likegozero/example/internal/middleware"
	"likegozero/example/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	// 全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.Info("func global middleware----pre")
			next(w, r)
			logx.Info("func global middleware----after")
		}
	})
	server.Use(middleware.GlobalMiddleware)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
