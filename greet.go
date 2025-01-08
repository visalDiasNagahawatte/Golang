package main

import "fmt"

func Greet(name string)string{
    message:= fmt.Sprintf("Welcome %v to To-do list",name)
	return message 
}