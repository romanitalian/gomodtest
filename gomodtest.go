package gomodtest

import (
	"fmt"
	"time"
)

func Hi(name string) string {
	fmt.Println("test")
	fmt.Println(time.Now().String())
	return fmt.Sprintf("Hi, %s! It is great Day!", name)
}
