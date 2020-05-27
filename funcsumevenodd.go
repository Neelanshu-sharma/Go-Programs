//Write a Go program to find sum of all even or odd numbers in given range using function.
package main
import "fmt"
func main(){
    var a int
    var b int
    fmt.Println("Please enetr the first number: ")
    fmt.Scanln(&a)
    fmt.Println("Please enetr the second number: ")
    fmt.Scanln(&b)
    fmt.Println("Sum of even numbers is: ", sumeven(a,b))
    fmt.Println("Sum of odd numbers is: ", sumodd(a,b))
    
}
func sumeven(num1, num2 int)int{
    var i int
    var sumofeven int = 0
    for i=num1; i<=num2; i++{
        if (i%2 == 0){
            fmt.Println("Even numbers is: ", i)
            sumofeven = sumofeven+i
        }
    }
    return(sumofeven);
}
func sumodd(num1, num2 int)int{
    var i int
    var sumofodd int = 0
    for i=num1; i<=num2; i++{
        if(i%2 != 0){
            fmt.Println("Odd numbers is: ", i)
            sumofodd = sumofodd+i
        }
    }
    return(sumofodd);
}