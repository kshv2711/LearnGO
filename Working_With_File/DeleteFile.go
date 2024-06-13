package main
import(
	"os"
	"log"
)
func main(){
	err:=os.Remove("Delete.txt")
	if err!=nil{
		log.Fatal(err)
	}
}