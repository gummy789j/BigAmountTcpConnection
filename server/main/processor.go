package main

import (
	"go_code/MultiusersChatRoom/common/message"
	"go_code/MultiusersChatRoom/server/processes"
	"go_code/MultiusersChatRoom/server/utils"
	"io"
	"log"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// This is a function for making a decision for which server response should be sent back
// According the client message type
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {

	case message.LoginMesType:
		// Handle the log in
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// Handle the sign up
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		log.Println("Message Type doesn't exit.....")
	}
	return
}

func (this *Processor) processing() error {

	for {

		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()

		if err != nil {

			if err == io.EOF {
				log.Println("the client has exited, server exit....")
				return err
			}

			log.Println("reading pkg err = ", err)
			return err
		}

		//fmt.Println("mes =", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
