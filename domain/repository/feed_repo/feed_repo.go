package feed_repo

import (
	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/entity"
)

type FeedRepository interface {
	CreateNewFeed(userId int, caption string) (*entity.Feed, error)
	AddFeedImage(feed *entity.Feed, image string) (*entity.Feed, error)
	FetchPaginatedFeeds(limit, page int) (*[]dto.FeedWithUserData, *int, error)
}
