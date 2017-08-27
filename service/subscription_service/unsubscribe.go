package service

import (
	"github.com/guregu/null"

	"github.com/torinos-io/api/type/model"
)

// UnSubscribeRequest holds target project uuid
type UnSubscribeRequest struct {
	User        *model.User `json:"-"`
	ProjectUUID string
}

// UnSubscribe performs unsubscribe project
func (s *service) UnSubscribe(req *UnSubscribeRequest) error {
	var userID null.Int

	if user := req.User; user != nil {
		userID = null.IntFrom(int64(user.ID))
	}

	return s.SubscriptionStore.DeleteSubscription(userID, req.ProjectUUID)
}
