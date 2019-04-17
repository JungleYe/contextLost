package foo

import "fmt"

type ContextKey struct{}
var CK = ContextKey{}

type Bar struct {
	Name string
}

type Bif interface{
	PrintB()
}

func (b *Bar)PrintB(){
	fmt.Println("b.Name = ",b.Name)
}
