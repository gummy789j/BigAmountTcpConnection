package model

import (
	"encoding/json"

	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/gummy789j/Multi-Users_ChatRoom/common/message"
)

// we can use this gobal variable to imporve our efficiency when we access redis
// When the server process start running, we use the pool which is build before to new a UserDao
// And put into this global variable, we can save a lot of time from accessing redis and build a UserDao
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

// Factory mode, build a constructor(建構式)
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}

	return

}

// Read (R in CRUD)
// return an User struct and error
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	res, err := redis.String(conn.Do("HGET", "users", id))

	if err != nil {
		if err == redis.ErrNil { // redis.ErrNil means we cannot find the corresponding id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		log.Println("json.Unmarshal Fail err =", err)
		return
	}

	return

}

// To verify the userid and userpwd for Login
// if userid and userpwd both are correct , return an User struct
// if userid or userpwd either one is incorrect , return error message

func (this *UserDao) Login(userid int, userpwd string) (user *User, err error) {

	conn := this.pool.Get()

	defer conn.Close()

	user, err = this.getUserById(conn, userid)
	if err != nil {
		return
	}

	if user.UserPwd != userpwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	conn := this.pool.Get()

	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)

	if err == nil {
		err = ERROR_USER_EXISTS
		return

	}

	data, err := json.Marshal(user)
	if err != nil {

		return
	}

	_, err = conn.Do("HSET", "users", user.UserId, data)
	if err != nil {

		log.Println("Redis HSET error=", err)

		return
	}

	return

}
