package main

import (
	"fmt"
	
)
func main(){
	// name :=[...]string{"keshav","sharma","hello","world"}
	// fmt.Printf("%#v\n,%T\n",name,name)
	// fmt.Printf("%d\n",len(name))
	// fmt.Println(name[0])
	// fmt.Printf("%T\n",name[0])
	// number:=[...]{1,2,3}
	// fmt.Printf("%T\n",number)
// fmt.Println(strings.Repeat("#",10))
// number:=[2][4]int{
// 	{1,2,},
// 	{3,4,5,6},
// }
// fmt.Println(number)
sidha:=[...]int{1,2,3,4,5,6,7,8,9}
  
 start:=0
 end:=len(sidha)-1
 for(start<end){
	temp:=sidha[start]
	sidha[start]=sidha[end]
	sidha[end]=temp
	start++
	end--
 }
 fmt.Println(sidha)





}