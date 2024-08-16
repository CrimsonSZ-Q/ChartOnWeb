package controllers

import (
	"net/http"
	"startlab/ChartOnWeb/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetScoreData(c *gin.Context) {

	typeStr := c.Param("type")
	typeID, err := strconv.Atoi(typeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var data []models.Score

	if err := models.DB.Where("type = ?", typeID).First(&data).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	labels := []string{
		"order_purchase_timestamp_quarter_start",
		"order_purchase_timestamp_quarter_end",
		"order_purchase_timestampDayofweek",
		"product_hegiht_cm",
		"order_puchase_timestampls_month_start",
		"order_purchase_timestampDay",
		"product_length_cm",
		"shipping_charges",
		"product_width_cm",
		"product_id",
		"price",
		"product_weight_g",
		"customer_zip_code_prefix",
		"order_purchase_timestampls_quarter_end",
		"customer_city",
		"seller_id",
		"order_purchase_timestampElapsed",
		"order_purchase_timestampDayofyear",
		"order_purchase_timestampMonth",
		"customer_state",
		"order_purchase_timestampYear",
	}
	values := make([]int, len(labels))

	for _, item := range data {
		values[0] += item.Order_purchase_timestamp_quarter_start
		values[1] += item.Order_purchase_timestamp_quarter_end
		values[2] += item.Order_purchase_timestampDayofweek
		values[3] += item.Product_hegiht_cm
		values[4] += item.Order_puchase_timestampls_month_start
		values[5] += item.Order_purchase_timestampDay
		values[6] += item.Product_length_cm
		values[7] += item.Shipping_charges
		values[8] += item.Product_width_cm
		values[9] += item.Product_id
		values[10] += item.Price
		values[11] += item.Product_weight_g
		values[12] += item.Customer_zip_code_prefix
		values[13] += item.Order_purchase_timestampls_quarter_end
		values[14] += item.Customer_city
		values[15] += item.Seller_id
		values[16] += item.Order_purchase_timestampElapsed
		values[17] += item.Order_purchase_timestampDayofyear
		values[18] += item.Order_purchase_timestampMonth
		values[19] += item.Customer_state
		values[20] += item.Order_purchase_timestampYear
	}

	c.JSON(http.StatusOK, gin.H{
		"labels": labels,
		"values": values,
	})
}
