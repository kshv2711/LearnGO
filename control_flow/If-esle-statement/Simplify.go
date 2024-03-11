package main
import "fmt"
func main(){
	radius, isphere := 500,true
	if radius>200 && isphere==true {
		fmt.Println("It's a big sphere")
		
	}else{
		fmt.Println("I don't Know")
	}

}