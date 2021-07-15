package processes

import (
	"fmt"
	"go_code/MultiusersChatRoom/client/utils"
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
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("1. Show the list of users on-line-")
		case 2:
			fmt.Println("2. Send message")
		case 3:
			fmt.Println("3. List of messages")
		case 4:
			fmt.Println("4. You choose to exit system....")
			os.Exit(0)
		default:
			fmt.Println("Unavailable choice....")

		}
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
			log.Println("tf.Reading err=", err)
		}

		fmt.Printf("mes=%v\n", mes)
	}
}
