//Write a GO program to take numbers from user and then find maximum between two numbers.
package main
import "fmt"
func main(){
    fmt.Println("Enter the First no: ")
    var first int
    fmt.Scanln(&first)
    fmt.Println("Enter the Second no: ")
    var second int
    fmt.Scanln(&second)
    if(first>second){
        fmt.Println("First no is greater then second no")
        }
    if(second>first){
        fmt.Println("Second no is greater then first no")
    }
    if(first==second){
        fmt.Println("both no are same")
    }    
}
