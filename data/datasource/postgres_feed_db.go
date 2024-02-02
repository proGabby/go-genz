package postgressDatasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/proGabby/4genz/data/dto"
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

func (feedDb *PostgresFeedDBStore) FetchFeedWithPagination(limit, offset int) (*[]dto.FeedWithUserData, *int, error) {

	feeds := []dto.FeedWithUserData{}
	feedMap := make(map[int]*dto.FeedWithUserData)

	var totalUniqueFeedId int
	err := feedDb.DB.QueryRow("SELECT COUNT(DISTINCT Id) FROM Feeds").Scan(&totalUniqueFeedId)
	if err != nil {
		return nil, nil, err
	}

	query := fmt.Sprintf(`
						SELECT
							f.Id AS FeedId,
							f.Caption,
							f.UserId,
							f.Likes,
							f.CreatedAt AS FeedCreatedAt,
							f.UpdatedAt AS FeedUpdatedAt,
							fi.ImageUrl,
							u.Name AS UserName,
							u.Email AS UserEmail,
							u.profile_image_url AS UserProfileImageUrl
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
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var feedId int
		var caption string
		var userId int
		var likes int
		var feedCreatedAt string
		var feedUpdatedAt string
		var imageUrl sql.NullString
		var userName sql.NullString
		var userEmail sql.NullString
		var userProfileImageUrl sql.NullString
		err := rows.Scan(&feedId, &caption, &userId, &likes, &feedCreatedAt, &feedUpdatedAt, &imageUrl, &userName, &userEmail, &userProfileImageUrl)
		if err != nil {
			return nil, nil, err
		}

		if feed, ok := feedMap[feedId]; ok {
			if imageUrl.Valid {
				feed.Images = append(feed.Images, imageUrl.String)
			}
		} else {
			feed := &entity.Feed{
				Id:        feedId,
				Caption:   caption,
				UserId:    userId,
				Likes:     likes,
				CreatedAt: feedCreatedAt,
				UpdatedAt: feedUpdatedAt,
				Images:    []string{},
			}

			if imageUrl.Valid {
				feed.Images = append(feed.Images, imageUrl.String)
			}

			feedUserData := &dto.FeedWithUserData{
				Feed:         *feed,
				UserName:     userName.String,
				UserPhone:    "",
				UserEmail:    userEmail.String,
				UserImageUrl: userProfileImageUrl.String,
			}

			feedMap[feedId] = feedUserData

		}
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("error on feed rows")
	}

	for _, element := range feedMap {

		feeds = append(feeds, *element)
	}

	return &feeds, &totalUniqueFeedId, nil

}
