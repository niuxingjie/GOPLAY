## W3Cschool

[Go 语言教程](https://www.w3cschool.cn/go/go-tutorial.html)
[Go online 菜鸟工具](https://c.runoob.com/compile/21/)

## Go Hello World 实例

- code
```go
package main /* 包声明 */

import "fmt" /* 引入包 */

/* 函数 */
func main (){
    // 语句
    fmt.PrintLn("Hello World!")
}

```
- 说明
```text
- 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包 TODO:难道包名和文件名不一样？
- 每个 Go 应用程序都包含一个名为 main 的包

- 下一行 import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素

- 下一行 func main() 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的
- main一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。

-  大写开头的标识符：
    - 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）

- 小写开头的标识符：
    - 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。

```


## 基础语法


- 使用 + 进行字符串连接
```go
package main

import "fmt"

func main(){
    fmt.PrintLn("Hello" + "World!")
}

```

- 关键字: 顺序、分支、循环；函数、对象、类、包
```text
break	default	func	interface	select
case	defer	go	map	struct
chan	else	goto	package	switch
const	fallthrough	if	range	type
continue	for	import	return	var

```

- 预定义标识符: 数据类型
```text
append	bool	byte	cap	close	complex	complex64	complex128	uint16
copy	false	float32	float64	imag	int	int8	int16	uint32
int32	int64	iota	len	make	new	nil	panic	uint64
print	println	real	recover	string	true	uint	uint8	uintptr

```

- 运算符： 数学计算、逻辑运算
```text


```

- 分隔符
```text
空格 [] () {}
```

- 标点符号
```text
. : ...

```


## 数据类型


- 布尔型
```go
var b bool = true  // bool 是关键字， true，false 小写开头

```

- 数字类型
```go

整形：
    uint8  无符号 8 位整型 (0 到 255)  2*2*2*2*2*2*2*2
    uint16
    uint32
    uint64

    int8 有符号 8 位整型 (-128 到 127)
    int16
    int32
    int64

浮点型：
    float32 IEEE-754 32位浮点型数
    float64

    complex64 32 位实数和虚数
    complex128

其他：
    byte 类似 uint8
    rune 类似 uint32
    uint  32 或 64 位
    int 与 uint 一样大小
    uintptr  无符号整型，用于存放一个指针
```


- 字符串类型
```text
string
字符串就是一串固定长度的字符连接起来的字符序列
Go的字符串是由单个字节连接起来的
Go语言的字符串的字节使用UTF-8编码标识Unicode文本。
```

- 派生类型
```
- 指针类型（Pointer）
- 数组类型
- 结构体类型(struct)
- 联合体类型 (union)
- 函数类型
- 切片类型
- 接口类型（interface）
- Map 类型
- Channel 类型
```


## 变量

- 变量声明
```go
// var identifier-标识符 type-类型

// 一种，指定变量类型，声明后若不赋值，使用默认值
var name string
name = "Hello World!"

// 根据值自行判定变量类型。
var name = "Hello World!"


// 省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误。
// //这种不带声明格式的只能在函数体中出现
name := "Hello World!"
```

- code
```go
package main

import "fmt"


var a = "Hello World!"
var b string = "w3school.com"
var c bool // 默认值为false



func main(){
    fmt.PrintLn(a, b, c)
}

```

- 多变量声明
```go
//类型相同多个变量, 非全局变量  TODO: 全局变量，局部变量，变量的作用域是怎样的？
var first_name, last_name, family_name  string
first_name, last_name, family_name = "Json", "Snow", "Stack"

var first_name, last_name, family_name = "Json", "Snow", "Stack"

first_name, last_name, family_name := "Json", "Snow", "Stack"

//类型不同多个变量, 全局变量, 局部变量不能使用这种方式
var (
    name string
    age uint8
)
```


- code
```go
package main

import "fmt"

var x, y int = 1, 2

var (
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "Hello World!"

// 报错:./main.go:9:1: syntax error: non-declaration statement outside function body
// g, h := 123, "Hello World!"


func main(){
    g, h := 123, "Hello World!"
    fmt.Println(a, b, c, d, e, f, g, h)
}

```

- 值类型和引用类型
```text
- 值类型
    - 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
    - 当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝,内存中存了两份
    - 你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。

- 引用类型:指针
    - 一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个值所在的位置。
    - 这个内存地址为称之为指针，这个指针实际上也被存在另外的某一个值中。
    - 当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。  TODO： Python什么是否传的是值，什么是否传的是引用（指针）？
    - 如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。
```

- 声明变量注意
```text
- 使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
- 这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
- 可以函数体内部避免重复声明同一个变量

- 如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误 -- 全局变量是允许声明但不使用
- 如果你在定义变量 a 之前使用它，则会得到编译错误
```

- 并行变量声明与复制
```go

a, b = b, a
——,c = 5, 7  // _ 实际上是一个只写变量，你不能得到它的值

val, err = Func1(var1) // 类似python也是可以的

```


## 常量


- 常量 TODO:什么是常量
```text
- 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
- 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
```

- code
```go
package main

import "fmt"


func main(){
    const LENGTH int = 10
    const WIDTH int = 5
    var area int
    const a, b, c = 1, false, "Hello World!"

    area = LENGTH * WIDTH

    fmt.Printf("面积为: %d", area)
    fmt.Println()
    fmt.Println(a, b, c)
}

```

- code 2
```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
// 枚举：数字 0、1 和 2 分别代表未知性别、女性和男性。
```

- code 3
```go
// 常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不通过：TODO: 内置函数啥意思？
package main

import "unsafe"

const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}

```


- iota，特殊常量:自增长变量  TODO: 奇怪的东西？
```go
// iota，特殊常量，可以认为是一个可以被编译器修改的常量。

const (
    a = iota
    b = iota
    c = iota
)

const (
    a = iota
    b 
    c 
)

package main

import "fmt"

func main (){
    const (
        a = iota  // 0
        b  // 1
        c  // 2
        d = "d" // d
        e  // d
        f  // d
        g = 100  // 100
        h = iota // 7
        i // 8

    )

    fmt.Println(a, b, c, d, e, f, g, h, i)

}


```


## 运算符


- 算术运算符
```text
运算符	描述	实例
+	相加	A + B 输出结果 30
-	相减	A - B 输出结果 -10
*	相乘	A * B 输出结果 200
/	相除	B / A 输出结果 2
%	求余	B % A 输出结果 0
++	自增	A++ 输出结果 11
--	自减	A-- 输出结果 9
```

- 关系运算符-比较运算符
```text
==	检查两个值是否相等，如果相等返回 True 否则返回 False。	(A == B) 为 False
!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。	(A != B) 为 True
>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。	(A > B) 为 False
<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。	(A < B) 为 True
>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。
```

- 逻辑运算符 
```text
&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。	(A && B) 为 False
||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	(A || B) 为 True
!	逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。	!(A && B) 为 True
```

- 位运算符
```text
&：与运算，全真为真； 
|：或运算，全假才假；  
^:异或运算，相同为假，不同为真  

<<	左移运算符"<<"是双目运算符。
>>	右移运算符">>"是双目运算符。
```

- 赋值运算符
```text
运算符	描述	实例
=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
+=	相加后再赋值	C += A 等于 C = C + A
-=	相减后再赋值	C -= A 等于 C = C - A
*=	相乘后再赋值	C *= A 等于 C = C * A
/=	相除后再赋值	C /= A 等于 C = C / A
%=	求余后再赋值	C %= A 等于 C = C % A
<<=	左移后赋值	C <<= 2 等于 C = C << 2
>>=	右移后赋值	C >>= 2 等于 C = C >> 2
&=	按位与后赋值	C &= 2 等于 C = C & 2
^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
|=	按位或后赋值	C |= 2 等于 C = C | 2
```

- 其他运算符
```
&	返回变量存储地址	&a; 将给出变量的实际地址。
*	指针变量。	*a; 是一个指针变量
```

- code
```go
package main
import "fmt"
func main(){
    var a int  = 4
    var b int32
    var c float32
    var pointer *int


    fmt.Printf("第一行 - a 变量的数据类型为 =  %T\n", a)  // 第一行 - a 变量的数据类型为 =  int
    fmt.Printf("第二行 - b 变量的数据类型为 =  %T\n", b)  // 第二行 - b 变量的数据类型为 =  int32
    fmt.Printf("第三行 - c 变量的数据类型为 =  %T\n", c)  // 第三行 - c 变量的数据类型为 =  float32

    pointer = &a
    fmt.Printf("a的值为 ：%d\n", a)  // a的值为 ：4
    fmt.Printf("&a的值为 ：%d\n", &a)  // &a的值为 ：824633811080

    fmt.Printf("pointer的值为 ：%d\n", pointer)  // pointer的值为 ：824633811080
    fmt.Printf("*pointer的值为 ：%d\n", *pointer)  // *pointer的值为 ：4
    fmt.Printf("&pointer的值为 ：%d\n", &pointer)  // &pointer的值为 ：824633778216

    a = 100
    fmt.Printf("a的值为 ：%d\n", a)  // a的值为 ：100
    fmt.Printf("&a的值为 ：%d\n", &a)  // &a的值为 ：824633811080

    fmt.Printf("pointer的值为 ：%d\n", pointer)  //pointer的值为 ：824633811080
    fmt.Printf("*pointer的值为 ：%d\n", *pointer)  // *pointer的值为 ：100
    fmt.Printf("&pointer的值为 ：%d\n", &pointer)  // &pointer的值为 ：824633778216
}
```


## 条件语句


- if
```go
package main

import "fmt"

func main(){
    var a int = 10

    if a < 20 {
        fmt.Printf("a的值是：%d小于20", a)
        fmt.Println()
    }
    fmt.Printf("a的值是: %d", a)
    fmt.Println()
}

```
- if else
```go
package main

import "fmt"

func main(){
    var a int = 30

    if a < 20 {
        fmt.Printf("a的值是：%d小于20", a)
        fmt.Println()
    } else {
        fmt.Printf("a的值是：%d大于等于20", a)
        fmt.Println()
    }
    fmt.Printf("a的值是: %d", a)
    fmt.Println()
}

```
- if 嵌套
```go
package main

import "fmt"

func main(){
    var a int = 50

    if a < 20 {
        fmt.Printf("a的值是：%d小于20", a)
        fmt.Println()
    } else {
        fmt.Printf("a的值是：%d大于等于20", a)
        fmt.Println()
        if a >= 30 {
            fmt.Printf("a的值是：%d大于等于30", a)
            fmt.Println()
        }
    }
    fmt.Printf("a的值是: %d", a)
    fmt.Println()
}
```
- switch
```go
// switch 语句用于基于不同条件执行不同动作。
package main

import "fmt"


func main(){
    var grade string = "B"
    var marks int = 90

    // 第一种：变量匹配
    switch marks {
        case 90:
            grade = "A"
        case 80: 
            grade = "B"
        case 70, 60, 50:
            grade = "c"
        default:
            grade = "D"
    }
    fmt.Printf("grade: %s", grade)
    fmt.Println()

    // 第二种：条件匹配
    switch {
        case grade == "A":
            fmt.Printf("grade: %s-优秀", grade)
        case grade == "B", grade == "c":  // 或
            fmt.Printf("grade: %s-良好", grade)
        case grade == "D":
            fmt.Printf("grade: %s-及格", grade)
        case grade == "E":
            fmt.Printf("grade: %s-不及格", grade)
        default:
            fmt.Printf("grade: %s-差", grade)
    }
    fmt.Println()
}

```

- [Type Switch](https://www.w3cschool.cn/go/go-switch-statement.html)
```go
// switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。TODO: interfaces是个啥？



```

- [select](https://www.w3cschool.cn/go/go-select-statement.html)
```text
- select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
- select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收
- 这里的通信，可以简单的理解为IO（输入输出）  TODO: channel是个啥东西？

select {
    case communication clause  :
       statement(s)      
    case communication clause  :
       statement(s)
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s)
}

- 每个case都必须是一个通信
- 所有channel表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，它就执行；其他被忽略。
- 如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
- 如果有default子句，则执行该语句。
- 如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
```


## 循环语句

- [for 循环](w3cschool.cn/go/go-for-loop.html)
```go
// 三种for循环：for,while,?
package main

import "fmt"

func main(){
    var b int = 15
    var a int

    numbers := [6]int{1, 2, 3, 5}  // TODO:数组类型？长度为6，剩余的两个元素为默认值0

    // 和 C 语言的 for 一样：
    for a :=0; a < 10; a++ {
        fmt.Printf("a的值为： %d\n", a)
    }

    // 和 C 的 while 一样：
    for a < b {
        a++
        fmt.Printf("a的值为： %d\n", a)
    }

    // 和 C 的 for(;;) 一样：
    for i,x:= range numbers {  // TODO:range是个啥？内置函数？
        fmt.Printf("第%d位的值为： %d\n", i, x)
    }
}

```

- for循环嵌套
```go
package main

import "fmt"

func main(){
    /* 定义局部变量 */
    var i, j int

    for i = 2; i < 100; i++{
        for j = 2; j <= (i/j); j++ {
            if i%j == 0 {
                break  // 如果发现因子，则不是素数
            }
        }
        if j > (i / j) {
            fmt.Printf("j的值是： %d\n", i)
        }
    }
}

```

- 循环控制语句
```text
- break： 经常用于中断当前 for 循环或跳出 switch 语句(单层)
- continue： 跳过当前循环的剩余语句，然后继续进行下一轮循环。
- goto： 
    - 语句可以无条件地转移到过程中指定的行
    - goto 退出多层循环
```

- goto 退出多层循环
```go
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

do_something:  // TODO：标签啥意思？语法奇怪？
    fmt.Printf("i只能到%d的\n", i)
}
```
- 使用 goto 集中处理错误(异常)
```go


```

- 无限循环
```go
package main

import "fmt"

func main() {
    for true  {
        fmt.Printf("这是无限循环。\n");
    }
}

```


## 函数


- 基本
```text
- Go 语言最少有1个 main() 函数。
- 函数声明告诉了编译器函数的名称，返回类型和参数。
- Go 语言标准库提供了多种可动用的内置的函数
    - len
    - 
```

- 函数定义与调用
```go
package main

import "fmt"

func main(){
    var num1 int = 100
    var num2 int = 200
    var ret int

    ret = max(num1, num2)  // 下面定义的max

    fmt.Printf("最大值为：%d\n", ret)
}

// 关键字func, 定义参数类型外，也需要定义返回值的类型int
func max (num1, num2 int) int {
    var result int

    if num1 >= num2 {
        result = num1
    } else {
         result = num2
    }
    return result
}

```

- 函数返回多个值
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Mahesh", "Kumar")  // := 可以在函数体内方便接收任意类型参数
   fmt.Println(a, b)
}

```

- 函数传参：值传递和引用传递
```go
// 值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
package main

import "fmt"

func main(){
    var x int = 100
    var y int = 200

    fmt.Println("swap(x,y)前,main函数中：")
    fmt.Printf("x=%d\n", x)
    fmt.Printf("y=%d\n", y)

    swap(x, y)

    fmt.Println("swap(x,y)后,main函数中：")
    fmt.Printf("x=%d\n", x)
    fmt.Printf("y=%d\n", y)
}

func swap(x, y int){
    var temp int

    temp = x
    x = y
    y = temp

    fmt.Println("swap(x,y)函数内部中交换后的x和y：")
    fmt.Printf("x=%d\n", x)
    fmt.Printf("y=%d\n", y)
}
```


- 函数传参：引用传递
```go
// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
// 无论python还是go，函数传参，本质上和=赋值操作是一样的，实参是否会改变取决于，实参是可变对象还是不可变对象
package main

import "fmt"

func main(){
    var a int = 100
    var b int = 200

    fmt.Println("swap(&a,&b)前,main函数中：")
    fmt.Printf("地址：&a=%d  指针指向的变量a的值：a=%d\n", &a, a)
    fmt.Printf("地址：&b=%d  指针指向的变量b的值：b=%d\n", &b, b)

    swap(&a, &b)  // 关键所在：传的是变量的地址，而不是值

    fmt.Println("swap(&a,&b)后,main函数中：")
    fmt.Printf("地址：&a=%d  指针指向的变量a的值：a=%d\n", &a, a)
    fmt.Printf("地址：&b=%d  指针指向的变量b的值：b=%d\n", &b, b)

    fmt.Println("swap(&a,&b)后,main函数中：ab的地址没有变，但是值变了")
}

// 定义函数的参数是指向int类型的指针
func swap(x *int, y *int){
    var temp int

    temp = *x  // *x是拿到指针指向数据的值
    *x = *y  // *x是拿到指针指向数据的值，然后赋值给y指向的数据
    *y = temp

    fmt.Println("swap(x,y)函数内部中交换后的x和y：")
    fmt.Printf("地址：x=%d  指针指向的变量a的值：*x=%d\n", x, *x)
    fmt.Printf("地址：y=%d  指针指向的变量b的值：*x=%d\n", y, *y)
}

```

- defer语句
```go
// defer语句会将其后面跟随的语句进行延迟处理
// 在defer所属的函数即将返回时，才会将延迟处理的语句按照defer定义的顺序逆序执行，即先进后出

package main

import "fmt"

func main() {
	fmt.Println("开始")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("结束")
}



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
```


## 变量作用域

- 局部变量： 它们的作用域只在函数体内

- 全局变量：可以在整个包甚至外部包（被导出后）使用
```go
package main

import "fmt"

/* 声明全局变量：函数内部可读写的全局变量 */
var g int = 100

func main() {

    /* 声明局部变量 */
    var a, b int

    print_g()  // 100

    /* 初始化参数 */
    a = 10
    b = 20
    g = a + b  // 30

    fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)

    print_g()  // 30 改变了全局变量的值，与python中不同：全局变量只读，除非global 关键词，才是可读写全局变量
}

func print_g(){
    fmt.Printf("结果： g = %d\n", g)
}
```
- 局部变量：函数内声明了与全局变量相同的参数时，覆盖全局变量。
```go
// Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑

package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g)
}
```

- 形式参数: 形式参数会作为函数的局部变量来使用

- 先声明后使用的参数默认值：初始化局部和全局变量
```text
数据类型	初始化默认值
int	          0
float32	      0
pointer	     nil  // TODO: 预定义的数据类型？
```


## 数组


- 基本结构
```text
- 数组是具有【相同唯一】类型的一组已编号且长度固定的数据项序列，
- 元素的类型可以是任意的原始类型例如整型、字符串或者自定义类型。
- Go 数组的长度不可改变
```

- 声明数组
```go
// 数组名 长度 类型
var balance [10] float32

// 初始化数组中 {} 中的元素个数不能大于 [] 中的数字,但是可以小于。
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

```

- code
```go
package main

import "fmt"

func main(){
    var numbers [10]int
    var i int // for中使用的i需要预定义，否则./main.go:8:9: undefined: i 

    for i=0; i<10; i++{
        numbers[i] = i + 100
    }

    for i=0; i<10; i++{
        fmt.Printf("numbers[%d]=%d\n", i, numbers[i])
    }
}
```

- 多维数据


- 向函数传递数组：
```text
- 请保证签名中形参与实参类型注解完全一致
- 
```

- code1
```go
package main

import "fmt"

func main(){
    var numbers = [5]int{1, 2, 3, 4, 5}  // get_average中形参类型注解需与这里完全一致，包括size的有无
    var average float32

    average = get_average(numbers, 5)
    fmt.Printf("numbers的平均值为:%f\n", average)
}


// func get_average(array []int, size int) float32{
// 不指定数组长度会报错：cannot use numbers (variable of type [5]int) as type []int in argument to get_average
func get_average(array [5]int, size int) float32{
    var i int
    var sum int = 0
    var average float32

    for i=0; i<size; i++{
        sum += array[i]
    }

    average = float32(sum) / float32(size)
    return average

}
```
- code2
```go
package main

import "fmt"

func main(){
    var numbers = []int{1, 2, 3, 4, 5}
    var average float32

    average = get_average(numbers, 5)
    fmt.Printf("numbers的平均值为:%f\n", average)
}


// func get_average(array [5]int, size int) float32{
// 不指定数组长度会报错：cannot use numbers (variable of type []int) as type [5]int in argument to get_average
func get_average(array []int, size int) float32{
    var i int
    var sum int = 0
    var average float32

    for i=0; i<size; i++{
        sum += array[i]
    }

    average = float32(sum) / float32(size)
    return average
}

```

- code3: 传参为可变对象时，内部可以变
```go
// 无论python还是go，函数传参，本质上和=赋值操作是一样的，实参是否会改变取决于，实参是可变对象还是不可变对象
package main

import "fmt"

func main(){
    var numbers = []int{1, 2, 3, 4, 5, 6}
    var i int
    var numbers_copy = numbers
    // 改变 numbers[0]、numbers[1]
    change_0(numbers)

    // 改变 numbers[2]、numbers[3]
    numbers_copy[2] = 0
    numbers_copy[3] = 0

    for i=0; i<6; i++{
        fmt.Printf("numbers[%d]的平均值为:%d\n", i, numbers[i])
    }
    
}

func change_0(array []int) {
    array[0] = 0
    array[1] = 0
}
```


## 指针

- 什么是指针
```text

- 取内存地址的符号是 &

- 一个指针变量，存储的是一个内存地址，定义时，需要定义存储的内存地址里存储的数据类型
- 一个指针变量，可以用*ptr 直接获取，内存地址里的值
- 非常方便的操作，内存地址和内存地址里存储的值
```


- 空指针
```text
- 当一个指针被定义后没有分配到任何变量时，它的值为 nil。
- nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值 TODO: nil == nil ?

```

- [指针数组](https://www.w3cschool.cn/go/go-array-of-pointers.html)
```go
package main

import "fmt"

const MAX int = 3

func main(){
    var numbers = [MAX]int{1, 2, 3}
    var ptr [MAX]*int
    var i int

    for i=0; i<MAX; i++{
        ptr[i] = &numbers[i]
    }

    for i=0; i<MAX; i++{
        fmt.Printf("第%d个数的值为%d\n", i, *ptr[i])
    }
}
```

- [指向指针的指针](https://www.w3cschool.cn/go/go-pointer-to-pointer.html)
```go
// 指向指针的指针变量值需要使用两个 * 号,定义和取值

package main

import "fmt"

func main(){
    var a int
    var ptr *int
    var pptr **int

    a = 3000
    ptr = &a
    pptr = &ptr

    /* 获取 pptr 的值 */
    fmt.Printf("变量 a = %d   &a=%d\n", a, &a)
    fmt.Printf("指针变量 *ptr = %d   ptr=%d   &ptr=%d\n", *ptr, ptr, &ptr)
    fmt.Printf("指向指针的指针变量 **pptr = %d   pptr=%d\n", **pptr, pptr)
}

```

- [指针作为函数参数](https://www.w3cschool.cn/go/go-passing-pointers-to-functions.html)
```go


```


## 结构体

- 基本信息
```text
- 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合

```

- 定义结构体
```text
- type 是定义结构体的关键字
- struct 是结构体的数据类型标识符

- 定义后的结构 --> 新的数据类型 --> 用于定义变量
```

- 结构体作为函数参数
```go
// 点号可以用于访问结构体成员
package main

import "fmt"


type Book struct {
    title string  // 结构体的成员定义：名字 类型
    author string
    subject string
    book_id int
}

func main(){
    var book1, book2 Book

    book1.title = "go book"
    book1.author = "w3cschool"
    book1.subject = "master go"
    book1.book_id = 1

    book2.title = "Python book"
    book2.author = "w3cschool"
    book2.subject = "master python"
    book2.book_id = 2

    /* 打印 Book1 信息 */
    fmt.Printf( "Book 1 title : %s\n", book1.title)
    fmt.Printf( "Book 1 author : %s\n", book1.author)
    fmt.Printf( "Book 1 subject : %s\n", book1.subject)
    fmt.Printf( "Book 1 book_id : %d\n", book1.book_id)

    fmt.Println()
    
    /* 打印 Book2 信息 */
    print_book(book2)
}


func print_book(book *Book){
    fmt.Printf( "Book  title : %s\n", book.title)    // TODO:奇怪，这里不该是*book拿到结构体吗？
    fmt.Printf( "Book  author : %s\n", book.author)
    fmt.Printf( "Book  subject : %s\n", book.subject)
    fmt.Printf( "Book  book_id : %d\n", book.book_id)
}
```


- 结构体指针: 指向结构体的指针
```go
package main

import "fmt"


type Book struct {
    title string  // 结构体的成员定义：名字 类型
    author string
    subject string
    book_id int
}

func main(){
    var book1, book2 Book  

    book1.title = "go book"   // TODO：使用结构体定义和赋值，就实现了类-->对象
    book1.author = "w3cschool"
    book1.subject = "master go"
    book1.book_id = 1

    book2.title = "Python book"
    book2.author = "w3cschool"
    book2.subject = "master python"
    book2.book_id = 2

    /* 打印 Book1 信息 */
    fmt.Printf( "Book 1 title : %s\n", book1.title)
    fmt.Printf( "Book 1 author : %s\n", book1.author)
    fmt.Printf( "Book 1 subject : %s\n", book1.subject)
    fmt.Printf( "Book 1 book_id : %d\n", book1.book_id)

    fmt.Println()
    
    /* 打印 Book2 信息 */
    print_book(&book2)
}


func print_book(book *Book){
    fmt.Printf( "Book  title : %s\n", *book.title)  // 奇怪，TODO：这里不该是*book拿到结构体吗？
    fmt.Printf( "Book  author : %s\n", *book.author)
    fmt.Printf( "Book  subject : %s\n", *book.subject)
    fmt.Printf( "Book  book_id : %d\n", *book.book_id)
}

```

## 切片(Slice)


- 基本信息
```text
- Go 数组的长度不可改变
- 切片(Slice):动态数组,长度可变
```

- 定义切片
```go
// TODO:不定义长度额数组就是切片？
var s [] int  
s :=[] int {1,2,3} 

// make([]type, length, capacity) 创建切片：capacity参数可选，为切片最大容量，length为初始长度
```

- len()求长度，cap()求最大容量
```go
package main

import "fmt"

func main(){
    // var numbers = []int{1, 2, 3}
    numbers := make([]int, 3, 5)
    
    printSlice(numbers)
}

func printSlice(nums []int){
    var i int
    fmt.Printf("len=%d cap=%d nums=%v\n", len(nums), cap(nums), nums)

    for i=0; i<len(nums); i++{
        fmt.Printf("第%d个元素是：%d\n", i, nums[i])
    }
}
```

- 空(nil)切片：定义但是未初始化
```go
// 一个切片在未初始化之前默认为 nil，长度为 0
package main

import "fmt"

func main(){
    var numbers[]int
    
    printSlice(numbers)
}

func printSlice(nums []int){
    if nums == nil {
        fmt.Println("空切片")
        fmt.Printf("len=%d cap=%d nums=%v\n", len(nums), cap(nums), nums)  // len=0 cap=0 nums=[]
    } else {
        fmt.Printf("len=%d cap=%d nums=%v\n", len(nums), cap(nums), nums)
    }
}

```

- 切片截取: 
```go
// slick[:] 类似python的切片
```

- append() 和 copy() 函数
```go
package main

import "fmt"

func main(){
    var numbers []int
    printSlice(numbers)

    // 空切片添加一个0
    // append(numbers, 0)   // TODO：为啥会报错？：append(numbers, 0) (value of type []int) is not used
    numbers = append(numbers, 0)
    printSlice(numbers)

    // 空切片添加多个
    numbers = append(numbers, 1, 2, 3, 4, 5)
    printSlice(numbers)

    var double = make([]int, len(numbers), cap(numbers) * 2)
    fmt.Printf("len=%d cap=%d double=%v\n", len(double), cap(double), double)

    copy(double, numbers)  // TODO: 没有重新赋值,为啥不报错？
    fmt.Printf("len=%d cap=%d double=%v\n", len(double), cap(double), double)  // len=6 cap=12 double=[0 1 2 3 4 5]

    copy(double, numbers)  // TODO: 为啥没有再增加6个值？并不类似python：list.extend()
    fmt.Printf("len=%d cap=%d double=%v\n", len(double), cap(double), double)  // len=6 cap=12 double=[0 1 2 3 4 5]
}

func printSlice(nums []int){
    if nums == nil {
        fmt.Println("空切片")
    }
    fmt.Printf("len=%d cap=%d nums=%v\n", len(nums), cap(nums), nums)
}
```


## 范围(Range)


- 基本信息
```text
- range 是关键字
- 用于for循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
- range：
    - 后面跟：数组(array)、切片(slice)、通道(channel)或集合(map) 
    - 返回：两个值
```

- code
```go
package main

import "fmt"

func main(){
    var numbers = [5]int{1, 2, 3, 4, 5}
    var sum int

    for _, num := range numbers {
        sum += num
    }
    fmt.Printf("len=%d cap=%d sum=%v\n", len(numbers), cap(numbers), sum)

    for i, num := range numbers {
        fmt.Printf("第%d个元素：%d", i, num)
    }
    
}

```

## Map(集合)


- 基本信息
```text
- 无序的键值对的集合
- 依靠key取值


```

- 定义 map
```go
// 如果不初始化 map，那么就会创建一个 nil map
// 可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable = make(map[key_data_type]value_data_type)

```
- cdoe
```go
package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string
    
    // nil map 不能用来存放键值对
    if countryCapitalMap == nil {
        fmt.Println("a nil map\n")
    }
    // countryCapitalMap["China"] = "Beijing"  // panic: assignment to entry in nil map

    // 初始化map
    countryCapitalMap = make(map[string]string)

    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"

    /* 使用 key 输出 map 值 */
    for country := range countryCapitalMap {
        fmt.Printf("key is %s, value is %s.\n", country, countryCapitalMap[country])
    }

    // capital, ok = countryCapitalMap["China"]  // undefined: capital
    /* 查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap["China"]  // :=方便使用未定义变量接收返回值
    /* 如果 ok 是 true, 则存在，否则不存在 */
    if (ok) {
        fmt.Printf("key is China, value is %s.\n", capital)
    } else {
        fmt.Println("Capital of China is not present") 
    }
}
```

- delete() 函数
```go
package main

import "fmt"

func main() {
    // 变量名 := 类型 初始值
    countryCapitalMap := map[string]string {"France":"Paris", "Italy":"Rome","Japan":"Tokyo", "India":"New Delhi"}
    
    fmt.Println("打印原始map")
    printMap(countryCapitalMap)

    // 使用内置函数delete通过key删除
    // delete(countryCapitalMap, 'France')  // more than one character in rune literal  TODO: 单引号和双引号含义不一样？
    delete(countryCapitalMap, "France")

    fmt.Println("打印新map")
    printMap(countryCapitalMap)
}

func printMap(country_map map[string]string){
    for country := range country_map {
        fmt.Printf("key is %s, value is %s.\n", country, country_map[country])
    }
}
```

## 递归函数


- 阶乘
```go
package main 

import "fmt"

func main(){
    var total int
    var num int = 5
    
    total = jie_cheng(num)

    fmt.Printf("%d 的阶乘是 %d\n", num, total)
}


// func jie_cheng(num int) {   // 函数必须指定返回值类型，否则报错：jie_cheng(num) (no value) used as value
func jie_cheng(num int) int {
    if num != 1 {
        fmt.Printf("%d * ", num)
        return num * jie_cheng(num - 1)
    } else {
        fmt.Printf("1\n")
        return 1
    }
}

```

- 斐波那契数列
```go


```


## 类型转换

- 基本信息
```text
- 将一种数据类型的变量转换为另外一种类型的变量
```

- go不支持隐式转换类型
```go
package main

import "fmt"

// 函数没有定义返回值类型，即使有return 也不会返回
func main(){
    var sum int = 32
    var num int = 5
    var mean float32

    // mean = sum / num  // cannot use sum / num (value of type int) as type float32 in assignment
    mean = float32(sum / num)
    
    fmt.Printf("%d / %d = %f \n", sum, num, mean)
}
```


## 接口【难点】

- 基本信息
```text
- interface 数据类型，内置标识符 TODO：使用场景是啥？
- 所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

- 用例： TODO: 用interface和struct实现了python中类似的class类？
    /* 定义接口 */
    type interface_name interface {
    method_name1() [return_type]
    method_name2() [return_type]
    method_name3() [return_type]
    ...
    method_namen [return_type]
    }

    /* 定义结构体 */
    type struct_name struct {
    /* variables */
    }

    /* 实现接口方法 */
    func (struct_name_variable struct_name) method_name1() [return_type] {
    /* 方法实现 */
    }
    ...
    func (struct_name_variable struct_name) method_namen() [return_type] {
    /* 方法实现*/
    }

```


- code
```go
package main

import (
    "fmt"
)

// var type func 三个关键字貌似都与定义有关
type  Phone interface {
    call()  // TODO：这里用了小括号？
}

// NokiaPhone 结构体一
type NokiaPhone struct {
}

// TODO：func不仅仅是定义函数的？还可以这样用？
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

// IPhone 结构体二
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


```


## 错误处理

- 基本信息
```text
- 内置的错误接口 error

    type error interface {
        Error() string
    }
- 
```

- code 
```go
package main

import (
    "fmt"
)

// 定义了一个机构体
type DivideError struct {
    dividee int
    divider int
}


// 实现了错误接口 TODO：Error() 是语言内置接口？
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, de.dividee)  // TODO：de是指针，指针.属性？
}


func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
    if varDivider == 0 {
        dData := DivideError{
            dividee: varDividee,  // syntax error: unexpected newline, expecting comma or }
            divider: varDivider,
        }
        errorMsg = dData.Error()  // TODO:错误会提前返回错误信息？不然errorMsg怎么返回的？
        return
    } else {
        return varDividee / varDivider, ""
    }
}


func main(){

    // TODO: if后面可以用语句statement，而不是表达式expression？
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
        fmt.Println("100/10= ", result)
    }

    if _, errorMsg := Divide(100, 0); errorMsg != "" {
        fmt.Println("errorMsg is: ", errorMsg)
    }

}

```

## 反射(Reflect)

- 基本概念
```text
- 不知道具体类型的情况下，可以用反射来更新变量值，查看变量类型
```

- TypeOf 和 ValueOf
```go
package main

import (
    "fmt"
    "reflect"
)

func main(){
    var booknum float32 = 6
    var isbook bool = true

    bookauthor := "w3cschool"
    bookdetail := make(map[string]string)
    bookdetail["go教程"]="www.w3cdchool.cn"

    fmt.Println(reflect.TypeOf(booknum), ":", reflect.ValueOf(booknum))
    fmt.Println(reflect.TypeOf(isbook), ":", reflect.ValueOf(isbook))
    fmt.Println(reflect.TypeOf(bookauthor), ":", reflect.ValueOf(bookauthor))
    fmt.Println(reflect.TypeOf(bookdetail), ":", reflect.ValueOf(bookdetail))  // map[string]string : map[go教程:www.w3cdchool.cn]
    fmt.Println(reflect.TypeOf(nil), ":", reflect.ValueOf(nil))  // <nil> : <invalid reflect.Value>
}
```

- 通过反射设置值
```go
/* 使用建议:
1、大量使用反射的代码通常会变得难以理解
2、反射的性能低下，基于反射的代码会比正常的代码运行速度慢一到两个数量级
*/
package main

import (
    "fmt"
    "reflect"
)

// TODO: interface{}这里为什么有大括号？
func reflectsetvalue1(x interface{}) {
    value := reflect.ValueOf(x)
    if value.Kind() == reflect.String {
        value.SetString("欢迎来到W3Cschool")
    }
}

func reflectsetvalue2(x interface{}) {
    value := reflect.ValueOf(x)

    // 反射中使用Elem()方法获取指针所指向的值
    if value.Elem().Kind() == reflect.String {
        value.Elem().SetString("欢迎来到W3Cschool")
    }
}


func main(){
    address := "www.w3cshool.cn"
    // reflectsetvalue1(address)  // panic: reflect: reflect.Value.SetString using unaddressable value
    fmt.Println(address)

    reflectsetvalue1(&address)
    fmt.Println(address)  // www.w3cshool.cn

    reflectsetvalue2(&address)
    fmt.Println(address)  // 欢迎来到W3Cschool
}
```

## 并发

- 基本概念
```text
- 并发: 同一时间段内执行多个任务（你早上在编程狮学习Java和Python）
- 并行: 同一时刻执行多个任务（你和你的网友早上都在使用编程狮学习Go）

Go语言中的并发程序
    - 主要是通过基于CSP（communicating sequential processes）的goroutine和channel来实现  TODO: 多进程？
    - 当然也支持使用传统的多线程共享内存的并发方式  # TODO： 多线程？

```

### goroutine

- goroutine
```go
// 函数或者方法前面加上go关键字就可以创建一个goroutine，从而让该函数或者方法在新的goroutine中执行
package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello!")
}

func main() {
    go hello()  // 没有打印出来
    hello()
    fmt.Println("main中！")
    time.Sleep(time.Second) // 等待子goroutine执行
}
```

- sync.WaitGroup
```go
// WaitGroup是实现等待一组并发操作完成的好方法
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func hello() {
    fmt.Println("Hello")
    defer wg.Done()  // goroutine结束计数器-1  TODO: 不懂啥意思？
}

func main() {
    wg.Add(1) // 把计算器加1
    go hello()

    fmt.Println("main中！“)
    wg.Wait() // 阻塞代码的运行，直到计算器为0
}
```

- 启动多个goroutine
```go
// WaitGroup是实现等待一组并发操作完成的好方法
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func hello(i int) {
    fmt.Printf("Hello-%d \n", i)
    defer wg.Done()  // goroutine结束计数器-1
}

// 并发每次执行结果顺序都不一样,10个goroutine都是并发执行的，而goroutine的调度是随机的
func main() {
    for i:=0; i<10; i++ {
        wg.Add(1) // 把计算器加1
        go hello(i)
    }
    fmt.Println("main中！")
    wg.Wait() // 阻塞代码的运行，直到计算器为0
}

```

- 动态栈
```text
- 操作系统的线程一般都有固定的栈内存（通常为2MB），
- 而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），所以在 Go 语言中一次创建数万个 goroutine 也是可能的。
- 并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小， Go 的 runtime 会自动为 goroutine 分配合适的栈空间。

```


- goroutine调度
```text
G: 表示goroutine，存储了goroutine的执行stack信息、goroutine状态以及goroutine的任务函数等；另外G对象是可以重用的。
P: 表示逻辑processor，P的数量决定了系统内最大可并行的G的数量（前提：系统的物理cpu核数>=P的数量）；P的最大作用还是其拥有的各种G对象队列、链表、一些cache和状态。
M: M代表着真正的执行计算资源。在绑定有效的p后，进入schedule循环；而schedule循环的机制大致是从各种队列、p的本地队列中获取G，切换到G的执行栈上并执行G的函数，调用goexit做清理工作并回到m，如此反复。M并不保留G状态，这是G可以跨M调度的基础。

```

- GOMAXPROCS 参数
```text
- Go运行时，调度器使用GOMAXPROCS的参数来决定需要使用多少个OS线程来同时执行Go代码。
- 默认值是当前计算机的CPU核心数。
- Go语言中可以使用runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU核心数
```


### channel

- channel 基本概念
```text
- 传统共享内存通信：
    - 单纯地将函数并发执行是没有意义的，函数与函数间需要交换数据才能体现并发执行函数的意义
    - 共享内存在不同的 goroutine 中容易发生竞态问题
    - 为了保证数据交换的正确性，很多并发模型中必须使用互斥锁对内存进行加锁，这种做法势必造成性能问题

- Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存，而不是通过共享内存而实现通信
    - 通道（channel）是一种特殊的类型。
    - 通道像一个传送带或者队列，总是遵循先入先出的规则，保证收发数据的顺序。
    - 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型
```


- channel类型
```text
- chan是关键字，元素类型指通道中传递的元素的类型
- 未经初始化的通道默认值为nil
- 声明的通道类型变量需要使用内置的make函数初始化之后才能使用
```

- channel操作-code
```go
// 通道共有发送，接收，关闭三种操作，而发送和接收操作均用​<-​符号


// 声明通道并初始化
a := make(chan int)

// 给一个通道发送值
a <- 10

// 从一个通道中取值
x := <-a //x从a通道中取值
<-a      //从a通道中取值，忽略结果

// 关闭通道: TODO: close内置函数？
close(a)

```

- channel关闭注意
```
- 一个通道值是可以被垃圾回收掉的。
- 通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
- 它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的
- 关闭后的通道有以下特点：
    - 对一个关闭的通道再发送值就会导致 panic。
    - 对一个关闭的通道进行接收会一直获取值直到通道为空。
    - 对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    - 关闭一个已经关闭的通道会导致 panic。
```

- 无缓冲的通道（阻塞的通道）
```go
// 无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段
// 如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。

package main

import (
    "fmt"
)

func main() {
    var channel chan int
    // channel = make(chan int)  // syntax error: missing channel element type
    channel = make(chan int)

    go receive(channel)

    // 无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段
    // deadlock表示我们程序中所有的goroutine都被挂起导致程序死锁了
    channel <- 10  // atal error: all goroutines are asleep - deadlock!

    fmt.Println("Hello!")

    // go receive(channel)  // 且必须先创建一个创建一个goroutine去接收值 TODO：需要使用go关键字在新的并发中去创建接收？
}


func receive(ch chan int) {
    result := <- ch
    fmt.Printf("接收成功：%d \n", result)
}
```

- 有缓冲区的通道
```text
- 另外还有一种方法解决上述死锁的问题，那就是使用有缓冲区的通道。
- 只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
- 当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
- 内置的len函数获取通道的长度，使用cap函数获取通道的容量

package main

import "fmt"

func main() {
	a := make(chan int,1)
	a <- 10
	fmt.Println("发送成功")
}

```


- 判断通道关闭
```go
// value, ok := <-ch
// value：表示从通道中所取得的值; ok：若通道已关闭，返回false，否则返回true

package main

import "fmt"

func receive(ch chan int) {
    
    // TODO： for 后面啥也不用加？
    for {  
        v,ok := <-ch
        if !ok {
            fmt.Println("通道已关闭！")
            break
        }
        fmt.Printf("v:%#v ok:%#v \n", v, ok)
    }
}

func main(){
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)
    receive(ch)  // TODO:这里没有用go关键字
}

```


- for range接收通道中的值
```go
package main

import "fmt"

func receive(ch chan int){
    for i:=range ch{
        fmt.Printf("v:%#v \n", i)
    }
}

func main(){
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    receive(ch)
}
```


- 单向通道
```text

- 在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递

<- chan int // 只接收通道，只能接收不能发送
chan <- int // 只发送通道，只能发送不能接收
```


- select多路复用
```text
- 同时从多个通道接收数据
- 如果没有数据可以被接收那么当前 goroutine 将会发生阻塞
- select关键字，使用它可以同时响应多个通道的操作 TODO: 那个通道有返回值则先执行，没有则阻塞？ select
```

- select 语句
```text
select {  
case <-ch1:
    //...
case data := <-ch2:
    //...
case ch3 <- 10:
    //...
default:
    //默认操作
}


- 可处理一个或多个channel的发送/接收操作
- 如果多个case同时满足，select会随机选择一个执行
- 对于没有case的select会一直阻塞，可用于阻塞 main 函数，防止退出

```

- code
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 1)  //创建一个类型为int，缓冲区大小为1的通道
    for i := 1; i<= 10; i++ {
        select {
            case x := <- ch:  //第一次循环由于没有值，所以该分支不满足
                fmt.Println(x)
            case ch <- i:  //将i发送给通道(由于缓冲区大小为1，缓冲区已满，第二次不会走该分支)
        }
    }
}
```

- 并发安全和互斥锁
```go
package main

import (
    "fmt"
    "sync"
)

var (
    x int64
    wg sync.WaitGroup
)

// add 对全局变量x执行5000次加1操作
func add() {
    for i:=0;i<5000;i++{
        x += 1
    }
    wg.Done()
}

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}

// 我们开启了2个goroutine去执行add函数，某个goroutine对全局变量x的修改可能会覆盖掉另外一个goroutine中的操作，所以导致结果与预期不符
```

- 互斥锁
```go
// sync包中提供的Mutex类型来实现互斥锁
// 用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；
// 当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的

package main

import (
    "fmt"
    "sync"
)

var (
    x int64
    wg sync.WaitGroup
    m sync.Mutex
)

// add 对全局变量x执行5000次加1操作
func add() {
    for i:=0;i<5000;i++{
        m.Lock()  // 使用互斥锁限制每次只有一个goroutine能修改全局变量x
        x += 1
        m.Unlock()
    }
    wg.Done()
}

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```


- 读写互斥锁
```text
- 当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的
- 读写锁分为两种：读锁和写锁
    - 当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
    - 而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待
```

- code
```go
// TODO：这个例子不对吧？先写后读？有并发读写吗？
package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    x = 0
    wg sync.WaitGroup
    rwlock sync.RWMutex
)

func read() {
    defer wg.Done()

    rwlock.RLock()
    fmt.Println("Read x= ", x)
    time.Sleep(time.Millisecond)

    rwlock.RUnlock()
}

func write() {
    defer wg.Done()
    rwlock.Lock()
    fmt.Println("Write x=", x)
    x += 1
    time.Sleep(time.Millisecond)
    fmt.Println("Write x+1=", x)
    rwlock.Unlock()
}


func main() {
    start := time.Now()
    for i := 0; i < 10; i++ {
        go write()
        wg.Add(1)
    }
    time.Sleep(time.Second)
    for i := 0; i < 20; i++ {
        go read()
        wg.Add(1)
    }
    wg.Wait()
    fmt.Println("----------------")
    fmt.Println(time.Since(start))
}

```




