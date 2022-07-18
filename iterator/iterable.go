package iterator

type IterableWithCallback interface {
    ForEach(do func(interface{}))
}

type iterator interface {
     HasNext() bool
	 Next() interface{}
}


type IterableCollection interface {
    CreateIterator() iterator
}

type BookIterator struct {
	current int
	Books []Book
}

func (b *BookIterator) HasNext() bool {
	return b.current < len(b.Books)
}

func (b *BookIterator) Next() interface{} {
    b.current ++
   return &b.Books[b.current-1]
}