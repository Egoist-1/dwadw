package integration

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"io/ioutil"
	"naming/naming/internal/config"
	"naming/naming/internal/integration/startup"
	"naming/naming/internal/server"
	"naming/naming/internal/svc"
	"naming/naming/pb/naming"
	"net/http"
	"testing"
)

type HostSuite struct {
	suite.Suite
	server    *server.NamingServer
	caddyHost string
	db        *gorm.DB
}

//	func (s *HostSuite) SetupSubTest() {
//		err := s.db.Exec("TRUNCATE TABLE `namings`").Error
//		require.NoError(s.T(), err)
//	}
func (s *HostSuite) SetupSuite() {
	cfg := svc.NewServiceContext(config.Config{
		DSN:       "root:root@tcp(localhost:3306)/naming",
		CaddyHost: "http://localhost:2019",
	})
	s.db = startup.InitGorm("root:root@tcp(localhost:3306)/naming")
	s.server = server.NewNamingServer(cfg)
	s.caddyHost = "http://localhost:2019"
}

func (s *HostSuite) TestIncrReadCnt() {
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
		res    *naming.AddHostRes
		req    *naming.AddHostReq
	}{
		{
			name: "添加成功",
			before: func(t *testing.T) {
				// 创建请求
				req, err := http.NewRequest("POST", s.caddyHost+"/load", bytes.NewBuffer([]byte(string(loadCfg))))
				require.NoError(t, err)
				// 设置请求头
				req.Header.Set("Content-Type", "application/json")
				// 发送请求
				client := &http.Client{}
				_, err = client.Do(req)
				require.NoError(t, err)
			},
			after: func(t *testing.T) {
				// 创建请求
				req, err := http.NewRequest("GET",
					s.caddyHost+"/config/apps/http/servers/myserver/routes/0/match/0/host", nil)
				require.NoError(t, err)
				client := &http.Client{}
				res, err := client.Do(req)
				resp, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				actual := make([]string, 10)
				err = json.Unmarshal(resp, &actual)
				assert.NoError(t, err)
				expected := []string{
					"example.com",
					"example2.com",
				}
				assert.Equal(s.T(), expected, actual)
			},
			res: &naming.AddHostRes{},
			req: &naming.AddHostReq{
				Host: []string{"example2.com"},
			},
		},
	}

	// 不需要每个测试都创建
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.before(s.T())
			res, err := s.server.AddHost(context.Background(), tc.req)
			assert.NoError(s.T(), err)
			assert.Equal(s.T(), tc.res, res)
			tc.after(s.T())
		})
	}
}

func TestAddHostService(t *testing.T) {
	suite.Run(t, &HostSuite{})
}
