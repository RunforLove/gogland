package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	price float64
}
func main () {
	var book01 Books
	var book02 Books

	book01.title = "shuanghun"
	book01.author = "jushi"
	book01.subject = "caobao"
	book01.price = 250.0

	book02.title = "shuanghun2"
	book02.author = "jushi2"
	book02.subject = "caobao2"
	book02.price = 250.2

	fmt.Println(book01.title,book01.author,book01.subject,book01.price)
	fmt.Println(book02.title,book02.author,book02.subject,book02.price)

	printBooks(book01)
	printBooks(book02)
}

func printBooks (book Books) {
	fmt.Printf("book title is %s, book author is %s, book subject is %s, book price is %f\n", book.title,book.author,book.subject,book.price)
}

