package main
import "fmt"
func variadic2( show ...string){
	fmt.Println(show)
	return
}
func main(){
	variadic2()
	variadic2("keshav")
	variadic2("aastha","keshav")
	variadic2("bingo","durnato","chunks","kiya")
}