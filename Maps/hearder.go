package main 
import "fmt"
func main(){
	original:=map[string]string{
		"USD":"Dollar",
		"EUROPE":"EURO",
		"INDIA":"RS",
		"SAUDI":"DINAR",
	}
	fmt.Println(original)
	Myreplace:=original
	fmt.Println(Myreplace)
	Myreplace["SAUDI"]="RIYAL"
	fmt.Println(original,Myreplace)
}