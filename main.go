package main

import (
	"fmt"
	"syscall/js"
)

func test(this js.Value, args []js.Value) interface{} {
	fmt.Println("test() is called.")
	return nil
}
func registerCallback(){
	js.Global().Set("test", js.FuncOf(test))
}

var c chan bool
func init()  {
	fmt.Println("Init is called")
	c = make(chan bool)
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	registerCallback()
	//wait until finish
	<-c
}
