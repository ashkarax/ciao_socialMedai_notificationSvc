package interface_repo_notifSvc

import (
	requestmodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/request_models"
	responsemodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/response_models"
)

type INotifRepo interface {
	CreateNewNotification(msg *requestmodels_notifSvc.KafkaNotificationTopicModel) error
	GetNotificationsForUser(userId, limit, offset *string) (*[]responsemodels_notifSvc.NotificationModel, error)
}
