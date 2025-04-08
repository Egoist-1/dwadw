package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewNamingGORM(db *gorm.DB) NamingDAO {
	return &namingGorm{db: db}
}

type NamingDAO interface {
	StoreHost(ctx context.Context, hosts []string) error
	FindByHost(ctx context.Context, host string) (int64, error)
	UpdateHost(ctx context.Context, org string, exp string) error
}
type namingGorm struct {
	db *gorm.DB
}

func (dao *namingGorm) UpdateHost(ctx context.Context, org string, exp string) error {
	err := dao.db.WithContext(ctx).Model(&Naming{}).
		Select("naming").
		Where("naming = ?", org).
		Updates(map[string]interface{}{
			"naming": exp,
		}).Error
	return err
}

func (dao *namingGorm) FindByHost(ctx context.Context, host string) (int64, error) {
	var naming Naming
	err := dao.db.WithContext(ctx).Where("naming = ?", host).First(&naming).Error
	return naming.Id, err
}

func (dao *namingGorm) StoreHost(ctx context.Context, hosts []string) error {
	now := time.Now().UnixMilli()
	var err error
	for _, v := range hosts {
		err = dao.db.WithContext(ctx).Create(&Naming{
			Naming: v,
			Ctime:  now,
			Utime:  now,
		}).Error
		if err != nil {
			return dao.duplicate(err)
		}
	}
	return err
}
func (dao *namingGorm) duplicate(err error) error {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		// MySQL错误码1062表示唯一冲突
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 返回自定义的唯一冲突错误
			return errors.New("域名已经存在")
		}
	}
	return err
}

type Naming struct {
	Id     int64  `gorm:"primaryKey;autoIncrement;"`
	Naming string `gorm:"varchar(255);unique"`
	//域名在caddy的下标
	Ctime int64
	Utime int64
}
