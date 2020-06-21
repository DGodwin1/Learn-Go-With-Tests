package main

import "fmt"

func main(){
	fmt.Println(Hello("David"))
}

func Hello(name string) string{
	return "Hello, " + name
}