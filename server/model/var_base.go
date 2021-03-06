package model

import (
	"github.com/Treblex/go-web-start/server/utils"
	"github.com/gin-gonic/gin"
)

// GetUserOrLogin GetUser
func GetUserOrLogin(c *gin.Context) *User {
	_user, exists := c.Get("user")
	if !exists {
		utils.Error(utils.JSON(utils.AuthedError, "请先登录1", nil))
	}
	user, ok := _user.(*User)
	if !ok {
		utils.Error(utils.JSON(utils.AuthedError, "请先登录", nil))
	}
	return user
}

// GetUserOrEmpty GetUser
func GetUserOrEmpty(c *gin.Context) *User {
	_user, exists := c.Get("user")
	if exists {
		user, ok := _user.(*User)
		if ok {
			return user
		}
	}
	return &User{}
}
