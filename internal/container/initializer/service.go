package initializer

import (
	"github.com/Skliar-Il/People-API/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceList struct {
	PeopleService service.PeopleServiceInterface
}

func NewServiceList(repository *RepositoryList, client *ClientList, dbPool *pgxpool.Pool) *ServiceList {
	return &ServiceList{
		PeopleService: service.NewPeopleService(dbPool, client.PeopleClient, repository.PeopleRepository),
	}
}
