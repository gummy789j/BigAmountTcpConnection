package processes

import (
	"encoding/json"
	"fmt"
	"go_code/MultiusersChatRoom/client/utils"
	"go_code/MultiusersChatRoom/common/message"
	"log"
	"net"
	"os"
)

func ShowMenu() {
	for {
		fmt.Println("-------------Congradulation xxx Log in SUCCESS----------")
		fmt.Println("-------------1. Show the list of users on-line----------")
		fmt.Println("-------------2. Send message----------")
		fmt.Println("-------------3. List of messages----------")
		fmt.Println("-------------4. Exit system----------")
		fmt.Println("Plese choose(1-4): ")

		var key int
		var content string
		fmt.Scanf("%d\n", &key)

		smsProcess := &SmsProcess{}

		switch key {
		case 1:
			//fmt.Println("1. Show the list of users on-line-")
			outputOnlineUser()
		case 2:
			fmt.Println("What do you want to say to all users:")

			fmt.Scanf("%s\n", &content)

			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("3. List of messages")
		case 4:
			fmt.Println("4. You choose to exit system....")
			os.Exit(0)
		default:
			fmt.Println("Unavailable choice....")

		}
		//os.Exit(0)
	}
}

func serverProcessMes(conn net.Conn) {

	// Build a Transfer structure to read message from server
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("Client keep waiting for the message from server")

		mes, err := tf.ReadPkg()
		if err != nil {
			log.Println("serverProcessMes tf.Reading err=", err)
		}

		//fmt.Printf("mes=%v\n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: // somobody on-line
			// 1. get the NotifyUserStatusMes

			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				log.Println("notifyUserStatusMes json.Unmarshal Fail err =", err)
			}

			// 2. store this message to the all users map sturture which is handled by client themself
			update(&notifyUserStatusMes)

		case message.SmsMesType:

			var smsMes message.SmsMes
			err = json.Unmarshal([]byte(mes.Data), &smsMes)
			if err != nil {
				log.Println("smsMes json.Unmarshal Fail err =", err)
			}
			outputGroupMes(&smsMes)

		default:
			log.Println("Server return unkonwn message type.....")
		}
	}
}
