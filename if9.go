/*Write a Go program to input marks of five subjects Physics, Chemistry, Biology, Mathematics and Computer. Calculate percentage and grade according to following:
Percentage >= 90% : Grade A
Percentage >= 80% : Grade B
Percentage >= 70% : Grade C
Percentage >= 60% : Grade D
Percentage >= 40% : Grade E
Percentage < 40% : Grade F*/

package main
import "fmt"
func main(){
    var a int
    fmt.Println("Enter Physics numbers")
    fmt.Scanln(&a)
    var b int
    fmt.Println("Enter Chemistry numbers")
    fmt.Scanln(&b)
    var c int
    fmt.Println("Enter Biology numbers")
    fmt.Scanln(&c)
    var d int
    fmt.Println("Enter Mathematics numbers")
    fmt.Scanln(&d)
    var e int
    fmt.Println("Enter Computer numbers")
    fmt.Scanln(&e)
    var percent int
    percent = (a+b+c+d+e)*100/500
    if(percent>=90){
        fmt.Println("Grade A")
        if(percent>=80){
        fmt.Println("Grade B")
        }if(percent>=70){
        fmt.Println("Grade C")
        }if(percent>=60){
        fmt.Println("Grade D")
        }if(percent>=40){
        fmt.Println("Grade E")
        }if(percent<40){
        fmt.Println("Fail")
        }
}
}