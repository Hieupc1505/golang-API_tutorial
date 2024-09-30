package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hieupc09/simple_api/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/ping", server.Ping)
	router.GET("/todos", server.getListTodoTask)
	router.POST("/add", server.createTodoTask)
	router.POST("/update", server.updateTodoTask)
	router.GET("/delete/:id", server.deleteTodoTask)

	router.POST("/register", server.registerUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
