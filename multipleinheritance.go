package main 

import (
	"fmt"
)

type Camera struct { 
	model string
}

func (_ Camera) takePicture() string {
	//not using the type, so discard it with _
	return "Click"
}

type Phone struct { 
	number string
	}

func (_ Phone ) call() string {
	//not using the type so discard it with _
	return "Ring Ring"
}

//mutliple inheritance
type CameraPhone struct {
	Camera //has anonymous camera
	Phone //has anonymous phone
}

func main() {
	cp := new (CameraPhone)
	cp.number = "123456"
	cp.model = "Canon"
	fmt.Println("Our new CameraPhone exhibits multiple behaviours ....")
	fmt.Println("It can take a picture: ", cp.takePicture()) //exhibits camera behaviour
	fmt.Println("It can also make calls: ", cp.call()) //and behaviour of a phone
	fmt.Println("Number", cp.number)
	fmt.Println("Number", cp.model)
}