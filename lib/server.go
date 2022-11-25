package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func NewServer() Server {
	engine := gin.New()
	return Server{Gin: engine}
}

func (s *Server) Run() {
	err := s.Gin.Run()
	if err != nil {
		panic(fmt.Sprintf("Server Failed to run, %v", err))
	}
}
