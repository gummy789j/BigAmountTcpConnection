package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"` // Message Type
	Data string `json:"data"` // Message Data
}

type LoginMes struct {
	UserId   int    `json:"userId"`   // user id
	UserPwd  string `json:"userPwd"`  // user password
	UserName string `json:"userName"` // user name
}

type LoginResMes struct {
	Code int `json:"code"`
	// return status code 500 => user doesn't sign up
	// 200 => log in Success
	Error string `json:"error"` // return error message
}

type RegisterMes struct {
}
