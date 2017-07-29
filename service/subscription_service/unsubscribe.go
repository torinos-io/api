package service

import "github.com/torinos-io/api/type/model"

// UnSubscribeRequest holds target project uuid
type UnSubscribeRequest struct {
	ProjectUUID string
}

// UnSubscribe performs unsubscribe project
func (s *service) UnSubscribe(req *UnSubscribeRequest, user *model.User) error {
	return s.SubscriptionStore.DeleteSubscription(user, req.ProjectUUID)
}
