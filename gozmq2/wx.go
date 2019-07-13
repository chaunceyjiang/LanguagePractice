package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/base"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/datacube"
	"github.com/chanxuehong/wechat/mp/menu"
	_ "github.com/chanxuehong/wechat/mp/message"
	"net/http"
	"time"
)

const (
	wxAppId     = "wx97ad2427ebe81b1a"
	wxAppSecret = "8aded44ddd247acba8e7689b5f4ada82"

	wxOriId         = "oriid"
	wxToken         = "token"
	wxEncodedAESKey = "aeskey"
)

func main() {
	client := *http.DefaultClient
	client.Timeout = time.Second * 50
	var (
		accessTokenServer core.AccessTokenServer = core.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, &client)
		wechatClient      *core.Client           = core.NewClient(accessTokenServer, nil)
	)
	fmt.Println(accessTokenServer.Token())
	fmt.Println(base.GetCallbackIP(wechatClient))
	b := menu.Button{
		Name: "抽奖",
	}

	b.SetAsClickButton("a", "b")
	m := menu.Menu{Buttons: []menu.Button{{Name: "抽奖"}}, MenuId: 1}
	t1, _ := time.Parse("2006-01-02", "2013-10-05")
	t2, _ := time.Parse("2006-01-02", "2018-10-05")
	t := datacube.NewRequest(t1, t2)
	fmt.Println(datacube.GetUserSummary(wechatClient, t))
	fmt.Println(menu.Create(wechatClient, &m))

}
