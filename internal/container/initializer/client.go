package initializer

import (
	"github.com/Skliar-Il/People-API/internal/transport/http/client"
	fiberclient "github.com/gofiber/fiber/v3/client"
)

type ClientList struct {
	PeopleClient client.PeopleClientInterface
}

func NewClientList(cfg *client.Config) *ClientList {
	client_ := fiberclient.New()
	
	return &ClientList{
		PeopleClient: client.NewPeopleClient(client_, &cfg.PeopleConfig),
	}
}
