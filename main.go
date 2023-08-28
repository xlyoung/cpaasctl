package main

import (
	_ "cpaasctl/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cpaasctl/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
