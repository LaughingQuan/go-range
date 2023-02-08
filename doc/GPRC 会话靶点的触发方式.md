# GPRC 会话八点的触发方式

在 XMirror GO Agent v4.10.0.3.alpha 及以上版本的探针中，我们实现了对 GRPC 简单模式（GRPC 4 种请求/响应模式详情请参见[此处](https://www.jianshu.com/p/c703ce510f26)）的漏洞监控。

鉴于 GRPC 不同于 HTTP 请求，因此，目前我们的实现还未能完成如以往一般，可以通过一个简单的前端页面实现对 GRPC 靶点 Transaction 的激活。因此，有必要简单就 grpc 靶点激发方式进行简单说明。

具体可划分为如下几步：

1) 同以往一样，在完成探针的部署以及靶机源码解压后，执行如下操作对 GPRC 服务端进行注入：

```shell
❯ cd $PATH_TO_GO_GOAT_DIR

❯ ls
bin  config      db.sql  docker-compose.yml  front   go.sum  Makefile  README.md  util
cmd  config.yml  doc     Dockerfile          go.mod  http    model     targets

❯ $XMIRROR_HOME/xmirror-build -o bin/grpc_server ./cmd/grpc/server
getting necessary dependencies for the instrumented ...
getting necessary dependencies for the instrumented ...
building package for the instrumented ...

❯ ls bin
grpc_server
```

2) 同以往一样，利用 `xmirror-config -b` 来生成被注入应用所必须的 `xmirror-config.json`;

3) 如下，调用 `grpc_server` 来启动 GRPC 服务器，并利用 `go run` 来启动客户端触发 GRPC 靶点：

```shell
# 服务端
❯ cd $PATH_TO_GO_GOAT_DIR/bin

❯ ls
grpc_server  xmirror-config.yaml

goat/bin on  master [?]
❯ ./grpc_server -port=9090
2022/06/15 16:23:10 sqlbase.go:116: open /home/tjan/Lab/go/src/xmirror.cn/iast/goat/bin/db.sql: no such file or directory
2022/06/15 16:23:10 main.go:63: server listening at [::]:9090

# 客户端
❯ cd $PATH_TO_GO_GOAT_DIR

❯ go run ./cmd/grpc/client/main.go -addr=localhost:9090
2022/06/15 16:23:28 reply: (1)	[level=1]	不安全的加密算法	应用使用了不安全的或加密强度弱的加密算法，使被加密数据有可能被攻击者破解。

❯ go run ./cmd/grpc/client/main.go -addr=localhost:9090 -content=3
2022/06/15 16:23:47 reply: (2)	[level=3]	XStream反序列化	XStream可以将对象序列化成XML或将XML反序列化为对象。在使用XStream进行反序列化时，如果程序在对外部数据反序列化时，没有校验，会导致反序列化漏洞。
```
 
3) 最后，登录 IAST 后台 Web 端，查看是否靶机上报了一个“SQL注入”漏洞，其漏洞信息包含：

```
... ...
漏洞地址： /grpc.GoatGrpcSvc/SendMessage
漏洞来源： “GoGrpcTest1”应用“GoTest”节点下发现
漏洞详情： 在grpc方法的请求/grpc.GoatGrpcSvc/SendMessage中，参数*grpc.GoatRequest存在sql注入
... ...
```
