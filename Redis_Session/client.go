package Redis_Session

import (
	"fmt"
	"log"
	rpc "net/rpc"
)

//var conn rpc.
func Set_session(sess Session) bool {

	conn, err := rpc.DialHTTP("tcp", Server_address)
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	defer conn.Close()
	var t bool
	err = conn.Call("Session_struct.Set_session", sess, &t)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Println(t)
	return t
}

func Get_session(sessid string) string {

	conn, err := rpc.DialHTTP("tcp", Server_address)
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}
	defer conn.Close()
	var res string
	err = conn.Call("Session_struct.Get_session", string(sessid), &res)
	return res
}
