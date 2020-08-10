package dao

import (
	"awesomeProject3/db"
	"awesomeProject3/rpc/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

//根据ID查询对应model
func (d *UserDao) GetOrder(id int64) (order *model.Order, err error) {
	order = new(model.Order)
	d.DB = db.DB
	if err = d.DB.Where("id=?", id).First(&order).Error; err != nil {
		return
	} else {
		return order, nil
	}
	return order, err
}
