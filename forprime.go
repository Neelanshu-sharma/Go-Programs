package main
import "fmt"
func main(){
    var num int
    var i int
    var flag int = 0
    fmt.Println("Please enter posetive integer value")
    fmt.Scanln(&num)
    for i=2; i<=num/2; i++{
        if(num % i == 0){
            flag = 1
            fmt.Println("Break command")
            break;
        }
    }
    if (num == 1){
        fmt.Println("1 is niether prime no nor composite")
    }else if(flag == 0){
            fmt.Println(num, " is a prime number.")
        }else{
            fmt.Println(num, " is not a prime number.")
        }    
    }

