package repository

import (
	"context"
	"fmt"
	"github.com/Skliar-Il/People-API/internal/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PeopleRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, people *dto.PeopleDTO) (uuid.UUID, error)
	GetById(ctx context.Context, tx pgx.Tx, id uuid.UUID) (*dto.PeopleDTO, error)
	GetList(ctx context.Context, tx pgx.Tx, filter *dto.GetPeoplesDTO) ([]*dto.PeopleDTO, error)
	Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
	Update(ctx context.Context, tx pgx.Tx, id uuid.UUID, people *dto.PeopleDTO) error
}

type PeopleRepository struct{}

func NewPeopleRepository() *PeopleRepository {
	return &PeopleRepository{}
}

func (PeopleRepository) Create(ctx context.Context, tx pgx.Tx, people *dto.PeopleDTO) (uuid.UUID, error) {
	query := `
		insert into people(name, last_name, patronymic, age, gender, nationalize)
		values($1, $2, $3, $4, $5, $6)
		returning id
	`

	var id uuid.UUID
	err := tx.QueryRow(ctx, query,
		people.Name,
		people.LastName,
		people.Patronymic,
		people.Age,
		people.Gender,
		people.Nationalize,
	).Scan(&id)

	return id, err
}

func (PeopleRepository) GetById(ctx context.Context, tx pgx.Tx, id uuid.UUID) (*dto.PeopleDTO, error) {
	query := `
		select name, last_name, coalesce(patronymic, ''), age, gender, nationalize
		from people
		where id = $1
	`

	var people dto.PeopleDTO
	err := tx.QueryRow(ctx, query, id).Scan(
		&people.Name,
		&people.LastName,
		&people.Patronymic,
		&people.Age,
		&people.Gender,
		&people.Nationalize)

	return &people, err
}

func (PeopleRepository) GetList(ctx context.Context, tx pgx.Tx, filter *dto.GetPeoplesDTO) ([]*dto.PeopleDTO, error) {
	query := `
		select name, last_name, patronymic, age, gender, nationalize
		from people
		where 1=1`
	var args []interface{}
	var argsCount int

	addCondition := func(condition string, value interface{}) {
		if value != "" && value != 0 {
			argsCount++
			query += fmt.Sprintf(" and %s = $%d", condition, argsCount)
			args = append(args, value)
		}
	}

	addConditionLike := func(field, value string) {
		if value != "" {
			argsCount++
			query += fmt.Sprintf(" %s like $%d", field, argsCount)
			args = append(args, "%"+value+"%")
		}
	}

	addConditionLike("name", filter.Name)
	addConditionLike("last_name", filter.LastName)
	addConditionLike("patronymic", filter.Patronymic)
	addConditionLike("nationality", filter.Nationalize)
	addConditionLike("gender", filter.Gender)
	addCondition("age", filter.Age)

	if filter.Limit > 0 {
		argsCount++
		query += fmt.Sprintf(" limit $%d", argsCount)
		args = append(args, filter.Limit)

		if filter.Page > 0 {
			argsCount++
			query += fmt.Sprintf(" offset $%d", argsCount)
			args = append(args, (filter.Page-1)*filter.Limit)
		}
	}

	var peoples []*dto.PeopleDTO
	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p dto.PeopleDTO
		if err := rows.Scan(&p.Name, &p.LastName, &p.Patronymic, &p.Age, &p.Gender, &p.Nationalize); err != nil {
			return nil, err
		}
		peoples = append(peoples, &p)
	}

	return peoples, nil
}

func (PeopleRepository) Delete(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {
	query := `
		delete from people
		where id = $1`

	_, err := tx.Exec(ctx, query, id)
	return err
}

func (PeopleRepository) Update(ctx context.Context, tx pgx.Tx, id uuid.UUID, people *dto.PeopleDTO) error {
	query := `
		update people
		set name=$1, last_name=$2, patronymic=$3, age=$4, gender=$5, nationalize=$6
		where id = $7`

	_, err := tx.Exec(ctx, query,
		people.Name,
		people.LastName,
		people.Patronymic,
		people.Age,
		people.Gender,
		people.Nationalize,
		id,
	)
	return err
}
