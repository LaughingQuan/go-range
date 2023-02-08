package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"xmirror.cn/iast/goat/pem"
	pb "xmirror.cn/iast/goat/rpc"
)

const (
	defaultContent = "1"
)

var (
	flagAddr = flag.String("addr", "localhost:8080", "the server address")
	flagType = flag.String("t", "", "one of [sql, cmd, gen, rem]")

	flagTls = flag.Bool("tls", false, "use TLS instead of plain TCP")
	flagCa  = flag.String("ca", "", "the CA root certification file path")

	cmdInputTable = map[string]string{
		"cmd": "ifconfig",
		"gen": "sample.txt",
		"rem": "sample.txt",
		"sql": "1 or 1=1",
	}
)

func createReq(cmd string) *pb.GoatRequest {
	return &pb.GoatRequest{
		Cmd:   cmd,
		Input: cmdInputTable[cmd],
	}
}

func runUnary(cli pb.GoatServiceClient, req *pb.GoatRequest) {
	log.Printf("start a grpc unary session with request(%s:\"%s\")", req.Cmd, req.Input)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := cli.SampleUnary(ctx, req)
	if err != nil {
		log.Fatalf("run SampleUnary failed: %v", err)
	}
	log.Printf("server echo: %s", res.Content)
}

func runStream(cli pb.GoatServiceClient, reqs []*pb.GoatRequest) {
	log.Printf("start a grpc stream session with %d requests", len(reqs))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	strm, err := cli.SampleStream(ctx)
	if err != nil {
		log.Fatalf("get SampleStream stream failed: %v", err)
	}

	waitc := make(chan struct{})
	// receive
	go func() {
		for {
			res, err := strm.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("SampleStream receiving failed: %v", err)
			}
			log.Printf("got echo: %s", res.GetContent())
		}
	}()
	// send
	for _, req := range reqs {
		if err := strm.Send(req); err != nil {
			log.Fatalf("SampleStream sending failed: %v", err)
		}
	}
	strm.CloseSend()
	<-waitc
}

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *flagTls {
		if *flagCa == "" {
			*flagCa = pem.LoadPath("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*flagCa, "")
		if err != nil {
			log.Fatalf("failed to load TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*flagAddr, opts...)
	if err != nil {
		log.Fatalf("failed to dial grpc addr %s: %v", *flagAddr, err)
		return
	}
	defer conn.Close()
	cli := pb.NewGoatServiceClient(conn)

	if *flagType == "" {
		reqs := []*pb.GoatRequest{
			createReq("cmd"),
			createReq("gen"),
			createReq("rem"),
			createReq("sql"),
		}
		runStream(cli, reqs)
	} else {
		_, ok := cmdInputTable[*flagType]
		if !ok {
			log.Fatalf("unknown unary command %s", *flagType)
			return
		}
		runUnary(cli, createReq(*flagType))
	}
	log.Println("client operation done!")
}
