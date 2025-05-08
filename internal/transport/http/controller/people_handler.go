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

// CreatePeople godoc
// @Summary Создать нового человека
// @Tags People
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param input body dto.CreatePeopleDTO true "Данные для создания"
// @Success 201 {object} uuid.UUID "ID созданного человека"
// @Failure 400
// @Failure 422
// @Router /people [post]
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

// GetPeople godoc
// @Summary Получить человека по ID
// @Tags People
// @Produce json
// @Param id path string true "UUID человека" format(uuid)
// @Success 200 {object} dto.PeopleFullDTO
// @Failure 400
// @Failure 404
// @Router /people/{id} [get]
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

// GetPeopleList godoc
// @Summary Получить список людей с фильтрацией
// @Tags People
// @Produce json
// @Param name query string false "Фильтр по имени"
// @Param last_name query string false "Фильтр по фамилии"
// @Param patronymic query string false "Фильтр по отчеству"
// @Param age query int false "Фильтр по возрасту"
// @Param nationalize query string false "Фильтр по национальности"
// @Param gender query string false "Фильтр по полу"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Лимит записей" default(10)
// @Success 200 {array} dto.PeopleFullDTO
// @Failure 400
// @Router /people [get]
func (p *PeopleHandler) GetPeopleList(c fiber.Ctx) error {
	localLogger := logger.GetLoggerFromCtx(c.Context())

	var filters dto.GetPeoplesDTO
	if err := c.Bind().Query(&filters); err != nil {
		localLogger.Info(c.Context(), "invalid query params", zap.Error(err))
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	peoples, err := p.PeopleService.GetPeoples(c.Context(), &filters)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(peoples)
}

// UpdatePeople godoc
// @Summary Обновить данные человека
// @Tags People
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "UUID человека" format(uuid)
// @Param input body dto.PeopleDTO true "Обновляемые данные"
// @Success 204
// @Failure 400
// @Failure 404
// @Router /people/{id} [put]
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

// DeletePeople godoc
// @Summary Удалить человека
// @Tags People
// @Security ApiKeyAuth
// @Param id path string true "UUID человека" format(uuid)
// @Success 204
// @Failure 400
// @Failure 404
// @Router /people/{id} [delete]
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
