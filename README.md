# APIGateway-service


开发需要安装 goctl工具

go install github.com/zeromicro/go-zero/tools/goctl@latest

查看版本号
goctl -version
goctl version 1.5.2 windows/amd64


使用 go-zero提供的便携安装 protoC & protoc-gen-go
goctl env check -i -f --verbose

生成路由文件
goctl api go -api gateway.api -dir .
