package naming

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NamingModel = (*customNamingModel)(nil)

type (
	// NamingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNamingModel.
	NamingModel interface {
		namingModel
	}

	customNamingModel struct {
		*defaultNamingModel
	}
)

// NewNamingModel returns a model for the database table.
func NewNamingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NamingModel {
	return &customNamingModel{
		defaultNamingModel: newNamingModel(conn, c, opts...),
	}
}
