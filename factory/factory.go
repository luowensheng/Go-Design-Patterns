package factory

import "fmt"


func NewPublication( pg int, pubType, name, pub string)(IPublication, error){
     
	if pubType == "newspaper" {
      return createNewspaper(pg, pub, name), nil
	}

	if pubType == "magazine" {
		return createMagazine(pg, pub, name), nil
	}

	return nil, fmt.Errorf("No such publication type")

}

