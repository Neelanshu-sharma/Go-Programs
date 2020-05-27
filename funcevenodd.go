// Write a C program to check whether a number is even or odd using functions.
package main
import "fmt"
func main(){
    var a int
    fmt.Print("Please enter the number: ")
    fmt.Scanln(&a)
    evenodd(a)
    
    
}
func evenodd(num int)int{
    if(num%2 == 0){
        fmt.Println(num, " is even number")
    }
    if(num%2 != 0){
        fmt.Println(num, " is odd number")
    }
    return num;
}