package httpservice

import (
	"github.com/dliakhov/db-query-analyzer/internal/httpservice/rest"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		return ctx.Status(e.Code).JSON(fiber.Map{
			"message": e.Message,
		})
	}

	if e, ok := err.(*rest.ValidationError); ok {
		return ctx.Status(http.StatusBadRequest).JSON(e.ErrorResponses)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Internal Server error",
	})
}
