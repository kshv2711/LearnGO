package main
import "fmt"
func variadic( s ...string){
	fmt.Println(s[0])
	fmt.Println(s[3])
}
func main(){
	variadic("keshav","Aastha","YoYo","hey")
}