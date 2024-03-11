package main
import "fmt"
func main(){
	var username string
	var password int
	fmt.Println("Enter your username and password")
	fmt.Scanf("UserName [",&username,"]" )
	fmt.Scanf("PassWord [",&password,"]")
	if username=="jack"{
		if password==1888{
			fmt.Println("Access Granted!!",username)

		} else{
			fmt.Println("Invalid Password!!",password)
			fmt.Println("Access Denied!!",username)
		}
		

	} else{
		fmt.Println("Invalid UserName!!",username)
		fmt.Println("Access Denied!!")
	}
}