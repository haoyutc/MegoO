package string_service

import (
	"context"
	"github.com/megoo/grpc/protos/pb"
	"io"
	"log"
	"strings"
)

type StringService struct {
}

func (s *StringService) LotsOfServerStream(request *pb.StringRequest, server pb.StringService_LotsOfServerStreamServer) error {
	response := pb.StringResponse{Ret: request.A + request.B}
	for i := 0; i < 10; i++ {
		server.Send(&response)
	}
	return nil
}

func (s *StringService) LotsOfClientStream(server pb.StringService_LotsOfClientStreamServer) error {
	var params []string
	for {
		in, err := server.Recv()
		if err == io.EOF {
			server.SendAndClose(&pb.StringResponse{Ret: strings.Join(params, "")})
			return nil
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
			return err
		}
		params = append(params, in.A, in.B)
	}
}

func (s *StringService) LotsOfServerAndClientStream(server pb.StringService_LotsOfServerAndClientStreamServer) error {
	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("failed to recv %v", err)
			return err
		}
		server.Send(&pb.StringResponse{Ret: in.A + in.B})
	}
	return nil
}

const StrMaxSize = 65533

func (s *StringService) Concat(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A)+len(req.B) > StrMaxSize {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	response := pb.StringResponse{Ret: req.B + req.A}
	return &response, nil
}

func (s *StringService) Diff(ctx context.Context, req *pb.StringRequest) (*pb.StringResponse, error) {
	if len(req.A) < 1 || len(req.B) < 1 {
		response := pb.StringResponse{Ret: ""}
		return &response, nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	response := pb.StringResponse{Ret: res}
	return &response, nil
}
