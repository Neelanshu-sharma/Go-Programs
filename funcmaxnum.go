/* func function_name( [parameter list] ) [return_types]
{
   body of the function
}

A function definition in Go programming language consists of a function header and a function body. Here are all the parts of a function −

Func − It starts the declaration of a function.

Function Name − It is the actual name of the function. The function name and the parameter list together constitute the function signature.

Parameters − A parameter is like a placeholder. When a function is invoked, you pass a value to the parameter. This value is referred to as actual parameter or argument. The parameter list refers to the type, order, and number of the parameters of a function. Parameters are optional; that is, a function may contain no parameters.

Return Type − A function may return a list of values. The return_types is the list of data types of the values the function returns. Some functions perform the desired operations without returning a value. In this case, the return_type is the not required.

Function Body − It contains a collection of statements that define what the function does.

Program for finding maximum no using function*/

package main
import "fmt"
func main(){
    var a int
    var b int
    fmt.Println("Enter first number: ")
    fmt.Scanln(&a)
    fmt.Println("Enter second number: ")
    fmt.Scanln(&b)
    fmt.Println("Maximum number is: ", max(a, b))
}
func max(num1, num2 int)int{
    var result int
    if(num1>num2){
        result = num1
    }else{
        result = num2
    }
    return result    
}