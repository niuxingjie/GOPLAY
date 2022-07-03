## W3Cschool

[Go 语言教程](https://www.w3cschool.cn/go/go-tutorial.html)

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









