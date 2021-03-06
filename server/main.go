package main

import (
	"fmt"

	"github.com/Treblex/go-web-start/server/app"
	"github.com/Treblex/go-web-start/server/config"
	"github.com/Treblex/go-web-start/server/model"
)

func main() {

	g := app.New()

	//初始化数据链接
	if err := model.MysqlConn(config.Global.Mysql.ToString()); err != nil {
		panic(err)
	}

	// 启动
	err := g.Run(fmt.Sprintf(":%d", config.Global.Port))
	if err != nil {
		panic(err)
	}
}
