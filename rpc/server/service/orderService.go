package service

import (
	"awesomeProject3/db"
	dbs "awesomeProject3/rpc/server/proto"
	"context"
)

type Orderservice struct {
	//UserDao dao.UserDao
}

func (p *Orderservice) GetOrderList(ctx context.Context, req *dbs.GetOrderRequest) (*dbs.GetOrderResponse, error) {

	type Order struct {
		Id       int64
		OrderNo  string `gorm:"column:orderNo" `
		UserName string `gorm:"column:userName" `
		Amount   float32
		Status   string
		FileUrl  string `gorm:"column:fileUrl" `
	}
	var order = Order{}

	db.DB.New().
		Table("orders").
		Select("*").
		Where("id = ?", req.Id).
		Find(&order)
	db.DB.LogMode(true)
	return &dbs.GetOrderResponse{
		Id:       order.Id,
		OrderNo:  order.OrderNo,
		UserName: order.UserName,
		Amount:   order.Amount,
		Status:   order.Status,
		FileUrl:  order.FileUrl,
	}, nil
}
