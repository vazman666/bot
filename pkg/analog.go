package pkg

import (
	pb "bot/pkg/gen"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

const address = "vazman.ru:50051"

//const address = "localhost:50051"

func Analogi(firm, num string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSqlRequestClient(conn)
	rect := &pb.Request{Number: num, Firm: firm}
	fmt.Printf("Looking for features within %v\n", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Analogs(ctx, rect)
	if err != nil {
		log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	type Analog struct {
		Firm   string
		Number string
	}

	var Analogs []Analog
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		fmt.Printf("Feature: name: %q\n", feature)
		Analogs = append(Analogs, Analog{Number: feature.Number, Firm: feature.Firm})
	}
	fmt.Printf("Аналоги для %v:  %v\n", rect, Analogs)
} /*
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("is connected")
	defer conn.Close()
	client := pb.NewSqlRequestClient(conn)
	printFeatures(client, &pb.Request{Number: "90915yzze2", Firm: "Toyota"})

}*/
