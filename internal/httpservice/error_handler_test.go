package httpservice

import (
	"encoding/json"
	"github.com/dliakhov/db-query-analyzer/internal/httpservice/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/http"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		responseCode int
		responseBody any
	}{
		{
			name:         "should return error standard error",
			err:          errors.New("error"),
			responseCode: http.StatusInternalServerError,
			responseBody: fiber.Map{
				"message": "Internal Server error",
			},
		},
		{
			name:         "should return error when error is fiber error",
			err:          &fiber.Error{Message: "fiber error", Code: http.StatusBadGateway},
			responseCode: http.StatusBadGateway,
			responseBody: fiber.Map{
				"message": "fiber error",
			},
		},
		{
			name: "should return error when error is validation errors",
			err: &rest.ValidationError{
				ErrorResponses: []*rest.ErrorResponse{
					{
						FailedField: "Query",
						Tag:         "tag",
						Value:       "20",
					},
				},
			},
			responseCode: http.StatusBadRequest,
			responseBody: []interface{}{
				map[string]interface{}{
					"failed_field": "Query",
					"tag":          "tag",
					"value":        "20",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ctx fasthttp.RequestCtx
			ctx.Init(new(fasthttp.Request), nil, nil)

			err := ErrorHandler(fiber.New().AcquireCtx(&ctx), tt.err)
			if err != nil {
				assert.IsType(t, tt.err, err)
				assert.Equal(t, tt.err.Error(), err.Error())

				return
			}

			var result any
			err = json.Unmarshal(ctx.Response.Body(), &result)
			if err != nil {
				t.Fatal(err)
			}

			assert.EqualValues(t, tt.responseBody, result)
			assert.Equal(t, tt.responseCode, ctx.Response.StatusCode())
		})
	}
}
