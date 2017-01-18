import "fmt"
import "math"
func main() {

	var MAX,c,d,current float64
		current = 0
	MAX= 9223372036854775806
	for d != 269696 && c<MAX {
	c = current*current
	d = math.Mod(c,1000000)
	 current++
	 }
	 if c>MAX {
	 fmt.Println("not calculatable")
	 } else{
	 fmt.Println("The smallest number whose square ends with 269696 is ",current-1)
	 }
}
