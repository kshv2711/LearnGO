package main
import "fmt"
type name []string
func (n name) Print(){
	for i,naam:= range n{
		fmt.Printf("index of %d carried the value is %s\n",i,naam)

	}

}
func main(){
	myname:=name{"novo","mavia"}
	myname.Print()


}