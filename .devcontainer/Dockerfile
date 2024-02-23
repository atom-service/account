FROM --platform=linux/x86_64 alpine:3.19.1

# 指定工作目录
WORKDIR /workspace

# 安装基本运行环境
RUN apk add bash make util-linux curl unzip go openssh-client git nodejs npm \
    && node --version \
    && npm --version \
    && go version

# 设置 go 的全局 bin 变量
ENV PATH=$PATH:/root/go/bin

# 安装 protoc
RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v25.3/protoc-25.3-linux-x86_64.zip \
    && unzip protoc-25.3-linux-x86_64.zip -d protoc-25.3-linux-x86_64 \
    && cp -r protoc-25.3-linux-x86_64/bin/* /usr/bin/ \
    && rm -rf protoc-25.3-linux-x86_64.zip \
    && rm -rf protoc-25.3-linux-x86_64 \
    && protoc --version

# 安装 go 开发需要的套件
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
    && go install github.com/cosmtrek/air@v1.49
