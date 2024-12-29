package Routes

import (
	"encoding/json"
	"net/http"
	"todo/model"
	"todo/repository"

	"github.com/gin-gonic/gin"
)

var err error

// create content
func CreateContent(c *gin.Context) {
	var TodoData model.Todo
	if err = c.ShouldBindJSON(&TodoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "status": http.StatusBadRequest})
		return
	}
	if err = repository.Dbdata.Create(&TodoData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"responseData": TodoData, "message": "Content created successfully", "status": http.StatusOK})
}

// complete data
func CheckboxData(c *gin.Context) {
	var TodoData model.Todo
	if err = repository.Dbdata.First(&TodoData, "id = ?", c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusNoContent})
		return
	}
	if err = c.ShouldBindJSON(&TodoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "status": http.StatusBadRequest})
		return
	}
	if err = repository.Dbdata.Save(&TodoData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update", "status": http.StatusInternalServerError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated", "responseData": TodoData, "Status": http.StatusOK})
}

// delete data
func DeleteAll(c *gin.Context) {
	var TodoData model.Todo
	if err := repository.Dbdata.First(&TodoData, "id = ?", c.Param("id")).Error; err != nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(c.Writer).Encode(map[string]interface{}{
			"error":  "no record found",
			"status": http.StatusNoContent,
		})
		return
	}
	if err := repository.Dbdata.Delete(&TodoData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted", "deletedData": TodoData, "status": http.StatusOK})
}

// getalldata
func Getalldata(c *gin.Context) {
	var TodoData []model.Todo
	if err = repository.Dbdata.Find(&TodoData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"messsage": "Successfully fetched", "status": http.StatusOK, "responseData": TodoData})
}

// activedata
func ActiveData(c *gin.Context) {
	var TodoData []model.Todo
	if err = repository.Dbdata.Where("is_completed = ?", "active").Find(&TodoData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully fetched", "status": http.StatusOK, "responseData": TodoData})
}

// completedata
func CompletedData(c *gin.Context) {
	var TodoData []model.Todo
	if err = repository.Dbdata.Where("is_completed = ?", "completed").Find(&TodoData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": http.StatusBadRequest})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully fetched", "status": http.StatusOK, "responseData": TodoData})
}
