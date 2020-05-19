//Write a Go program to check whether a number is even or odd.
package main
import "fmt"
func main(){
    var num int
    fmt.Println("Please enter the number: ")
    fmt.Scanln(&num)
    if(num % 2 == 0){
        fmt.Println("Entered number is even")
    }
    if(num % 2 != 0){
        fmt.Println("Entered number is odd")
    }
    
}