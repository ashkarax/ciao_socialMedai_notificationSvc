package requestmodels_notifSvc

import "time"

type KafkaNotificationTopicModel struct {
	UserID      string
	ActorID     string
	ActionType  string
	TargetID    string
	TargetType  string
	CommentText string
	CreatedAt   time.Time
}
