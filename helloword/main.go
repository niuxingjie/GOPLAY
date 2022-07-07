package main

import (
    "fmt"
)

// var type func 三个关键字貌似都与定义有关
type  Phone interface {
    call()  // TODO：这里用了小括号？
}

type NokiaPhone struct {
}

// TODO：func不仅仅是定义函数的？还可以这样用？
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am IPhone, I can call you!")
}

func main(){

    // TODO：var 使用自定义的数据类型？
    var phone Phone

    phone = new(NokiaPhone)  // TODO：new关键字啥用？
    phone.call()

    phone = new(IPhone)
    phone.call()

}