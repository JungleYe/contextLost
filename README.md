# contextLost
跨goroutine时，通过context传递数据，接收方通过ctx.Value()获取到的结果为nil。

# types from different packages

问题的出现：
rpc需要通过ctx传递opentracing的span到业务上，传递的过程中通过业务注册好的controller对象以reflect.call的方式来执行对应的方法。
在方法中以opentracing.SpanFromContext(ctx)来获取span。结果是：通过ctx透传的其他信息都能够借助ctx.Value()正常还原回来，就是这个span对象怎么都拿不到！

问题调查：
1.在调用reflect.call之前是否能够拿到span？ A：可以的。
2.如果只透传spanContext，可以拿到么？A：可以的。
3.delve调试在controller中能够看到透传的信息？A：可以的。

4.自己重新实现opentracing.SpanFromContext以及ContextWithSpan后，业务controller中能够正常接收到span了。
5.在只透传spanContext的时候，还原成jaeger.SpanContext时，失败了。且有错误信息：opentracing.SpanContext is jaeger.SpanContext ,not jaeger.SpanContext(types from different packages)

问题的原因：
在执行reflect.call的时候，还是使用的rpc框架层的代码，opentracing的目录为："github.com/opentracing/opentracing-go"，其实际找到的时候是rpc目录的rpc/vendor/github....
而业务代码跟rpc框架在不同的src下面，业务controller中直接调用opentracing的方法时是从 business/vendor/github.....中找的。
也就是，通过context传递了A包的某对象，通过B包的某类型来接收。接收到的就是nil。
