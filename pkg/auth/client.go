package auth

import (
	"fmt"
	"github.com/Markovk1n/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/Markovk1n/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Cloud not connect: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}
