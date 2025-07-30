package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ariiiiph/Url-Shortner/api/database"
	"github.com/ariiiiph/Url-Shortner/api/models"
	"github.com/gin-gonic/gin"
)

func AddTag(c *gin.Context) {
	var tagRequest models.TagRequest
	if err := c.ShouldBindJSON(&tagRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	shortID := tagRequest.ShortID
	tag := tagRequest.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found for the given shortID",
		})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		// if the data is not a JSON object, assume it as plain string
		data = make(map[string]interface{})
		data["data"] = val
	}

	//check if "tags" field already exists and it's a slice of strings
	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	//check for duplicate tags
	for _, existingTags := range tags {
		if existingTags == tag {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Tag already Exists",
			})
			return
		}
	}

	//Add the new tag to the tags slice
	tags = append(tags, tag)
	data["tags"] = tags

	//Marshal the updated daata back to JSON
	updatedData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Marshal updated data",
		})
		return
	}

	err = r.Set(database.Ctx, shortID, updatedData, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Update the Database",
		})
		return
	}
	//Response with the updated data
	c.JSON(http.StatusOK, data)

}
