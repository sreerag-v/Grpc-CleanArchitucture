package http

import (
	handlerInterface "Clean-Grpc/pkg/api/handler/interface"
	"Clean-Grpc/pkg/api/routes"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}


func NewServerHTTP(userHandler handlerInterface.UserHandler,)*ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())
	// Auth middleware
	// For Routes

	routes.UserRoutes(engine.Group("/user"), userHandler)
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
