package database
import "github.com/gomodule/redigo/redis"
//import "time"

var Pool *redis.Pool

func init() {
  Pool = &redis.Pool{
  	//MaxActive: 0,
    //MaxIdle: 5,
    //IdleTimeout: 5 * time.Second,
    //Wait: true,
    Dial: func () (redis.Conn, error) {return redis.Dial("tcp", ":6379")},
  }
}

