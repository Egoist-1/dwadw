package repo

import (
	"context"
	"naming/naming/internal/repo/dao"
)

func NewNamingRepo(dao dao.NamingDAO) NamingRepo {
	return &namingRepo{dao: dao}
}

type NamingRepo interface {
	StoreHost(ctx context.Context, hosts string) error
	FindByHost(ctx context.Context, host string) (int64, error)
	UpdateHost(ctx context.Context, org string, exp string) error
}
type namingRepo struct {
	dao dao.NamingDAO
}

func (repo *namingRepo) UpdateHost(ctx context.Context, org string, exp string) error {
	return repo.dao.UpdateHost(ctx, org, exp)
}

func (repo *namingRepo) FindByHost(ctx context.Context, host string) (int64, error) {
	return repo.dao.FindByHost(ctx, host)
}

func (repo *namingRepo) StoreHost(ctx context.Context, hosts string) error {
	return repo.dao.StoreHost(ctx, hosts)
}
