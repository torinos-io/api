package service

import (
	"github.com/guregu/null"

	"github.com/torinos-io/api/type/model"
)

// SubscribeRequest holds target project uuid and user's email
type SubscribeRequest struct {
	ProjectUUID string
	Email       string
	User        *model.User `json:"-"`
}

// Subscribe performs subscribe project update
func (s *service) Subscribe(req *SubscribeRequest) (*model.Subscription, error) {
	var userID null.Int

	if user := req.User; user != nil {
		userID = null.IntFrom(int64(user.ID))
	}

	return s.SubscriptionStore.CreateSubscription(userID, req.Email, req.ProjectUUID)
}
