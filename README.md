# GoGoat

A sample vulnerable target site for xmirror [go-agent](http://192.168.172.218/root/GoAgent)

## 1. HTTP 靶机部署方式

### a. 将探针的xmirro-go文件拷贝到靶机根目录

### b. 在灵脉前端界面注册结点并修改靶机根目录下面 Dockerfile 文件中第 15 行 `-b` 后面的字符串

### c. 在靶机根目录下执行 `make all`

看见如下输出,则表示靶机部署成功
```bash
rm -f grpc_server grpc_cli xmirror-config.yaml
docker stop GoGoatApp GoGoatOpenLdap GoGoatMysql 2>/dev/null || true
docker rm GoGoatApp GoGoatOpenLdap GoGoatMysql 2>/dev/null || true
docker network rm goat-network 2>/dev/null || true
goat-network
docker-compose up -d
Creating network "goat-network" with the default driver
Pulling openldap (osixia/openldap:latest)...
latest: Pulling from osixia/openldap
.....
Digest: sha256:3f68751292b43564a2586fc29fb7337573e2dad692b92d4e78e49ad5c22e567b
Status: Downloaded newer image for osixia/openldap:latest
Pulling mysql (mysql:5.6)...
5.6: Pulling from library/mysql
....
Digest: sha256:20575ecebe6216036d25dab5903808211f1e9ba63dc7825ac20cb975e34cfcae
Status: Downloaded newer image for mysql:5.6
Building app
Sending build context to Docker daemon  55.24MB
Step 1/11 : FROM  golang:1.17.8 AS builder
1.17.8: Pulling from library/golang
......
Digest: sha256:f675106e44f205a7284e15cd75c41b241329f3c03ac30b0ba07b14a6ea7c99d9
Status: Downloaded newer image for golang:1.17.8
......
Step 9/11 : RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime     && echo 'Asia/Shanghai' >/etc/timezone     && ./xmirror-go build -o app ./cmd/gin     && ./xmirror-go config -b OzsxOTIuMTY4LjE3Mi4zMDs5MDkwO1BZQ0Y2UUoyR0MxUkcyUjM7RkdFUVkzUFdLMzgxNlhKRDtnb2F0Q0k7Z29hdENJOztwbWFhV0NSM3NIbEhtcjV6U1l6TXlDMTZnUERLYnlYVlNOVHMxYzAwMjZoSVczcw==
 ---> Running in 57c528e835e7
getting necessary dependencies for the instrumented ...
getting necessary dependencies for the instrumented ...
building package for the instrumented ...
config saved successfully in xmirror-config.yaml
Removing intermediate container 57c528e835e7
......
Successfully built 84fb8c8ff58a
Successfully tagged injectee_app:latest
Creating GoGoatOpenLdap ... done
Creating GoGoatMysql    ... done
Creating GoGoatApp      ... done
```

## 2. Grpc 靶机部署方式

### a. 在机器中部署好探针（必要）

### b. 在靶机根目录下执行 `make grpc` 构建 grpc 靶机服务器与对应的客户端

运行效果如下（需要将探针文件xmirrot-go文件拷贝到靶机根目录）：

```bash
❯ make grpc
rm -f grpc_server grpc_cli xmirror-config.yaml
rm -rf xmirror.cn
mkdir -p xmirror.cn/bin
/home/tjan/Lab/go/src/xmirror.cn/iast/agent/bin/xmirror-build -a -o ./grpc_server ./cmd/grpc/server/
getting necessary dependencies for the instrumented ...
getting necessary dependencies for the instrumented ...
building package for the instrumented ...
/home/tjan/Lab/go/src/xmirror.cn/iast/agent/bin/xmirror-build -a -o ./grpc_cli ./cmd/grpc/client/
getting necessary dependencies for the instrumented ...
getting necessary dependencies for the instrumented ...
building package for the instrumented ...
```

执行完毕后，确认目录下确实生成 `grpc_server` 与 `grpc_cli` 两个二进制文件。

### c. 依照《[xmirror-config 应用的使用说明](http://192.168.172.85:8090/pages/viewpage.action?pageId=142574160)》操作生成探针运行配置文件

### d. 分别启动 grpc 靶机服务器，并使用客户端仅进行访问模拟

将前一部生成的 `xmirror-config.yaml` 探针配置置于 `grpc_server` 同一目录下，然后即可按照如下操作启动 grpc 服务器：

```bash
❯ ./grpc_server
2022/08/16 17:17:13.384 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: no such file or directory # 由靶点的 beego 关联导致，可忽略
2022/08/16 17:17:13 sqlbase.go:120: Generate database files...
2022/08/16 17:17:14 main.go:77: server listening at [::]:8080
```

服务器即开始监听本地 8080 端口。如果测试环境需要，用户可通过 `-p` 参数指定其他服务器监听端口。

之后，用户可在另一个命令中断，通过调用 `grpc_cli` 实现基于 grpc 的访问模拟。执行效果如下：

```bash
# 不带任何参数，默认访问 localhost:8080，通信方式为 grpc 双向流模式，向服务器发起一次包含：命令执行、文件创建、文件读写、文件删除 以及 SQL 注入的访问流；并通过服务器的回复流获取访问结果
❯ ./grpc_cli 
2022/08/16 17:23:48 main.go:58: start a grpc stream session with 4 requests
... ...
2022/08/16 17:23:49 main.go:134: client operation done!

# -addr 可用于指定服务器地址（用于访问非本地服务器），-t 可用于选择可触发的漏洞，目前提供 cmd（命令执行），gen（文件创建、文件读写），rem（文件删除，注意需要服务器先响应 gen 或者用户手动在服务器根目录创建 sample.txt 文件），sql（SQL 注入）
# 注意，当 -t 标识开启时，服务器与客户端会以简单模式（一元模式）进行通讯
❯ ./grpc_cli -addr 127.0.0.1:8080 -t sql 
2022/08/16 17:28:57 main.go:45: start a grpc unary session with request(sql:"1 or 1=1")
... ...
2022/08/16 17:28:57 main.go:134: client operation done!
```
