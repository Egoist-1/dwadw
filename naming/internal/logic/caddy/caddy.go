package caddy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"strconv"
)

func NewCaddy(host string) Caddy {
	return &caddy{
		host: host,
	}
}

type Caddy interface {
	Load(ctx context.Context, l logx.Logger, cfg string) error
	AddHost(ctx context.Context, l logx.Logger, host string) error
	GetAllInfo(ctx context.Context, l logx.Logger) (string, error)
	GetHostInfo(ctx context.Context, l logx.Logger) ([]string, error)
	UpdateHost(ctx context.Context, index int64, host string) error
}

type caddy struct {
	host string
}

func (c *caddy) UpdateHost(ctx context.Context, index int64, host string) error {
	// 创建请求
	j, err := json.Marshal(host)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", c.hostByIndex(index), bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	//defer resp.Body.Close()
	//
	//// 读取响应
	//body, _ := ioutil.ReadAll(resp.Body)
	return err
}

func (c *caddy) GetHostInfo(ctx context.Context, l logx.Logger) ([]string, error) {
	// 基本 GET 请求
	str := make([]string, 0, 10)
	resp, err := http.Get(c.hostUrl())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &str)

	if err != nil {
		fmt.Println("解析失败:", err)
		return nil, err
	}
	return str, err
}

func (c *caddy) Load(ctx context.Context, l logx.Logger, cfg string) error {
	// 创建请求

	req, err := http.NewRequest("POST", c.hostUrl()+"/load", bytes.NewBuffer([]byte(cfg)))
	if err != nil {
		l.Error("caddy 发送http请求失败", err)
		return err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		l.Error("caddy 发送http请求失败", err)
		return err
	}
	//defer resp.Body.Close()
	//
	//// 读取响应
	//body, _ := ioutil.ReadAll(resp.Body)
	return err
}

func (c *caddy) AddHost(ctx context.Context, l logx.Logger, host string) error {

	json, err := json.Marshal(host)
	addHostUrl := c.hostUrl()
	req, err := http.NewRequest("POST", addHostUrl, bytes.NewBuffer(json))
	if err != nil {
		l.Error("caddy 发送http请求失败", err)
		return err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		l.Error("caddy 发送http请求失败", err)
		return err
	}
	//defer resp.Body.Close()
	//
	//// 读取响应
	//body, _ := ioutil.ReadAll(resp.Body)
	return err
}

func (c *caddy) GetAllInfo(ctx context.Context, l logx.Logger) (string, error) {
	// 基本 GET 请求
	resp, err := http.Get(c.host + "/config")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}
func (c *caddy) hostUrl() string {
	return "http://" + c.host + "/config/apps/http/servers/myserver/routes/0/match/0/host"
}
func (c *caddy) hostByIndex(index int64) string {
	i := strconv.FormatInt(index, 10)
	return c.hostUrl() + "/" + i
}
