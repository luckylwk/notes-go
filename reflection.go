package main

import (
    "fmt"
    "reflect"
)

type Car struct {
    NumberOfWheels int
    model string
}

type Motorbike struct {
    NumberOfWheels int
    brand string
}

func fct(i interface{}) {
    s := reflect.ValueOf(i).Elem()
    typeOfT := s.Type()
    if typeOfT.Name() == "Car" {
        for i := 0; i < s.NumField(); i++ {
            f := s.Field(i)
            if typeOfT.Field(i).Name == "NumberOfWheels" {
                fmt.Printf("I am a %s and I have a field B: %d\n", typeOfT.Name(), f.Interface())
            }
        }
    }
}

func main() {
    c := Car{23, "Ford Focus"}
    m := Motorbike{23, "Ducati"}
    fct(&c)
    fct(&m)
}