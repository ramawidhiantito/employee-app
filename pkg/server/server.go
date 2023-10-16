package server

import (
	"github.com/gin-gonic/gin"
)

// New creates and returns a new Gin server.
func New() *gin.Engine {
	r := gin.Default()
	return r
}

// Run starts the server.
func Run(r *gin.Engine, address string) error {
	return r.Run(address)
}
