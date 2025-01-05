package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"poc-testcontainers/internal/adapters/controllers/user"
	"poc-testcontainers/internal/application/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateUseCase struct {
	mock.Mock
}

func (m *MockCreateUseCase) Create(user *dto.CreateUserReqDTO) (*dto.CreateUserResDTO, error) {
	args := m.Called(user)
	if user, ok := args.Get(0).(*dto.CreateUserResDTO); ok {
		return user, args.Error(1)
	}

	return nil, args.Error(1)
}

func TestCreateHandle(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedBody   string
		repositoryMock func(repo *MockCreateUseCase)
	}{
		{
			name:           "Missing request body",
			requestBody:    nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Request body is required"}`,
			repositoryMock: func(repo *MockCreateUseCase) {},
		},
		{
			name:           "Invalid JSON body",
			requestBody:    "{invalid}",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"invalid character 'i' looking for beginning of object key string"}`,
			repositoryMock: func(repo *MockCreateUseCase) {},
		},
		{
			name: "Valid request body",
			requestBody: dto.CreateUserReqDTO{
				Name: "John Doe",
				Age:  30,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"John Doe","age":30}`,
			repositoryMock: func(repo *MockCreateUseCase) {
				repo.On("Create", &dto.CreateUserReqDTO{Name: "John Doe", Age: 30}).
					Return(&dto.CreateUserResDTO{ID: 1, Name: "John Doe", Age: 30}, nil).
					Once()
			},
		},
		{
			name: "Repository returns error",
			requestBody: dto.CreateUserReqDTO{
				Name: "John Doe",
				Age:  30,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"Failed to create user"}`,
			repositoryMock: func(repo *MockCreateUseCase) {
				repo.On("Create", &dto.CreateUserReqDTO{Name: "John Doe", Age: 30}).
					Return(nil, errors.New("database error")).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := new(MockCreateUseCase)
			ctrl := user.NewCreateUserController(repo)

			tt.repositoryMock(repo)

			r := gin.Default()
			r.POST("/user", ctrl.Handle)

			var requestBody []byte
			if tt.requestBody != nil {
				if strBody, ok := tt.requestBody.(string); ok {
					requestBody = []byte(strBody)
				} else {
					requestBody, _ = json.Marshal(tt.requestBody)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())

			repo.AssertExpectations(t)
		})
	}
}
