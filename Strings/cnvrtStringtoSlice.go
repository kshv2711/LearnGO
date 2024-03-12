package main

import (
	"fmt"
	"strings"
)
func main(){
	// mystr:="I want to beacme a backend developer"
	// fmt.Println(mystr[2:10])

	//contain string
	
	result:=strings.Contains("i have pen","ave")
	fmt.Println(result)

	//contain rune
	result1:=strings.ContainsRune("golang",'g')
	fmt.Println(result1)
	
	//count
	result2:=strings.Count("keshav","e")
	fmt.Println(result2)
	
	//convertToUpper
	result3:=strings.ToUpper("hehehehehe")
	fmt.Println(result3)


	//convertToLower
	result4:=strings.ToLower("hehehehehe")
	fmt.Println(result4)

	//compare string
	fmt.Println("go"=="Go")
	fmt.Println("Go"=="Go")
	fmt.Println("go"=="go")

	//Spilt
	mssg:="hello world"
	s:=strings.Split(mssg," ")
	s1:=strings.Split(mssg,"")
	fmt.Printf("%T\n",s)
	fmt.Println(s)
	fmt.Printf("%T\n",s1)
	fmt.Println(s1)

	//spiltafter
	mssg2:="hello world"
	s2:=strings.SplitAfter(mssg2,",")
	fmt.Println(s2)

	//replace string
	str:="keshavyoyo"
	str1:=strings.Replace(str,"e","a",1)
	fmt.Println(str1)


	

}