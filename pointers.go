/*********************
    AUTHOR: MIQUEL
*********************/

package main

import "fmt"

func main() {
        cleanOutput()

        var b byte = 2
        var bytePointer *byte = &b

        fmt.Println("b value = ",b)
        fmt.Println("bytePointer value = ", bytePointer)
        fmt.Printf("Content of %p is %d\n",bytePointer,*bytePointer)
        cleanOutput()
}

func cleanOutput(){
    for i:=0;i<3;i++{
        fmt.Println()
    }
}
