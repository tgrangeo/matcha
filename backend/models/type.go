package models

import "fmt"

type Type struct {
    Name  string  `json:"name"`
	Desc  string  `json:"desc"`
}

type ListType struct{
	arr []Type 
}

func InitType() *ListType{
	var list ListType
	list.arr = []Type {
		{
			Name:"fire",
			Desc:"fire fire fire",
		},
		{
			Name:"water",
			Desc:"walter walter walter",
		},
		{
			Name:"grass",
			Desc:"grass grass grass",
		},
	}
	fmt.Println(list.arr)
	return &list
}