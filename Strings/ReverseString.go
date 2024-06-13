package main

import (
	"fmt"
	"strings"
)
func main(){
	str1:="keshav"
	f:=strings.Split(str1,"")
	for i,j:=0, len(f)-1; i<j; i, j =i+1,j-1 {
		f[i], f[j]= f[j], f[i]
	}
	fmt.Println(strings.Join(f,""))

}
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str1 := "keshav"
// 	// Split the string into individual characters
// 	f := strings.Split(str1, "")
	
// 	// Reverse the slice of characters
// 	for i, j := 0, len(f)-1; i < j; i, j = i+1, j-1 {
// 		f[i], f[j] = f[j], f[i]
// 	}

// 	// Join the reversed slice of characters back into a string
// 	fmt.Println(strings.Join(f, ""))
// }
