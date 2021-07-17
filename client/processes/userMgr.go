package processes

import (
	"fmt"
	"go_code/MultiusersChatRoom/client/model"
	"go_code/MultiusersChatRoom/common/message"
)

// Client have to handle these map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func update(mes *message.NotifyUserStatusMes) {

	// optimize effectiveness
	if user, ok := onlineUsers[mes.UserId]; !ok {

		user = &message.User{
			UserId:     mes.UserId,
			UserStatus: mes.Status,
		}

		onlineUsers[mes.UserId] = user

	} else {
		onlineUsers[mes.UserId].UserStatus = mes.Status

	}

	outputOnlineUser()

	return

}

func outputOnlineUser() {
	fmt.Println("On-line Users List: ")

	//range can also iterate over just the keys of a map.
	for id := range onlineUsers {
		fmt.Println("User ID:\t", id)
	}
}
