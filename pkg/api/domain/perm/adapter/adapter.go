package adapter

import (
	"github.com/casbin/casbin/model"
	"zeus/pkg/api/domain/perm/adapter/mysql"
)

type Adapter interface {
	LoadPolicy(model model.Model) error
	SavePolicy(model model.Model) error
	AddPolicy(sec string, ptype string, rule []string) error
	RemovePolicy(sec string, ptype string, rule []string) error
	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
}

// NewMysqlAdapter : delegate of adapter
func NewSqlAdapter() *sqlAdapter {
	ad := &sqlAdapter{
		a: mysql.NewGormAdapter(),
	}
	return ad
}

type sqlAdapter struct {
	a Adapter
}

func (ca *sqlAdapter) LoadPolicy(model model.Model) error { return ca.a.LoadPolicy(model) }
func (ca *sqlAdapter) SavePolicy(model model.Model) error { return ca.a.SavePolicy(model) }
func (ca *sqlAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return ca.a.AddPolicy(sec, ptype, rule)
}
func (ca *sqlAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return ca.a.RemovePolicy(sec, ptype, rule)
}
func (ca *sqlAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return ca.a.RemoveFilteredPolicy(sec, ptype, fieldIndex, fieldValues...)
}
