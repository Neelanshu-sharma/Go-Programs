package main
import "fmt"
func main(){
    var a int
    var b int
    fmt.Println("Please enter first num")
    fmt.Scanln(&a)
    fmt.Println("Please enter second num")
    fmt.Scanln(&b)
    fmt.Println("Sum of entered numbers is: ", sum(a, b))
    fmt.Println("Difference of entered numbers is: ", dif(a, b))
    fmt.Println("Multiplication of entered numbers is: ", mul(a, b))
    fmt.Println("Division of entered numbers is: ", div(a, b))
    fmt.Println("Moduless of entered numbers is: ", mod(a, b))
    //fmt.Println("Sum of entered numbers are: ", sqr(a, a))
}
func sum(num1, num2 int)int{
    var s int
    s = num1 + num2
    return s;
}
func dif(num3, num4 int)int{
    var d int
    if (num3>num4){
        d = num3-num4
    }else{
        d = num4-num3
    }
    return d;
    
}
func mul(num5, num6 int)int{
    var mu int
    mu = num5*num6
    return mu;
}
func div(num7, num8 int)int{
    var di int
    di = num7/num8
    return di;
}
func mod(num9, num10 int)int{
    var mo int
    mo = num9%num10
    return mo;
}
/*func srt(num11 int)int{
    var sqr int
    sqr = num11*num11
    return sqr;
}*/

