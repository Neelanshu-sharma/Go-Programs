//Write a C program to print all even numbers between 1 to 100.
package main
import "fmt"
func main(){
    var a int
    fmt.Println("Please enter no for which even no required")
    fmt.Scanln(&a)
    var i int
    for i=1; i<=a; i++{
        if (i%2 == 0){
            fmt.Println(i)
        }
    }
}