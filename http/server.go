package http

import (
	"context"
	"fmt"
	"kenSevLeb/go-framework/component/event"
	"kenSevLeb/go-framework/http/middleware"
	"kenSevLeb/go-framework/http/response"
	"kenSevLeb/go-framework/http/validator"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

// Server Web服务管理器
type Server struct {
	// gin的路由
	Router *gin.Engine

	instance *http.Server

	conf *Config
}

// NewServer 实例化server
func NewServer(conf *Config) *Server {
	router := gin.New()
	// use global trace middleware
	router.Use(gin.Recovery(), middleware.Trace(conf.TraceHeader))

	// 404
	router.NoRoute(notFoundHandler)
	router.NoMethod(notFoundHandler)

	instance := &Server{
		Router: router,
		conf:   conf,
	}

	return instance
}

// run before start
func (server *Server) beforeStart() {
	// 设置自定义验证器,支持字段命名
	binding.Validator = validator.New()

	// http启动前
	_ = event.Trigger(event.BeforeHttpStart, nil)

	if server.conf.LogDebug { // 启动debug日志
		logrus.SetLevel(logrus.DebugLevel)
	}

	if server.conf.Pprof { // 启动pprof
		pprof.Register(server.Router)
	}
}

// Run http server
func (server *Server) Start() {
	server.beforeStart()

	completeHost := net.JoinHostPort("", strconv.Itoa(server.conf.Port))

	// 平滑关闭
	srv := &http.Server{
		Addr:    completeHost,
		Handler: server.Router,
	}
	go func() {
		log.Printf("Listening and serving HTTP on %s\n", completeHost)

		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("%s\n", err)
		}

		// after start
		_ = event.Trigger(event.AfterHttpStart, nil)
	}()
}

// Shutdown 关闭http
func (server *Server) Shutdown() {
	if server.instance == nil {
		return
	}
	_ = event.Trigger(event.BeforeHttpShutdown, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.instance.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exited")
}

// notFoundHandler 404
func notFoundHandler(ctx *gin.Context) {
	response.Wrap(ctx).Error(404, fmt.Sprintf("404 Not Found: %s", ctx.Request.RequestURI))
}
