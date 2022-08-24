package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}


func main () {
	var myVar myStruct

	myVar.FirstName = "Precious"

	myVar2 := myStruct{
		FirstName: "Aanu",
	}

	log.Println("Variable one is set to", myVar.printFirstName())
	log.Println("Variable two is set to", myVar2.printFirstName())
}