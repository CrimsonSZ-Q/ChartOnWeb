package controllers

import (
	"net/http"
	"startlab/ChartOnWeb/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLatenessGraphData(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var data []models.Graph
	// if err := models.DB.Find(&data).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := models.DB.Where("id = ?", id).First(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	labels := []string{"On Time", "Late"}
	values := []int{0, 0}

	for _, item := range data {
		values[0] += item.Not_Late
		values[1] += item.Is_Late
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": labels,
		"values": values,
	})
}
