package main

import "log"

type logger struct{
	Name string
	isOn bool
}

func (l *logger) log(format string, v ...any){
	if l.isOn{
		log.Printf(format, v...)
	}
}


func (l *logger) warn(format string, v ...any){
	if l.isOn{
		log.Fatalf(format, v...)
	}
}
