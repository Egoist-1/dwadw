package e

type Err struct {
	//错误码
	code errCode
	//日志定位
	msg string
}

func (e Err) Error() string {
	return string(e.code) + e.msg

}

// 日志定位
func (e Err) Log() string {
	return e.msg
}
func (e Err) Code() errCode {
	return e.code
}

// msg 定位作用
func NewErr(code errCode, msg string, errToString string) error {
	return Err{
		code: code,
		msg:  msg + errToString,
	}
}
func NewServerErr(msg string, errToString string) error {
	return Err{
		code: ServerErr,
		msg:  msg + errToString,
	}
}

type errCode int

func (e errCode) ToInt() int {
	return int(e)
}

//go:generate stringer -type errCode -linecomment
const (
	ServerErr errCode = 50001 //系统错误
)

// Caddy
const (

)

