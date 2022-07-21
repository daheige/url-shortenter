package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/houseme/url-shortenter/app/api/internal/packed"

	_ "github.com/houseme/url-shortenter/app/api/internal/logic"

	"github.com/houseme/url-shortenter/app/api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
