package main

/* import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	zmq "github.com/alecthomas/gozmq"
	"github.com/clbanning/mxj"
	//  "gopkg.in/vmihailenco/msgpack.v2"
)

type Map map[string]interface{}

const DefaultHighWaterMark = 100000

func main() {
	file, err := os.Open("default.xml")
	if err != nil {
		log.Fatalln(err)
	}
	xmlVal, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	m, err := mxj.NewMapXml(xmlVal)
	if err != nil {
		log.Fatalln(err)
	}
	high_water_mark, _ := m.ValueForPath("switch.high_water_mark")
	fmt.Println(high_water_mark)
	s, err := m.ValuesForPath("switch.sender")
	if err != nil {
		log.Fatalln(err)
	}
	for _, ss := range s {
		v := ss.(map[string]interface{})
		fmt.Println(v["endpoint"])
	}
	// mv := mxj.Map(s[0].(map[string]interface{}))
	// endpoint, _ := mv.ValuesForKey("endpoint")
	// fmt.Printf("%s\n", endpoint[0])
	sw := NewSwitchConfig("default.xml")
	sends, _ := sw.Senders()
	hwm, _ := sw.HWM()
	fmt.Println(sends, hwm)
}

type SwitchConfig struct {
	ConfigFile string
	Config     mxj.Map
	namespace  string
}

func NewSwitchConfig(xmlFileName string) *SwitchConfig {

	if !strings.HasSuffix(xmlFileName, ".xml") {
		log.Println("Not XML file")
		return nil
	}

	xmlData, err := ioutil.ReadFile(xmlFileName)
	d, err := mxj.NewMapXml(xmlData)
	if err != nil {
		log.Fatalln(err)
	}
	return &SwitchConfig{
		Config:     d,
		ConfigFile: xmlFileName,
		namespace:  "switch",
	}
}

func (sw *SwitchConfig) HWM() (int, error) {
	highWaterMark, err := sw.Config.ValueForPath(sw.namespace + "." + "high_water_mark")
	if err != nil {
		return DefaultHighWaterMark, err
	}
	return strconv.Atoi(highWaterMark.(string))
}

func (sw *SwitchConfig) Senders() ([]map[string]string, error) {
	var senders []map[string]string
	mv, err := sw.Config.ValuesForPath(sw.namespace + "." + "sender")
	if err != nil {
		return nil, err
	}
	for _, sender := range mv {
		v := sender.(map[string]interface{})
		senders = append(senders, map[string]string{
			"endpoint": v["endpoint"].(string),
			"hwm":      v["high_water_mark"].(string),
		})
	}
	return senders, nil
}

type Sreload interface {
	reload() error
}
type SigReloader struct {
	os.Signal
	ProcessName string
	sync.Mutex
	Sreload
}

func (sr *SigReloader) load(sigch chan os.Signal) {

	log.Printf("signal recieved %d, %s reload start", sr.ProcessName)

	for {
		<-sigch
		sr.Lock()
		if err := sr.reload(); err == nil {
			log.Printf("%s reload finished!", sr.ProcessName)
		} else {
			log.Printf("process %s reload fail: %s", sr.ProcessName, err.Error())
		}
		sr.Unlock()
	}

}

type MsgSwitch struct {
	SigReloader
	Context   *zmq.Context
	receiver  string
	config    *SwitchConfig
	Senders   []*MsgSender
	SenderNum int
}

func (msw *MsgSwitch) initConfig() {
	msw.SenderNum = len(msw.Senders)
	senders, err := msw.config.Senders()
	if err != nil {
		log.Panicln(err)
	}
	for _, endpoint := range senders {
		msw.Senders = append(msw.Senders, NewMsgSender(endpoint["endpoint"], endpoint["hwm"]))
	}

}

func (msw *MsgSwitch) RUN() {

}

type MsgSender struct {
	endpoint string
	hwm      uint64
	ZmqSender
}
type MsgRecver struct {
	endpoint string
	hwm      uint64
	ZmqRecver
}

func NewMsgSender(endpoint, hwm string) *MsgSender {
	context, _ := zmq.NewContext()
	i, _ := strconv.ParseUint(hwm, 10, 64)
	socket, _ := context.NewSocket(zmq.PUSH)
	socket.SetHWM(i)
	socket.Connect(endpoint)
	return &MsgSender{
		endpoint:  endpoint,
		hwm:       i,
		ZmqSender: socket,
	}
}

func NewMsgRecver(endpoint, hwm string) *MsgRecver {
	context, _ := zmq.NewContext()
	i, _ := strconv.ParseUint(hwm, 10, 64)
	socket, _ := context.NewSocket(zmq.PULL)
	socket.SetHWM(i)
	socket.Bind(endpoint)
	return &MsgRecver{
		endpoint:  endpoint,
		hwm:       i,
		ZmqRecver: socket,
	}
}

type ZmqSender interface {
	SendMultipart(parts [][]byte, flags zmq.SendRecvOption) (err error)
}

type ZmqRecver interface {
	RecvMultipart(flags zmq.SendRecvOption) (parts [][]byte, err error)
}
 */