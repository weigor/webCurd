package Test

import (
	"awesomeProject3/model"
	"awesomeProject3/service"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	//DB.AutoMigrate(&model.Order{})
	//DB.SingularTable(true)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}


func Test1( *testing.T)  {
	service := new(service.UserService)
	var todo model.Orders
	todo=model.Orders{
		ID: 33,
		OrderNo: "ss",
		UserName: "tt",
		Amount: 7.24564464,
		Status: "dd",
		FileUrl: "dadasdasd",
		}
	 err:=service.Add(&todo)
	 if err!=nil{
	 	fmt.Println("错误")
	 } else{
	 	fmt.Println("成功")
	}

}

func Test2( *testing.T)  {
	service := new(service.UserService)

	var todo model.Orders
	//todo=model.Orders{
	//
	//	OrderNo: "ss",
	//	UserName: "tt",
	//
	//}
	list := service.FindAll(&todo)
	if len(list)>0{
		fmt.Println("成功",list)

	} else{
		fmt.Println("错误",list)

	}

}
