package routes

import (
	"net/http"
	"time"

	"github.com/ariiiiph/Url-Shortner/api/database"
	"github.com/ariiiiph/Url-Shortner/api/models"
	"github.com/gin-gonic/gin"
)

func EditURL(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot Parse JSON",
		})
		return
	}

	//check if the shortID exists in the DB or not

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ShortID doesn't exists",
		})
		return
	}

	//Update the content of the URL, expiry time with the shortID
	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to update the shortend content",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "The Content Has been Updated!",
	})

}
