package main

import (
	"encoding/binary"
	"encoding/json"
	"go_code/BigAmountTcpConnection/common/message"
	"log"
	"net"
)

func login(userId int, userPwd string) (err error) {

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

	// 7. Now, mes is a complete message we want to send

	// 7-1. **Before sending the message, we have to send the length of message
	// as the identification to make sure the sending data success or not
	// So that's a question !
	// How to change the len(data) it's a integer to a slice of byte
	// bcuz conn.Write need to send data of []byte type
	// we have to use the package encoding/binary
	// tips : The UInt32 value type represents unsigned integers
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buffer [4]byte
	binary.BigEndian.PutUint32(buffer[:4], pkgLen)
	n, err := conn.Write(buffer[:4])
	if n != 4 || err != nil {
		log.Println("conn.Write(buffer) Fail")
	}

	log.Println("Length of message has been sent, the length is", len(data), "content is", string(data))

	return

}
