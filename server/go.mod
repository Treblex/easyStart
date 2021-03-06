module github.com/Treblex/go-web-start/server

go 1.13

require (
	github.com/Treblex/go-web-template v1.0.5
	github.com/Treblex/simple-daily v0.0.0-20201214014444-2874a70d28a6
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/static v0.0.0-20200916080430-d45d9a37d28e
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.1.4
	github.com/gorilla/websocket v1.4.2
	gorm.io/gorm v1.20.9
)

// replace github.com/Treblex/go-web-template => ../../go-web-template
