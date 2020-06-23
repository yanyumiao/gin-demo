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
    Dial: func () (redis.Conn, error) {
      c, err := redis.Dial("tcp", ":6379")
      if err != nil {
        return nil, err
      }
      if _, err := c.Do("AUTH", "123456"); err != nil {
        c.Close()
        return nil, err
      }
      return c, nil
    },
  }
}


