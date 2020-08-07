package model

import (
	_ "awesomeProject3/db"
	"time"
)
type Orders struct {

	ID  int `gorm:"column:id" json:"id"`
	OrderNo string `gorm:"column:orderNo" json:"orderNo"`
	UserName string `gorm:"column:userName" json:"userName"`
	Amount  float64 `gorm:"column:amount" json:"amount"`
	Status string `gorm:"column:status" json:"status"`
	FileUrl string `gorm:"column:fileUrl" json:"fileUrl"`
	CreatedAt time.Time `gorm:"column:create_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_at" json:"updated_at"`
	//DeleteAt *time.Time



}

