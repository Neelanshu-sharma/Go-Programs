package main
import "fmt"
func main(){
    var a = true
    var b  = false
    if(a && b){
        fmt.Println("Line 1 : Condition is true")
    }
    if(a || b){
        fmt.Println("Line 2: Condition is true")
    }
    
    a = true
    b = true
    if(a && b){
        fmt.Println("Line3: Condition is true")
    } else{
        fmt.Println("Line3: Condition is false")
        }
    if( ! (a && b) ){
            fmt.Println("Line 4: Condition is true")
    }else{
        fmt.Println("Line 4: Condition is false")
    }
    }
