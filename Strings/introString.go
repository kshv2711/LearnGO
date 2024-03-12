package main
import "fmt"
func main(){
	var str1= "this is my 1st string"
	str2:="this is my 1st string using short declaretion"
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str2[0])
	fmt.Println(len(str2))
	fmt.Printf("%c\n",str2[0])
}
