package interface_usecase_notifSvc

import responsemodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/response_models"

type INotifUseCase interface {
	KafkaMessageConsumer()
	GetNotificationsForUser(userId, limit, offset *string) (*[]responsemodels_notifSvc.NotificationModel, error)
}
