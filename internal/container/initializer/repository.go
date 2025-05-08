package initializer

import "github.com/Skliar-Il/People-API/internal/repository"

type RepositoryList struct {
	PeopleRepository repository.PeopleRepositoryInterface
}

func NewRepositoryList() *RepositoryList {
	return &RepositoryList{
		PeopleRepository: repository.NewPeopleRepository(),
	}
}
