/*********************
    AUTHOR: MIQUEL
*********************/

package main

import "fmt"

func main() {
    separeOutput()
    var mySlice []int
    mySlice = make([]int,5)

    for i:=0; i<len(mySlice);i++{
        mySlice[i] = i
    }
    fmt.Println("mySlice: ",mySlice)
    fmt.Println("mySlice[3:5]",mySlice[3:5])
    fmt.Println("Also:")
    newSlice:=[]int{10,20,30}
    fmt.Println("newSlice: ",newSlice)
    fmt.Println("newSlice[2]: ",newSlice[2])

    separeOutput()
}

func separeOutput(){

    for i:=0;i<3;i++{
        fmt.Println()
    }
}
