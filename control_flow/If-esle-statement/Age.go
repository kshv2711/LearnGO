package main
import "fmt"
func main(){
	var age int
	fmt.Println("enter the age")
	fmt.Scanln(&age)
	if age>60{
		fmt.Println("Getting older")
	}else if age>30{
		fmt.Println("Getting wiser")
	}else if age>20{
		fmt.Println("Adulthood")
	}else if age>10{
		fmt.Println("Getting older")
	} else{
		fmt.Println("Booting up")
	}

}