/*********************
    AUTHOR: MIQUEL
*********************/

package main

import "fmt"

type Gopher struct{
    name string
    id int
}

func newGopher(name string, id int)(*Gopher){
    g:= new(Gopher) // g:=Gopher{name,id}
    g.name = name   // return &g
    g.id = id
    return g
}

func (g *Gopher)sayHi(){
    fmt.Println("\n* Gopher ",g.name," says: hi! :) *\n")
}

func main(){

    kevin := newGopher("Kevin",1)
    kevin.sayHi()

}
