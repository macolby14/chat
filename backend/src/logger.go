package main

import "log"

type Logger struct{
	Name string
	IsOn bool
}

func (l *Logger) log(format string, v ...any){
	if l.IsOn{
		log.Printf(format, v...)
	}
}


func (l *Logger) warn(format string, v ...any){
	if l.IsOn{
		log.Fatalf(format, v...)
	}
}
