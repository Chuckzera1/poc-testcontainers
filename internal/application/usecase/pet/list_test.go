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

type PetRepository struct {
	mock.Mock
}

func (m *PetRepository) List(pet *model.Pet, page int) ([]model.Pet, error) {
	args := m.Called(pet, page)
	return args.Get(0).([]model.Pet), args.Error(1)
}

func TestListPetUsecase_List(t *testing.T) {
	type args struct {
		name string
		page int
	}

	type mockBehavior struct {
		inputModel *model.Pet
		page       int
		result     []model.Pet
		err        error
	}

	tests := []struct {
		name          string
		args          args
		mockBehavior  mockBehavior
		expectedDtos  []*dto.PetListDTO
		expectedError error
	}{
		{
			name: "success - valid input with pets",
			args: args{
				name: "Buddy",
				page: 1,
			},
			mockBehavior: mockBehavior{
				inputModel: &model.Pet{Name: "Buddy"},
				page:       1,
				result: []model.Pet{
					{
						ID:   1,
						Name: "Buddy",
						Age:  5,
						UserResponsible: &model.User{
							ID:   10,
							Name: "Alice",
							Age:  30,
						},
					},
				},
				err: nil,
			},
			expectedDtos: []*dto.PetListDTO{
				{
					ID:   1,
					Name: "Buddy",
					Age:  5,
					UserResponsible: &dto.PetUserResponsibleDTO{
						ID:   10,
						Name: "Alice",
						Age:  30,
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "error - repository returns an error",
			args: args{
				name: "Max",
				page: 2,
			},
			mockBehavior: mockBehavior{
				inputModel: &model.Pet{Name: "Max"},
				page:       2,
				result:     nil,
				err:        errors.New("repository error"),
			},
			expectedDtos:  nil,
			expectedError: errors.New("repository error"),
		},
		{
			name: "success - no pets found",
			args: args{
				name: "Unknown",
				page: 1,
			},
			mockBehavior: mockBehavior{
				inputModel: &model.Pet{Name: "Unknown"},
				page:       1,
				result:     []model.Pet{},
				err:        nil,
			},
			expectedDtos:  []*dto.PetListDTO{},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(PetRepository)
			mockRepo.On("List", tt.mockBehavior.inputModel, tt.mockBehavior.page).Return(tt.mockBehavior.result, tt.mockBehavior.err)

			uc := pet.NewListPetUseCase(mockRepo)

			dtos, err := uc.List(tt.args.name, tt.args.page)

			assert.Equal(t, tt.expectedDtos, dtos)
			assert.Equal(t, tt.expectedError, err)

			mockRepo.AssertCalled(t, "List", tt.mockBehavior.inputModel, tt.mockBehavior.page)
		})
	}
}
