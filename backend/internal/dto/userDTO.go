package dto

type UserDTO struct {
	ID       uint   `json:"id" binding:"required"`
	Dep      string `json:"dep" binding:"required"`
	Lvl      string `json:"lvl" binding:"required"`
	Course   uint   `json:"course" binding:"required"`
	Sport    string `json:"sport"`
	Club     string `json:"club"`
	Mer      string `json:"mer"`
	AvatarID uint   `json:"avatar_id" binding:"required"`
}

type CreateUserDTO struct {
	Dep      string `json:"dep" binding:"required"`
	Lvl      string `json:"lvl" binding:"required"`
	Course   uint   `json:"course" binding:"required"`
	Sport    string `json:"sport"`
	Club     string `json:"club"`
	Mer      string `json:"mer"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
}

type FindUserDTO struct {
	ID uint `json:"user_id" binding:"required"`
}

type UpdateUserDTO struct {
	ID       uint   `json:"id" binding:"required"`
	Dep      string `json:"dep" binding:"required"`
	Lvl      string `json:"lvl" binding:"required"`
	Course   uint   `json:"course" binding:"required"`
	Sport    string `json:"sport"`
	Club     string `json:"club"`
	Mer      string `json:"mer"`
	AvatarId uint   `json:"avatar_id" binding:"required"`
}

type DeleteUserDTO struct {
	UserID uint `json:"user_id" binding:"required"`
}
