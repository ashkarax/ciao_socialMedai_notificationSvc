package usecase_notifSvc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	config_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/config"
	requestmodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/request_models"
	responsemodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/response_models"
	"github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/pb"
	interface_repo_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/repository/interface"
	interface_usecase_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/usecase/interface"
	interface_dateToAge "github.com/ashkarax/ciao_socialMedai_notificationSvc/utils/DateToAge/interface"
)

type NotifUseCase struct {
	NotifRepo   interface_repo_notifSvc.INotifRepo
	KafkaConfig config_notifSvc.KafkaConfigs
	AuthClient  pb.AuthServiceClient
	DateToAge   interface_dateToAge.IDateToAge
}

func NewNotifUseCase(notifRepo interface_repo_notifSvc.INotifRepo,
	config config_notifSvc.KafkaConfigs,
	authClient *pb.AuthServiceClient,
	dateToAge interface_dateToAge.IDateToAge) interface_usecase_notifSvc.INotifUseCase {
	return &NotifUseCase{
		NotifRepo:   notifRepo,
		KafkaConfig: config,
		AuthClient:  *authClient,
		DateToAge:   dateToAge,
	}
}

func (r *NotifUseCase) KafkaMessageConsumer() {
	fmt.Println("---------kafka consumer initiated")
	configs := sarama.NewConfig()

	consumer, err := sarama.NewConsumer([]string{r.KafkaConfig.KafkaPort}, configs)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("----", consumer)
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(r.KafkaConfig.KafkaTopicNotification, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to get partitions: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		message := <-partitionConsumer.Messages()
		msg, _ := unmarshalChatMessage(message.Value)
		fmt.Println("===", msg)
		err := r.NotifRepo.CreateNewNotification(msg)
		if err != nil {
			log.Fatalln("-------err kafka consumer----------", err)
		}
	}
}

func unmarshalChatMessage(data []byte) (*requestmodels_notifSvc.KafkaNotificationTopicModel, error) {
	var store requestmodels_notifSvc.KafkaNotificationTopicModel

	err := json.Unmarshal(data, &store)
	if err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *NotifUseCase) GetNotificationsForUser(userId, limit, offset *string) (*[]responsemodels_notifSvc.NotificationModel, error) {

	notifData, err := r.NotifRepo.GetNotificationsForUser(userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range *notifData {
		context, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		userData, err := r.AuthClient.GetUserDetailsLiteForPostView(context, &pb.RequestUserId{UserId: fmt.Sprint((*notifData)[i].ActorID)})
		if err != nil || userData.ErrorMessage != "" {
			return nil, errors.New(fmt.Sprint(err) + userData.ErrorMessage)
		}
		(*notifData)[i].ActorUserName = userData.UserName
		(*notifData)[i].ActorProfileImgURL = userData.UserProfileImgURL

		(*notifData)[i].NotificationAge = *r.DateToAge.DateTOAge(&(*notifData)[i].CreatedAt)
	}

	return notifData, nil
}
