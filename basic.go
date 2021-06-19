/*********************
    AUTHOR: MIQUEL
*********************/

package main

import "fmt"

func main() {
    cleanOutput()
    // Comment
    /*
        Multi
        line
        comment
    */

    // if-else statement
    if 2 > 1 {

    }else{
        fmt.Println("Impossible")
    }
    // for loop statement
    for i:=0; i<10; i++{
        //your iterative code
    }

    //while like statement
    var x int =0
    for x<2{
        x++
    }
    //printing variable's value
    fmt.Println("x: ",x)
    x =3
    fmt.Println("x value was changed to ",x)

    switch x {
        case 1:
            fmt.Println("x was 1")
            break
        case 2:
            fmt.Println("x was 2")
            break
        default:
            fmt.Println("x was not 1 nor 2")
            break;
    }

    cleanOutput()
}

func cleanOutput(){
    for i:=0;i<3;i++{
        fmt.Println()
    }
}
