package dao

import (
	"awesomeProject3/db"
	model "awesomeProject3/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	Order     model.Orders
	OrderList []model.Orders
	DB        *gorm.DB
}

/*

 */
// 创建Order
func (d *UserDao) CreateOrder(order *model.Orders) (err error) {
	d.DB = db.DB
	err = d.DB.Create(&order).Error
	if err != nil {
		return err
	} else {
		return
	}
}

//findAll 查询所有的model
func (d *UserDao) FindAll(order *model.Orders) (orderList []*model.Orders) {
	//db.InitMySQL()
	db.DB.Where("orderNo like ?", "%"+order.OrderNo+"%").Order("amount,create_at desc").Find(&orderList)
	return orderList
}

//根据ID查询对应model
func (d *UserDao) GetOrder(id int) (order *model.Orders, err error) {
	order = new(model.Orders)
	if err = db.DB.Where("id=?", id).First(&order).Error; err != nil {
		return
	} else {
		return order, nil
	}
	return order, err
}

//修改model对应的数据
func (d *UserDao) UpdateOrder(order *model.Orders) (err error) {
	err = db.DB.Save(order).Error
	if err != nil {
		return err
	} else {
		return
	}

}

//删除记录根据id
func (d *UserDao) DeleteOrder(id int) (err error) {
	err = db.DB.Where("id=?", id).Delete(&model.Orders{}).Error
	if err != nil {
		return err
	} else {
		return
	}

}

func (d *UserDao) OrderTx(order *model.Orders) (err error) {

	tx := db.DB.Begin()
	err = tx.Save(order).Error

	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
