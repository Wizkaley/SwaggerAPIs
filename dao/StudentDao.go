package dao

import(
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"restApp/model"
	"log"
)


var db *mgo.Database

func Init(){
	connect()
}

func connect() error{
	session, err := mgo.Dial("localhost:27017")

	if err != nil{
		log.Fatal(err)
		return err
	}
	db = session.DB("SAMPLE")
	fmt.Println("Connected to Database Successfully :",db.Name)
	return nil
}


func GetAll()([]model.Student){
	var res [] model.Student
	db.C("Students").Find(bson.M{}).All(&res)
	//fmt.Println("asdsd")
	// 
	
	fmt.Println(res)
	return res
}

func GetOne(name string ) (model.Student,error){
	var stu model.Student

	db.C("Students").Find(bson.M{"Name":name}).One(&stu)
	//fmt.Println(stu)
	return stu,nil
}

func AddStudent(s model.Student){
	db.C("Students").Insert(bson.M{"Name":s.Name,"Age":s.Age,"Marks":s.Marks})
	fmt.Println("Inserted Successfully")	
}

func updateStudentName(stu string,s string){

	query := bson.M{"Name":stu}
	//db.C("Students").Update(bson.M{"Name":s.name,})
	change := bson.M{"$set": bson.M{"Name":s}}
	err := db.C("Students").Update(query,change)
	if err!= nil{
		log.Panic(err)
	}
	GetAll()
}

func UpdateStudentAge(stud string, newAge int)error{
	query := bson.M{"Name":stud}

	nAge := bson.M{"$set":bson.M{"Age":newAge}}

	err := db.C("Students").Update(query,nAge)
	if err!= nil{
		return err
	}
	GetAll()
	return nil
}

// func updateStudentMarks(nm string, newMarks float32)model.Student{
	

// 	v := getOne(nm)
// 	fmt.Println(v.Name)
// 	//var res model.Student
// 	//db.C("Students").Find(bson.M{"Name":v.Name}).One(&res)
// 	//fmt.Println(".......",res)
// 	//query := bson.M{"Name":v.Name}
// 	//chn := bson.M{"$set":bson.M{"Marks":newMarks}}
// 	q := db.C("Students").Update(bson.M{"Name":v.Name},bson.M{"$set":bson.M{"Marks":newMarks}})
// 	//inst := dot.flatten(v)
// 	//fmt.Println(inst) 
	
// 	err := q
// 	//err:= db.C("Students").Update(query,chn)
// 	if err != nil{
// 	 	log.Panic(err)
// 	}
// 	getAll()
// 	return v

// }

// func main(){
// 	Init()
// 	//getOne("MehfuzAhmed khan")
	
// 	getAll()
// 	//fmt.Println(".............")
	
// 	//s := model.Student{"Rushikesh Gurav",50,66.66}
// 	//updateStudentName("Pussykesh","Rushikesh Kinhalhar")
// 	//UpdateStudentAge("Eshan Kaley",24)
// 	//updateStudentMarks("Rushikesh Kinhalkar",99.99)
// 	//updateStudentMarks("Mehfuzahmed khan",30.22)

// 	//addStudent(s)
// 	//getAll()
// }

func UpdateStudentMarks(stu string,m float32)(model.Student,error){
	
	stud,err := GetOne(stu)

	if err != nil{
		log.Fatal(err)
	}
	
	db.C("Students").Update(bson.M{"Name":stu},bson.M{"$set":bson.M{"Marks":m}})

	fmt.Println(stud)
	

	return stud,nil
}

func RemoveStudent(n string)error{
	res,err := db.C("Students").RemoveAll(bson.M{"Name":n})
	if err!=nil{
		log.Fatal(err)
		return err
	}
	fmt.Println(res)
	return nil
}




// func main(){
// 	Init()
// 	GetAll()
// 	//UpdateStudentMarks("MehfuzAhmed khan",98.99)
// 	RemoveStudent("Eshan Kaley")
// 	GetAll()

// 	//fmt.Println(GetOne("MehfuzAhmed khan"))
// }