package model

// BEGIN __INCLUDE_EXAMPLE__
import "github.com/muhammad-fakhri/archetype-be/internal/constant"

type SystemConfig struct {
	Name      constant.SystemConfig `json:"name"` //primary key, max length 20 chars
	Config    string                `json:"config"`
	CreatedBy string                `json:"created_by"`
	UpdatedBy string                `json:"updated_by"`
	CreatedAt int64                 `json:"created_at"`
	UpdatedAt int64                 `json:"updated_at"`
}

func (m *SystemConfig) GetByName() []interface{} {
	return []interface{}{
		&m.Config,
	}
}

func (m *SystemConfig) GetAll() []interface{} {
	return []interface{}{
		&m.Name,
		&m.Config,
		&m.CreatedAt,
		&m.UpdatedAt,
	}
}

// END __INCLUDE_EXAMPLE__
