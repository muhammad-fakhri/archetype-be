package repository

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"
	"fmt"

	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/go-libs/email"
)

// example to send simple email alert to predefined addresses
func (r *repository) SendSystemConfigReport(ctx context.Context, details string) error {
	if len(r.config.cmsAuditEmail) == 0 {
		return nil
	}

	go func() {
		adminID, _ := ctx.Value(constant.AuthorizedAdminIDKey).(string)

		// construct specific email format
		m := &email.MailDetail{
			SendTo: r.config.cmsAuditEmail,
			Title:  "[INFO] System Config Updated",
			Body:   fmt.Sprintf("System config %s updated by %s", details, adminID),
		}
		err := r.bridge.SendEmail(ctx, m)
		if err != nil {
			r.logger.Errorf(ctx, "failed to send email. err:%v", err)
		}
	}()

	return nil
}

// END __INCLUDE_EXAMPLE__
