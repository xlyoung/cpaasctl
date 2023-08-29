package controller


import (
    "github.com/gogf/gf/net/ghttp"
    "gitlab.hycgy.com/paas-tools/cpaasctl/internal/service"
)

func GetSql(r *ghttp.Request) {
    // 读取配置文件
    cfg := g.Cfg().GetMap("sql")

    // 调用 Service 层的方法，传递配置信息
    service.ProcessGitSql(cfg)

    r.Response.WriteJson(g.Map{
        "success": true,
        "message": "Git repositories processed successfully",
    })
}