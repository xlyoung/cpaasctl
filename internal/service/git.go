// service/git_service.go

package service

import (
    "github.com/gogf/gf/os/gfile"
    "github.com/gogf/gf/os/glog"
    "github.com/gogf/gf/os/gcfg"
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing/transport/http"
)

func ProcessGitSql(serviceName string, config interface{}) {
    // 读取配置文件中的账号名和密码
    cfg := gcfg.New()
    if err != nil {
        glog.Error(err)
        return
    }
    username := cfg.GetString("git.username")
    password := cfg.GetString("git.password")

    // 转换配置为Map类型
    cfgMap := config.(map[string]interface{})

    repoURL := cfgMap["gitlab"].(string)
    branch := cfgMap["branch"].(string)
    path := cfgMap["path"].(string)

    localPath := "sql/" + serviceName // 本地目录，例如：sql/sso

    // 检查本地目录是否存在，如果不存在则创建它
    if !gfile.Exists(localPath) {
        if err := gfile.Mkdir(localPath); err != nil {
            glog.Error("Failed to create local directory:", err)
            return
        }
    }

    auth := &http.BasicAuth{
        Username: username,
        Password: password,
    }

    r, err := git.PlainClone(localPath, false, &git.CloneOptions{
        Auth: auth,
        URL:  repoURL,
        ReferenceName: plumbing.NewBranchReferenceName(branch),
    })

    if err != nil {
        glog.Error(err)
        return
    }

    glog.Println("Repository cloned to:", r.Worktree.Filesystem.Root())

    // 这里可以处理SQL目录，例如复制到指定位置、解析SQL文件等
    // path 变量指定的是你在配置文件中定义的 SQL 文件目录
    // localPath 变量指定的是 Git 仓库本地克隆的目录
    // 可以在这里进行相应的处理

    glog.Println("SQL directory fetched successfully:", serviceName)
}
