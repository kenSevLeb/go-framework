package main

import (
	"fmt"
	"kenSevLeb/go-framework/app"
	"kenSevLeb/go-framework/component/i18n"
	"kenSevLeb/go-framework/component/log"
	"kenSevLeb/go-framework/component/paginate"
	"kenSevLeb/go-framework/errors"
	"kenSevLeb/go-framework/http"
	"kenSevLeb/go-framework/http/middleware"
	"kenSevLeb/go-framework/http/response"
	"kenSevLeb/go-framework/util/secure"
	"github.com/gin-gonic/gin"
)

func main() {
	cmd := app.New()
	//secure.Panic(cmd.LoadConfigFile("/usr/app.yaml"))

	secure.Panic(cmd.Container().Invoke(loadRouter))
	cmd.ServeHttp()

	cmd.Run()
}

func loadRouter(router *gin.Engine, config *http.Config, i18n *i18n.Bundle) {
	router.Use(middleware.Recover, middleware.FaviconFilter, middleware.RequestLog)

	router.GET("/swagger/*any", middleware.SwaggerHandler(config.Swagger))
	router.GET("/check", func(ctx *gin.Context) {
		fmt.Println(ctx.ClientIP())
		response.Wrap(ctx).Success("hello")
	})
	router.GET("/paginate", func(ctx *gin.Context) {
		pagination := paginate.Paginate(10, 1, 5)
		response.Wrap(ctx).Paginate(nil, &pagination)
	})
	router.GET("/i18n", func(ctx *gin.Context) {
		lang := ctx.Query("lang")
		response.Wrap(ctx).Success(i18n.MustLocalize(lang, "Hello"))
	})
	router.GET("/panic", func(context *gin.Context) {
		panic(errors.New(1001, "test").Hide("some error"))
	})
	router.GET("log", func(context *gin.Context) {
		log.Info("info", nil)
		log.Debug("debug", log.Content{"other": "aa"})
		log.Error("error", log.Content{"xx": 123})
	})
}

