package main

import (
	"fmt"
	"syscall/js"
	ezimage "github.com/bangau1/ezimage/pkg/image"
)

func addWatermark(this js.value, args[]js.Value) interface{}{
	fmt.Println("Trying to add watermark...")

	js.Global().Get()
	return nil
}

func registerCallback(){
	js.Global().Set("addWatermark", js.FuncOf(addWatermark))
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
