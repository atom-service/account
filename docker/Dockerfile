# 使用基础镜像
FROM golang:alpine

# 设置工作目录
WORKDIR /app

# 拷贝 Go 代码到容器中
COPY . .

# 构建 Go 应用程序
RUN go build -o main .

# 暴露应用程序监听的端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]
