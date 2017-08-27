package service

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-errors/errors"
	"github.com/guregu/null"

	"github.com/torinos-io/api/type/model"
)

// SubscribeRequest holds target project uuid and user's email
type SubscribeRequest struct {
	ProjectUUID string
	Email       string      `valid:"email"`
	User        *model.User `json:"-"`
}

// Subscribe performs subscribe project update
func (s *service) Subscribe(req *SubscribeRequest) (*model.Subscription, error) {
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	var userID null.Int

	if user := req.User; user != nil {
		userID = null.IntFrom(int64(user.ID))
	}

	return s.SubscriptionStore.CreateSubscription(userID, req.Email, req.ProjectUUID)
}
