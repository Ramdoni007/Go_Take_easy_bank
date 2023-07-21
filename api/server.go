package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/ramdoni007/Take_Easy_Bank/db/sqlc"
)

// Server serves HTTP requestfor our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a New HTTP server and SetUp Routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/tranfers", server.createTransfer)

	server.router = router
	return server
}

// Start Server the HTTP serveron a specific addres

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}

// error function gin.H
func errorResponse(err error) gin.H {

	return gin.H{"Error": err.Error()}
}
