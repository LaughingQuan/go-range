package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pem "xmirror.cn/iast/goat/pem"
	pb "xmirror.cn/iast/goat/rpc"
)

var (
	flagPort         = flag.Int("p", 8080, "the port of server")
	flagTls          = flag.Bool("tls", false, "the TLS switch")
	flagCert         = flag.String("cert", "", "the cerfitication file path")
	flagKey          = flag.String("key", "", "the key file for TLS connection")
	flagPythonServer = flag.String("ps", "", "the python grpc server addr ,such as 192.168.172.180:81")
)

type Server struct {
	pb.UnimplementedGoatServiceServer
}

func (s *Server) SampleUnary(ctx context.Context, req *pb.GoatRequest) (res *pb.GoatResponse, err error) {
	requestPythonGrpc()
	return handle(req.GetCmd(), req.GetInput())
}

func (s *Server) SampleStream(srv pb.GoatService_SampleStreamServer) error {
	requestPythonGrpc()
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if res, err := handle(req.GetCmd(), req.GetInput()); err == nil {
			srv.Send(res)
		} else {
			log.Printf("Sample stream failed for %s input %s: %v", req.GetCmd(), req.GetInput(), err)
			return err
		}
	}
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *flagPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *flagTls {
		if *flagCert == "" {
			*flagCert = pem.LoadPath("x509/server_cert.pem")
		}
		if *flagKey == "" {
			*flagKey = pem.LoadPath("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*flagCert, *flagKey)
		if err != nil {
			log.Fatalf("Failed to init credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGoatServiceServer(s, &Server{})
	log.Printf("server listening at %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func requestPythonGrpc() {
	if len(*flagPythonServer) == 0 {
		return
	}
	conn, err := grpc.Dial(*flagPythonServer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial grpc addr %s: %v", *flagPythonServer, err)
		return
	}
	defer conn.Close()
	cli := pb.NewGRPCDemoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := cli.SimpleMethod(ctx, &pb.Request{ClientId: 100086, RequestData: "cat /etc/passwd"})
	if err != nil {
		log.Printf("failed to request python grpc err %s", err)
		return
	}
	fmt.Printf("request python node request id %d, data %s", response.ServerId, response.ResponseData)
}
