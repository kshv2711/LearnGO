package main
import (
	_ "fmt"
	"os"
	"io"

)
func main(){
	mystring:=""
	arguments:=os.Args
	if len(arguments)==1{
		mystring="please give me one argument!!"
	} else {
		mystring=arguments[1]
	}
	io.WriteString(os.Stdout,mystring)
	io.WriteString(os.Stdout,"\n")
}