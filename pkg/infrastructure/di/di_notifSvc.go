package di_notifSvc

import (
	"fmt"

	client_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/client"
	config_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/config"
	db_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/db"
	server_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/server"
	repository_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/repository"
	usecase_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/usecase"
	datetoage "github.com/ashkarax/ciao_socialMedai_notificationSvc/utils/DateToAge"
	hashpassword_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/utils/hash_password"
)

func InitializeNotificationServer(config *config_notifSvc.Config) (*server_notifSvc.NotifService, error) {

	hashUtil := hashpassword_notifSvc.NewHashUtil()

	DB, err := db_notifSvc.ConnectDatabase(&config.DB, hashUtil)
	if err != nil {
		fmt.Println("ERROR CONNECTING DB FROM DI.GO")
		return nil, err
	}

	authClient, err := client_notifSvc.InitAuthServiceClient(config)
	if err != nil {
		fmt.Println("--------err--------", err)
		return nil, err
	}

	dateToAge := datetoage.NewDateToAgeUtil()

	notifRepo := repository_notifSvc.NewNotifRepo(DB)
	notifUseCase := usecase_notifSvc.NewNotifUseCase(notifRepo, config.KafkaConfig, authClient,dateToAge)

	go notifUseCase.KafkaMessageConsumer()

	return server_notifSvc.NewNotifServiceServer(notifUseCase), nil

}
