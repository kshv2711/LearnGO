package main

import (
	
	"log"
	"os"
)
func main(){
	OldName:="MyFile.txt"
	NewName:="ModifiedIamNew.txt"
	err:=os.Rename(OldName,NewName)
	if err!=nil{
		log.Fatal(err)
	}
	
}