package main

import (
	"github.com/nautilis/netease_music_bot/logging"
	"github.com/nautilis/netease_music_bot/setting"
)

func init() {
	setting.Setup()
	logging.Setup()
}
func main() {
	LoginNetEase()
	StartBot()
}
