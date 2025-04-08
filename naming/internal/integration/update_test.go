package integration

import (
	"context"
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"naming/naming/internal/config"
	"naming/naming/internal/integration/startup"
	"naming/naming/internal/server"
	"naming/naming/internal/svc"
	"naming/naming/pb/naming"
	"testing"
)

type UpdateHost struct {
	suite.Suite
	server    *server.NamingServer
	caddyHost string
	db        *gorm.DB
}

func (s *UpdateHost) SetupSuite() {
	cfg := svc.NewServiceContext(config.Config{
		DSN:       "root:root@tcp(localhost:3306)/naming",
		CaddyHost: "http://localhost:2019",
	})
	s.db = startup.InitGorm("root:root@tcp(localhost:3306)/naming")
	s.server = server.NewNamingServer(cfg)
	s.caddyHost = "http://localhost:2019"
}

func (s *UpdateHost) TestIncrReadCnt() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
		res    *naming.UpdateHostRes
		req    *naming.UpdateHostReq
	}{
		{
			name: "添加成功",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
			res: &naming.UpdateHostRes{},
			req: &naming.UpdateHostReq{
				OrgHost: "example2.com",
				ExpHost: "23123,com",
			},
		},
	}

	// 不需要每个测试都创建
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.before(s.T())
			res, err := s.server.UpdateHost(context.Background(), tc.req)
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tc.res, res)
			tc.after(s.T())
		})
	}
}

func TestUpdateService(t *testing.T) {
	suite.Run(t, &UpdateHost{})
}
