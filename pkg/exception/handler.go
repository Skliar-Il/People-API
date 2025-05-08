package exception

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func Handler(ctx fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = ctx.Status(code).JSON(fiber.Map{"message": e.Error()})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "internal server error"})
	}

	return nil
}
