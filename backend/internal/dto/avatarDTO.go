package dto

type AvatarDTO struct {
	AvatarID  uint     `json:"id"`
	UserID    uint     `json:"user_Id"`
	Name      string   `json:"name" binding:"required"`
	ImageURL  string   `json:"image_url" binding:"required"`
	Interests []string `json:"interests" binding:"required"`
}

type CreateAvatarDTO struct {
	Name      string   `json:"name" binding:"required"`
	UserID    uint     `json:"user_id" binding:"required"`
	ImageURL  string   `json:"image_url" binding:"required"`
	Interests []string `json:"interests" binding:"required"`
}

type FindAvatarDTO struct {
	AvatarID uint `json:"avatar_id" binding:"required"`
}

type UpdateAvatarDTO struct {
	Name      string   `json:"name" binding:"required"`
	AvatarID  uint     `json:"avatar_id" binding:"required"`
	UserID    uint     `json:"user_id" binding:"required"`
	ImageURL  string   `json:"image_url" binding:"required"`
	Interests []string `json:"interests" binding:"required"`
}

type DeleteAvatarDTO struct {
	AvatarID uint `json:"avatar_id" binding:"required"`
}
