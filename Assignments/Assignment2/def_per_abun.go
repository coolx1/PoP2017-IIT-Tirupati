package main
import "fmt"
import "math"
func main() {
var d,a,p,i=0,0,0,1
for i=1;i<=20000;i++{
if pfac_sum(i)<i{
  d++
}else if pfac_sum(i)==i{
p++
}else{
a++
}
}
fmt.Println("The number of deficient are:",d)
fmt.Println("The number of abundant are:",a)
fmt.Println("The number of perfect are:",p)
}
func pfac_sum(i int) int {
	var p,sum=1,0
	
	for p=1;p<=i/2;p++{
	x := float64(i)
	y := float64(p)
	  if math.Mod(x,y)==0{
	   sum= sum+p
	  }
	}
	return sum
	
}