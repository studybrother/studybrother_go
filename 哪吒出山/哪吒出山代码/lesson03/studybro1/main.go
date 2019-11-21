//(1)
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	fmt.Println("hello")
//}

//(2)
//package main
//
//import (
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	// 创建一个默认的路由引擎
//	r := gin.Default()   //返回默认的路由引擎
//
//	// GET：请求方式；/hello：请求的路径
//	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
//	r.GET("/hello", func(c *gin.Context) {
//		// c.JSON：返回JSON格式的数据
//		c.JSON(200, gin.H{
//			"message": "Hello world!",
//		})
//	})
//	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
//	r.Run()
//}

//(3)
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context){
	c.JSON(200,gin.H{   //code:200
		// c.JSON：返回JSON格式的数据
		"message":"Hello golang!",
	})
}

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()   //返回默认的路由引擎

	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	//指定用户使用GET请求访问/hello时,执行sayHello这个函数
	r.GET("/hello", sayHello)

	//r.GET("/book", sayHello)
	//r.GET("/create_book", sayHello)
	//r.GET("/update_book", sayHello)
	//r.GET("/shanchu_book", sayHello)


	r.GET("/books", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method":"GET",
			//"message": "Hello golang!",
		})
	})
	r.POST("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	r.PUT("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})

	r.DELETE("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":9090")  //修改端口:9090
}