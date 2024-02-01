package feed_repo

import "github.com/proGabby/4genz/domain/entity"

type FeedRepository interface {
	CreateNewFeed(userId int, caption string) (*entity.Feed, error)
	AddFeedImage(feed *entity.Feed, image string) (*entity.Feed, error)
}
