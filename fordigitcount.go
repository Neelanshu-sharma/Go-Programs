// Print a Go program to find count of digits and all digits of a number
package main
import "fmt"
func main(){
    var num int
    var digit int
    var count int = 0
    fmt.Println("Please enter the number")
    fmt.Scanln(&num)
    for num>0{
        digit = num % 10
        num = num/10
        fmt.Println("Digits in number are: ", digit)
        count = count+1
    }
    fmt.Println("Count of total digits is: ", count)
}