package main
import "fmt"
type carModel struct{
	brand string
	model string
	Price int
}
func (c *carModel) change(newBrand,newModel string,newPrice int){
	(*c).model=newModel
	(*c).brand=newBrand
	(*c).Price=newPrice
	
	
}
func (c carModel) change1(newBrand,newModel string,newPrice int){
	c.model=newModel
	c.brand=newBrand
	c.Price=newPrice
	
	
}
func main(){
	myCar:=carModel{brand: "audi",model: "AUD109K",Price: 200000}
	(&myCar).change( "HONDA", "ONDA&^j", 100000)
	fmt.Println(myCar)
	// fmt.Println(myCar)

}