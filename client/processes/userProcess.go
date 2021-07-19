package processes

import (
	"encoding/json"
	"fmt"

	"log"
	"net"

	"github.com/gummy789j/Multi-Users_ChatRoom/client/utils"
	"github.com/gummy789j/Multi-Users_ChatRoom/common/message"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	// starting set the protocol
	// 	fmt.Printf("userId=%d userPwd= %s", userId, userPwd)

	// 	return nil

	// 1. Build the connection with server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		log.Println("net.Dial err=", err)
		return
	}
	// close conn in the end
	defer conn.Close()

	// 2. Ready to send the message to server through conn
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3. Build a LoginMes structure

	loginMes := message.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}

	// 4. let logMes serialize
	// data is type []byte originally
	data, err := json.Marshal(loginMes)
	if err != nil {
		log.Println("json.Marshal loginMes error", err)
		return
	}

	// 5. change data type to string
	mes.Data = string(data)

	// 6. let mes serialize
	data, err = json.Marshal(mes)
	if err != nil {
		log.Println("json.Marshal mes error", err)
		return
	}

	// 7.8. writing the data to the server
	// Build a transfer structure

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data) // writePkg function in the utils.go
	if err != nil {
		log.Println("writePkg Fail err=", err)
	}

	//fmt.Printf("The message from client is %v\n", string(data)) //"{\"userId\":100,\"userPwd\":\"abc\",\"userName\":\"Steven\"}"

	// 9. We have to process the return message(LoginResMes)
	mes, err = tf.ReadPkg() // readPkg function in the utils.go
	if err != nil {
		log.Println("readPkg Fail err=", err)
	}
	// 10. decoding the mes.Data
	var loginResMes message.LoginResMes

	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		log.Println("json.Unmarshal(mes.Data) Fail err=", err)
	}

	// 11. identify the status code to know whether log in success
	if loginResMes.Code == 200 {
		//log.Println("Login Success")

		// initialize CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		fmt.Println("Users are on-line in current on the following list:")

		for _, id := range loginResMes.UserIds {

			if id == userId {
				continue
			}

			fmt.Println("User ID:\t", id)

			user := &message.User{
				UserId:     id,
				UserStatus: message.UserOnline,
			}

			onlineUsers[id] = user
		}
		fmt.Print("\n\n")

		go serverProcessMes(conn)

		ShowMenu()

	} else {
		log.Println(loginResMes.Error)
	}

	return

}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	// starting set the protocol
	// 	fmt.Printf("userId=%d userPwd= %s", userId, userPwd)

	// 	return nil

	// 1. Build the connection with server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		log.Println("net.Dial err=", err)
		return
	}
	// close conn in the end
	defer conn.Close()

	// 2. Ready to send the message to server through conn
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3. Build a registerMes structure

	userMes := message.User{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}

	registerMes := message.RegisterMes{
		User: userMes,
	}

	// 4. let logMes serialize
	// data is type []byte originally
	data, err := json.Marshal(registerMes)
	if err != nil {
		log.Println("json.Marshal registerMes error", err)
		return
	}

	// 5. change data type to string
	mes.Data = string(data)

	// 6. let mes serialize
	data, err = json.Marshal(mes)
	if err != nil {
		log.Println("json.Marshal mes error", err)
		return
	}

	// 7.8. writing the data to the server
	// Build a transfer structure

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data) // writePkg function in the utils.go
	if err != nil {
		log.Println("writePkg Fail err=", err)
	}

	fmt.Printf("The message from client is %v\n", string(data))

	// 9. We have to process the return message(LoginResMes)
	mes, err = tf.ReadPkg() // readPkg function in the utils.go
	if err != nil {
		log.Println("readPkg Fail err=", err)
	}
	// 10. decoding the mes.Data
	var registerResMes message.RegisterResMes

	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		log.Println("json.Unmarshal(mes.Data) Fail err=", err)
	}

	// 11. identify the status code to know whether log in success
	if registerResMes.Code == 200 {
		//log.Println("Login Success")
		log.Println("Registration Success. Please re-login")

	} else {
		log.Println(registerResMes.Error)
	}

	return

}
