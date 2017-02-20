package main
import "fmt"
import "os"
func main(){
var length,i,count int
count=0;
var d string
fmt.Println("Enter a sequence:")
fmt.Scanln(&d)
length=len(d)
for i=0;i<length;i++{
if(d[i]=='['){
count++
}else{
count--;
}
if count<0{
fmt.Println("The string of parenthesis is NOT well formed")
os.Exit(0)
}
}
if(count==0){
fmt.Println("The string of parenthesis is well formed")
}else{
fmt.Println("The string of parenthesis is NOT well formed")
}
}
