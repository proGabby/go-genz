package feeds_usecase

type FeedUsecases struct {
	CreateFeed CreateFeedUsecase
	FetchFeed  FetchFeedsUsecase
}

func NewFeedUsecases(createFeed CreateFeedUsecase, fetchFeed FetchFeedsUsecase) *FeedUsecases {

	return &FeedUsecases{
		CreateFeed: createFeed,
		FetchFeed:  fetchFeed,
	}
}
