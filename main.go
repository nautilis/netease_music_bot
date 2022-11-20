package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nautilis/netease_music_bot/logging"
	"github.com/nautilis/netease_music_bot/netEaseMusic/modules/login"
	"github.com/nautilis/netease_music_bot/setting"
)

var bot *tgbotapi.BotAPI
var botData *BotData
var NetEaseCookie map[string]string

func LoginNetEase() {
	cookie, resp, err := login.Query(&login.LoginData{
		Username:      setting.AppSetting.NetEaseAccount,
		Password:      setting.AppSetting.NetEasePwd,
		RememberLogin: "true",
		CsrfToken:     "",
	}, NetEaseCookie)
	if err != nil {
		logging.Errorf("login err: %s", err.Error())
	} else {
		if resp.Code == 200 {
			logging.Infof("loging netease success , cookie => %v", cookie)
			NetEaseCookie = cookie
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
