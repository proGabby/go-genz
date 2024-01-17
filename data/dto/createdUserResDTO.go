package dto

type CreateUserResponse struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"profile_image_url"`
}
