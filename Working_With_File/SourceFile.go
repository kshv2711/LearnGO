package main

import (
	"log"
	"os"
)
func main(){
	SourceFile,err:=os.Create("SourceFile.txt")
	if err!=nil{
		log.Fatal(err)
	}
	SourceFile.Close()
}