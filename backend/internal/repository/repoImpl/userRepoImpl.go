package repoImpl

import (
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/dto"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/entity"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/mapper"
	"github.com/Bl00mGuy/VK-Mini-Apps-Hackaton/blob/main/backend/internal/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
	mp *mapper.UserMapper
}

func NewUserRepository(db *gorm.DB, mp *mapper.UserMapper) repository.UserRepository {
	return &userRepository{
		db: db,
		mp: mp,
	}
}

func (r *userRepository) Create(createUserDTO *dto.CreateUserDTO) error {
	user := &entity.User{
		Dep:      createUserDTO.Dep,
		Lvl:      createUserDTO.Lvl,
		Course:   createUserDTO.Course,
		Sport:    createUserDTO.Sport,
		Club:     createUserDTO.Club,
		MER:      createUserDTO.Mer,
		AvatarID: createUserDTO.AvatarId,
	}
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*dto.UserDTO, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToDTO(&user), nil
}

func (r *userRepository) FindAll() ([]dto.UserDTO, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return r.mp.ConvertToSliceDTO(users), nil
}

func (r *userRepository) Update(updateUserDTO *dto.UpdateUserDTO) error {
	var user entity.User
	if err := r.db.First(&user, updateUserDTO.ID).Error; err != nil {
		return err
	}

	user.Dep = updateUserDTO.Dep
	user.Lvl = updateUserDTO.Lvl
	user.Course = updateUserDTO.Course
	user.Sport = updateUserDTO.Sport
	user.Club = updateUserDTO.Club
	user.MER = updateUserDTO.Mer
	user.AvatarID = updateUserDTO.AvatarId

	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}
