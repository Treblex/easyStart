package wechat

import (
	"EK-Server/config"
	"EK-Server/model"
	"EK-Server/util"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var mp = &config.Global.Wechat

// Init 初始化
func Init(g *echo.Group) {
	baseURL := "/wechat"
	wechat := g.Group(baseURL)

	wechat.GET("/jsApiConfig", jsAPIConfig)
	wechat.GET("/wechat_redirect", wechatRedirect)
	wechat.GET("/login", login)

	wechat.GET("/info", userInfo)
	wechat.GET("/signature", signatureCheck)               //配置接口token验证
	wechat.POST("/signature", handleWechatMessage)         //服务token验证，验证成功之后微信会post用户消息 和 事件到这个接口
	wechat.GET("/sendTemplateMsg", sendTemplateMsgHandler) //发送模版消息

}

// 微信登陆
func login(c echo.Context) (err error) {
	code := c.QueryParam("code")
	if code == "" {
		return util.JSONErr(c, nil, "code不可空")
	}

	// 请求微信登陆返回信息
	WeChatLogin, msg := sendCodeToWechatServer(code)
	if msg == "" {
		return util.JSONErr(c, WeChatLogin, msg)
	}
	// 更新过期时间
	WeChatLogin.ExpiresIn += time.Now().Unix()
	// 查询微信用户数据库表
	WeChatUser, msg, errCode := wechatDoLogin(WeChatLogin)
	if msg != "" {
		return util.JSON(c, WeChatUser, msg, errCode)
	}
	// 如果uid存在则查询用户 返回用户信息 生产jwt token
	// ...

	return util.JSONSuccess(c, WeChatUser, "登陆成功")
}

// 获取微信用户信息  是否关注
func userInfo(c echo.Context) (err error) {
	wechatID := c.QueryParam("id")
	if wechatID == "" {
		return util.JSONErr(c, nil, "用户id不可空")
	}
	newID, err := strconv.Atoi(wechatID)
	if err != nil {
		return util.JSONErr(c, err, "参数错误")
	}
	if newID < 1 {
		return util.JSONErr(c, nil, "用户id不可为空")
	}
	user := model.WechatOauth{BaseControll: model.BaseControll{ID: uint(newID)}}

	db := model.DB
	if db.Find(&user).RecordNotFound() {
		return util.JSONErr(c, err, "未找到用户")
	}
	// token := user.AccessToken
	if time.Now().Unix()-user.CreatedAt.Unix() > 3600*24*10 || user.Nickname == "" || user.Headimgurl == "" {
		info, msg := updateWechatInfo(&user, false)
		if msg != "" {
			return util.JSONErr(c, nil, msg)
		}

		db.Model(&user).Updates(&info)
	}

	return util.JSONSuccess(c, user, "获取成功")
}

// 换取微信网页登陆授权链接
func wechatRedirect(c echo.Context) (err error) {
	host := "http://"
	if c.IsTLS() {
		host = "https://"
	}
	host += c.Request().Host
	redirectURI := host + "/wechat/login"

	callbackURL := c.QueryParam("callback")
	if callbackURL != "" {
		redirectURI = callbackURL
	}

	redirectURI = url.PathEscape(redirectURI)
	urlStr := mp.LoginRedirect(redirectURI)
	return c.Redirect(http.StatusMovedPermanently, urlStr)
}

// 换取微信分享 jsapi授权
func jsAPIConfig(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return util.JSONErr(c, nil, "url不可空")
	}
	conf, err := mp.JsAPIConfig(url)
	if err != nil {
		return util.JSONErr(c, err, "配置错误")
	}
	return util.JSONSuccess(c, conf, "请求成功")
}
