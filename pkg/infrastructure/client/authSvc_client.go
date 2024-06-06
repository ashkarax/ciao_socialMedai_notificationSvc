package client_notifSvc

import (
	"fmt"

	config_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/config"
	"github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitAuthServiceClient(config *config_notifSvc.Config) (*pb.AuthServiceClient, error) {
	cc, err := grpc.Dial(config.PortMngr.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("-------", err)
		return nil, err
	}

	Client := pb.NewAuthServiceClient(cc)

	return &Client, nil
}
