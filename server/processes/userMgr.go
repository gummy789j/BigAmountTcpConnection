package processes

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// Add online user (also a Update operation)
func (this *UserMgr) AddonlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// Delete
func (this *UserMgr) DelOnlineUser(up *UserProcess) {
	delete(this.onlineUsers, up.UserId)
}

// Return current online users
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// According to userId to return the specified UserProcess
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	if up, ok := this.onlineUsers[userId]; !ok {
		err = fmt.Errorf("userId: %d doesn't exsit", userId)
		return up, err
	}
	return up, err

}
