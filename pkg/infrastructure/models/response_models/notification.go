package responsemodels_notifSvc

import "time"

type NotificationModel struct {
	NotificaitonID     uint64
	UserID             uint64
	ActorID            uint64
	ActorUserName      string
	ActorProfileImgURL string
	ActionType         string
	TargetID           uint64
	TargetType         string
	CreatedAt          time.Time
	CommentText        string
	NotificationAge    string
}
