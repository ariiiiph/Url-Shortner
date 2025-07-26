package routes

import (
	"net/http"

	"github.com/ariiiiph/Url-Shortner/api/database"
	"github.com/gin-gonic/gin"
)

func DeleteURL(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	err := r.Del(database.Ctx, shortID).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete shortened Link",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Shortened URL Deleted Successfully",
	})
}
