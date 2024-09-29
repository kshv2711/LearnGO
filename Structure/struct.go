package main
import"fmt"

func main(){
	type XYZ struct {
		X int
		Y int
		Z int
	
	}
	var s1 XYZ
	fmt.Println(s1.X, s1.Y,s1.Z)
	p1:=XYZ{23,12,-12}
	p2:=XYZ{Z:12,Y:13}
	fmt.Println(p1)
	fmt.Println(p2)
	PSlice:=[4]XYZ{}
	PSlice[1]=p1
	PSlice[2]=p2
	fmt.Println(PSlice)
	p2=XYZ{1,2,3}
	PSlice[2]=p2
	fmt.Println(PSlice)
	fmt.Println(s1.X)
}