package pet_test

import (
	"errors"
	"poc-testcontainers/internal/application/dto"
	"poc-testcontainers/internal/application/usecase/pet"
	"poc-testcontainers/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPetRepo struct {
	mock.Mock
}

func (m *MockPetRepo) Create(pet *model.Pet) (*model.Pet, error) {
	args := m.Called(pet)
	return args.Get(0).(*model.Pet), args.Error(1)
}

func TestCreatePetUsecase_Create(t *testing.T) {
	mockRepo := new(MockPetRepo)
	usecase := pet.NewCreatePetUseCase(mockRepo)

	tests := []struct {
		name          string
		input         *dto.CreatePetReqDTO
		mockReturn    *model.Pet
		mockError     error
		expectedRes   *dto.CreatePetResDTO
		expectedError string
	}{
		{
			name: "should successfully create a pet",
			input: &dto.CreatePetReqDTO{
				Name:              "Buddy",
				Age:               3,
				UserResponsibleID: 1,
			},
			mockReturn: &model.Pet{
				ID:                1,
				Name:              "Buddy",
				Age:               3,
				UserResponsibleID: 1,
			},
			mockError: nil,
			expectedRes: &dto.CreatePetResDTO{
				ID:                1,
				Name:              "Buddy",
				Age:               3,
				UserResponsibleID: 1,
			},
			expectedError: "",
		},
		{
			name: "should return an error when repo.Create fails",
			input: &dto.CreatePetReqDTO{
				Name:              "Buddy",
				Age:               3,
				UserResponsibleID: 1,
			},
			mockReturn:    nil,
			mockError:     errors.New("database error"),
			expectedRes:   nil,
			expectedError: "database error",
		},
		{
			name:          "should handle nil input gracefully",
			input:         nil,
			mockReturn:    nil,
			mockError:     nil,
			expectedRes:   nil,
			expectedError: "invalid input: pet cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.ExpectedCalls = nil
			mockRepo.Calls = nil

			if tt.input != nil {
				mockRepo.On("Create", mock.MatchedBy(func(p *model.Pet) bool {
					return p.Name == tt.input.Name && p.Age == tt.input.Age && p.UserResponsibleID == tt.input.UserResponsibleID
				})).Return(tt.mockReturn, tt.mockError).Once()
			}

			result, err := usecase.Create(tt.input)

			if tt.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedError)
			}
			assert.Equal(t, tt.expectedRes, result)

			if tt.input != nil {
				mockRepo.AssertNumberOfCalls(t, "Create", 1)
			} else {
				mockRepo.AssertNotCalled(t, "Create")
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
