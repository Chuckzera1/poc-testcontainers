package user_test

import (
	"errors"
	"poc-testcontainers/internal/application/usecase/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	DeleteFn func(id uint64) error
}

func (m *MockRepository) Delete(id uint64) error {
	return m.DeleteFn(id)
}

func TestDeleteUserUseCase(t *testing.T) {
	tests := []struct {
		name          string
		inputID       uint64
		mockError     error
		expectedError error
	}{
		{
			name:          "Successful deletion",
			inputID:       123,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Repository returns error",
			inputID:       456,
			mockError:     errors.New("user not found"),
			expectedError: errors.New("user not found"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := &MockRepository{
				DeleteFn: func(id uint64) error {
					assert.Equal(t, tt.inputID, id)
					return tt.mockError
				},
			}

			useCase := user.NewDeleteUserUseCase(mockRepo)

			err := useCase.Delete(tt.inputID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
