package repository

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"

	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

func (r *repository) PublishSystemConfig(ctx context.Context, data map[constant.SystemConfig]interface{}) error {
	adminID, ok := ctx.Value(constant.AuthorizedAdminIDKey).(string)
	if !ok || len(adminID) == 0 {
		return errors.Wrap(errors.ErrBadRequest, "invalid admin id")
	}

	return r.publisher.UpdateSystemConfig.Publish(ctx, data, map[string]string{
		string(constant.AuthorizedAdminIDKey): adminID,
	})
}

// END __INCLUDE_EXAMPLE__
