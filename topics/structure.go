//Structures as Function Arguments
package main
import ("fmt")

type Books struct{
  title string
  author string
} 

func main(){
  var Book1 Books
  var Book2 Books

  Book1.title = "GOT"
  Book1.author = "AAAA"

  Book2.title = "Telecom Billing"
  Book2.author = "Zara Ali"

  /* print Book1 info */
  /*fmt.Printf( "Book 1 title : %s\n", Book1.title)
  fmt.Printf( "Book 1 author : %s\n", Book1.author)*/
  /* print Book1 info */
  /*fmt.Printf( "Book 2 title : %s\n", Book2.title)
  fmt.Printf( "Book 2 author : %s\n", Book2.author)*/

  printBook(Book1);
}

func printBook(book Books){
  fmt.Println(book.title)
  fmt.Println(book.author)
}
