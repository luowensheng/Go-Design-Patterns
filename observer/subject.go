package observer

type observable interface {
	registerObserver(o observer)
	unregisterObserver(o observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener 
	field string
}