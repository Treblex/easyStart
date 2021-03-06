package ws

import (
	"github.com/Treblex/go-web-start/server/router/api/ws/chat"
	"github.com/gin-gonic/gin"
)

// Init Init
func Init(g *gin.RouterGroup) {
	baseURL := "/ws"
	app := g.Group(baseURL)
	app.GET("", chat.WsServer)
}
