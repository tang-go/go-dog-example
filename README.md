<h1 align="center">go-dog-example</h1>
<div align="center">
go-dog对应的案例
</div>

### 使用
- 启动注册中心
此处使用go-dog默认注册中心演示
```bash
git clone https://github.com/tang-go/go-dog-tool.git
cd go-dog-tool/go-dog-find
go build
./go-dog-find
```

- 启动自动化网管
```bash
git clone https://github.com/tang-go/go-dog-tool.git
cd go-dog-tool/go-dog-gw
go build
./go-dog-gw
```


- 启动服务1
```bash
git clone https://github.com/tang-go/go-dog-example.git
cd go-dog-example
go mod tidy
cd service2
go build
./service2
```

- 启动服务2
```bash
git clone https://github.com/tang-go/go-dog-example.git
cd go-dog-example
go mod tidy
cd service2
go build
./service2
```

- 测试服务
```bash
浏览器输入 http://127.0.0.1/swagger/index.html 可访问API文档
```
- mysql为初始化mysql案例

- redis为使用redis案例


### 其他

- go-dog 框架 https://github.com/tang-go/go-dog
- go-dog-vue 前端 https://github.com/tang-go/go-dog-vue
- go-dog-example 例子 https://github.com/tang-go/go-dog-example

