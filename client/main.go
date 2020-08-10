package main

import (
	"context"
	"log"

	cc "github.com/iamanujvrma/currency-converter-app/pkg/proto/currency_conversion"
	"google.golang.org/grpc"
)

func main() {
	log.Println("starting grpc client...")

	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := cc.NewCurrencyConversionServiceClient(conn)

	request := &cc.CurrencyConversionRequest{
		From:  "INR",
		To:    "USD",
		Value: 1,
	}

	response, err := client.Credit(context.Background(), request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("converted currency value is: ", response.GetConvertedValue())
}
