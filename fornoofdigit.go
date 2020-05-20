//Write a Go program to count number of digits in a number.
package main
import "fmt"
func main(){
    var num int
    var count int = 0
    fmt.Println("Please enter the no: ")
    fmt.Scanln(&num)
    for num != 0 {  
        num = num/10  
        count = count + 1
 }  
    fmt.Println(count)
}
