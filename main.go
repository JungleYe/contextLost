package main

import (
	"glgo/testcode/ctx/bar"
	"context"
	"fmt"
)

func main(){
	b := new(bar.Bar)
	b.Name = "world"
	ctx := newContextWithValue(context.TODO(),bar.CK,b)
	Pit(ctx)
}

//通过context来解析
func Pit(ctx context.Context){
	//使用foo.CK接收就会有异常。
	v := ctx.Value(foo.CK)
	re,ok := v.(foo.Bif)
	//v := ctx.Value(bar.CK)
	//re,ok := v.(bar.Bif)
	if ok{
		re.PrintB()
	}
	fmt.Println(v)
}

//创建一个新的context
func newContextWithValue(ctx context.Context,ck bar.ContextKey,bif bar.Bif)context.Context{
	return context.WithValue(context.TODO(),bar.CK,bif)
}
