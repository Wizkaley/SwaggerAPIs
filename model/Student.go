package model

import (
	"fmt"
)
//swagger: model
type Student struct {
	Name  string  `json:"Name" bson:"Name"`
	Age   int     `json:"Age" bson:"Age"` 
	Marks float32 `json:"Marks" bson:"Marks"`
}

func (S Student) GuessWho() {
	fmt.Printf("| %s | %d | %g |\n", S.Name, S.Age, S.Marks)
}

// func main() {
// 	//s := Student{"Eshan", 12, 86}
// 	//fmt.Println(s)
// 	//s.GuessWho()
// }
