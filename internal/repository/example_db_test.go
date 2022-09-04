package repository_test

// BEGIN __INCLUDE_EXAMPLE__
import (
	"context"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/muhammad-fakhri/archetype-be/internal/constant"
	"github.com/muhammad-fakhri/archetype-be/internal/model"
	"github.com/muhammad-fakhri/archetype-be/internal/test/testutil"
	"github.com/muhammad-fakhri/archetype-be/pkg/dto"
	commonconstant "github.com/muhammad-fakhri/go-libs/constant"
	"github.com/stretchr/testify/assert"
)

func TestGetSystemConfigAll(t *testing.T) {
	m := testutil.InitMock(t)
	defer func() {
		m.DB.Master.Close()
		m.Controller.Finish()
	}()

	repo := testutil.InitMockRepository(m)

	var (
		query = "SELECT name, config, created_at, updated_at FROM system_config"
	)
	type params struct {
		data []*model.SystemConfig
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result []*model.SystemConfig, err error)
	}{
		{
			caseName: "success",
			params: params{
				data: []*model.SystemConfig{
					{
						Name: constant.CountryConfig,
					},
					{
						Name: constant.SystemConfig("other config"),
					},
				},
			},
			expectations: func(params params) {
				rows := sqlmock.NewRows([]string{"name", "config", "created_at", "updated_at"})

				for _, d := range params.data {
					rows.AddRow(d.Name, d.Config, d.CreatedAt, d.UpdatedAt)
				}

				m.DBMock.ExpectQuery(query).WillReturnRows(rows)
			},
			results: func(t *testing.T, result []*model.SystemConfig, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 2, len(result))
			},
		},
		{
			caseName: "failed",
			params: params{
				data: []*model.SystemConfig{
					{
						Name: constant.CountryConfig,
					},
					{
						Name: constant.SystemConfig("other config"),
					},
				},
			},
			expectations: func(params params) {
				rows := sqlmock.NewRows([]string{"name", "config", "created_at", "updated_at"})

				for _, d := range params.data {
					rows.AddRow(d.Name, d.Config, d.CreatedAt, d.UpdatedAt)
				}

				m.DBMock.ExpectQuery(query).WillReturnError(testutil.DefaultErr)
			},
			results: func(t *testing.T, result []*model.SystemConfig, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - scan",
			params: params{
				data: []*model.SystemConfig{
					{
						Name: constant.CountryConfig,
					},
					{
						Name: constant.SystemConfig("other config"),
					},
				},
			},
			expectations: func(params params) {
				rows := sqlmock.NewRows([]string{"name", "config", "created_at", "updated_at", "additional_column"})

				for _, d := range params.data {
					rows.AddRow(d.Name, d.Config, d.CreatedAt, d.UpdatedAt, "string")
				}

				m.DBMock.ExpectQuery(query).WillReturnRows(rows)
			},
			results: func(t *testing.T, result []*model.SystemConfig, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)
			res, err := repo.GetSystemConfigAll(context.Background())
			testCase.results(t, res, err)
		})
	}
}

func TestGetSystemConfigByName(t *testing.T) {
	m := testutil.InitMock(t)
	defer func() {
		m.DB.Master.Close()
		m.Controller.Finish()
	}()

	repo := testutil.InitMockRepository(m)

	var (
		query = "SELECT config FROM system_config WHERE name = \\?"
	)
	type params struct {
		out  interface{}
		name constant.SystemConfig
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result interface{}, err error)
	}{
		{
			caseName: "success",
			params: params{
				name: constant.CountryConfig,
				out:  &dto.CountryConfig{},
			},
			expectations: func(params params) {
				c, _ := json.Marshal(&dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				})
				rows := sqlmock.NewRows([]string{"config"})
				rows.AddRow(string(c))

				m.DBMock.ExpectQuery(query).WithArgs(params.name).WillReturnRows(rows)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.Nil(t, err)
				value, ok := result.(*dto.CountryConfig)
				assert.True(t, ok)
				assert.NotNil(t, value)
				assert.Equal(t, 2, len(value.Whitelist))
			},
		},
		{
			caseName: "failed - db",
			params: params{
				name: constant.CountryConfig,
				out:  &dto.CountryConfig{},
			},
			expectations: func(params params) {
				c, _ := json.Marshal(&dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				})
				rows := sqlmock.NewRows([]string{"config"})
				rows.AddRow(string(c))

				m.DBMock.ExpectQuery(query).WithArgs(params.name).WillReturnError(testutil.DefaultErr)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - convert config",
			params: params{
				name: constant.CountryConfig,
				out:  &dto.CountryConfig{},
			},
			expectations: func(params params) {
				rows := sqlmock.NewRows([]string{"config"})
				rows.AddRow("not json")

				m.DBMock.ExpectQuery(query).WithArgs(params.name).WillReturnRows(rows)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)

			err := repo.GetSystemConfigByName(context.Background(), testCase.params.name, &testCase.params.out)
			testCase.results(t, testCase.params.out, err)
		})
	}
}

func TestUpdateSystemConfig(t *testing.T) {
	m := testutil.InitMock(t)
	defer func() {
		m.DB.Master.Close()
		m.Controller.Finish()
	}()

	repo := testutil.InitMockRepository(m)

	var (
		query = "UPDATE system_config SET config = \\?, updated_by = \\?, updated_at = \\? WHERE name = \\?"
	)
	type params struct {
		in      interface{}
		name    constant.SystemConfig
		adminID interface{}
	}

	var testCases = []struct {
		caseName     string
		params       params
		expectations func(params)
		results      func(t *testing.T, result interface{}, err error)
	}{
		{
			caseName: "success",
			params: params{
				name: constant.CountryConfig,
				in: &dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				},
				adminID: testutil.AdminID,
			},
			expectations: func(params params) {
				c, _ := json.Marshal(params.in)
				updatedBy := params.adminID.(string)
				m.DBMock.ExpectExec(query).WithArgs(string(c), updatedBy, testutil.AnyNumber{}, params.name).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.Nil(t, err)
				value, ok := result.(*dto.CountryConfig)
				assert.True(t, ok)
				assert.NotNil(t, value)
				assert.Equal(t, 2, len(value.Whitelist))
			},
		},
		{
			caseName: "failed - db",
			params: params{
				name: constant.CountryConfig,
				in: &dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				},
				adminID: testutil.AdminID,
			},
			expectations: func(params params) {
				c, _ := json.Marshal(params.in)
				updatedBy := params.adminID.(string)
				m.DBMock.ExpectExec(query).WithArgs(string(c), updatedBy, testutil.AnyNumber{}, params.name).WillReturnError(testutil.DefaultErr)
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - update result",
			params: params{
				name: constant.CountryConfig,
				in: &dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				},
				adminID: testutil.AdminID,
			},
			expectations: func(params params) {
				c, _ := json.Marshal(params.in)
				m.DBMock.ExpectExec(query).WithArgs(string(c), testutil.AnyNumber{}, params.name).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - rows",
			params: params{
				name: constant.CountryConfig,
				in: &dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				},
				adminID: testutil.AdminID,
			},
			expectations: func(params params) {
				c, _ := json.Marshal(params.in)
				m.DBMock.ExpectExec(query).WithArgs(string(c), testutil.AnyNumber{}, params.name).WillReturnResult(sqlmock.NewErrorResult(testutil.DefaultErr))
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - empty admin id",
			params: params{
				name: constant.CountryConfig,
				in: &dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				},
				adminID: "",
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			caseName: "failed - invalid admin id type",
			params: params{
				name: constant.CountryConfig,
				in: &dto.CountryConfig{
					Whitelist: []commonconstant.Country{commonconstant.MY, commonconstant.BR},
				},
				adminID: 1,
			},
			expectations: func(params params) {
			},
			results: func(t *testing.T, result interface{}, err error) {
				assert.NotNil(t, err)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseName, func(t *testing.T) {
			testCase.expectations(testCase.params)

			ctx := context.WithValue(context.Background(), constant.AuthorizedAdminIDKey, testCase.params.adminID)
			err := repo.UpdateSystemConfig(ctx, testCase.params.name, testCase.params.in)
			testCase.results(t, testCase.params.in, err)
		})
	}

}

// END __INCLUDE_EXAMPLE__
