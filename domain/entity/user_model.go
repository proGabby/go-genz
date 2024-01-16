package entity


type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ProfileImageUrl string `json:"profile_image_url"`
}




