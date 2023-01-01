package setting

import (
	"os"

	"github.com/go-ini/ini"
)

type App struct {
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
	NetEaseAccount  string
	NetEasePwd      string
	TelegramToken   string
}

var AppSetting = &App{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	AppSetting = &App{
		RuntimeRootPath: "runtime/",
		LogSavePath:     "log",
		LogSaveName:     "log",
		LogFileExt:      "log",
		TimeFormat:      "20060102",
		NetEaseAccount:  os.Getenv("NetEaseAccount"),
		NetEasePwd:      os.Getenv("NetEasePwd"),
		TelegramToken:   os.Getenv("TelegramToken"),
	}
}
