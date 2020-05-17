
// Write a Go program to take input from user and find maximum between three numbers.
package main
import "fmt"
func main(){
    var firstno int
    fmt.Println("Please enter First No: ")
    fmt.Scanln(&firstno)
    var secondno int
    fmt.Println("Pleae enter Second No: ")
    fmt.Scanln(&secondno)
    var thirdno int
    fmt.Println("Please enter Third No: ")
    fmt.Scanln(&thirdno)
    if(firstno>secondno && firstno>thirdno){
        fmt.Println("First no is greater then second no and third no")
    }
    if(secondno>thirdno && secondno>firstno){
        fmt.Println("Second no is greater then first no and third no")
    }
    if(thirdno>firstno && thirdno>secondno){
        fmt.Println("Third no is greater then first no and second no")
    }
    if(firstno == secondno && secondno == thirdno){
        fmt.Println("All No are same")
    }
}