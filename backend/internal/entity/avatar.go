package entity

type Avatar struct {
	ID        uint     `json:"id"`
	UserID    uint     `json:"User_id"`
	Name      string   `json:"name"`
	ImageURL  string   `json:"image_url"`
	Interests []string `json:"interests"`
}
