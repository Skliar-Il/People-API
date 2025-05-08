package service

import (
	"context"
	"github.com/Skliar-Il/People-API/internal/dto"
	"github.com/Skliar-Il/People-API/internal/repository"
	"github.com/Skliar-Il/People-API/internal/transport/http/client"
	"github.com/Skliar-Il/People-API/pkg/database"
	"github.com/Skliar-Il/People-API/pkg/logger"
	"github.com/Skliar-Il/People-API/pkg/render"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PeopleServiceInterface interface {
	CreatePeople(ctx context.Context, people *dto.CreatePeopleDTO) (*dto.PeopleIdDTO, error)
	GetPeople(ctx context.Context, id uuid.UUID) (*dto.PeopleDTO, error)
	UpdatePeople(ctx context.Context, id uuid.UUID, people *dto.PeopleDTO) error
	DeletePeople(ctx context.Context, id uuid.UUID) error
}

type PeopleService struct {
	DBPool           *pgxpool.Pool
	PeopleRepository repository.PeopleRepositoryInterface
	PeopleClient     client.PeopleClientInterface
}

func NewPeopleService(
	dbPool *pgxpool.Pool,
	client_ client.PeopleClientInterface,
	peopleRepository repository.PeopleRepositoryInterface,
) *PeopleService {
	return &PeopleService{
		DBPool:           dbPool,
		PeopleClient:     client_,
		PeopleRepository: peopleRepository,
	}
}

func (p *PeopleService) CreatePeople(ctx context.Context, people *dto.CreatePeopleDTO) (*dto.PeopleIdDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	tx, err := p.DBPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "failed start tx")
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	defer database.RollbackTx(ctx, tx)

	age, err := p.PeopleClient.GetAge(ctx, people.Name)
	if err != nil {
		localLogger.Error(ctx, "get people age error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	gender, err := p.PeopleClient.GetGender(ctx, people.Name)
	if err != nil {
		localLogger.Error(ctx, "get people gender error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	nationalize, err := p.PeopleClient.GetNationalize(ctx, people.Name)
	if err != nil {
		localLogger.Error(ctx, "get people nationalize error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}

	fullPeopleDTO := dto.PeopleDTO{
		Name:        people.Name,
		LastName:    people.LastName,
		Patronymic:  people.Patronymic,
		Gender:      gender,
		Nationalize: nationalize,
		Age:         age,
	}

	id, err := p.PeopleRepository.Create(ctx, tx, &fullPeopleDTO)
	if err != nil {
		localLogger.Error(ctx, "create people database error", zap.Error(err))
	}
	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit database changes error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}

	return &dto.PeopleIdDTO{ID: id}, nil
}

func (p *PeopleService) GetPeople(ctx context.Context, id uuid.UUID) (*dto.PeopleDTO, error) {
	localLogger := logger.GetLoggerFromCtx(ctx)
	tx, err := p.DBPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "failed start tx")
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}
	defer database.RollbackTx(ctx, tx)

	people, err := p.PeopleRepository.GetById(ctx, tx, id)
	if err != nil {
		pgError := database.ValidatePgxError(err)
		if pgError != nil && pgError.Type == database.TypeNoRows {
			localLogger.Info(ctx, "people not found", zap.Error(err))
			return nil, render.Error(fiber.ErrNotFound, "people not found")
		}
		localLogger.Error(ctx, "get people database error", zap.Error(err))
		return nil, render.Error(fiber.ErrInternalServerError, "")
	}

	return people, err
}

func (p *PeopleService) UpdatePeople(ctx context.Context, id uuid.UUID, people *dto.PeopleDTO) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	tx, err := p.DBPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "failed start tx")
		return render.Error(fiber.ErrInternalServerError, "")
	}
	defer database.RollbackTx(ctx, tx)

	if err := p.PeopleRepository.Update(ctx, tx, id, people); err != nil {
		localLogger.Error(ctx, "update people database error", zap.Error(err))
		return render.Error(fiber.ErrInternalServerError, "")
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit database changes error", zap.Error(err))
		return render.Error(fiber.ErrInternalServerError, "")
	}

	return nil
}

func (p *PeopleService) DeletePeople(ctx context.Context, id uuid.UUID) error {
	localLogger := logger.GetLoggerFromCtx(ctx)
	tx, err := p.DBPool.Begin(ctx)
	if err != nil {
		localLogger.Error(ctx, "failed start tx")
		return render.Error(fiber.ErrInternalServerError, "")
	}
	defer database.RollbackTx(ctx, tx)

	if err := p.PeopleRepository.Delete(ctx, tx, id); err != nil {
		localLogger.Error(ctx, "delete people database error", zap.Error(err))
		return render.Error(fiber.ErrInternalServerError, "")
	}

	if err := tx.Commit(ctx); err != nil {
		localLogger.Error(ctx, "commit database changes error", zap.Error(err))
		return render.Error(fiber.ErrInternalServerError, "")
	}

	return nil
}
