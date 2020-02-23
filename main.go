package main

import (
	"bytes"
	"fmt"
	"syscall/js"
	ezimage "github.com/bangau1/ezimage/pkg/image"
)

func addWatermark(this js.Value, args[]js.Value) interface{}{
	fmt.Println("Trying to add watermark...%T", args[0].Type())
	fmt.Println("type baseImage %T", args[0].Get("baseImage").Type())
	fmt.Println("type watermarkImage %T", args[0].Get("watermarkImage").Type())
	fmt.Println("type watermarkImage %T", args[0].Get("w").Type())
	baseImageJs := args[0].Get("baseImage")
	watermarkImageJs :=  args[0].Get("watermarkImage")
	baseImg := make([]byte, baseImageJs.Length())


	n := js.CopyBytesToGo(baseImg, baseImageJs)
	fmt.Printf("n is %d\n", n)
	fmt.Printf("baseImg len: %d, cap: %d\n", len(baseImg), cap(baseImg))

	waterImg := make([]byte, watermarkImageJs.Length())
	n = js.CopyBytesToGo(waterImg, watermarkImageJs)
	fmt.Printf("n is %d\n", n)
	fmt.Printf("waterImg len: %d, cap: %d\n", len(waterImg), cap(waterImg))

	result := make(map[string]interface{})

	baseImgData, err := ezimage.NewImageFromReader(bytes.NewReader(baseImg))
	if err != nil{
		result["error"] = err.Error()
		return js.ValueOf(result)
	}

	watermarkImgData, err := ezimage.NewImageFromReader(bytes.NewReader(waterImg))
	if err != nil{
		result["error"] = err.Error()
		return js.ValueOf(result)
	}

	resultData := ezimage.NewWaterMarkProcessing(watermarkImgData).Apply(baseImgData)
	if resultData.Error != nil{
		result["error"] = resultData.Error.Error()
		return js.ValueOf(result)
	}

	buf := new(bytes.Buffer)
	err = resultData.Data.Save(buf, ezimage.MIME_TYPE_JPEG, 90)
	if err != nil{
		fmt.Println("error saving data", err)
		result["error"] = err.Error()
		return js.ValueOf(result)
	}
	fmt.Println("success, now return the data")
	dataJs := js.Global().Get("Uint8Array").New(len(buf.Bytes()))
	fmt.Println(dataJs)

	js.CopyBytesToJS(dataJs, buf.Bytes())
	result["data"] = dataJs
	return js.ValueOf(result)
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
