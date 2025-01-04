package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"poc-testcontainers/internal/adapters/controllers/user"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Delete(userID uint64) error {
	args := m.Called(userID)
	return args.Error(0)
}

func TestDeleteUserController_Handle(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name               string
		userID             string
		mockDeleteError    error
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "successful delete",
			userID:             "1",
			mockDeleteError:    nil,
			expectedStatusCode: http.StatusNoContent,
			expectedBody:       "",
		},
		{
			name:               "missing userID param",
			userID:             "",
			mockDeleteError:    nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"message":"route param userID is missing"}`,
		},
		{
			name:               "non-numeric userID param",
			userID:             "abc",
			mockDeleteError:    nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       `{"message":"route param userID is not a number"}`,
		},
		{
			name:               "repository delete error",
			userID:             "1",
			mockDeleteError:    errors.New("repository error"),
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       `{"message":"repository error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			if tt.userID != "" && tt.userID != "abc" {
				userIDInt, _ := strconv.Atoi(tt.userID)
				mockRepo.On("Delete", uint64(userIDInt)).Return(tt.mockDeleteError)
			}

			controller := user.NewDeleteUserController(mockRepo)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{{Key: "id", Value: tt.userID}}

			controller.Handle(c)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, strings.TrimSpace(w.Body.String()))
			} else {
				assert.Empty(t, w.Body.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
