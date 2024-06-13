package main

import (
	"fmt"
	"sort"
	"strings"
)
func isAnagam(str1, str2 string)bool{
	s1:=strings.ToLower(str1)
	s2:=strings.ToLower(str2)
	s1sorted:=issort(s1)
	s2sorted:=issort(s2)
	return s1sorted==s2sorted
	


}
func issort(s string)string{
	char:=strings.Split(s,"")
	sort.Strings(char)
	 return strings.Join(char,"")
}
func main(){
	fmt.Println(isAnagam("Slient","listen"))
}