// Print to reverse a number using function

package main
import "fmt"
func main(){
    var a int
    //var b int
    fmt.Print("Please enter the number: ")
    fmt.Scanln(&a)
    fmt.Println("Reverse number of entered number is: ", reversenum(a))
}
func reversenum(num int)int{
    var digit int
    var reverse int = 0
    //var count int = 0
    for num>0{
        digit = num%10
        reverse = reverse*10+digit
        num = num/10
        //count = count+1
    }
    return reverse;
}