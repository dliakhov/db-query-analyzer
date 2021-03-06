package httpservice

import (
	"github.com/dliakhov/db-query-analyzer/internal/httpservice/rest"
	"github.com/dliakhov/db-query-analyzer/internal/models"
	"github.com/dliakhov/db-query-analyzer/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type QueryAnalyzerHandler struct {
	repo     repository.QueryAnalyzerRepository
	validate *validator.Validate
}

func NewQueryAnalyzer(repo repository.QueryAnalyzerRepository) *QueryAnalyzerHandler {
	return &QueryAnalyzerHandler{
		repo:     repo,
		validate: validator.New(),
	}
}

func (h *QueryAnalyzerHandler) GetQueries(ctx *fiber.Ctx) error {
	var queryRequest models.QueryRequest
	err := ctx.QueryParser(&queryRequest)
	if err != nil {
		return err
	}

	errors := h.validateStruct(queryRequest)
	if errors != nil {
		return rest.NewValidationError(errors)
	}

	queries, err := h.repo.GetDatabaseQueryInfo(queryRequest)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(queries)
}

var validate = validator.New()

func (h *QueryAnalyzerHandler) validateStruct(queryRequest models.QueryRequest) []*rest.ErrorResponse {
	var errors []*rest.ErrorResponse
	err := validate.Struct(queryRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element rest.ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
