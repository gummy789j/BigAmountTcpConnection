package processes

import (
	"fmt"
	"go_code/MultiusersChatRoom/common/message"
)

func outputGroupMes(mes *message.SmsMes) {
	info := fmt.Sprintf("User ID:\t%d Talking to Group: \t%s", mes.UserId, mes.Content)

	fmt.Println(info)
	fmt.Println()
}
