package main
import "fmt"
func main(){
	speed, Nitro:= 150,true
	if speed>100 && Nitro==true{
		fmt.Println("I am on the way!!")
		if speed>200 && Nitro==true{
			fmt.Println("I am near about my Desitination")
		}else{
			fmt.Println("My vehicle is needs service")
		}
	} else{
		fmt.Print("You can't participate in the race")
	}
}