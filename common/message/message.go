package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

// User Status constant
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
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

	UserIds []int `json:"userIds"` // For client to know which users are on-line
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// Server can send the message of change of user status
// Server send message "Actively"
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
