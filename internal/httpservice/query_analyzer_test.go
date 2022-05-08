package httpservice

import (
	"encoding/json"
	"github.com/dliakhov/db-query-analyzer/internal/httpservice/rest"
	"github.com/dliakhov/db-query-analyzer/internal/models"
	"github.com/dliakhov/db-query-analyzer/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/url"
	"testing"
)

func TestQueryAnalyzerHandler_GetQueries(t *testing.T) {
	type fields struct {
		repo func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository
	}
	type args struct {
		queryParams map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.DatabaseQueryInfo
		wantErr error
	}{
		{
			name: "should return validation error when page is missing",
			fields: fields{
				repo: func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository {
					analyzerRepository := repository.NewMockQueryAnalyzerRepository(ctrl)
					return analyzerRepository
				},
			},
			args: args{
				queryParams: map[string]string{
					"query_type": "select",
				},
			},
			wantErr: rest.NewValidationError([]*rest.ErrorResponse{
				{
					FailedField: "Page",
					Tag:         "required",
					Value:       "",
				},
				{
					FailedField: "Size",
					Tag:         "required",
					Value:       "",
				},
			}),
		},
		{
			name: "should return successfully one query successfully",
			fields: fields{
				repo: func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository {
					analyzerRepository := repository.NewMockQueryAnalyzerRepository(ctrl)
					analyzerRepository.EXPECT().GetDatabaseQueryInfo(gomock.Any()).Return([]models.DatabaseQueryInfo{
						{
							Model: models.Model{
								ID: 1,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 20,
						},
					}, nil)
					return analyzerRepository
				},
			},
			args: args{
				queryParams: map[string]string{
					"query_type": "select",
					"page":       "1",
					"size":       "1",
				},
			},
			want: []models.DatabaseQueryInfo{
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 20,
				},
			},
		},
		{
			name: "should return successfully one query successfully without query_type parameter",
			fields: fields{
				repo: func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository {
					analyzerRepository := repository.NewMockQueryAnalyzerRepository(ctrl)
					analyzerRepository.EXPECT().GetDatabaseQueryInfo(gomock.Any()).Return([]models.DatabaseQueryInfo{
						{
							Model: models.Model{
								ID: 1,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 20,
						},
					}, nil)
					return analyzerRepository
				},
			},
			args: args{
				queryParams: map[string]string{
					"page": "1",
					"size": "1",
				},
			},
			want: []models.DatabaseQueryInfo{
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 20,
				},
			},
		},
		{
			name: "should return successfully empty response when nothing was found",
			fields: fields{
				repo: func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository {
					analyzerRepository := repository.NewMockQueryAnalyzerRepository(ctrl)
					analyzerRepository.EXPECT().GetDatabaseQueryInfo(gomock.Any()).Return([]models.DatabaseQueryInfo{}, nil)
					return analyzerRepository
				},
			},
			args: args{
				queryParams: map[string]string{
					"page": "1",
					"size": "1",
				},
			},
			want: []models.DatabaseQueryInfo{},
		},
		{
			name: "should return error when analyzerRepository returns error",
			fields: fields{
				repo: func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository {
					analyzerRepository := repository.NewMockQueryAnalyzerRepository(ctrl)
					analyzerRepository.EXPECT().GetDatabaseQueryInfo(gomock.Any()).
						Return(nil, errors.New("error occurred"))
					return analyzerRepository
				},
			},
			args: args{
				queryParams: map[string]string{
					"page": "1",
					"size": "1",
				},
			},
			wantErr: errors.New("error occurred"),
		},
		{
			name: "should return successfully 5 queries successfully",
			fields: fields{
				repo: func(ctrl *gomock.Controller) repository.QueryAnalyzerRepository {
					analyzerRepository := repository.NewMockQueryAnalyzerRepository(ctrl)
					analyzerRepository.EXPECT().GetDatabaseQueryInfo(gomock.Any()).Return([]models.DatabaseQueryInfo{
						{
							Model: models.Model{
								ID: 1,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 20,
						},
						{
							Model: models.Model{
								ID: 2,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 30,
						},
						{
							Model: models.Model{
								ID: 3,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 50,
						},
						{
							Model: models.Model{
								ID: 4,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 100,
						},
						{
							Model: models.Model{
								ID: 5,
							},
							Query:           "SELECT * FROM TABLE",
							ExecutionTimeMs: 200,
						},
					}, nil)
					return analyzerRepository
				},
			},
			args: args{
				queryParams: map[string]string{
					"query_type": "select",
					"page":       "1",
					"size":       "5",
				},
			},
			want: []models.DatabaseQueryInfo{
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 20,
				},
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 30,
				},
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 50,
				},
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 100,
				},
				{
					Query:           "SELECT * FROM TABLE",
					ExecutionTimeMs: 200,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := NewQueryAnalyzer(tt.fields.repo(ctrl))
			q := make(url.Values)
			for key, value := range tt.args.queryParams {
				q.Add(key, value)
			}
			var ctx fasthttp.RequestCtx
			var req fasthttp.Request
			req.SetRequestURI("https://db.statistic/queries?" + q.Encode())
			ctx.Init(&req, nil, nil)

			err := h.GetQueries(fiber.New().AcquireCtx(&ctx))
			if tt.wantErr != nil {
				if err == nil {
					assert.Fail(t, "Expected error here, but got nil")
					return
				}
				assert.EqualError(t, tt.wantErr, err.Error())
				return
			}
			if err != nil {
				assert.Failf(t, "Expected no error, but got", err.Error())
				return
			}

			var result []models.DatabaseQueryInfo
			err = json.Unmarshal(ctx.Response.Body(), &result)
			if err != nil {
				assert.Fail(t, err.Error())
				return
			}

			assert.EqualValues(t, tt.want, result)
		})
	}
}
