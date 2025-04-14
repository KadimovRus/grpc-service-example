package main

import (
	pb "github.com/KadimovRus/grpc-service-example/proto"
)

type Server struct {
	pb.GeometryServiceServer // сервис из сгенерированного пакета
}

func NewServer() *Server {
	return &Server{}
}

func main() {

}
