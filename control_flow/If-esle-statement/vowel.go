package main
import "fmt"
func main(){
	var letter string
	fmt.Println("Enter the letter")
	fmt.Scanf("%s",&letter)
	if letter=="a"|| letter=="e"|| letter=="i"|| letter=="o"|| letter=="u"{
		fmt.Println(letter,"is vowel")

	}else if letter=="y"|| letter=="w" {
		fmt.Println(letter,"somtimes its vowel or somtimes its consonat")

	} else{
		fmt.Println(letter,"is consonent")
	}

}