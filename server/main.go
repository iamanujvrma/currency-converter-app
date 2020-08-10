package main

import (
	"context"
	"fmt"
	"log"
	"net"

	cc "github.com/iamanujvrma/currency-converter-app/pkg/proto/currency_conversion"
	"google.golang.org/grpc"
)

// this is the struct which implements the gRPC server
// client calls the methods implemented by this struct
// while registering the server we pass reference to this struct
// to indicate that this is the implementation we want
// for our CurrencyConversionService

type server struct {
}

func (s *server) Credit(ctx context.Context, request *cc.CurrencyConversionRequest) (*cc.CurrencyConversionResponse, error) {
	log.Println(fmt.Sprintf("request received to convert currency %+v", request))

	// NOTE: here you can write your logic to convert the currency

	return &cc.CurrencyConversionResponse{
		ConvertedValue: 74.02,
	}, nil
}

func main() {
	log.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	cc.RegisterCurrencyConversionServiceServer(srv, &server{})

	log.Fatalln(srv.Serve(lis))
}
