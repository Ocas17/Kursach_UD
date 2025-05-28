package Kursach_UD

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	// Create a new CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500", "http://localhost:5500"}, // Add your frontend origins here
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true, // Enable debugging for development
	})

	// Wrap your handler with the CORS middleware
	handlerWithCORS := c.Handler(handler)

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handlerWithCORS,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
