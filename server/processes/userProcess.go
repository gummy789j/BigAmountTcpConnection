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

	// Add a content to represent the User Id for this user Conn
	UserId int
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

		// Since users are log in Suceess
		// We have to record userProcess to the userMgr
		this.UserId = loginMes.UserId

		userMgr.AddonlineUser(this)

		// Notify other users that you are on-line
		this.NotifyOthersOnlineUser(loginMes.UserId)

		// return online users list to client
		for id, _ := range userMgr.onlineUsers {

			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

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

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	// decoding the mes.Data from mes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		log.Println("json.Unmarshal Fail err=", err)
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {

		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error() // err.Error is making error to string type
		} else {
			registerResMes.Code = 50
			registerResMes.Error = "Register occur unknown ERROR...."
		}

	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.User.UserId, "Register Success")
	}

	data, err := json.Marshal(registerResMes)
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

// Notify other users you are on-line
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	for id, up := range userMgr.onlineUsers {

		if id == userId {
			continue
		}

		up.Notify(userId)
	}
}

// Implement Notify process
func (this *UserProcess) Notify(userId int) {

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.Status = message.UserOnline
	notifyUserStatusMes.UserId = userId

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		log.Println("json.Marshal Fail err=", err)
		return
	}

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		log.Println("json.Marshal Fail err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		log.Println("writePkg Fail err =", err)
		return
	}

	return

}
