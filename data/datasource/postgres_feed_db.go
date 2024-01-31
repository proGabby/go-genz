package postgressDatasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/proGabby/4genz/domain/entity"
)

type PostgresFeedDBStore struct {
	DB *sql.DB
}

func NewPostgresFeedDBStore(db *sql.DB) *PostgresFeedDBStore {
	return &PostgresFeedDBStore{
		DB: db,
	}
}

func (feedDb *PostgresFeedDBStore) CreateNewFeed(postId int, userId int, caption string) (*entity.Feed, error) {

	var feed *entity.Feed

	query := "INSERT INTO Feeds(Caption, UserId) VALUES($1, $2) RETURNING id, likes, createdAt, updatedAt"
	err := feedDb.DB.QueryRow(query, caption, userId).Scan(&feed.Id, &feed.Likes, &feed.CreatedAt, &feed.UpdatedAt)
	if err != nil {
		return nil, err
	}

	newFeed := entity.NewFeed(feed.Id, userId, feed.Likes, caption, []string{})
	return newFeed, nil
}

func (feedDb *PostgresFeedDBStore) AddFeedImage(feed *entity.Feed, image string) (*entity.Feed, error) {

	if feed == nil {
		return nil, fmt.Errorf("feed is nil")
	}

	query := "INSERT INTO FeedImages(FeedId, ImageUrl) VALUES($1, $2)"
	err := feedDb.DB.QueryRow(query, feed.Id, image).Scan(&feed.Id, &feed.UserId, &feed.Caption, &feed.Likes, &feed.CreatedAt, &feed.UpdatedAt)
	if err != nil {
		fmt.Printf("error inserting feed image: %v\n", err)
		return nil, err
	}

	feed.Images = append(feed.Images, image)

	return feed, nil

}
