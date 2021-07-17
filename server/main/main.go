package main

import (
	"fmt"
	"go_code/MultiusersChatRoom/server/model"
	"log"
	"net"
	"time"
)

// // This is a function for reading the client/server message
// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	// build a buffer as the container for any readed data
// 	buffer := make([]byte, 8096)
// 	//fmt.Println("Reading the client's data....")

// 	// read the length of message
// 	n, err := conn.Read(buffer[:4])
// 	if n != 4 || err != nil {
// 		//err = errors.New("read pkg header error")
// 		return
// 	}

// 	// decoding the length of message from binary to uint32
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buffer[:4])

// 	// using the given length to read the real message
// 	n, err = conn.Read(buffer[:int(pkgLen)])
// 	if n != int(pkgLen) || err != nil {
// 		err = errors.New("read pkg body error")
// 		return
// 	}

// 	// decoding the real message from json to message.Message(it's a Structure defined in the common)
// 	err = json.Unmarshal(buffer[:int(pkgLen)], &mes)
// 	if err != nil {
// 		log.Println("json.Unmarshal Fail err =", err)
// 	}

// 	return
// }

// // This is a function for writing the client/server message
// func writePkg(conn net.Conn, data []byte) (err error) {

// 	// sending the length of response message to client
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))
// 	var buffer [4]byte
// 	binary.BigEndian.PutUint32(buffer[:4], pkgLen)
// 	n, err := conn.Write(buffer[:4])
// 	if n != 4 || err != nil {
// 		log.Println("conn.Write(buffer) Fail")
// 		return
// 	}

// 	// And then, we have to send the real response message to the client
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		log.Println("conn.Write(data) Fail")
// 		return
// 	}

// 	return
// }

// // This is a function for handling the request of log in
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

// 	// decoding the mes.Data from mes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		log.Println("json.Unmarshal Fail err=", err)
// 		return
// 	}

// 	// build a respose message
// 	var resMes message.Message
// 	resMes.Type = message.LoginResMesType

// 	// build a loginResMes
// 	var loginResMes message.LoginResMes

// 	// identify the id， if id = 100 & password = abc, it's legal
// 	if loginMes.UserId == 100 && loginMes.UserPwd == "abc" {
// 		// legal
// 		loginResMes.Code = 200 //status code 200 represent success
// 	} else {
// 		// illegal
// 		loginResMes.Code = 500 //status code 500 represent unsuccess
// 		loginResMes.Error = "This user doesn't exit. Please log in again after sign up...."
// 	}

// 	// serialize the loginResMes
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		log.Println("json.Marshal Fail err =", err)
// 		return
// 	}

// 	// store in the resMes.Data
// 	resMes.Data = string(data)

// 	// serialize the resMes and ready to send
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		log.Println("json.Marshal Fail err =", err)
// 		return
// 	}

// 	// writing the reponse message to client
// 	err = utils.WritePkg(conn, data)
// 	return
// }

// // This is a function for making a decision for which server response should be sent back
// // According the client message type
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

// 	switch mes.Type {

// 	case message.LoginMesType:
// 		// Handle the log in
// 		err = serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		// Handle the sign up
// 	default:
// 		log.Println("Message Type doesn't exit.....")
// 	}
// 	return
// }

func process(conn net.Conn) {
	// read the message form the client
	defer conn.Close()

	processor := &Processor{
		Conn: conn,
	}
	err := processor.processing()
	if err != nil {
		log.Println("This goroutine occur error =", err)
		return
	}
	// for {
	// 	mes, err := utils.ReadPkg(conn)

	// 	if err != nil {

	// 		if err == io.EOF {
	// 			log.Println("the client has exited, server exit....")
	// 			return
	// 		}

	// 		log.Println("reading pkg err = ", err)
	// 		return
	// 	}

	// 	//fmt.Println("mes =", mes)
	// 	err = serverProcessMes(conn, &mes)
	// 	if err != nil {
	// 		return
	// 	}
	// }
}

// To initialize an UserDao
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {

	// When server start running, we initialize our redis connection pool
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	fmt.Println("Server is listening port:8889....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		log.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	// Once listening success, waiting for the client to connect server

	for {
		fmt.Println("Waiting clients to connect the server.....")
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listen Accept err=", err)
		}

		// Once connection success, open a goroutine and keep communicating with the client

		fmt.Println("hello")
		go process(conn)
	}
}
