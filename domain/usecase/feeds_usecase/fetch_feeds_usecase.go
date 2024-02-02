package feeds_usecase

import (
	"fmt"

	"github.com/proGabby/4genz/domain/repository/feed_repo"
)

type FetchFeedsUsecase struct {
	feedRepo feed_repo.FeedRepository
}

func NewFetchFeedsUsecase(feedRepo feed_repo.FeedRepository) *FetchFeedsUsecase {

	return &FetchFeedsUsecase{
		feedRepo: feedRepo,
	}
}

func (ff *FetchFeedsUsecase) Execute(limit, page int) (*map[string]interface{}, error) {
	skip := (page - 1) * limit
	feeds, docCount, err := ff.feedRepo.FetchPaginatedFeeds(limit, skip)

	if err != nil {
		return nil, err
	}

	if feeds == nil || docCount == nil {
		return nil, fmt.Errorf("feeds fetching error")
	}

	var nextPage *int

	if skip+limit <= *docCount {
		nextPageValue := page + 1
		nextPage = &nextPageValue
	} else {
		nextPage = nil
	}

	var nextPageValue interface{}

	if nextPage != nil {
		nextPageValue = *nextPage
	} else {
		nextPageValue = nil
	}

	resMap := map[string]interface{}{
		"data":     feeds,
		"count":    docCount,
		"nextPage": nextPageValue,
	}

	return &resMap, nil

}
