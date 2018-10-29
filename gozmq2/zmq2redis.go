package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func init() {
	log.SetFlags(log.LUTC | log.Ldate | log.Lshortfile)
}

var redisDB *redis.Client

func main() {
	redisDB = redis.NewClient(&redis.Options{
		Addr: "192.168.1.119:6379",
		DB:   1,
	})
	pong, err := redisDB.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	if pong != "PONG" {
		os.Exit(1)
	}
	result, err := redisDB.RPop("queue:23220").Result()
	if err != nil {
		log.Fatal(err)
	}
	var head, body []byte
	var msghead []interface{}
	var msgbody map[string]interface{}
	// mh := new(codec.MsgpackHandle)
	// mh.RawToString = true
	result = strings.Trim(result, "[] ")
	infos := strings.Split(result, ",")
	infos[0] = strings.Trim(infos[0], `" '`)
	infos[1] = strings.Trim(infos[1], `"' `)
	fmt.Println(infos[1])
	head = []byte(infos[0])
	body = []byte(infos[1])
	// dec := codec.NewDecoderBytes(head, mh)
	// dec.Decode(&msghead)
	// dec = codec.NewDecoderBytes(body, mh)
	// dec.Decode(&msgbody)
	msgpack.Unmarshal(body, &msgbody)
	msgpack.Unmarshal(head, &msghead)

	// fmt.Println(msgbody, msghead)
	fmt.Printf("%v %v\n", msgbody,msghead)
}
