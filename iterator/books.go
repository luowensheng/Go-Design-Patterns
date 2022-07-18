package iterator

type Book struct {
    Title string
	Author string
}


type Library struct {
   Collection []Book
}

func (lib *Library) ForEach(do func(interface{})){

    for _, book := range lib.Collection {
        do(book)
	}
}

func (lib *Library) AddBook(book Book){
    lib.Collection = append(lib.Collection, book)
}

func (lib *Library) CreateIterator() BookIterator{
    return BookIterator{
		Books: lib.Collection,
		current: 0,
	}
}
