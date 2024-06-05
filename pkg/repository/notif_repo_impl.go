package repository_notifSvc

import (
	interface_repo_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/repository/interface"
	"gorm.io/gorm"
)

type NotifRepo struct {
	DB *gorm.DB
}

func NewNotifRepo(db *gorm.DB) interface_repo_notifSvc.INotifRepo {
	return &NotifRepo{DB: db}
}
