package factory

import "fmt"


type Newspaper struct {
	Publication
 }
 
func (n Newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

func createNewspaper(pages int , publisher, name string) IPublication {
  
	return &Newspaper{
		Publication: Publication{name: name, publisher: publisher, pages: pages},
	}
}