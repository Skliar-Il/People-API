package server

import (
	"context"
	"fmt"
	"github.com/Skliar-Il/People-API/internal/config"
	"github.com/Skliar-Il/People-API/internal/container/initializer"
	"github.com/Skliar-Il/People-API/internal/transport/http/controller"
	"github.com/Skliar-Il/People-API/pkg/exception"
	"github.com/Skliar-Il/People-API/pkg/logger"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/storage/redis/v3"
	"log"
	"os/signal"
	"syscall"
)

func Serve(cfg *config.Config, storage *redis.Storage, services *initializer.ServiceList) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server := fiber.New(
		fiber.Config{
			ErrorHandler: exception.Handler,
		},
	)

	server.Use(cors.New())
	server.Use(cache.New(cache.Config{
		Storage: storage,
	}))
	server.Use(logger.Middleware())

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
