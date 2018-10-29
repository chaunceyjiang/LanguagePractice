package goftp

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
)

type FTP struct {
	conn      net.Conn
	addr      string
	tlsConfig *tls.Config
	reader    *bufio.Reader
	writer    *bufio.Writer
	debug     bool
	sync.Mutex
	logger *log.Logger
}

const defaultBufSize = 8192

var RePwdPath = regexp.MustCompile(`\"(.*)\"`)

type (
	WalkFunc func(path string, info os.FileMode, filename string) error
	RetrFunc func(r io.Reader) error
)

func NewFtpClient(addr string) (*FTP, error) {
	var err error
	var conn net.Conn
	if conn, err = net.Dial("tcp", addr); err != nil {
		return nil, err
	}
	return &FTP{
		conn:   conn,
		addr:   addr,
		debug:  true,
		logger: log.New(os.Stdout, "ftp:", log.Ldate|log.LUTC|log.Lshortfile),
		reader: bufio.NewReaderSize(conn, defaultBufSize), //将FTP reader和writer与conn绑定
		writer: bufio.NewWriterSize(conn, defaultBufSize), //将FTP reader和writer与conn绑定
	}, nil
}
func parseLine(line string) (perm string, t string, filename string) {
	for _, v := range strings.Split(line, ";") {
		v2 := strings.Split(v, "=")
		switch v2[0] {
		case "perm":
			perm = v2[1]
		case "type":
			t = v2[1]
		default:
			filename = v[1 : len(v)-2]
		}
	}
	return
}
func (ftp *FTP) send(command string, arguments ...interface{}) error {
	if ftp.debug {
		ftp.log("> %s", fmt.Sprintf(command, arguments...))
	}

	command = fmt.Sprintf(command, arguments...)
	command += "\r\n"
	if _, err := ftp.writer.WriteString(command); err != nil {
		return err
	}
	if err := ftp.writer.Flush(); err != nil {
		return err
	}
	return nil

}
func (ftp *FTP) receiveLine() (string, error) {
	line, err := ftp.reader.ReadString('\n')
	if ftp.debug {
		ftp.log("< %s", line)
	}
	return line, err
}
func (ftp *FTP) receive() (string, error) {
	line, err := ftp.receiveLine()
	if err != nil {
		return line, err
	}
	if (len(line) >= 4) && (line[3] == '-') {
		closingCode := line[:3] + " "
		for {
			str, err := ftp.receiveLine()
			line = line + str
			if err != nil {
				return line, err
			}
			if len(str) < 4 {
				if ftp.debug {
					ftp.log("Uncorrectly terminated response")
				}
			} else {
				if str[:4] == closingCode {
					break
				}
			}

		}
	}
	ftp.ReadAndDiscard()
	return line, err
}
func (ftp *FTP) cmd(expects string, command string, args ...interface{}) (line string, err error) {
	if err = ftp.send(command, args...); err != nil { //发送指令
		return
	}
	if line, err = ftp.receive(); err != nil { //接受返回码
		return
	}
	if !strings.HasPrefix(line, expects) {
		err = errors.New(line)
		return
	}
	return

}
func (ftp *FTP) log(v ...interface{}) {
	if ftp.debug {
		if ftp.logger == nil {
			log.Println(v)
		} else {
			ftp.logger.Println(v)
		}

	}
}
func (ftp *FTP) SetLogger(logger *log.Logger) {
	ftp.Lock()
	defer ftp.Unlock()
	ftp.logger = logger
}
func (ftp *FTP) Close() error {
	return ftp.conn.Close()
}

func (ftp *FTP) Quit() error {
	if _, err := ftp.cmd(StatusConnectionClosing, "QUIT"); err != nil {
		return err
	}
	ftp.conn.Close()
	ftp.conn = nil
	return nil
}
func (ftp *FTP) Noop() error {
	_, err := ftp.cmd(StatusOK, "NOOP")
	return err
}
func (ftp *FTP) Walk(path string, walkFunc WalkFunc) (err error) {
	if ftp.debug {
		ftp.log("Walking:", path)
	}
	var lines []string
	return
}
func (ftp *FTP) ReadAndDiscard() (int, error) {
	bufferSize := ftp.reader.Buffered()
	var i int
	for i = 0; i < bufferSize; i++ {
		if _, err := ftp.reader.ReadByte(); err != nil {
			return i, err
		}
	}
	return i, nil
}
func (ftp *FTP)AuthTLS(config *tls.Config) error {
	if _,err:=ftp.cmd("234","AUTH TLS");err!=nil{
		return err
	}
	ftp.tlsConfig=config
	ftp.conn=tls.Client(ftp.conn,config)
	ftp.writer = bufio.NewWriter(ftp.conn)//基于TLS的conn
	ftp.reader = bufio.NewReader(ftp.conn)

	if _, err := ftp.cmd(StatusOK, "PBSZ 0"); err != nil {
		return err
	}

	if _, err := ftp.cmd(StatusOK, "PROT P"); err != nil {
		return err
	}

	return nil
}