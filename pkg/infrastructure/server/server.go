package server_notifSvc

import (
	"github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/pb"
	interface_usecase_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/usecase/interface"
)

type NotifService struct {
	NotifUseCase interface_usecase_notifSvc.INotifUseCase
	pb.NotificationServiceServer
}

func NewNotifServiceServer(notifUseCase interface_usecase_notifSvc.INotifUseCase) *NotifService {
	return &NotifService{NotifUseCase: notifUseCase}
}
