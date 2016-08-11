package main

import "fmt"

func zero(iPtr *int){
    *iPtr=0
}

func main() {
	fmt.Println("POINTERS")
    x:=5
    zero(&x)
    fmt.Println(x)
}
