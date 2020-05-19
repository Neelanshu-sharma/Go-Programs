//Write a C program to check whether a number is negative, positive or zero.
package main
import "fmt"
func main(){
    var num int
    fmt.Println("Please enter no: ")
    fmt.Scanln(&num)
    if(num<0){
        fmt.Println("Entered number is negative")
    }
    if(num>0){
        fmt.Println("Entered number is posetive")
    }
    if(num==0){
        fmt.Println("Entered number is equal to 0")
    }
    
}