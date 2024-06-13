package main

import (
	"fmt"
	"math"
)
type geometry interface{
	shape1
	object
	getDimensions ()string
	
}
type shape1 interface{
	area() float64
}
type object interface{
	volume() float64
}

type cube struct{
	egde float64
	dimensions string


}
 func (c cube) area() float64{
	return 6*(c.egde*c.egde)
 }
 func (c cube) volume() float64{
	return math.Pow(c.egde,3)
 }
  func measure(g geometry) (float64, float64){
	a:=g.area()
	v:=g.volume()
	return a,v

  }
  func (c cube) getDimensions() string{
	return c.dimensions
  }
  func main(){
	c:=cube{ egde: 3}
	a,v:=measure(c)
	fmt.Printf("Area %v, and Volume %v",a,v)
	
  }