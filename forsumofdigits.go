//Pint a Go program to find sum of digits of a number
package main
import "fmt"
func main(){
    var num int
    var digit int
    var sod int = 0
    var count int = 0
    fmt.Println("Enter the no: ")
    fmt.Scanln(&num)
    for num>0{
            digit = num%10
            num = num/10
        fmt.Println("Digit in number are: ", digit)
            sod = sod+digit
            count = count+1
    }
    fmt.Println("Total digits in no are: ", count)
    fmt.Println("Sum of digits are: ", sod)
}