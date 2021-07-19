package model

import (
	"net"

	"github.com/gummy789j/Multi-Users_ChatRoom/common/message"
)

type CurUser struct {
	Conn net.Conn

	message.User
}
