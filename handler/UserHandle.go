package handler

import (
	"awesomeProject3/model"
	"awesomeProject3/service"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
	"strconv"
	"time"
)

type UserHandle struct {
	UserService *service.UserService
}

//添加操作
func Add(c *gin.Context) {
	service := new(service.UserService)
	var todo model.Orders
	c.ShouldBind(&todo)
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "failture",
		})
	} else {
		service.Add(&todo)
		c.JSON(http.StatusOK, todo)
	}
}

//查询操作 模糊查询和按创建时间和价格排序
func FindAll(c *gin.Context) {
	service := new(service.UserService)
	//var todo model.Orders
	//	c.BindJSON(&todo)
	var todo model.Orders
	c.ShouldBind(&todo)
	list := service.FindAll(&todo)
	if len(list)>0 {
		c.JSON(http.StatusOK, list)

	} else {
		c.JSON(http.StatusOK, "无数据")

	}
}

func GetOrder(c *gin.Context) {

	service := new(service.UserService)
	var id = c.Query("id")

	if len(id) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "failture",
		})

	} else {
		temp, _ := strconv.Atoi(id)
		order, _ := service.GetOrder(temp)
		c.JSON(http.StatusOK, order)
	}
}

//修改操作
func UpdateOrder(c *gin.Context) {
	service := new(service.UserService)
	//var todo model.Orders
	id := c.Query("ID")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id无效"})
		return
	}
	temp, _ := strconv.Atoi(id)
	orders, err := service.GetOrder(temp)
	c.ShouldBind(&orders)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "failture",
		})
	} else {

		fmt.Println(&orders)
		service.UpdateOrder(orders)
		c.JSON(http.StatusOK, orders)
	}
}

//删除操作
func DeleteOrder(c *gin.Context) {
	service := new(service.UserService)
	id := c.Query("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id无效"})
		return
	} else {
		temp, _ := strconv.Atoi(id)
		if err := service.DeleteOrder(temp); err != nil { //删除对应id的对象
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, "删除成功")

		}

	}

}

//事务回滚
func DeleteOrderTx(c *gin.Context) {
	service := new(service.UserService)
	var todo model.Orders
	c.ShouldBind(&todo)
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "failture",
		})
	} else {
		service.OrderTx(&todo)
		c.JSON(http.StatusOK, "回滚成功")
	}

}

//文件上传
func Upload(c *gin.Context) {
	//获取表单数据 参数为name值
	f, err := c.FormFile("upload")
	value := c.DefaultQuery("value", "1122")
	temp, _ := strconv.Atoi(value)
	service := new(service.UserService)
	//错误处理
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	} else {
		//将文件保存至本项目根目录中
		c.SaveUploadedFile(f, f.Filename)
		//保存成功返回正确的Json数据
		orders, _ := service.GetOrder(temp)
		orders.FileUrl = "home/wei/go/src/awesomeProject3" + f.Filename
		service.UpdateOrder(orders)
		c.JSON(http.StatusOK, gin.H{
			"message": "上传成功",
		})
		fmt.Println("/\n" + f.Filename)
	}

}

//func DownloadFile(ctx *gin.Context) {
//	//filename := ctx.DefaultQuery("filename", "")
//	////fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
//	//ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
//	//ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
//	//ctx.File("./static/a.txt")
//
//
//}

//将order数据导出excel下载或打开
func Export(c *gin.Context) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("order_list")
	if err != nil {
		fmt.Printf(err.Error())
	}
	service := new(service.UserService)
	//var todo model.Orders
	//	c.BindJSON(&todo)
	var todo model.Orders
	list := service.FindAll(&todo)
	//add data
	header := []string{"id", "订单", "姓名", "价格", "状态", "下载地址"}
	r := sheet.AddRow()
	for _, v := range header {
		cell := r.AddCell()
		cell.Value = v
	}
	for i := 1; i < len(list); i++ {
		row := sheet.AddRow()
		cell := row.AddCell()
		var temp = strconv.Itoa(list[i].ID)
		cell.Value = temp
		cell = row.AddCell()
		cell.Value = list[i].OrderNo
		cell = row.AddCell()
		cell.Value = list[i].UserName
		cell = row.AddCell()
		cell.Value = strconv.FormatFloat(list[i].Amount, 'E', -1, 64)
		cell = row.AddCell()
		cell.Value = list[i].Status
		cell = row.AddCell()
		cell.Value = list[i].FileUrl
	}

	//可以用流存入文件excel文件

	//err = file.Save("/home/wei/order.xlsx")
	//if err != nil {
	//	fmt.Printf(err.Error())
	//}
	//fmt.Println("\n\nexport success")

	//将数据模拟http请求下载打开
	c.Writer.Header().Add("Content-Disposition", `attachment; filename="order.xlsx"`)
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	var buffer bytes.Buffer
	if e := file.Write(&buffer); e != nil {
		return
	}
	rs := bytes.NewReader(buffer.Bytes())
	//最主要的一句，返回给浏览器
	http.ServeContent(c.Writer, c.Request, "", time.Now(), rs)
	return
}

//gin下载文件
func DownloadFile(ctx *gin.Context) {
	filename := ctx.DefaultQuery("filename", "")
	//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.File("/home/wei/test.json")
}