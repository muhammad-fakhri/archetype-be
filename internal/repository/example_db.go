package repository

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"
	"database/sql"
	"fmt"

	json "github.com/json-iterator/go"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/model"
	"github.com/muhammad-fakhri/archetype-be/internal/util/timeutil"
	"github.com/muhammad-fakhri/archetype-be/pkg/errors"
)

const (
	systemConfigTableName = "system_config"
)

func (r *repository) GetSystemConfigAll(ctx context.Context) (cs []*model.SystemConfig, err error) {
	query := fmt.Sprintf(`SELECT name, config, created_at, updated_at FROM %s`, systemConfigTableName)
	rows, err := r.db.Master.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(errors.ErrDatabase, err.Error())
	}

	cs = []*model.SystemConfig{}
	for rows.Next() {
		c := &model.SystemConfig{}

		err = rows.Scan(c.GetAll()...)
		if err != nil {
			return nil, errors.Wrap(errors.ErrDatabase, err.Error())
		}

		cs = append(cs, c)
	}

	return
}

func (r *repository) GetSystemConfigByName(ctx context.Context, name constant.SystemConfig, out interface{}) (err error) {
	// get from db
	c := &model.SystemConfig{}

	query := fmt.Sprintf(`SELECT config FROM %s WHERE name = ?`, systemConfigTableName)
	err = r.db.Master.QueryRowContext(ctx, query, name).Scan(c.GetByName()...)
	if err != nil {
		return errors.Wrap(errors.ErrDatabase, err.Error())
	}

	err = json.Unmarshal([]byte(c.Config), out)
	if err != nil {
		return
	}

	return
}

func (r *repository) UpdateSystemConfig(ctx context.Context, name constant.SystemConfig, in interface{}) (err error) {
	adminID, ok := ctx.Value(constant.AuthorizedAdminIDKey).(string)
	if !ok || len(adminID) == 0 {
		return errors.Wrap(errors.ErrBadRequest, "invalid admin id")
	}

	config, err := json.Marshal(in)
	if err != nil {
		return
	}

	updatedAt := timeutil.NowMillis()

	query := fmt.Sprintf("UPDATE %s SET config = ?, updated_by = ?, updated_at = ? WHERE name = ?", systemConfigTableName)

	res, err := r.db.Master.Exec(query, string(config), adminID, updatedAt, name)
	if err != nil {
		return errors.Wrap(errors.ErrDatabase, err.Error())
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(errors.ErrDatabase, err.Error())
	}

	if affected == 0 {
		return errors.Wrap(errors.ErrDatabase, sql.ErrNoRows.Error())
	}

	return nil
}

// END __INCLUDE_EXAMPLE__
