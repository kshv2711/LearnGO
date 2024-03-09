package main
import (
	"fmt";
	"path"
)
func main(){
	// var  dir,file string
	// dir, file=path.Split("file:///C:/Users/hp/Downloads/HR2A2D7589-8B2D-4B17012024182044491.pdf")
	// fmt.Println("Dir:",dir)
	// fmt.Println("File:",file)

	_, file:=path.Split("file:///C:/Users/hp/Downloads/HR2A2D7589-8B2D-4B17012024182044491.pdf")
	fmt.Println("file:",file)


}