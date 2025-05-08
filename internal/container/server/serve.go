package server

import (
	"context"
	"fmt"
	"github.com/Skliar-Il/People-API/internal/config"
	"github.com/Skliar-Il/People-API/internal/container/initializer"
	"github.com/Skliar-Il/People-API/internal/transport/http/controller"
	"github.com/Skliar-Il/People-API/pkg/exception"
	"github.com/Skliar-Il/People-API/pkg/logger"
	pkgvalidator "github.com/Skliar-Il/People-API/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"log"
	"os/signal"
	"syscall"
)

func Serve(cfg *config.Config, services *initializer.ServiceList) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := fiber.New(
		fiber.Config{
			ErrorHandler:    exception.Middleware,
			StructValidator: pkgvalidator.Validator{Validator: validator.New()},
		},
	)

	server.Use(cors.New())
	server.Use(logger.Middleware(&cfg.Logger))

	controller.NewController(server, services)

	go func() {
		if err := server.Listen(fmt.Sprintf(":%s", cfg.Server.HttpPort)); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case <-ctx.Done():
		if err := server.Shutdown(); err != nil {
			log.Fatalf("stop server error: %v", err)
		}
		log.Println("Server is stopped")
	}
}
