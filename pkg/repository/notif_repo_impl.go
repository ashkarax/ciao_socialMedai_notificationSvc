package repository_notifSvc

import (
	requestmodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/request_models"
	responsemodels_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/infrastructure/models/response_models"
	interface_repo_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/repository/interface"
	"gorm.io/gorm"
)

type NotifRepo struct {
	DB *gorm.DB
}

func NewNotifRepo(db *gorm.DB) interface_repo_notifSvc.INotifRepo {
	return &NotifRepo{DB: db}
}

func (d *NotifRepo) CreateNewNotification(msg *requestmodels_notifSvc.KafkaNotificationTopicModel) error {
	query := "INSERT INTO notifications (user_id,actor_id,action_type,target_id,target_type,comment_text,created_at) VALUES($1,$2,$3,$4,$5,$6,$7)"
	err := d.DB.Exec(query, msg.UserID, msg.ActorID, msg.ActionType, msg.TargetID, msg.TargetType, msg.CommentText, msg.CreatedAt).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *NotifRepo) GetNotificationsForUser(userId, limit, offset *string) (*[]responsemodels_notifSvc.NotificationModel, error) {
	var respModel []responsemodels_notifSvc.NotificationModel

	query := "SELECT * FROM notifications WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
	err := d.DB.Raw(query, userId, limit, offset).Scan(&respModel).Error
	if err != nil {
		return nil, err
	}
	return &respModel, nil
}
