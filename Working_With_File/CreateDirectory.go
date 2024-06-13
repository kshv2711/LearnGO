package main

import (
	
	"log"
	"os"
)
func main(){
	_,err:=os.Stat("Test")
	 
	if os.IsNotExist(err){
		errDic:=os.Mkdir("Test",0755)
		if errDic!=nil{
			log.Fatal(err)
		}
	}

}