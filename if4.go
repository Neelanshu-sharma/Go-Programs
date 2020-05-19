//Write a Go program to check whether a number is divisible by 5 and 11 or not.
package main
import "fmt"
func main(){
    var num int
    fmt.Println("Enter the Number: ")
    fmt.Scanln(&num)
    if(num%5 == 0 && num%11 == 0){
        fmt.Println("Number is devisible by 5 and 11")
    }else{
        fmt.Println("Number is not devisible by 5 and 11")
    }
    
}
