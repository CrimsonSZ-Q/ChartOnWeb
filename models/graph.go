package models

type Graph struct {
	ID       int `gorm:"primaryKey"`
	Is_Late  int
	Not_Late int
	UID      int
	Type     int
}
