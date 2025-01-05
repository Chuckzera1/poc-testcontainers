package pet_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"poc-testcontainers/internal/adapters/controllers/pet"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateRepository struct {
	mock.Mock
}

func (m *MockCreateRepository) Create(pet *model.Pet) (*model.Pet, error) {
	args := m.Called(pet)
	if pet, ok := args.Get(0).(*model.Pet); ok {
		return pet, args.Error(1)
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
		repositoryMock func(repo *MockCreateRepository)
	}{
		{
			name:           "Missing request body",
			requestBody:    nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Request body is required"}`,
			repositoryMock: func(repo *MockCreateRepository) {},
		},
		{
			name:           "Invalid JSON body",
			requestBody:    "{invalid}",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"invalid character 'i' looking for beginning of object key string"}`,
			repositoryMock: func(repo *MockCreateRepository) {},
		},
		{
			name: "Missing UserResponsibleID body",
			requestBody: dto.CreatePetReqDTO{
				Name: "John Doe",
				Age:  30,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Key: 'CreatePetReqDTO.UserResponsibleID' Error:Field validation for 'UserResponsibleID' failed on the 'required' tag"}`,
			repositoryMock: func(repo *MockCreateRepository) {},
		},
		{
			name: "Valid request body",
			requestBody: dto.CreatePetReqDTO{
				Name:              "John Doe",
				Age:               30,
				UserResponsibleID: 1,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"John Doe","age":30, "userResponsibleId": 1}`,
			repositoryMock: func(repo *MockCreateRepository) {
				repo.On("Create", &model.Pet{Name: "John Doe", Age: 30, UserResponsibleID: 1}).
					Return(&model.Pet{ID: 1, Name: "John Doe", Age: 30, UserResponsibleID: 1}, nil).
					Once()
			},
		},
		{
			name: "Repository returns error",
			requestBody: dto.CreatePetReqDTO{
				Name:              "John Doe",
				Age:               30,
				UserResponsibleID: 1,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"Failed to create pet"}`,
			repositoryMock: func(repo *MockCreateRepository) {
				repo.On("Create", &model.Pet{Name: "John Doe", Age: 30, UserResponsibleID: 1}).
					Return(nil, errors.New("database error")).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := new(MockCreateRepository)
			ctrl := pet.NewCreatePetController(repo)

			tt.repositoryMock(repo)

			r := gin.Default()
			r.POST("/pet", ctrl.Handle)

			var requestBody []byte
			if tt.requestBody != nil {
				if strBody, ok := tt.requestBody.(string); ok {
					requestBody = []byte(strBody)
				} else {
					requestBody, _ = json.Marshal(tt.requestBody)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/pet", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			r.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())

			repo.AssertExpectations(t)
		})
	}
}
