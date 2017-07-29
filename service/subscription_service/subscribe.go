package service

import "github.com/torinos-io/api/type/model"

// SubscribeRequest holds target project uuid and user's email
type SubscribeRequest struct {
	ProjectUUID string
	Email       string
}

// Subscribe performs subscribe project update
func (s *service) Subscribe(req *SubscribeRequest, user *model.User) (*model.Subscription, error) {
	return s.SubscriptionStore.CreateSubscription(user.ID, req.Email, req.ProjectUUID)
}
