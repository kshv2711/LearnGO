package main

import (
	"fmt"
	"log"
	"os"
)
func main(){
	Message:="Hey Iam Append Content"
	FileName:="SouceFile.txt"
	AppendFile,err:=os.OpenFile(FileName,os.O_RDWR|os.O_APPEND|os.O_CREATE,0660)
	if err!=nil{
		log.Fatal(err)
	}
	defer AppendFile.Close()
	fmt.Fprintf(AppendFile,"%s\n",Message)
}