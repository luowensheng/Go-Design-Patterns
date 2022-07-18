package factory
import "fmt"


 
type Magazine struct {
	Publication
 }
 
func (m Magazine) String() string {
	return fmt.Sprintf("This is a newspaper named %s", m.name)
}

func createMagazine(pages int , publisher, name string) IPublication {
  
	return &Magazine{
		Publication: Publication{name: name, publisher: publisher, pages: pages},
	}
}