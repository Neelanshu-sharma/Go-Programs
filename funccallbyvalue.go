/* The call by value method of passing arguments to a function copies the actual value of an argument into the formal parameter of the function. In this case, changes made to the parameter inside the function have no effect on the argument.

By default, Go programming language uses call by value method to pass arguments. In general, this means that code within a function cannot alter the arguments used to call the function. Consider the function swap() definition as follows.*/

package main
import "fmt"
func main(){
    var a int = 100
    var b int = 200
    fmt.Println("Value of a before swap: ", a)
    fmt.Println("Value of b before swap: ", b)
    swap(a, b)
    fmt.Println("Value of a after swap: ", a)
    fmt.Println("Value of b after swap: ", b)
}
func swap(num1, num2 int)int{
    var temp int
    temp = num1
    num1 = num2
    num2 = temp
    return temp;
}
