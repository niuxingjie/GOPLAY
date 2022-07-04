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
    print_book(&book2)
}


func print_book(book *Book){
    fmt.Printf( "Book  title : %s\n", book.title)  // 奇怪，TODO：这里不该是*book拿到结构体吗？
    fmt.Printf( "Book  author : %s\n", book.author)
    fmt.Printf( "Book  subject : %s\n", book.subject)
    fmt.Printf( "Book  book_id : %d\n", book.book_id)
}