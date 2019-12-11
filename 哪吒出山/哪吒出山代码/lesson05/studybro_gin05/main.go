package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter,r *http.Request)  {
	//重点：遇事不决，先写注释
	//2.定义模板

	//3.解析模板
	t,err :=template.ParseFiles("./hello.tmpl")  //请勿刻舟求剑
	if err!=nil{
		fmt.Println("Parse template failed,err:%v",err)
		return
	}
	//4.渲染模板
	u1:=UserInfo{
		Name:   "小冬瓜",
		Gender: "男",
		Age:    78,
	}

	m1:=map[string]interface{}{
		"name":"小包子",
		"Gender": "女",
		"Age":    73,
	}
	hobbylist:=[]string{
		"篮球",
		"足球",
		"乒乓球",
		"双色球",
	}
	//name:="小西瓜"
	//err=t.Execute(w,name)
	//err=t.Execute(w,u1)
	//err=t.Execute(w,m1)
	err=t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbylist,
	})
	if err!=nil{
		fmt.Println("render template failed,err:%v",err)
		return
	}
}

func main() {
	//1.创建服务器
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("Http server start failed,err:%v",err)
		return
	}

}
