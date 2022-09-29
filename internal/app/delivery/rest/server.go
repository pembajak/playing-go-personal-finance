package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	httppkg "github.com/pembajak/personal-finance/internal/pkg/http"
)

type Server interface {
	Router(delivery Delivery) (w httppkg.Router)
	GetHTTPServer() *http.Server
	GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool)
}

type server struct {
	Addr         string
	Delivery     Delivery
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func NewServer(addr string, delivery Delivery, readTimeout time.Duration, writeTimeout time.Duration, idleTimeout time.Duration) Server {
	return &server{
		Addr:         addr,
		Delivery:     delivery,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

func (s *server) GetHTTPServer() *http.Server {
	return &http.Server{
		Addr:         s.Addr,
		Handler:      s.Router(s.Delivery),
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}
}

// GracefullShutdown ...
func (s *server) GracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}
