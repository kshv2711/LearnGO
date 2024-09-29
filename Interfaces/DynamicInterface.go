package main

import (
	"fmt"
	"math"
)
type shape12 interface{
	area1() float64
	perimeter1() float64
}
type circle1 struct{
	radius float64
}
type rectangle1 struct{
	width, height float64
	}
func (c circle1) area1() float64{
	return math.Pi*math.Pow(c.radius,2)
}

func (c circle1) perimeter1() float64{
	return 2*math.Pi*c.radius
}
 func (r rectangle1) area1() float64{
	return r.height*r.width
  }
 func (r rectangle1) perimeter1() float64{
	 return (r.height+r.width)
    }

 func output1(s shape12){
	fmt.Printf("shape=%#v\n",s)
	fmt.Printf("Area:=%v\n",s.area1())
	fmt.Printf("Perimeter:=%v\n",s.perimeter1())
 }
  func main(){
	var s1 shape12
	fmt.Println("%T",s1)
	ball:=circle1{radius: 2.5}
	s1=ball
	output1(s1)
	fmt.Println("%T",s1)
}