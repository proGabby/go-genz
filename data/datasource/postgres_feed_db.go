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

func (feedDb *PostgresFeedDBStore) CreateNewFeed(userId int, caption string) (*entity.Feed, error) {

	var feed entity.Feed

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

	query := "INSERT INTO FeedImages(feedid, imageurl) VALUES($1, $2)"
	_, err := feedDb.DB.Exec(query, feed.Id, image)
	if err != nil {
		fmt.Printf("error inserting feed image: %v\n", err)
		return nil, err
	}

	feed.Images = append(feed.Images, image)

	return feed, nil

}

func (feedDb *PostgresFeedDBStore) FetchFeedWithPagination(limit, page int) (*[]entity.Feed, error) {

	//calculate limit and offset
	offset := limit * page

	feeds := []entity.Feed{}

	query := fmt.Sprintf(`
						SELECT
							f.Id AS FeedId,
							f.Caption,
							f.UserId,
							f.Likes,
							f.CreatedAt AS FeedCreatedAt,
							f.UpdatedAt AS FeedUpdatedAt,
							fi.Id AS ImageId,
							fi.ImageUrl,
							u.Id AS UserId,
							u.Name AS UserName,
							u.Email AS UserEmail,
							u.ProfileImageUrl AS UserProfileImageUrl
						FROM
							Feeds f
						LEFT JOIN
							FeedImages fi ON f.Id = fi.FeedId
						LEFT JOIN
							Users u ON f.UserId = u.Id
						ORDER BY
							f.CreatedAt DESC
						LIMIT %d OFFSET %d;
						`, limit, offset,
	)
	rows, err := feedDb.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		feed := entity.Feed{}
		err := rows.Scan(&feed.Id, &feed.Caption, &feed.UserId, &feed.Likes, &feed.CreatedAt, &feed.UpdatedAt)
		if err != nil {
			return nil, err
		}

		feeds = append(feeds, feed)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("error on feed rows")
	}

	return &feeds, nil

}
