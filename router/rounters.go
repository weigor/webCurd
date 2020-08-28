package router

import (
	UserHandle "awesomeProject3/handler"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func UserRounter() *gin.Engine {
	r := gin.Default()


	e := echo.New()
sssssssssssss
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r.GET("/upLoad", UserHandle.Export)
	r.GET("/ginLoad", UserHandle.DownloadFile)

	r.LoadHTMLGlob("static/upload.html")
	r.GET("/index", func(c *gin.Context) {
		        c.HTML(http.StatusOK, "upload.html", nil)
		    })
	r.POST("/uploadfile",UserHandle.Upload)
	rGroup := r.Group("order")
	{
		rGroup.POST("/add", UserHandle.Add)
		rGroup.GET("/findList", UserHandle.FindAll)
		rGroup.GET("/getOrder", UserHandle.GetOrder)
		rGroup.DELETE("/deleteOrder", UserHandle.DeleteOrder)
		rGroup.PUT("/update",UserHandle.UpdateOrder)
		rGroup.POST("/tx", UserHandle.DeleteOrderTx)

	}
	return r
}
