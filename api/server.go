package api

import (
	db "github.com/bbbfishhh/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to router
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
