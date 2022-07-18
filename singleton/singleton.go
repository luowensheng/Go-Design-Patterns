package singleton

import (
	"fmt"
	"sync"
)

type myLogger struct {
	loglevel int
}

func (l *myLogger) Log(s string){
	fmt.Println(l.loglevel, ":", s)
}

func (l *myLogger) SetLogLevel(level int){
	l.loglevel = level
}

var logger *myLogger 

var once sync.Once

func GetLoggerInstance() *myLogger{

	once.Do(func() {
		if logger == nil {
			fmt.Println("Creating logger")
			logger = &myLogger{}
		} 	
	})
	fmt.Println("Returning logger")
   return logger
}

