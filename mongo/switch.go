package main

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	zmq "github.com/pebbe/zmq4"
	"github.com/ugorji/go/codec"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	syslog "log"
	. "mongo/protocol"
	"os"
)

var (
	protocolpath       string
	err                error
	client             *mongo.Client
	recvEndpoint       string
	sendEndpoint       string
	workerNum          int
	hwm                int
	instanceIPS        map[int]map[string][]string
	decodeIdProtoCache map[uint64]*ProtocolConfig
	instanceIdCache    map[uint64]*Instances
	instances          []*Instances
)

const (
	_default_host = "localhost"
	_default_port = 27017
)

func main() {

	app := cli.NewApp()
	app.Version = "0.0.1"
	app.Name = "switch"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "protocoldir,d",
			Value:       "",
			Usage:       "协议目录",
			Destination: &protocolpath,
		},
		cli.StringFlag{
			Name:        "connect,c",
			Value:       "tcp://*:23220",
			Usage:       "发送地址 `endpoint`",
			Destination: &sendEndpoint,
		},

		cli.StringFlag{
			Name:        "bind,b",
			Value:       "tcp://127.0.0.1:23200",
			Usage:       "接收地址 `endpoint`",
			Destination: &recvEndpoint,
		},
		cli.IntFlag{
			Name:        "hwm",
			Value:       100000,
			Usage:       "zmq hwm",
			Destination: &hwm,
		},
		cli.IntFlag{
			Name:        "n",
			Value:       1,
			Usage:       "worker的个数",
			Destination: &workerNum,
		},
	}
	app.Action = Run

	if err = app.Run(os.Args); err != nil {
		syslog.Fatal(err)
	}
}

func Run(c *cli.Context) error {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	if protocolpath == "" {
		protocolpath = "/root/protocol"
		//return cli.NewExitError("协议目录未指定", -1)
	}

	ctx := context.Background()

	instances = GetInstanceId(ctx)

	soc := newSoeckt(recvEndpoint, hwm, true, zmq.PULL)
	producerChan := make(chan map[string]interface{}, hwm)
	consumerChan := make(chan map[string]interface{}, hwm)
	go zmqHandler(soc, producerChan) //接收数据
	go protoHanlder(ctx, producerChan, consumerChan)
	//for v := range consumerChan {
	//	fmt.Println(v)
	//}
	for {
		select {
		case v := <-consumerChan:
			fmt.Println(v)
		}
	}
	return nil
}
func GetProtocolName(t *Topov, intfName string) string {
	// TODO 查找intf对应的协议,python版本中还不清除是怎么实现的.
	//for _,a:=range t.Areas{
	//	for _,l:=range a.Levels{
	//		for _,comp:=range l.Components{
	//			for _,conn:=range comp.Connprops{
	//				return conn.Protocol
	//			}
	//		}
	//	}
	//}
	return "cups"
}
func GetInstanceId(ctx context.Context) (result []*Instances) {

	coll := getMongdodb(ctx, _default_host, _default_port, "cbms", "instance_id")
	curs, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err)
	}
	var instanceId Instances
	for curs.Next(ctx) {
		if err = curs.Decode(&instanceId); err != nil {
			log.Error(err)
		}
		result = append(result, &instanceId)
	}
	return

}

func getMongdodb(ctx context.Context, host string, port int, dbName string, collectionName string) *mongo.Collection {
	if client == nil {
		client, err = mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)))
		if err != nil {
			log.Fatal(err)
		}
		if err = client.Connect(ctx); err != nil {
			log.Fatal(err)
		}
		if err = client.Ping(ctx, nil); err != nil {
			log.Fatal(err)
		}
	}

	var db *mongo.Database
	db = client.Database(dbName)

	coll := db.Collection(collectionName)
	return coll
}
func GetTopoV(ctx context.Context) *Topov {
	coll := getMongdodb(ctx, _default_host, _default_port, "cbms", "main_app_datapath")
	cur := coll.FindOne(ctx, bson.M{"state": "apply"}) //FIXME 这里应该用find
	var result Topov
	fmt.Println(cur)
	if err = cur.Decode(&result); err != nil {
		log.Error(err)
	}
	return &result

}
func newSoeckt(endpoint string, hwm int, bind bool, t zmq.Type) *zmq.Socket {

	ctx, err := zmq.NewContext()
	if err != nil {
		log.Error(err)
	}
	socket, err := ctx.NewSocket(t)
	if err != nil {
		log.Error(err)
	}

	if bind {
		err = socket.SetRcvhwm(hwm)
		if err != nil {
			log.Error(err)
		}
		if err = socket.Bind(endpoint); err != nil {
			log.Error(err)

		}
	} else {
		err = socket.SetSndhwm(hwm)
		if err != nil {
			log.Error(err)
		}
		if err = socket.Connect(endpoint); err != nil {
			log.Error(err)
		}

	}

	return socket
}
func zmqHandler(soc *zmq.Socket, msgch chan<- map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	var msgbody map[string]interface{}
	mh := new(codec.MsgpackHandle)
	mh.RawToString = true
	mh.WriteExt = true
	mh.MapValueReset = true
	mh.InterfaceReset  = true
	mh.SliceElementReset = true

	for {

		parts, err := soc.RecvMessageBytes(0)
		if err != nil {
			log.Error(err)
		}
		// head := parts[0]   头部信息,其实可以不需要,所以这里进行舍弃
		body := parts[1]

		dec := codec.NewDecoderBytes(body, mh)
		err = dec.Decode(&msgbody)
		//msgch <- DeepCopy(msgbody).(map[string]interface{}) // 这里进行了拷贝,性能较差.
		msgch <- msgbody
	}
}

// TODO 这里只做了协议的转码,但是判断双向交易目前还没有办法做
// protoHanlder  进行协议转码
func protoHanlder(ctx context.Context, bodych <-chan map[string]interface{}, msgch chan<- map[string]interface{}) {
	for {
		select {
		case body := <-bodych:
			DecodeId, ok := body["DecodeId"].(uint64)
			if !ok {
				continue
			}
			p, ok := decodeIdProtoCache[DecodeId]
			if !ok {
				for _, v := range instances {
					if v.InstanceID == DecodeId {
						p = NewProtocolConfig(protocolpath, GetProtocolName(GetTopoV(ctx), v.IntfName))
						if decodeIdProtoCache == nil {
							decodeIdProtoCache = make(map[uint64]*ProtocolConfig)
						}
						decodeIdProtoCache[DecodeId] = p
						break
					}
				}

			}

			msgch <- p.Process(body)
		}
	}
}
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}
