package main

import (
	"fmt"
	"github.com/st0012/metago"
)

type Bar struct {
	name string
}

func (b *Bar) Name() string {
	return b.name
}

func (b *Bar) SetName(n string) {
	b.name = n
}

func (b *Bar) Send(methodName string, args ...interface{}) interface{} {
	return metago.CallFunc(b, methodName, args...)
}

func main() {
	b := &Bar{}
	b.Send("SetName", "Stan")   // This is like Object#send in Ruby
	fmt.Println(b.name)         				  // Should be "Stan"
	fmt.Println(b.Send("Name")) 	  // Should also be "Stan"
}
