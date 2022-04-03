FROM golang:1.17

LABEL maintainer="Yan <471573617@qq.com>"

ENV GOPROXY https://goproxy.cn,direct

USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim \
    xdg-utils\
    links

# 安装 goctl
USER root
RUN GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/go-zero/tools/goctl@latest

# 安装 protoc
USER root
RUN curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
    unzip -d /tmp/protoc /tmp/protoc.zip && \
    mv /tmp/protoc/bin/protoc $GOPATH/bin

# 安装 protoc-gen-go  goctl rpc protoc ./rpc/user.proto --go_out=./rpc --go-grpc_out=./rpc 
USER root
RUN GOPROXY=https://goproxy.cn/,direct go install github.com/golang/protobuf/protoc-gen-go@latest

# 安装grpc 
RUN GOPROXY=https://goproxy.cn/,direct go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# 安装grpc调试
RUN GOPROXY=https://goproxy.cn/,direct go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

# $GOPATH/bin添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 设置工作目录
WORKDIR /usr/src/code


EXPOSE 9003
