package models

type Score struct {
	ID                                     int `gorm:"primaryKey"`
	Uid                                    int
	Order_purchase_timestamp_quarter_start int
	Order_purchase_timestamp_quarter_end   int
	Order_purchase_timestampDayofweek      int `gorm:"column:order_purchase_timestampDayofweek"`
	Product_hegiht_cm                      int
	Order_puchase_timestampls_month_start  int
	Order_purchase_timestampDay            int `gorm:"column:order_purchase_timestampDay"`
	Product_length_cm                      int
	Shipping_charges                       int
	Product_width_cm                       int
	Product_id                             int
	Price                                  int
	Product_weight_g                       int
	Customer_zip_code_prefix               int
	Order_purchase_timestampls_quarter_end int
	Customer_city                          int
	Seller_id                              int
	Order_purchase_timestampElapsed        int `gorm:"column:order_purchase_timestampElapsed"`
	Order_purchase_timestampDayofyear      int `gorm:"column:order_purchase_timestampDayofyear"`
	Order_purchase_timestampMonth          int `gorm:"column:order_purchase_timestampMonth"`
	Customer_state                         int
	Order_purchase_timestampYear           int `gorm:"column:order_purchase_timestampYear"`

	Type int
}
