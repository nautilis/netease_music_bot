package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nautilis/netease_music_bot/logging"
	"github.com/nautilis/netease_music_bot/netEaseMusic/modules/login"
	"github.com/nautilis/netease_music_bot/setting"
)

var bot *tgbotapi.BotAPI
var botData *BotData
var NetEaseCookie map[string]string

func LoginNetEase() {

	cookieFile := setting.AppSetting.RuntimeRootPath + "cookie.txt"
	cookieBytes, rerr := ioutil.ReadFile(cookieFile)
	if rerr != nil {
		logging.Errorf("fail to read cookie file, err: %s", rerr.Error())
	} else {
		cookieMap := map[string]string{}
		if err := json.Unmarshal(cookieBytes, &cookieMap); err == nil {
			expireAt, perr := time.Parse(time.RFC1123, cookieMap["Expires"])
			if perr == nil {
				if expireAt.Unix() > time.Now().Unix() {
					NetEaseCookie = cookieMap
					logging.Infof("reused cookie")
					return
				}
			}
		} else {
			logging.Errorf("fail to unmarshal cookie text, err: %s", err.Error())
		}
	}
	cookie, resp, err := login.Query(&login.LoginData{
		Username:      setting.AppSetting.NetEaseAccount,
		Password:      setting.AppSetting.NetEasePwd,
		RememberLogin: "true",
		CsrfToken:     "",
	}, nil)
	if err != nil {
		logging.Errorf("login err: %s", err.Error())
	} else {
		if resp.Code == 200 {
			logging.Infof("loging netease success , cookie => %v", cookie)
			NetEaseCookie = cookie
			cookieJson, err := json.Marshal(cookie)
			if err == nil {
				if werr := ioutil.WriteFile(cookieFile, cookieJson, 0655); werr != nil {
					logging.Errorf("fail to write down cookie, err: %v", werr.Error())
				}
			} else {
				logging.Errorf("fail to marshal cookie , err: %s", err.Error())
			}
		} else {
			j, _ := json.Marshal(resp)
			logging.Infof("loging fail, resp: %s", string(j))
		}
	}
}

func main() {
	setting.Setup()
	logging.Setup()
	LoginNetEase()
	botData = NewBotData()
	sigs := make(chan os.Signal, 1)
	doneCha := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("shutdown now")
		botData.DumpData()
		doneCha <- struct{}{}
		fmt.Println("shutdown done")
	}()
	go func() {
		StartBot()
	}()
	<-doneCha
}
