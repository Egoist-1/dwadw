package integration

import (
	"context"
	_ "embed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"io/ioutil"
	"naming/internal/config"
	"naming/internal/integration/startup"
	"naming/internal/server"
	"naming/internal/svc"
	"naming/pb/naming"
	"net/http"
	"testing"
)

type LoadSuite struct {
	suite.Suite
	server    *server.NamingServer
	caddyHost string
	db        *gorm.DB
}

//func (s *LoadSuite) SetupSubTest() {
//	err := s.db.Exec("TRUNCATE TABLE `namings`").Error
//	require.NoError(s.T(), err)
//}

func (s *LoadSuite) SetupSuite() {
	cfg := svc.NewServiceContext(config.Config{
		DSN:       "root:root@tcp(localhost:3306)/naming",
		CaddyHost: "http://localhost:2019",
	})
	s.db = startup.InitGorm("root:root@tcp(localhost:3306)/naming")
	s.server = server.
		NewNamingServer(cfg)
	s.caddyHost = "http://localhost:2019"
}
func (s *LoadSuite) TestTRUNCATE() {
	err := s.db.Exec("TRUNCATE TABLE `namings`").Error
	require.NoError(s.T(), err)
}
func (s *LoadSuite) TestIncrReadCnt() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
		res    *naming.LoadRes
		req    *naming.LoadReq
	}{
		{
			name: "配置成功",
			before: func(t *testing.T) {
			},
			after: func(t *testing.T) {
				req, err := http.NewRequest("GET",
					s.caddyHost+"/config", nil)
				require.NoError(t, err)
				client := &http.Client{}
				res, err := client.Do(req)
				_, err = ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				//assert.Equal(s.T(), loadCfg, string(body))
			},
			res: &naming.LoadRes{},
			req: &naming.LoadReq{
				Cfg: loadCfg,
			},
		},
	}

	// 不需要每个测试都创建
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.before(s.T())
			res, err := s.server.Load(context.Background(), tc.req)
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tc.res, res)
			tc.after(s.T())
		})
	}
}

func TestLoadService(t *testing.T) {
	suite.Run(t, &LoadSuite{})
}
