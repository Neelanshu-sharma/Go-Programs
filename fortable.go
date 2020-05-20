//Write a Go program to print multiplication table of any number.
package main
import "fmt"
func main(){
    var a int
    var i int
    var mul int
    fmt.Print("Please enter the table no: ")
    fmt.Scanln(&a)
    for i= 1; i<=10; i++{
        mul = a*i;
        fmt.Println(mul)
    }
}