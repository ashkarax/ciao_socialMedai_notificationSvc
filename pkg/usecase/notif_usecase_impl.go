package usecase_notifSvc

import (
	interface_repo_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/repository/interface"
	interface_usecase_notifSvc "github.com/ashkarax/ciao_socialMedai_notificationSvc/pkg/usecase/interface"
)

type NotifUseCase struct {
	NotifRepo interface_repo_notifSvc.INotifRepo
}

func NewNotifUseCase(notifRepo interface_repo_notifSvc.INotifRepo) interface_usecase_notifSvc.INotifUseCase {
	return &NotifUseCase{NotifRepo: notifRepo}
}
