package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"cpaasctl/internal/controller/hello"
	"gitlab.hycyg.com/paas-tools/cpaasctl/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})

		    // 注册 GitController
			s.Group("/git", func(group *ghttp.RouterGroup) {
				group.ALL("/", controllers.GetSql)
			})

			s.Run()
			return nil
		},
	}
)
