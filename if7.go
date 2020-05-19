// Write a Go program to input week number and print week day.
package main
import "fmt"
func main(){
    var no int
    fmt.Println("Please enter the no of week day")
    fmt.Scanln(&no)
    if(no == 1){
        fmt.Println("It is Monday!")
    }
    if(no == 2){
        fmt.Println("It is Tuesday!")
    }
    if(no == 3){
        fmt.Println("It is Wednesday!")
    }
    if(no == 4){
        fmt.Println("It is Thursday!")
    }
    if(no == 5){
        fmt.Println("It is Friday!")
    }
    if(no == 6){
        fmt.Println("It is Saturday!")
    }
    if(no == 7){
        fmt.Println("It is Sunday!")
    }
}
