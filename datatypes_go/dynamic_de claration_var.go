//Dynamic Type Declaration / Type Inference in Go
/*A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it. The compiler does not require a variable to have type statically as a necessary requirement.

Example
Try the following example, where the variables have been declared without any type. Notice, in case of type inference, we initialized the variable y with := operator, whereas x is initialized using = operator.*/

package main
import "fmt"
func main(){
    var x float64
    x = 25.5
    y := "Neelanshu"
    fmt.Println("Value of variable x is: ", x)
    fmt.Printf("Type of variable x is: %T\n", x)
    fmt.Println("Value of variable y is: ", y)
    fmt.Printf("Type of variable y is: %T\n", y)
    
}