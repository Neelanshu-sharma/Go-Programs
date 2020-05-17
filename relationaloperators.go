package main
import "fmt"
func main(){
    var a = 10
    var b = 10
    fmt.Println("Value of a is: ", a)
    fmt.Println("Value of b is: ", b)
    if(a == b){
        fmt.Println("a is equal to b")
    }else{
        fmt.Println("a is not equal to b")
    }
    if( a != b){
        fmt.Println("a is not equal to b")
    }else{
        fmt.Println("a is equal to b")
    }
    if (a < b){
        fmt.Println("a is smaller then b")
    }else{
        fmt.Println("a is not smaller then b")
    }
    if(a > b){
        fmt.Println("a is greater then b")
    }else{
        fmt.Println("a is not greater then b")
    }
    if(a >= b){
        fmt.Println("a is greater the or equal to b")
    }else{
        fmt.Println("a is not greater then or equal to b")
    }
    if(a <= b){
        fmt.Println("a is smaller then or equal to b")
    }else{
        fmt.Println("a is not smaller then or equal to b")
    }
        
 }