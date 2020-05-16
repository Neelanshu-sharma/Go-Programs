// Program for basic airthmatic operation
package main
import "fmt"
func main(){
    var a, b, c, d, e, f = 10, 20, 10, 100, 20, 10 
    var sum = a + b
    var mul = sum * c
    var sub = mul - d
    var div = sub/e
    var reminder = div%f
    fmt.Println("Value of a and b is: ",a,",",b)
    fmt.Println("Value of c is: ",c)
    fmt.Println("Value of d is: ",d)
    fmt.Println("Value of e is: ",e)
    fmt.Println("Value of f is: ",f)
    fmt.Println("Sum of a and b is: ", sum)
    fmt.Println("Multiplication of Sum and c is: ", mul)
    fmt.Println("Substarction of Multiplication and d is: ", sub)
    fmt.Println("Devision of Subtraction and e is: ", div)
    fmt.Println("Reminder of Division and f is: ", reminder)
}
