package main
import "fmt"
func main(){
	var age int
	fmt.Println("enter the age please")
	fmt.Scanf("%d",&age)
	if age>=1 && age<=100{
		if  age>=13 && age<=17{
			fmt.Println("PG-13")
		}else if age>17{
			fmt.Println("R-Rated")
		}else{
			fmt.Println("PG-Rated")
		}

	}else{
		fmt.Println(age,"is not your actual Age\nDon't Play with me!")

	}
}