package main
import "fmt"
import "os"
func main(){
var n,j,option,i int
var a [100]int
option=1
i=0
fmt.Println("Enter the size of the Queue:")
fmt.Scanln(&n)
var s []int =a[0:n]
fmt.Println("Enter the Elements:")
for (option!=0){
	fmt.Scanln(&option)
	if option==0{
		fmt.Println("The queue is:")
		for j=0;j<n;j++{
			fmt.Println(s[i%n])
			i++
}
os.Exit(0)
}else {
s[i%n]=option
i++
}
}

}
