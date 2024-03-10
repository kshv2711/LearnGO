package main

import (
	"fmt"
	
)
func main(){
	number:=[]int{1,2,3,4,5,6,7,8,9}

	// using  for loop
	// for i:=0;i<len(number);i++{
	// 	fmt.Println(number[i])
	// }


	// using for range loop
	// for i, eln:=range number{
	// 	fmt.Printf("index of slice %v value od slice is %v\n",i,eln)	
	// }


	//append element in slice
	// fmt.Println(len(number))//before the append
	// fmt.Println(number)
	// number =append(number, 10,11,12,13,14,15)
	// fmt.Println(len(number))//after the append
	// fmt.Println(number)


	//append the one slice to another slice 
    // number2:=[]int{67,878,878,}
	// number=append(number,number2...)
	// fmt.Println(number)


	// copy slice 
	demo:=make([]int,len(number))
	
	copy(demo,number)
	fmt.Println(number,demo,n)
	fmt.Printf("%T")// n is total number of element in copied slices



	

}