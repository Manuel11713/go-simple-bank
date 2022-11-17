package api

import (
	db "github.com/Manuel11713/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server HTTP request fot our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "ok",
			"message": " api is working",
		})
	})

	router.POST("/accounts", server.createAccount)

	server.router = router

	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
