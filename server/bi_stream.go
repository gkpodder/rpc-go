package main

import (
	pb "github.com/gkpodder/grpc-go/proto"
)

func (s *helloserver) SayHelloBidrectionalStream(stream pb.GreetService_SayHelloBidirectionalStreamServer) error {

}
