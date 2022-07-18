package main

import (
	"fmt"

	"github.com/luowensheng/design_patterns/iterator"
)

func iteratorWithCallback() {

    fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Iterator with callback") 	
    fmt.Println("_________________________________\n") 	

	library := &iterator.Library{}
	
	library.AddBook(iterator.Book{Title: "Harry Potter", Author: "JK Rowling"})
	library.AddBook(iterator.Book{Title: "Bible", Author: "God"})
	library.AddBook(iterator.Book{Title: "The School of Life", Author: "Botton"})
    

	library.ForEach(func(item interface{}) {

		   var book = item.(iterator.Book)

           fmt.Println("Title:", book.Title, "\nAuthor:", book.Author, "\n")
	})
}

func iteratorWithInterface() {
     
    fmt.Println("_________________________________\n") 	

    fmt.Println("Starting Iterator with interface") 	
    fmt.Println("_________________________________\n") 	

	library := &iterator.Library{}
	
	library.AddBook(iterator.Book{Title: "Harry Potter", Author: "JK Rowling"})
	library.AddBook(iterator.Book{Title: "Bible", Author: "God"})
	library.AddBook(iterator.Book{Title: "The School of Life", Author: "Botton"})
    
    libIter := library.CreateIterator()

	for libIter.HasNext() {

		book := libIter.Next().(*iterator.Book)
		fmt.Println("Title:", book.Title, "\nAuthor:", book.Author, "\n")
	}
}

func viewIterator(){
    
	
	iteratorWithCallback()
	iteratorWithInterface()

}