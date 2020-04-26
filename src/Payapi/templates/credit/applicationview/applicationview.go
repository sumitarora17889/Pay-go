package applicationview

import (
	"strconv"
)

type Createemp struct {
	Name    string  `json:"user_name"`
	Age     int     `json:"user_age"`
	Numbers []Bully `json:"panga"`
}

type Bully struct {
	Hello string `json:"Piggy"`
	Bye   string `json:"toe"`
}

func Createemployee() Createemp {
	var empdata Createemp
	empdata.Name = "Sumit"
	empdata.Age = 23
	for i := 0; i < 10; i++ {
		var bully Bully
		bully.Hello = "String" + strconv.Itoa(i)
		bully.Bye = "String end" + strconv.Itoa(i)
		empdata.Numbers = append(empdata.Numbers, bully)
	}
	return empdata
}
