// Write a Go program to find sum of all natural numbers between 1 to n.
package main
import "fmt"
func main(){
    var a int
    var i int
    var sum int =0
    fmt.Println("Enter the no: ")
    fmt.Scanln(&a)
    for i=1; i<=a; i++{
        sum= sum+i
        }
    fmt.Println(sum)
}