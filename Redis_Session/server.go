package Redis_Session

import (
	"fmt"
	"log"
	"net"
	"net/http"
	rpc "net/rpc"
	"os"

	redis "github.com/garyburd/redigo/redis"
)

//	"log"
//	"net/rpc"
var (
	pool *redis.Pool
)

// 这里是用于获取一个新的pool
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", Redis_address)
		},
	}
}

//这里是用于注册rpc服务
type Session_struct struct {
}

func (this *Session_struct) Set_session(res Session, t *bool) error {
	rc := pool.Get()
	defer rc.Close()
	_, e := rc.Do("setex", string(res.Session_ID[:]), res.Expires_time, res.Content)
	if e != nil {
		*t = false
		return e
	}
	*t = true
	return nil
}

func (this *Session_struct) Get_Session(sess_id string, content *string) error {
	rc := pool.Get()
	*content = ""
	defer rc.Close()
	var err error
	*content, err = redis.String(rc.Do("GET", sess_id))
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
func Server_Init() {
	rpc.Register(new(Session_struct)) //注册服务
	rpc.HandleHTTP()                  //采用HTTP作为载体

	lis, err := net.Listen("tcp", Server_address)
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection")
	http.Serve(lis, nil)
}
func init() {
	pool = NewPool()
}
