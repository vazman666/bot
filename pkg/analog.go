package pkg

import (
	"bot/models"
	pb "bot/pkg/gen"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

const address = "vazman.ru:50051"

//const address = "192.168.1.16:50051"

func Analogi(firm, num string) {
	models.Analogs = models.Analogs[:0]
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
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		fmt.Printf("Feature: name: %q\n", feature)
		models.Analogs = append(models.Analogs, models.Analog{Number: feature.Number, Firm: feature.Firm})
	}
	fmt.Printf("Аналоги для %v:  %v\n", rect, models.Analogs)
}

