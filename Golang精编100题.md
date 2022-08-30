
# Golang精编100题

[Golang精编100题](https://www.jianshu.com/p/f690203ff168)

通过练习题，反复学习[Golang-100-Days笔记](https://github.com/niuxingjie/Golang-100-Days)

## 能力模型

- 初级 primary：
  - 熟悉基本语法，能够看懂代码的意图；
  - 在他人指导下能够完成用户故事的开发，编写的代码符合CleanCode规范；
- 中级 intermediate：
  - 能够独立完成用户故事的开发和测试；
  - 能够嗅出代码的坏味道，并知道如何重构达成目标；
- 高级 senior：
  - 能够开发出高质量高性能的代码；
  - 能够熟练使用高级特性，开发编程框架或测试框架；
  

## 选择题

- 初级
```text
1. [primary] 下面属于关键字的是（）
A. func
B. def
C. struct
D. class

答案：AC
解析：注意区分“关键字”与内建“标识符”


2. [primary] 定义一个包内全局字符串变量，下面语法正确的是 （）
A. var str string
B. str := ""
C. str = ""
D. var str = ""

答案：AD
解析：
    - str := ""只能用于函数或方法内部
    - var str = "" 简写，go编译时，可以自动识别类型并赋给str

3. [primary] 通过指针变量 p 访问其成员变量 name，下面语法正确的是（）
A. p.name
B. (*p).name
C. (&p).name
D. p->name

答案：AB
解析:
    - go中指针变量可以直接获取指针指向的成员时，*号可以不写
    - 或许是因为，指针变量p和*p用法单一，基本等同，于是编译器增加处理此简写能力。

4. [primary] 关于接口和类的说法，下面说法正确的是（）
A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理
C. 类实现接口时，需要导入接口所在的包
D. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口

答案：ABD
解析：
疑问：TODO？不懂


5. [primary] 关于字符串连接，下面语法正确的是（）
A. str := ‘abc’ + ‘123’
B. str := "abc" + "123"
C. str ：= '123' + "abc"
D. fmt.Sprintf("abc%d", 123)

答案：BD
解析：
    - ‘’单引号：表示byte类型或rune类型；byte用来强调数据是raw data；rune用来表示Unicode的code point；
    - “”双引号才是真正字符串
    - ``反引号(Back quote)表示字符串字面量，但不支持任何转义序列。
    - 题目中貌似：只有字符串才能进行+运算


6. [primary] 关于协程，下面说法正确是（）
A. 协程和线程都可以实现程序的并发执行
B. 线程比协程更轻量级
C. 协程不存在死锁问题
D. 通过channel来进行协程间的通信

答案：AD


7. [intermediate] 关于init函数，下面说法正确的是（）
A. 一个包中，可以包含多个init函数
B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数
C. main包中，不能有init函数
D. init函数可以被其他函数调用

答案：AB
解析：
    - 注意A

8. [primary] 关于循环语句，下面说法正确的有（）
A. 循环语句既支持for关键字，也支持while和do-while
B. 关键字for的基本使用方法与C/C++中没有任何差异
C. for循环支持continue和break来控制循环，但是它提供了一个更高级的break，可以选择中断哪一个循环
D. for循环不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多
个变量

答案：CD
解析：
    - go 多种形式：
        - for init; condition; post { }
        - for condition { }
        - for { }
        - for key, value := range oldMap { newMap[key] = value }


9. [intermediate] 对于函数定义：
func add(args ...int) int {
        sum := 0
        for _, arg := range args {
            sum += arg
        }
        return sum
}

下面对add函数调用正确的是（）
A. add(1, 2)
B. add(1, 3, 7)
C. add([]int{1, 2})
D. add([]int{1, 3, 7}...)

答案：ABD
解析：
    - 不定长参数的定义与解析：类似python中 *args
        - ...int 不定长的整数
        - []int{1, 3, 7}... 将数组解包成多个整数值


10. [primary] 关于类型转化，下面语法正确的是（）
A.
type MyInt int
var i int = 1
var j MyInt = i

B. 
type MyInt int
var i int = 1
var j MyInt = (MyInt)i

C. 
type MyInt int
var i int = 1
var j MyInt = MyInt(i)

D. 
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)


解析：
    - type 类型名 Type ：使用type，可以定义新类型
    - type 别名 = Type ：TypeAlias 只是 Type 的别名，本质上 TypeAlias 与 Type 是同一个类型。
    - 类型转换语法：valueOfTypeB = typeB(valueOfTypeA)
疑问：
    - 结构体struct interface 都是类型，还可以自定义类型 TODO:go中的类型到底是什么？

类型断言：
    - 用法一：
        package main

        import "fmt"

        func main() {
            var a interface{} =10
            switch a.(type){
            case int:
                    fmt.Println("int")
            case float32:
                    fmt.Println("string")
            }
        }

    - 用法二：
        package main

        import "fmt"

        func main() {
            var a interface{} =10
            t,ok:= a.(int)
            if ok{
                fmt.Println("int",t)
            }
            t2,ok:= a.(float32)
            if ok{
                fmt.Println("float32",t2)
            }
        }


11. [primary] 关于局部变量的初始化，下面正确的使用方式是（）
A. var i int = 10
B. var i = 10
C. i := 10
D. i = 10

答案：ABC


[primary] 关于const常量定义，下面正确的使用方式是（）
A.
const Pi float64 = 3.14159265358979323846
const zero = 0.0

B.
const (
        size int64 = 1024
        eof = -1 
)
C.
const (
        ERR_ELEM_EXIST error = errors.New("element already exists")
        ERR_ELEM_NT_EXIST error = errors.New("element not exists")
)
D. 
const u, v float32 = 0, 3 
const a, b, c = 3, 4, "foo"

答案：ABD
解析：
    - 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
    - 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型


13. [primary] 关于布尔变量b的赋值，下面【错误】的用法是（）
A. b = true
B. b = 1
C. b = bool(1)
D. b = (1 == 2)

答案：BC
解析：
    - bool是关键字，不是标识符，不是函数，
    - go不支持隐形类型转换


14. [intermediate] 下面的程序的运行结果是（）
func main() {  
        if (true) {
            defer fmt.Printf("1")
        } else {
            defer fmt.Printf("2")
        }
        fmt.Printf("3")
}

A. 321
B. 32
C. 31
D. 13

答案：C
解析：
    - defer语句执行顺序，先进后执行
    - defer语句内变量是快照还是实时的？:答案:快照
        package main

        import "fmt"

        func main() {
            var a int
            fmt.Println("开始")
            defer fmt.Println(a)
            a++
            defer fmt.Println(a)
            a++
            defer fmt.Println(a)
            a++
            fmt.Println("结束")
        }


        开始
        结束
        2
        1
        0


15. [primary] 关于switch语句，下面说法正确的有（）
A. 条件表达式必须为常量或者整数
B. 单个case中，可以出现多个结果选项
C. 需要用break来明确退出一个case
D. 只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case


答案：BD
解析：
    - 1. case后的常量值不能重复
    - 2. case后可以有多个常量值
    - 3. fallthrough应该是某个case的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误。
    - switch 两种匹配：
        - 常量匹配：switch marks:=5;marks {
        - 条件匹配：switch { case grade == "A" :

16. [intermediate] golang中没有隐藏的this指针，这句话的含义是（）
A. 方法施加的对象显式传递，没有被隐藏起来
B. golang沿袭了传统面向对象编程中的诸多概念，比如继承、虚函数和构造函数
C. golang的面向对象表达更直观，对于面向过程只是换了一种语法形式来表达
D. 方法施加的对象不需要非得是指针，也不用非得叫this

答案：ACD
解析：
    - 方法的本质是函数，请参考：geek 教程

17. [intermediate] golang中的引用类型包括（）
A. 数组切片
B. map
C. channel
D. interface

答案：ABCD
解析:
    - 值类型：值类型分别有：int系列、float系列、bool、string、数组和结构体
    - 引用类型：指针、slice切片、管道channel、接口interface、map、函数等

18. [intermediate] golang中的指针运算包括（）
A. 可以对指针进行自增或自减运算
B. 可以通过“&”取指针的地址
C. 可以通过“*”取指针指向的数据
D. 可以对指针进行下标运算

答案：BC
解析：
    - p.name是获取指针变量*p的成员的简写方式


19. [primary] 关于main函数（可执行程序的执行起点），下面说法正确的是（）
A. main函数不能带参数
B. main函数不能定义返回值
C. main函数所在的包必须为main包
D. main函数中可以使用flag包来获取和解析命令行参数

答案：ABCD
```

- 中级
```
20. [intermediate] 下面赋值正确的是（）
A. var x = nil
B. var x interface{} = nil
C. var x string = nil
D. var x error = nil

答案：BD
解析：
    - 不同类型的变量，go会赋予默认的初始值（默认O值）
        类型	默认值
        整数	0
        浮动	0
        复数	0 个实部和 0 个虚部
        字节	0
        符文	0
        字串	“”
        布尔	错误
        数组	每个数组的值都为其默认值
        结构	每个字段均为默认值
        地图	无（nil）	
        频道	无（nil）	
        界面	无（nil）	
        切片	无（nil）	
        指针    无（nil）	
    - 如何定义一个值为nil的变量？


21. [intermediate] 关于整型切片的初始化，下面正确的是（）
A. s := make([]int)
B. s := make([]int, 0)
C. s := make([]int, 5, 10)
D. s := []int{1, 2, 3, 4, 5}

答案：BCD
解析：
    - 定义切片：var identifier []type 切片不需要说明长度
    - go语言make的用法有：
        - 初始化映射：make(map[string]string) 这种用法只能用在类型为map或chan的场景
        - 初始化切片：make([]int, 2) 一个长度为2的slice
        - make([]int, 2, 4)  切片a的最大长度为4（容量），当前只初始化2个。



```

## 填空题





## 判断题


