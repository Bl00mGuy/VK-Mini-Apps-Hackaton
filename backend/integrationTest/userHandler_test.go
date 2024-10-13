package controllerTest

import (
	mocks "github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/integrationTest/mock"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/services/serviceImpl"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	userService := serviceImpl.NewUserService(mockRepo)

	createUserDTO := &dto.CreateUserDTO{
		Dep:      "IT",
		Lvl:      "Junior",
		Course:   3,
		Sport:    "Football",
		Club:     "Chess Club",
		Mer:      "mer",
		AvatarId: 1,
	}

	mockRepo.On("Create", createUserDTO).Return(nil)

	err := userService.CreateUser(createUserDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindUserByID(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	userService := serviceImpl.NewUserService(mockRepo)

	userID := uint(1)
	userDTO := &dto.UserDTO{
		ID:       userID,
		Dep:      "IT",
		Lvl:      "Senior",
		Course:   4,
		Sport:    "Basketball",
		Club:     "Drama Club",
		Mer:      "mer",
		AvatarID: 2,
	}

	mockRepo.On("FindByID", userID).Return(userDTO, nil)

	result, err := userService.FindUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, userDTO, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	userService := serviceImpl.NewUserService(mockRepo)

	userID := uint(1)

	mockRepo.On("Delete", userID).Return(nil)

	err := userService.DeleteUser(userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	userService := serviceImpl.NewUserService(mockRepo)

	updateUserDTO := &dto.UpdateUserDTO{
		ID:       1,
		Dep:      "HR",
		Lvl:      "Senior",
		Course:   4,
		Sport:    "Tennis",
		Club:     "Music Club",
		Mer:      "updated_mer",
		AvatarId: 2,
	}

	mockRepo.On("Update", updateUserDTO).Return(nil)

	err := userService.UpdateUser(updateUserDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
