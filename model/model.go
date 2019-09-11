package model

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BaseModel struct {
	gorm.Model
}

type UserInfo struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"say_hello"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
