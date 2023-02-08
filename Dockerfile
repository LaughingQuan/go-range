FROM  golang:1.17.8 AS builder


ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux 
ENV GOARCH=amd64
ENV GOPROXY=https://goproxy.cn,direct
    
WORKDIR /build 

COPY . .

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' >/etc/timezone \
    && ./xmirror-go build -o app ./cmd/gin \
    && ./xmirror-go config -b OzsxOTIuMTY4LjE3Mi4zMDs5MDkwO1BZQ0Y2UUoyR0MxUkcyUjM7RkdFUVkzUFdLMzgxNlhKRDtnb2F0Q0k7Z29hdENJOztwbWFhV0NSM3NIbEhtcjV6U1l6TXlDMTZnUERLYnlYVlNOVHMxYzAwMjZoSVczcw==

EXPOSE 8888

ENTRYPOINT ["/build/app", "-addr=:8888"] 
