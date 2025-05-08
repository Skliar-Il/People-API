package container

import (
	"context"
	"github.com/Skliar-Il/People-API/internal/config"
	"github.com/Skliar-Il/People-API/internal/container/initializer"
	"github.com/Skliar-Il/People-API/internal/container/server"
	"github.com/Skliar-Il/People-API/pkg/database"
	"log"
)

func NewApp() {
	cfg, err := config.New()
	ctx := context.Background()
	if err != nil {
		log.Fatalf("failed load config: %v", err)
	}

	dbPool, err := database.New(ctx, cfg.DataBase)
	if err != nil {
		log.Fatalf("init database error: %v", err)
	}

	repositoryList := initializer.NewRepositoryList()
	clientList := initializer.NewClientList(&cfg.Client)
	serviceList := initializer.NewServiceList(repositoryList, clientList, dbPool)
	server.Serve(cfg, serviceList)
}
