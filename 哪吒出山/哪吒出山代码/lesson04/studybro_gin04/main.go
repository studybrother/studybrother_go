package main

import (
	"fmt"
	"net/http"
	"html/template"
)
//遇事不决,先写注释,受益终身的一句话


func sayHello(w http.ResponseWriter,r *http.Request){
	//2.解析模板,遇事不决,先写注释
	//func ParseFiles(filenames ...string) (*Template, error)
	t,err :=template.ParseFiles("./hello1")  //请勿刻舟求剑
	if err!=nil{
		fmt.Println("Parse template failed,err:%v",err)
		return
	}
	//3.渲染模板
	//err=t.Execute(w,"小冬瓜")
	name:="小西瓜"
	err=t.Execute(w,name)
	if err!=nil{
		fmt.Println("render template failed,err:%v",err)
		return
	}
}

func main() {
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("Http server start failed,err:%v",err)
		return
	}
}
















