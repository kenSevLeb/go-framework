package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.yingxiong.com/platform/go-framework/component/auth"
	"git.yingxiong.com/platform/go-framework/component/config"
	"git.yingxiong.com/platform/go-framework/component/curl"
	"git.yingxiong.com/platform/go-framework/component/es"
	"git.yingxiong.com/platform/go-framework/component/i18n"
	log2 "git.yingxiong.com/platform/go-framework/component/log"
	"git.yingxiong.com/platform/go-framework/component/mongo"
	"git.yingxiong.com/platform/go-framework/component/mysql"
	"git.yingxiong.com/platform/go-framework/component/rbac"
	"git.yingxiong.com/platform/go-framework/component/redis"
	"git.yingxiong.com/platform/go-framework/component/task"
	"git.yingxiong.com/platform/go-framework/http"
	"git.yingxiong.com/platform/go-framework/rpc"
	"go.uber.org/dig"
)

type App struct {
	container *dig.Container
}

func New() *App {
	return &App{container: buildContainer()}
}

func (app *App) Container() *dig.Container {
	return app.container
}

// buildContainer
func buildContainer() *dig.Container {
	container := dig.New()

	// 注入配置项
	if err := config.Inject(container); err != nil {
		log.Fatalf("inject config failed: %v\n", err)
	}

	// 注入定时任务
	if err := container.Provide(task.New); err != nil {
		log.Fatalf("inject task failed: %v\n", err)
	}

	// http server
	if err := http.Inject(container); err != nil {
		log.Fatalf("inject http server failed: %v\n", err)
	}

	// elasticsearch
	if err := es.Inject(container); err != nil {
		log.Fatalf("inject elasticsearch failed: %v\n", err)
	}

	// redis
	if err := redis.Inject(container); err != nil {
		log.Fatalf("inject redis failed: %v\n", err)
	}

	// mysql
	if err := mysql.Inject(container); err != nil {
		log.Fatalf("inject mysql failed: %v\n", err)
	}

	// mongo
	if err := mongo.Inject(container); err != nil {
		log.Fatalf("inject mongo failed: %v\n", err)
	}

	// rbac
	if err := rbac.Inject(container); err != nil {
		log.Fatalf("inject rbac failed: %v\n", err)
	}

	// i18n
	if err := i18n.Inject(container); err != nil {
		log.Fatalf("inject i18n failed: %v\n", err)
	}

	// rpc
	if err := rpc.Inject(container); err != nil {
		log.Fatalf("inject rpc failed: %v\n", err)
	}

	// jwt
	if err := container.Provide(func(conf *http.Config) *auth.Jwt {
		return auth.New(conf.JwtSign)
	}); err != nil {
		log.Fatalf("inject jwt failed: %v\n", err)
	}

	// curl
	if err := container.Provide(func(conf *http.Config) *curl.Client {
		return curl.New(
			curl.WithTimeout(time.Second*time.Duration(conf.HttpRequestTimeOut)),
			curl.WithTraceHeader(conf.TraceHeader),
		)
	}); err != nil {
		log.Fatalf("inject curl client failed: %v\n", err)
	}
	return container
}

// LoadConfigFile 加载配置文件
func (app *App) LoadConfigFile(path ...string) error {
	return app.container.Invoke(func(conf *config.Config) error {
		return conf.LoadFile(path...)
	})
}

// ServeHttp 开启http服务
func (app *App) ServeHttp() {
	if err := app.container.Invoke(func(server *http.Server) {
		server.Start()
	}); err != nil {
		log.Fatalf("Serve HTTP: %v\n", err)
	}
}

// ServeTask 启动定时任务
func (app *App) serveTask() {
	_ = app.container.Invoke(func(t task.Task, conf *http.Config) {
		if conf.Task {
			t.Start()
		}
	})
}

// ServeRpc 启动rpc
func (app *App) ServeRpc() {
	if err := app.container.Invoke(func(server *rpc.Server) error {
		return server.Start()
	}); err != nil {
		log.Fatalf("Serve RPC: %v\n", err)
	}
}

// serveWS websocket
func (app *App) serveWS() {

}

// Run 启动
func (app *App) Run() {
	app.serveTask()

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// 关闭http
	_ = app.container.Invoke(func(server *http.Server) {
		server.Shutdown()
	})
	// 服务关闭后，关闭日志系统
	log2.Close()
}
