package pet_test

import (
	"net/http"
	"net/http/httptest"
	"poc-testcontainers/internal/adapters/controllers/pet"
	"poc-testcontainers/internal/application/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockListUsecase struct {
	mock.Mock
}

func (m *MockListUsecase) List(name string, page int) ([]*dto.PetListDTO, error) {
	args := m.Called(name, page)
	return args.Get(0).([]*dto.PetListDTO), args.Error(1)
}

func TestHandleWithQueryName(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryName      string
		queryPage      string
		mockResult     []*dto.PetListDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "valid request",
			queryName: "John",
			queryPage: "1",
			mockResult: []*dto.PetListDTO{
				{
					ID:   1,
					Name: "John",
					Age:  30,
					UserResponsible: &dto.PetUserResponsibleDTO{
						ID:   1,
						Name: "James",
						Age:  30,
					},
				}},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":[{ "id":1, "name":"John", "age":30, "userResponsible": {"id": 1, "name": "James", "age": 30}}]}`,
		},
		{
			name:           "invalid page query",
			queryName:      "John",
			queryPage:      "invalid",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"page query params is wrong"}`,
		},
		{
			name:           "repository error",
			queryName:      "John",
			queryPage:      "1",
			mockResult:     nil,
			mockError:      assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"Failed to list pets"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockList := new(MockListUsecase)
			controller := pet.NewListPetController(mockList)

			if tt.mockResult != nil || tt.mockError != nil {
				mockList.On("List", tt.queryName, mock.AnythingOfType("int")).
					Return(tt.mockResult, tt.mockError)
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/pets", controller.Handle)

			req := httptest.NewRequest(http.MethodGet, "/pets?name="+tt.queryName+"&page="+tt.queryPage, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
