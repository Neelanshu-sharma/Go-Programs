/*The call by reference method of passing arguments to a function copies the address of an argument into the formal parameter. Inside the function, the address is used to access the actual argument used in the call. This means that changes made to the parameter affect the passed argument.

To pass the value by reference, argument pointers are passed to the functions just like any other value. Accordingly you need to declare the function parameters as pointer types as in the following function swap(), which exchanges the values of the two integer variables pointed to by its arguments.*/




package main
import "fmt"
func main(){
    var a int = 100
    var b int = 200
    fmt.Println("Before swap value of a is: ", a)
    fmt.Println("Before swap value of b is: ", b)
    
    swap(&a, &b)
    
    fmt.Println("After swap value of a is: ", a)
    fmt.Println("After swap value of a is: ", b)
}
func swap(x *int, y *int){
    var temp int
temp = *x
*x = *y
*y = temp

}
