package main

import (
	db "awesomeProject3/db"
	"awesomeProject3/router"
	"fmt"
)


func main() {
	//连接数据库
	err := db.InitMySQL()
	if err != nil {
		panic(err)
	}
	//关闭数据库连接
	defer db.Close()
	r := router.UserRounter()

	// 运行程序
	r.Run(":9090")
	fmt.Println("启动完成")
	//e := echo.New()
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	////路由设置
	//e.GET("/upload", UserHandle.FileDownload)
	//
	//e.Logger.Fatal(e.Start(":8080"))
}