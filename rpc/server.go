package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

// Server grpc服务
type Server struct {
	*grpc.Server
	conf *Config
}

func NewServer(conf *Config) *Server {
	s := &Server{conf: conf}
	s.Server = grpc.NewServer()
	return s
}

// 启动
func (s *Server) Start() error {
	addr := net.JoinHostPort("", strconv.Itoa(s.conf.Port))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	reflection.Register(s.Server)
	log.Printf("Listening and serving RPC on %s\n", addr)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("serve rpc failed: %v\n", err)
		}
	}()

	return nil
}
