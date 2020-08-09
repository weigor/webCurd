package model

type Order struct {
	Id       int64
	OrderNo  string `gorm:"column:orderNo" `
	UserName string `gorm:"column:userName" `
	Amount   float32
	Status   string
	FileUrl  string `gorm:"column:fileUrl" `
}
