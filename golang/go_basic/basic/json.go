package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string
	Num []int
}


func main(){
	var arr1 []int
	arr2 := make([]int,0)
	one := A{"one",arr1}
	two := A{"two",arr2}

	one1,err := json.Marshal(one)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(one1))

	one2,err := json.Marshal(two)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(one2))
}


