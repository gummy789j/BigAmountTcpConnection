package processes

import (
	"encoding/json"
	"go_code/MultiusersChatRoom/common/message"
	"go_code/MultiusersChatRoom/server/utils"
	"log"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content

	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	data, err := json.Marshal(smsMes)
	if err != nil {
		log.Println("SendGroupMes json.Marshal Fail err =", err.Error())
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		log.Println("SendGroupMes json.Marshal Fail err =", err.Error())
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		log.Println("writePkg Fail err=", err.Error())
		return
	}
	return
}
