package handlers

import (
	"fmt"
	"main/apierror"
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Get Posts
func GetPostsHandler(c *gin.Context) {
	var query GetPostsQuery
	err := c.ShouldBind(&query)
	if err != nil {
		apierror.SendErrorResp(http.StatusUnprocessableEntity, c, err)
		return
	}

	fmt.Println(query)

	posts, err := repository.GetPosts(repository.GetPostsParams{
		Take: query.Take,
		Skip: query.Skip,
	})
	if err != nil {
		apierror.SendErrorResp(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, posts)
}

type GetPostsQuery struct {
	Take int `form:"take"`
	Skip int `form:"skip"`
}

// Create Post
func CreatePostHandler(c *gin.Context) {
	var payload CreatePostPayload
	err := c.Bind(&payload)
	if err != nil {
		apierror.SendErrorResp(http.StatusUnprocessableEntity, c, err)
		return
	}

	validationErrors := payload.Validate()
	if validationErrors != nil {
		apierror.SendErrorResp(http.StatusUnprocessableEntity, c, validationErrors...)
		return
	}

	createPostId, err := repository.CreatePost(repository.PostData(payload))
	if err != nil {
		apierror.SendErrorResp(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post", "id": createPostId})
}

type CreatePostPayload struct {
	Description string `json:"description"`
	Url         string `json:"url" binding:"required,url"`
}

func (p CreatePostPayload) Validate() []error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		var errorMessages []error

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			switch err.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Errorf("%s is required", field))
			case "url":
				errorMessages = append(errorMessages, fmt.Errorf("%s is not a valid URL", field))
			default:
				errorMessages = append(errorMessages, fmt.Errorf("%s is not valid", field))
			}
		}

		return errorMessages
	}

	return nil
}
