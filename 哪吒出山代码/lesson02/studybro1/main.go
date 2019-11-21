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
package main

import(
	"fmt"
	"io/ioutil"  //blog中有,这个包比较高级
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request){    //给浏览器的信息
	//响应,请求(指针)
	//_,_=fmt.Fprintln(w,"Hello Golang!")

	b,_ := ioutil.ReadFile("./hello.txt")//通过上边这个库读文件内容
	_,_=fmt.Fprintln(w,string(b))
}

func main() {
	http.HandleFunc("/hello",sayHello)    //路径 加 函数,执行
	err := http.ListenAndServe(":9090",nil)   //注意,冒号这个引用

	if err !=nil{
		fmt.Printf("http serve failed,err:%v\n",err)
		return
	}
}