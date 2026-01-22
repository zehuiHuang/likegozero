https://go-zero.dev/
https://github.com/zeromicro/go-zero/blob/master/readme-cn.md

## 引用官网的描述:
go-zero 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

## goctl 工具
goctl 是 go-zero 的内置脚手架，是提升开发效率的一大利器，可以一键生成代码、文档、部署 k8s yaml、dockerfile 等。

## protoc 工具
protoc 是一个用于生成代码的工具，它可以根据 proto 文件生成C++、Java、Python、Go、PHP 等多重语言的代码，而 gRPC 的代码生成还依赖 protoc-gen-go，protoc-gen-go-grpc 插件来配合生成 Go 语言的 gRPC 代码。



## 创建简单的http服务案例
```
goctl api new demo
```

## 创建简单的rpc服务案例
```
goctl rpc new demo
```

