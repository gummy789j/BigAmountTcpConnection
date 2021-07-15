package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("User doesn't exit...")
	ERROR_USER_EXISTS    = errors.New("User already exit...")
	ERROR_USER_PWD       = errors.New("User Password incorrect...")
)
