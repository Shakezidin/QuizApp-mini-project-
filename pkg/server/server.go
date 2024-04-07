package server

import "github.com/gin-gonic/gin"

// ServerStruct represents the server structure.
type ServerStruct struct {
	R *gin.Engine // Gin engine instance.
}

// StartServer starts the server on the specified port.
func (s *ServerStruct) StartServer(port string) {
	s.R.Run("localhost:" + port) // Run the server on the specified port.
}

// Server creates a new server instance.
func Server() *ServerStruct {
	engine := gin.Default() // Create a new Gin engine instance.

	return &ServerStruct{
		R: engine, // Initialize the server structure with the Gin engine.
	}
}
