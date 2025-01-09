package pet_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"poc-testcontainers/internal/adapters/controllers/pet"
	"poc-testcontainers/internal/application/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCreateUseCase struct {
	mock.Mock
}

func (m *MockCreateUseCase) Create(pet *dto.CreatePetReqDTO) (*dto.CreatePetResDTO, error) {
	args := m.Called(pet)
	if pet, ok := args.Get(0).(*dto.CreatePetResDTO); ok {
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
		useCaseMock    func(usecase *MockCreateUseCase)
	}{
		{
			name:           "Missing request body",
			requestBody:    nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Request body is required"}`,
			useCaseMock:    func(usecase *MockCreateUseCase) {},
		},
		{
			name:           "Invalid JSON body",
			requestBody:    "{invalid}",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"invalid character 'i' looking for beginning of object key string"}`,
			useCaseMock:    func(usecase *MockCreateUseCase) {},
		},
		{
			name: "Missing UserResponsibleID body",
			requestBody: dto.CreatePetReqDTO{
				Name: "John Doe",
				Age:  30,
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"message":"Key: 'CreatePetReqDTO.UserResponsibleID' Error:Field validation for 'UserResponsibleID' failed on the 'required' tag"}`,
			useCaseMock:    func(usecase *MockCreateUseCase) {},
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
			useCaseMock: func(usecase *MockCreateUseCase) {
				usecase.On("Create", &dto.CreatePetReqDTO{Name: "John Doe", Age: 30, UserResponsibleID: 1}).
					Return(&dto.CreatePetResDTO{ID: 1, Name: "John Doe", Age: 30, UserResponsibleID: 1}, nil).
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
			useCaseMock: func(usecase *MockCreateUseCase) {
				usecase.On("Create", &dto.CreatePetReqDTO{Name: "John Doe", Age: 30, UserResponsibleID: 1}).
					Return(nil, errors.New("database error")).
					Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usecase := new(MockCreateUseCase)
			ctrl := pet.NewCreatePetController(usecase)

			tt.useCaseMock(usecase)

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

			usecase.AssertExpectations(t)
		})
	}
}
