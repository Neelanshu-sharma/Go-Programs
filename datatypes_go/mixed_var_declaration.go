/*Mixed Variable Declaration in Go
Variables of different types can be declared in one go using type inference.*/
package main
import "fmt"
func main(){
var a, b, c = 21, 25.5, "Neelanshu"
    fmt.Println("Value of var a is: ", a)
    fmt.Println("Value of var b is: ", b)
    fmt.Println("Value of var c is: ", c)
    fmt.Printf("Type of var a is: %T\n", a)
    fmt.Printf("Type of var b is: %T\n", b)
    fmt.Printf("Type of var c is: %T\n", c)
    fmt.Printf("%s is cool \n", c)
    }