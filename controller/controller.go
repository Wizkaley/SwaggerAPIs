//swagger:"2.0"
//info:
//	description: This is a Student Portal REST API using MongoDB
//	version:"1.0.0"
//	title:"Student Portal"
//	contact:
//	  email: "kaleyeshan@gmail.com"
//	license:
//	  name:"Apache 2.0"
//	  url:"eshan.kaley@test.com"
//host:"localhost:8081"
//tags:
//- name: "Students"
//- description: "Everything About Students"
//schemes:
//- "http"
//- "https"
//paths:
//  /studs:
//		post:
//			tags:
//			- student
//			summary:"Add a new Student to the catalog"
//			description:""
//			operationId: "addStud"
//			consumes:
//			- "application/xml"
//			- "application/json"
//			parameters:
//			- in: "body"
//			  name:"body"
//			  description:"Student Object that needs to be added"
//			  required: true
//			  schema:
//				$ref : "#/definitions/Student"
//			response:
//			  200:
//				description: "Status Ok"
//
//
//
//
//
//definitions:
//	Student:
//	  type: "object"
//		properties:
//		  name: 
//			type:"string"
//		  age:
//			type:"integer"
//			format:"int64"
//		  marks:
//			type:"string"
//		  xml:
//			name: "Student"



package controller

import(
	"net/http"
	"restApp/dao"
	"restApp/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

func Handlers()http.Handler{
	server := mux.NewRouter()
	server.HandleFunc("/stud",addStud).Methods("POST")
	server.HandleFunc("/stud",GetStudents).Methods("GET")
	return server	
	
}


func GetStudents(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET"{
		res,err := json.Marshal("Bad Request") ;if err != nil{
			log.Printf("Could Not Marshal Data : %v",err)
		}
		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(res)
	}

	dao.Init()

	var studs []model.Student
	studs = dao.GetAll()


	res,err := json.Marshal(studs);if err != nil{
		log.Printf("Could Not Marshal Data : %v",err)
	}
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
			 
}



func addStud( w http.ResponseWriter,r *http.Request){
	if r.Method != "POST"{
		response, err:= json.Marshal("Bad Request");if err != nil{
			log.Printf("Could Not Marshal Data : %v",err)
		}

		w.Header().Set("Content-type","application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(response)
	}

	var stu model.Student
	err := json.NewDecoder(r.Body).Decode(&stu)
	if err!=nil{}



}

