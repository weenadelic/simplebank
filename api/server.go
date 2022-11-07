package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/weenadelic/simplebank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Account routes
	router.GET("/accounts", server.listAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.POST("/accounts", server.createAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PATCH("/accounts", server.updateAccount)

	// Next routes

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
