package main

import "fmt"

func main(){
    var i, j int

    for i=0; i<10; i++{
        for j=0; j<i; j++{
            if i == 6 {
                goto do_something
            }
            fmt.Printf("%d-%d\n", i, j)
        }
        fmt.Printf("-----------\n")
    }

    return

do_something:
    fmt.Printf("i只能到%d的\n", i)
}
