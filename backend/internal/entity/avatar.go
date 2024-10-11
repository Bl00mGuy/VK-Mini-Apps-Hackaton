package entity

type Avatar struct {
	Name      string   `json:"name"`
	ImageURL  string   `json:"image_url"`
	Interests []string `json:"interests"`
}
