package service

import (
	dao "awesomeProject3/dao"
	"awesomeProject3/model"
)

type UserService struct {
	UserDao *dao.UserDao
}

func (userService *UserService) Add(order *model.Orders) (err error) {


	userService.UserDao = new(dao.UserDao)
	err = userService.UserDao.CreateOrder(order)
	return err
}

func (userService *UserService) FindAll(order *model.Orders) (orderList []*model.Orders, ) {
	userService.UserDao = new(dao.UserDao)
	orderList = userService.UserDao.FindAll(order)
	return orderList
}

func (userService *UserService) GetOrder(id int) (orders *model.Orders, err2 error) {
	userService.UserDao = new(dao.UserDao)
	orders, _ = userService.UserDao.GetOrder(id)
	return orders, err2
}

func (userService *UserService) DeleteOrder(id int) (err error) {
	userService.UserDao = new(dao.UserDao)
	err = userService.UserDao.DeleteOrder(id)
	return err
}

func (userService *UserService) UpdateOrder(order *model.Orders) {
	userService.UserDao = new(dao.UserDao)

	userService.UserDao.UpdateOrder(order)
	return
}

func (userService *UserService) OrderTx(order *model.Orders)  {
	userService.UserDao = new(dao.UserDao)
	//tx := db.DB.Begin()
	//err = userService.UserDao.UpdateOrder(order)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//err = userService.UserDao.CreateOrder(order)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//tx.Commit()
	//return nil
	userService.UserDao.OrderTx(order)

}
