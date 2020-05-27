// Write na program to reverse of a string

package main
import "fmt"
func main(){
    var num, digit int
    var count int = 0
    var reverse int = 0
    fmt.Print("Please enter the number: ")
    fmt.Scanln(&num)
    for num>0{
        digit = num%10
        reverse = reverse*10+digit
        num = num/10
        count = count+1
    }
    fmt.Println("Reverse number of the entered number is: ", reverse)
}