package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Inject(container *dig.Container) error {
	if err := container.Provide(NewServer); err != nil {
		return err
	}

	return container.Provide(func(server *Server) *gin.Engine {
		return server.Router
	})
}
