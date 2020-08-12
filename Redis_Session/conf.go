package Redis_Session

var Redis_address string
var Consumer_address string
var Server_address string

func init() {
	Redis_address = "127.0.0.1:6379"
	Server_address = "127.0.0.1:2048"
}

func Set_RedisAddress(redis_address string) {
	Redis_address = redis_address
}
func Set_ConsumerAddress(consumer_address string) {
	Consumer_address = consumer_address
}
func Set_ServerAddress(sever_address string) {
	Server_address = Server_address
}
