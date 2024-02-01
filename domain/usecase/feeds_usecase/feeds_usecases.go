package feeds_usecase

type FeedUsecases struct {
	CreateFeed CreateFeedUsecase
}

func NewFeedUsecases(createFeed CreateFeedUsecase) *FeedUsecases {

	return &FeedUsecases{
		CreateFeed: createFeed,
	}
}
