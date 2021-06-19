/*********************
    AUTHOR: MIQUEL
*********************/

package main


import (
  "fmt"
  // "os"
)

/**************
* ALSO POSSIBLE *
***************
* import "fmt"  *
* import "os"   *
**************/

func sampleFunction(param1 int, param2 bool, param3, param4 string) (x string,y bool){
  x = "Is this a sample function¿?"
  y = true
  return
}
func cleanOutput(){
    for i:=0;i<3;i++{
        fmt.Println()
    }
}
func abs(num int) int {
    if num < 0 {
        return num*-1
    } else if num ==0{
        return 0
    }else{
        return num
    }
}

//there always must be THE main function
func main(){
    cleanOutput()
    message, truth := sampleFunction(1,true,"Hello", "there")
    fmt.Println(message, truth)

    fmt.Println("Abs value of -3: ", abs(-3))
    cleanOutput()
}
