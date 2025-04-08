// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package types

type AddHostReq struct {
	Host []string `json:"host"`
}

type AddHostRes struct {
	Res
}

type LoadByjsonReq struct {
	LoadJson string `json:"loadJson"` // 或 `json:"loadJson,omitempty"`（如果字段可选）
}

type LoadByjsonRes struct {
	Res
}

type Res struct {
	Code int
	Data string
	Msg  string
}

type UpdateHostReq struct {
	OrgHost string `json:"orgHost"`
	ExpHost string `json:"expHost"`
}

type UpdateHostRes struct {
	Res
}
