package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("you get return value[err]", B())
	fmt.Println("you get return value[name]", D())
}

//
func B() (err error) {

	name, err := C()
	fmt.Println(name, err)
	return err
}

func D() (name string) {

	name, err := C()
	fmt.Println(name, err)
	return name
}

func C() (name string, err error) {
	err = errors.New("wo shi zhu")
	name = "123"
	return
}
