package feed_repo_impl

import (
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/domain/entity"
)

type FeedRepositoryImpl struct {
	psql postgressDatasource.PostgresFeedDBStore
}

func NewFeedRepoImpl(psql postgressDatasource.PostgresFeedDBStore) *FeedRepositoryImpl {
	return &FeedRepositoryImpl{
		psql: psql,
	}
}

func (feedRepoImpl *FeedRepositoryImpl) CreateNewFeed(userId int, caption string) (*entity.Feed, error) {
	return feedRepoImpl.psql.CreateNewFeed(userId, caption)
}

func (feedRepoImpl *FeedRepositoryImpl) AddFeedImage(feed *entity.Feed, imageUrl string) (*entity.Feed, error) {
	return feedRepoImpl.psql.AddFeedImage(feed, imageUrl)
}
