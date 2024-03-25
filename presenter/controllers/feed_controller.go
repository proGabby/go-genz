package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/proGabby/4genz/domain/entity"
	"github.com/proGabby/4genz/domain/usecase/feeds_usecase"
	"github.com/proGabby/4genz/utils"
)

type FeedsController struct {
	feedUsecases feeds_usecase.FeedUsecases
	feedChan     *chan *entity.Feed
}

func NewFeedsController(feedUsecases feeds_usecase.FeedUsecases, feedChan *chan *entity.Feed) *FeedsController {
	return &FeedsController{
		feedUsecases: feedUsecases,
		feedChan:     feedChan,
	}
}

func (f *FeedsController) CreateFeed(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*entity.User)

	if user == nil || !ok {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}

	type requestBody struct {
		Caption string `validate:"required,min=3,max=120"`
	}

	caption := r.FormValue("caption")

	c := requestBody{Caption: caption}

	err := validate.Struct(c)
	if err != nil {
		fmt.Printf("error validation: %v", err)
		// Handle validation error
		errors := err.(validator.ValidationErrors)
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("Validation failed: %v", errors),
		}, http.StatusBadRequest, w)

		return
	}

	er := r.ParseMultipartForm(30 << 20)
	if er != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("file parsing error: %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	files := r.MultipartForm.File["images"]

	if len(files) > 3 {
		utils.HandleError(map[string]interface{}{
			"error": "you can only attach 3 images max",
		}, http.StatusBadRequest, w)
		return
	}

	feed, err := f.feedUsecases.CreateFeed.Execute(user.Id, caption, files)

	if err != nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("feed creation failed: %v", err),
		}, http.StatusBadRequest, w)
		return
	}

	if feed == nil {
		utils.HandleError(map[string]interface{}{
			"error": fmt.Sprintf("feed creation failed"),
		}, http.StatusBadRequest, w)
	}

	if f.feedChan != nil {

		*f.feedChan <- feed
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*feed)

}

func (f *FeedsController) FetchFeeds(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(*entity.User)

	if user == nil || !ok {
		utils.HandleError(map[string]interface{}{
			"error": "user not authenticated",
		}, http.StatusBadRequest, w)
		return
	}
	var limit int = 10
	var page int = 1
	limt, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err == nil {
		limit = limt
	}

	pg, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err == nil {
		page = pg
	}

	feedRes, err := f.feedUsecases.FetchFeed.Execute(limit, page)

	if err != nil || feedRes == nil {
		fmt.Printf("error fetching feeds: %v \n", err)
		utils.HandleError(map[string]interface{}{
			"error": "something went wrong fetching feeds",
		}, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*feedRes)

}
