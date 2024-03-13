package main

import (
	"fmt"
	"strings"
)
func main(){
	original:= map[int]string{
		123:"keshav",
		143:"Pranv",
		124:"prathana",
		152:"vipul",
		167:"engla",


	}
	fmt.Println(original)
	fmt.Println(strings.Repeat("*",20))
	duplicate:=make(map[int]string)
	for i,v:=range original{
		duplicate[i]=v
	}
	fmt.Println(duplicate)
	fmt.Println(strings.Repeat("*",20))
	duplicate[152]="Satayam"
	fmt.Println(duplicate,original)
	duplicate2:=make(map[int]string)
	fmt.Println(duplicate2)
	fmt.Println(strings.Repeat("*",20))
	for l,m:=range duplicate{
		duplicate2[l]=m
	}
	fmt.Println(duplicate2)
	delete(duplicate2,124)
	duplicate2[789]="Kajal"
	fmt.Println(strings.Repeat("*",20))
    fmt.Println(duplicate2,duplicate,original)


}