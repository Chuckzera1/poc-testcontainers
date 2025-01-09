package user_test

import (
	"errors"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/application/usecase/user"
	"poc-testcontainers/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockListUserRepository struct {
	mock.Mock
}

func (m *mockListUserRepository) List(filter *model.User, page int) ([]model.User, error) {
	args := m.Called(filter, page)
	return args.Get(0).([]model.User), args.Error(1)
}

func TestListUserUseCase(t *testing.T) {
	tests := []struct {
		name           string
		inputName      string
		inputPage      int
		mockUsers      []model.User
		mockError      error
		expectedResult []dto.UserListResDTO
		expectedError  error
	}{
		{
			name:      "valid user list",
			inputName: "John",
			inputPage: 1,
			mockUsers: []model.User{
				{ID: 1, Name: "John", Age: 30},
				{ID: 2, Name: "Jane", Age: 25},
			},
			expectedResult: []dto.UserListResDTO{
				{ID: 1, Name: "John", Age: 30},
				{ID: 2, Name: "Jane", Age: 25},
			},
			expectedError: nil,
		},
		{
			name:           "repository returns error",
			inputName:      "Doe",
			inputPage:      0,
			mockUsers:      nil,
			mockError:      errors.New("repository error"),
			expectedResult: nil,
			expectedError:  errors.New("repository error"),
		},
		{
			name:           "empty user list",
			inputName:      "Nonexistent",
			inputPage:      1,
			mockUsers:      []model.User{},
			expectedResult: []dto.UserListResDTO{},
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(mockListUserRepository)
			useCase := user.NewListUserUseCase(mockRepo)

			mockRepo.On("List", &model.User{Name: tt.inputName}, tt.inputPage).
				Return(tt.mockUsers, tt.mockError)

			result, err := useCase.List(tt.inputName, tt.inputPage)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
