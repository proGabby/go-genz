package entity

type User struct {
	Id                          int    `json:"id"`
	Name                        string `json:"name" validate:"required,min=3,max=20"`
	Email                       string `json:"email" validate:"required,email"`
	Password                    string `json:"password" validate:"required: min6, max60"`
	ProfileImageUrl             string `json:"profile_image_url"`
	IsVerified                  bool   `json:"is_verified"`
	TokenVersion                int    `json:"token_version"`
	EmailOtp                    string `json:"email_otp"`
	ForgotPasswordOtp           string `json:"forgot_password_otp"`
	ForgetPasswordOtpExpiryTime int    `json:"forget_password_otp_expiry_time"`
}

func NewUser(id int, name, email, profileImgUrl string, isVerified bool) *User {
	return &User{
		Id:                          id,
		Name:                        name,
		Email:                       email,
		ProfileImageUrl:             profileImgUrl,
		IsVerified:                  isVerified,
		TokenVersion:                0,
		EmailOtp:                    "",
		ForgotPasswordOtp:           "",
		ForgetPasswordOtpExpiryTime: 0,
	}
}
