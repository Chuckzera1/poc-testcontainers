package user_test

import (
	"net/http"
	"net/http/httptest"
	"poc-testcontainers/internal/adapters/controllers/user"
	"poc-testcontainers/internal/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockListRepository struct {
	mock.Mock
}

func (m *MockListRepository) List(user *model.User, page int) ([]model.User, error) {
	args := m.Called(user, page)
	return args.Get(0).([]model.User), args.Error(1)
}

func TestHandleWithQueryName(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryName      string
		queryPage      string
		mockResult     []model.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid request",
			queryName:      "John",
			queryPage:      "1",
			mockResult:     []model.User{{ID: 1, Name: "John", Age: 30}},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"data":[{ "id":1, "name":"John", "age":30}]}`,
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
			expectedBody:   `{"message":"Failed to list users"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(MockListRepository)
			controller := user.NewListUserController(mockRepo)

			if tt.mockResult != nil || tt.mockError != nil {
				mockRepo.On("List", &model.User{Name: tt.queryName}, mock.AnythingOfType("int")).
					Return(tt.mockResult, tt.mockError)
			}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/users", controller.Handle)

			req := httptest.NewRequest(http.MethodGet, "/users?name="+tt.queryName+"&page="+tt.queryPage, nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
