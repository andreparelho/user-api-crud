package server

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/andreparelho/user-api-crud/internal/user"
)

func NewServer(userService user.UserService, logger *slog.Logger) *Server {
	app := CreateRouter(userService)

	return &Server{
		App:    app,
		Logger: logger,
	}
}

func (s *Server) Start(ctx context.Context, port string) error {
	go func() {
		<-ctx.Done()
		s.Logger.Info("shutting down http server")
		_ = s.App.Shutdown()
	}()

	s.Logger.Info("starting server", "port", port)
	return s.App.Listen(":" + port)
}

func Shutdown() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-ch
		cancel()
	}()

	return ctx
}
