package processes

import (
	"encoding/json"
	"fmt"
	"go_code/MultiusersChatRoom/common/message"
	"go_code/MultiusersChatRoom/server/model"
	"go_code/MultiusersChatRoom/server/utils"
	"log"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

// This is a function for handling the request of log in
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {

	// decoding the mes.Data from mes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		log.Println("json.Unmarshal Fail err=", err)
		return
	}

	// build a respose message
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// build a loginResMes
	var loginResMes message.LoginResMes

	// We need to veritfy userid and userpwd on redis
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "Server INSIDE ERROR...."
		}

	} else {
		loginResMes.Code = 200
		fmt.Println(user, "Login Success")
	}

	// // identify the id， if id = 100 & password = abc, it's legal
	// if loginMes.UserId == 100 && loginMes.UserPwd == "abc" {
	// 	// legal
	// 	loginResMes.Code = 200 //status code 200 represent success
	// } else {
	// 	// illegal
	// 	loginResMes.Code = 500 //status code 500 represent unsuccess
	// 	loginResMes.Error = "This user doesn't exit. Please log in again after sign up...."
	// }

	// serialize the loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		log.Println("json.Marshal Fail err =", err)
		return
	}

	// store in the resMes.Data
	resMes.Data = string(data)

	// serialize the resMes and ready to send
	data, err = json.Marshal(resMes)
	if err != nil {
		log.Println("json.Marshal Fail err =", err)
		return
	}

	// writing the reponse message to client

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	return
}
