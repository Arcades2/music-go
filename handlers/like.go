package handlers

import (
	"main/apierror"
	"main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLikeHandler(c *gin.Context) {
	var payload CreateLikePayload
	err := c.Bind(&payload)
	if err != nil {
		apierror.SendErrorResp(http.StatusUnprocessableEntity, c, err)
		return
	}

	createdLikeId, err := repository.CreateLike(repository.LikeData(payload))
	if err != nil {
		apierror.SendErrorResp(http.StatusInternalServerError, c, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"like_id": createdLikeId,
	})
}

type CreateLikePayload struct {
	PostId int `json:"postId" binding:"required"`
}
