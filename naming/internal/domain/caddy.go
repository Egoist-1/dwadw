package domain

type CaddyInfo struct {
	Apps Apps `json:"apps"`
}

type Apps struct {
	HTTP HTTP `json:"http"`
}

type HTTP struct {
	Servers map[string]Server `json:"servers"`
}

type Server struct {
	Listen []string `json:"listen"`
	Routes []Route  `json:"routes"`
}

type Route struct {
	Match  []Match  `json:"match"`
	Handle []Handle `json:"handle"`
}

type Match struct {
	Host []string `json:"host,omitempty"`
	// 可以添加其他匹配条件字段
}

type Handle struct {
	Handler string `json:"handler"`
	// 根据不同的handler可以添加额外字段
	// 例如 file_server 可能有 root 字段
	Root string `json:"root,omitempty"`
}
