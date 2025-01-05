package user_test

import (
	"errors"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/application/usecase/user"
	"poc-testcontainers/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCreateUserRepository struct {
	CreateFn func(*model.User) (*model.User, error)
}

func (m *mockCreateUserRepository) Create(user *model.User) (*model.User, error) {
	return m.CreateFn(user)
}

func TestCreateUserUseCase_Create(t *testing.T) {
	tests := []struct {
		name          string
		input         *dto.CreateUserReqDTO
		mockResponse  *model.User
		mockError     error
		expected      *dto.CreateUserResDTO
		expectedError error
	}{
		{
			name: "Successful creation",
			input: &dto.CreateUserReqDTO{
				Name: "John Doe",
				Age:  30,
			},
			mockResponse: &model.User{
				ID:   1,
				Name: "John Doe",
				Age:  30,
			},
			mockError: nil,
			expected: &dto.CreateUserResDTO{
				ID:   1,
				Name: "John Doe",
				Age:  30,
			},
			expectedError: nil,
		},
		{
			name: "Repository error",
			input: &dto.CreateUserReqDTO{
				Name: "Jane Doe",
				Age:  25,
			},
			mockResponse:  nil,
			mockError:     errors.New("repository error"),
			expected:      nil,
			expectedError: errors.New("repository error"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := &mockCreateUserRepository{
				CreateFn: func(user *model.User) (*model.User, error) {
					assert.Equal(t, tt.input.Name, user.Name)
					assert.Equal(t, tt.input.Age, user.Age)
					return tt.mockResponse, tt.mockError
				},
			}

			useCase := user.NewCreateUserUseCase(mockRepo)

			result, err := useCase.Create(tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
