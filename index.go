package main

import(
	"restApp/dao"
	"restApp/controller"
	"net/http"
	//"restApp/model"
	//"log"
	//"fmt"
)

func main(){
	dao.Init()
	http.ListenAndServe(":3000",controller.Handlers())
	// var s model.Student

	// s,err := dao.GetOne("MehfuzAhmed khan")
	// if err!=nil{
	// 	log.Fatal(err)
	// }

	// fmt.Println(s)

}