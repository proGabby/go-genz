package dto

import "github.com/proGabby/4genz/domain/entity"

type FeedWithUserData struct {
	entity.Feed
	UserName     string
	UserImageUrl string
	UserEmail    string
	UserPhone    string
}