package entity

import "time"

// A feed struct to hold the feed data it has a one to many relationship with the user struct and also list of images which should be 3 max
type Feed struct {
	Id        int      `json:"id"`
	Caption   string   `json:"caption" validate:"required,min=3,max=120"`
	UserId    int      `json:"user_id"`
	Images    []string `json:"images"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Likes     int      `json:"likes" `
}

func NewFeed(id, userId, likes int, caption string, images []string) *Feed {

	return &Feed{
		Id:        id,
		UserId:    userId,
		Likes:     likes,
		Caption:   caption,
		Images:    images,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}
