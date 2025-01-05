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

type createUserRepository struct {
	mock.Mock
}

func (m *createUserRepository) Create(user *model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestCreateUserUseCase_Create(t *testing.T) {
	t.Run("successfully creates a user", func(t *testing.T) {
		t.Parallel()

		mockRepo := &createUserRepository{}
		createUserUseCase := user.NewCreateUserUseCase(mockRepo)

		reqDTO := &dto.CreateUserReqDTO{
			Name: "John Doe",
			Age:  30,
		}

		expectedUser := &model.User{
			ID:   1,
			Name: "John Doe",
			Age:  30,
		}

		mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(expectedUser, nil)

		res, err := createUserUseCase.Create(reqDTO)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, expectedUser.ID, res.ID)
		assert.Equal(t, expectedUser.Name, res.Name)
		assert.Equal(t, expectedUser.Age, res.Age)

		mockRepo.AssertCalled(t, "Create", &model.User{
			Name: "John Doe",
			Age:  30,
		})
	})

	t.Run("fails to create user when repository returns error", func(t *testing.T) {
		t.Parallel()

		mockRepo := &createUserRepository{}
		createUserUseCase := user.NewCreateUserUseCase(mockRepo)

		reqDTO := &dto.CreateUserReqDTO{
			Name: "Jane Doe",
			Age:  25,
		}

		mockError := errors.New("repository error")
		mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return((*model.User)(nil), mockError)

		res, err := createUserUseCase.Create(reqDTO)

		assert.Nil(t, res)
		assert.EqualError(t, err, "repository error")

		mockRepo.AssertCalled(t, "Create", &model.User{
			Name: "Jane Doe",
			Age:  25,
		})
	})
}
