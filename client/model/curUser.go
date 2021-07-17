package model

import (
	"go_code/MultiusersChatRoom/common/message"

	"net"
)

type CurUser struct {
	Conn net.Conn

	message.User
}
