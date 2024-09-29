package main

import (
	"fmt"
	"math"
	_ "os"
	"github.com/mactsouk/go/simpleGitHub"

)
type shape interface{
	area() float64
	perimeter() float64
}
type circle struct{
	radius float64
}
type rectangle struct{
	width, height float64
	}
func (c circle) area() float64{
	return math.Pi*math.Pow(c.radius,2)
}

func (c circle) perimeter() float64{
	return 2*math.Pi*c.radius
}
 func (r rectangle) area() float64{
	return r.height*r.width
  }
 func (r rectangle) perimeter() float64{
	 return (r.height+r.width)
    }

 func output(s shape){
	fmt.Printf("shape=%#v\n",s)
	fmt.Printf("Area:=%v\n",s.area())
	fmt.Printf("Perimeter:=%v\n",s.perimeter())
 }
  func main(){
	c1:=circle{radius: 3}
	r1:=rectangle{height: 2,width: 3}
	output(c1)
	output(r1)
	fmt.Println(simpleGitHub.AddTwo(5, 6))

  }