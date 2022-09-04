package handler

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	ps "github.com/muhammad-fakhri/archetype-be/internal/component/pubsub"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

type UpdateSystemConfigHandler func(ctx context.Context, configMap map[constant.SystemConfig]interface{}) (err error)

func UpdateSystemConfig(f UpdateSystemConfigHandler) ps.SubscriberHandler {
	return func(ctx context.Context, m *pubsub.Message) error {
		var data map[constant.SystemConfig]interface{}

		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			return errors.Wrap(errors.ErrBadRequest, err.Error())
		}

		adminID, ok := m.Attributes[string(constant.AuthorizedAdminIDKey)]
		if !ok {
			return errors.ErrBadRequest
		}

		ctx = context.WithValue(ctx, constant.AuthorizedAdminIDKey, adminID)
		err = f(ctx, data)
		return err
	}
}

// END __INCLUDE_EXAMPLE__
