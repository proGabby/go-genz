package feed_repo_impl

import (
	"github.com/googollee/go-socket.io"
	
	postgressDatasource "github.com/proGabby/4genz/data/datasource"
	"github.com/proGabby/4genz/data/dto"
	"github.com/proGabby/4genz/domain/entity"
)

type FeedRepositoryImpl struct {
	psql       postgressDatasource.PostgresFeedDBStore
	socketServ *socketio.Server
}

func NewFeedRepoImpl(psql postgressDatasource.PostgresFeedDBStore, socketServ *socketio.Server) *FeedRepositoryImpl {
	return &FeedRepositoryImpl{
		psql:       psql,
		socketServ: socketServ,
	}
}

func (feedRepoImpl *FeedRepositoryImpl) CreateNewFeed(userId int, caption string) (*entity.Feed, error) {
	return feedRepoImpl.psql.CreateNewFeed(userId, caption)
}

func (feedRepoImpl *FeedRepositoryImpl) AddFeedImage(feed *entity.Feed, imageUrl string) (*entity.Feed, error) {
	return feedRepoImpl.psql.AddFeedImage(feed, imageUrl)
}

func (feedRepoImpl *FeedRepositoryImpl) FetchPaginatedFeeds(limit, page int) (*[]dto.FeedWithUserData, *int, error) {
	return feedRepoImpl.psql.FetchFeedWithPagination(limit, page)
}

func (feedRepoImpl *FeedRepositoryImpl) SendFeedToSocket(namespace string, event string, args ...interface{}) {
	feedRepoImpl.socketServ.BroadcastToNamespace(namespace, event, args...)
}
