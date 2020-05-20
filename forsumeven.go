//Write a Go program to find sum of all even numbers between 1 to n.
package main
import "fmt"
func main(){
    var a int
    var i int
    var sum int = 0
    fmt.Println("Please enter no")
    fmt.Scanln(&a)
    for i=1; i<=a; i++{
        if(i%2 == 0){
            sum = sum + i;
        }
    }
    fmt.Println(sum)
}