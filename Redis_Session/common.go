package Redis_Session

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	Session_ID   [16]byte //用于a
	Expires_time int
	Content      string
}

func (sess *Session) Get_SessionID() {
	sess.Session_ID = uuid.NewV4()
	fmt.Println(string(sess.Session_ID[:]))
}

func (sess *Session) Set_expires(expires int) {
	sess.Expires_time = expires
}

func (sess *Session) Set_Content(content string) {
	sess.Content = content
}

//func (sess *Session) Get_Session() bool {
//	rc := pool.Get()
//	_, c := rc.Do("set", "ttts", "wes")
//	if c != nil {
//		return false
//	}
//	return true
//}
