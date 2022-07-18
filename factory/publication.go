package factory


type IPublication interface {
	SetName(name string)
	SetPages(pages int)
	SetPublisher(publisher string)
	GetName() string
	GetPages() int
	GetPublisher() string
}

type Publication struct {
   name string
   pages int 
   publisher string
}

func (publication *Publication)	SetName(name string){
    publication.name = name
}

func (publication *Publication)	SetPages(pages int){
   publication.pages = pages
}

func (publication *Publication)	SetPublisher(publisher string){
  publication.publisher = publisher
}

func (publication *Publication)	GetName() string{
  return publication.name
}

func (publication *Publication)	GetPages() int{
  return publication.pages
}

func (publication *Publication)	GetPublisher() string {
  return publication.publisher
}