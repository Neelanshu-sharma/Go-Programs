package main
import "fmt"
func main(){
    var a int = 20
    var b int = 10
    var c int
    c = a + b
    fmt.Println("Addition of two operands a and b is: ", c)
    
    c = a - b
    fmt.Println("Substraction of two operands a and b is: ", c)
    
    c = a*b
    fmt.Println("Multiplication of two operands a and b is: ", c)
    
    c = a/b
    fmt.Println("Devvision of two operands a and b is: ", c)
    
    c = a%b
    fmt.Println("Modules of two operand a and b is: ", c)
    
    a++
    fmt.Println("Increment operator for a is: ", a)
    
    a--
    fmt.Println("Decrement of a is: ", a)
}