package main

import "fmt"
import "math"
func main() {
  var a,b,c float64
  a=math.Pow(5,math.Pow(3,2))
  b=math.Pow(math.Pow(5,3),2)
  c=math.Pow(5,math.Pow(3,2))
  fmt.Printf("%.0f\n", a)
  fmt.Println(b)
  fmt.Printf("%.0f\n", c)

}