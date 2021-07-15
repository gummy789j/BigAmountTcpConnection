package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"go_code/MultiusersChatRoom/common/message"
	"log"
	"net"
)

// Implenment OOP
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

// This is a function for reading the client/server message
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	// build a buffer as the container for any readed data
	//buffer := make([]byte, 8096)
	//fmt.Println("Reading the client's data....")

	// read the length of message
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	// decoding the length of message from binary to uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	// using the given length to read the real message
	n, err = this.Conn.Read(this.Buf[:int(pkgLen)])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	// decoding the real message from json to message.Message(it's a Structure defined in the common)
	err = json.Unmarshal(this.Buf[:int(pkgLen)], &mes)
	if err != nil {
		log.Println("json.Unmarshal Fail err =", err)
	}

	return
}

// This is a function for writing the client/server message
func (this *Transfer) WritePkg(data []byte) (err error) {

	// sending the length of response message to client
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buffer [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		log.Println("conn.Write(buffer) Fail")
		return
	}

	// And then, we have to send the real response message to the client
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		log.Println("conn.Write(data) Fail")
		return
	}

	return
}
