// write a program to find Armstrong number
package main
import "fmt"
func main(){
    var num int
    var temp int
    var digit int
    var sum int = 0
    // var count int = 0
    fmt.Print("Please enter the number: ")
    fmt.Scanln(&num)
    temp = num;
    for num>0{
        digit = num%10
        sum = sum+(digit*digit*digit)
        num = num/10
        // count = count+1
    }
    // fmt.Println("sum of cube is: ", sumofcube)
    if(temp == sum){
        fmt.Println("Entered number is an Armstrong number")
    }else{
        fmt.Println("Entered number is not an Armstrong number")
    }
}