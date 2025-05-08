package controller

import (
	"fmt"
	"github.com/Skliar-Il/People-API/internal/dto"
	"github.com/Skliar-Il/People-API/internal/service"
	"github.com/Skliar-Il/People-API/pkg/logger"
	"github.com/Skliar-Il/People-API/pkg/render"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PeopleHandler struct {
	PeopleService service.PeopleServiceInterface
}

func NewPeopleHandler(peopleService service.PeopleServiceInterface) *PeopleHandler {
	return &PeopleHandler{
		PeopleService: peopleService,
	}
}

func (p *PeopleHandler) CreatePeople(c fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(c.Context())

	var people dto.CreatePeopleDTO
	if err := c.Bind().JSON(&people); err != nil {
		localLogger.Info(c.Context(), "validation error", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, fmt.Sprintf("validation error: %v", err))
	}

	id, err := p.PeopleService.CreatePeople(c.Context(), &people)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(id)
}

func (p *PeopleHandler) GetPeople(c fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(c.Context())

	peopleIDStr := c.Params("id")
	peopleID, err := uuid.Parse(peopleIDStr)
	if err != nil {
		localLogger.Info(c.Context(), "pars id error", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, "path param id mast be uuid v4")
	}

	people, err := p.PeopleService.GetPeople(c.Context(), peopleID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(people)
}

func (p *PeopleHandler) UpdatePeople(c fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(c.Context())

	peopleIDStr := c.Params("id")
	peopleID, err := uuid.Parse(peopleIDStr)
	if err != nil {
		localLogger.Info(c.Context(), "pars id error", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, "path param id mast be uuid v4")
	}
	var people dto.PeopleDTO
	if err := c.Bind().JSON(&people); err != nil {
		localLogger.Info(c.Context(), "parse body error", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, err.Error())
	}

	if err := p.PeopleService.UpdatePeople(c.Context(), peopleID, &people); err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

func (p *PeopleHandler) DeletePeople(c fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(c.Context())

	peopleIDStr := c.Params("id")
	peopleID, err := uuid.Parse(peopleIDStr)
	if err != nil {
		localLogger.Info(c.Context(), "pars id error", zap.Error(err))
		return render.Error(fiber.ErrUnprocessableEntity, "path param id mast be uuid v4")
	}

	if err := p.PeopleService.DeletePeople(c.Context(), peopleID); err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}
