package routes

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ariiiiph/Url-Shortner/api/database"
	"github.com/ariiiiph/Url-Shortner/api/models"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot Parse JSON"})
		return
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	val, err := r2.Get(database.Ctx, c.ClientIP()).Result()

	if err != nil {
		_ = r2.Set(database.Ctx, c.ClientIP(), os.Getenv("API_QUOTA"), 30*60*time.Second).Err()
	} else {
		val, _ = r2.Get(database.Ctx, c.ClientIP()).Result()
		valInt, _ := strconv.Atoi(val)

		if valInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.ClientIP()).Result()
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error":            "rate limit exceeded",
				"rate_limit_reset": limit / time.Nanosecond / time.Minute,
			})
			return
		}
	}
	if !govalidator.IsURL(body.URL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	if !utils.IsDiffrentDomain(body.URL) {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "You Cannot Hack this system",
		})
	}
}
