package main

import (
	"context"
	"github.com/megoo/grpc/protos/pb"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	serviceAddress := "127.0.0.1:1234"
	conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
	if err != nil {
		panic("connect error")
	}
	defer conn.Close()

	//stringClient := pb.NewStringServiceClient(conn)
	//stringReq := &pb.StringRequest{
	//	A: "A",
	//	B: "B",
	//}
	//reply, _ := stringClient.Concat(context.Background(), stringReq)
	//fmt.Printf("StringService Concat : %s concat %s = %s",
	//	stringReq.A, stringReq.B, reply.Ret)

	//stringClient := pb.NewStringServiceClient(conn)
	//stringReq := &pb.StringRequest{
	//	A: "A",
	//	B: "B",
	//}
	//stream, _ := stringClient.LotsOfServerStream(context.Background(), stringReq)
	//for {
	//	item, streamErr := stream.Recv()
	//	if streamErr == io.EOF {
	//		break
	//	}
	//	if streamErr != nil {
	//		log.Printf("failed to recv: %v", streamErr)
	//	}
	//	fmt.Printf("StringService Concat : %s concat %s = %s\n",
	//		stringReq.A, stringReq.B, item.GetRet())
	//}

	//stringClient := pb.NewStringServiceClient(conn)
	//stream, err := stringClient.LotsOfClientStream(context.Background())
	//for i := 0; i < 10; i++ {
	//	if err != nil {
	//		log.Printf("failed to call: %v", err)
	//		break
	//	}
	//	stream.Send(&pb.StringRequest{
	//		A: strconv.Itoa(i),
	//		B: strconv.Itoa(i + 1),
	//	})
	//}
	//reply, err := stream.CloseAndRecv()
	//if err != nil {
	//	fmt.Printf("failed to recv: %v", err)
	//}
	//log.Printf("ret is : %s", reply.Ret)

	client := pb.NewStringServiceClient(conn)
	stream, err := client.LotsOfServerAndClientStream(context.Background())
	if err != nil {
		log.Printf("failed to call: %v", err)
		return
	}
	var i int
	for {
		err2 := stream.Send(&pb.StringRequest{
			A: strconv.Itoa(i),
			B: strconv.Itoa(i + 1),
		})
		if err2 != nil {
			log.Printf("failed to send: %v", err2)
			break
		}
		reply, err3 := stream.Recv()
		if err3 != nil {
			log.Printf("failed to recv: %v", err3)
			break
		}
		log.Printf("Ret is : %s", reply.Ret)
	}
}
