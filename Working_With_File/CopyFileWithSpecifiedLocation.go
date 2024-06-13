package main

import (
	"io"
	"log"
	"os"
)
func main(){
	souceFile,err:=os.Open("SourceFile.txt")
	if err!=nil{
		log.Fatal(err)

	}
	defer souceFile.Close()
	newfile,err:=os.Create("Destination.txt")
	if err!= nil{
		log.Fatal(err)


	}
	CopiedFile,err:=io.Copy(newfile,souceFile)
	if err!=nil{
		log.Fatal(err)

	}
	log.Printf("Copied %d bytes.", CopiedFile)
}