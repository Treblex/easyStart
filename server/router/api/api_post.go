package api

import (
	"EK-Server/model"
	"EK-Server/util"
	"EK-Server/util/middleware"
	"strings"

	"github.com/labstack/echo"
)

var modelPost model.Articles

// Init 初始化
func post(g *echo.Group) {
	modelPost.BaseControll.Model = &modelPost
	post := g.Group("/posts")

	//list
	post.GET("", modelPost.List)
	//detail
	post.GET("/:id", func(c echo.Context) error {
		return modelPost.BaseControll.Detail(c, "文章不存在")
	})
	//del
	post.GET("/:id/del", modelPost.BaseControll.Delete)

	//add todo:整理新建和更新到同一个方法
	post.POST("", func(c echo.Context) error {
		empty := func(err interface{}, msg string) error {
			return util.JSONErr(c, err, msg)
		}
		name := c.FormValue("name")
		if strings.Trim(name, " ") == "" {
			return empty(nil, "作者不可以空")
		}
		return util.JSONSuccess(c, name, "添加成功")
	}, middleware.AdminJWT)
	// Update
	post.PUT("/:id", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "")
	})

	// Actions
	// 点赞文章
	post.GET("/:id/like", func(c echo.Context) error {
		return util.JSONSuccess(c, nil, "点赞成功")
	}, middleware.AdminJWT)

}