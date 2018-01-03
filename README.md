# telegramd

## Chinese

### 简介
Go语言开源mtproto服务器，兼容telegram客户端

### 编译
首先，我们假设您已经回基本的golang基础，至少知道GOROOT，GOPATH这些概念。

#### 下载代码

    mkdir $GOPATH/src/github.com/nebulaim/
    cd $GOPATH/src/github.com/nebulaim/
    git clone https://github.com/nebulaim/telegramd.git

#### 编译代码

编译frontend

    cd $GOPATH/src/github.com/nebulaim/telegramd/frontend
    go get
    go build

编译sync_server

    cd $GOPATH/src/github.com/nebulaim/telegramd/sync_server
    go get
    go build

编译biz_server

    cd $GOPATH/src/github.com/nebulaim/telegramd/biz_server
    go get
    go build


### 运行

    cd $GOPATH/src/github.com/nebulaim/telegramd/frontend
    ./frontend

    cd $GOPATH/src/github.com/nebulaim/telegramd/sync_server
    ./sync_server
    
    cd $GOPATH/src/github.com/nebulaim/telegramd/biz_server
    ./biz_server

## English

### Introduce
open source mtproto server implement by golang, which compatible telegram client.

### Compile
