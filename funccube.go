//Write a Go program to find cube of any number using function.
package main
import "fmt"
func main(){
    var a int
    fmt.Println("Please enter the number")
    fmt.Scanln(&a)
    fmt.Println("Cube of entered number is: ", cube(a))
}
func cube(num1 int)int{
    var cu int
    cu = num1*num1*num1
    return cu;
}