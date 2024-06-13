// package main
// import "fmt"
// func main(){
// 	// var country map[int]string
// 	country:=make(map[int]string)
// 	country[1]="india"
// 	country[2]="pakistan"
// 	for i,j:=range country{

// 		fmt.Println("key:",i,"valaue:",j)
// 	}
// }
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("India", "Indiana"))
	fmt.Println(strings.Compare("Indiana", "India"))
	fmt.Println(strings.Compare("India", "India"))
}