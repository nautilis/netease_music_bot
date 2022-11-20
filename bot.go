package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/nautilis/netease_music_bot/netEaseMusic/modules/search"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nautilis/netease_music_bot/logging"
	"github.com/nautilis/netease_music_bot/netEaseMusic/modules/song_detail"
	"github.com/nautilis/netease_music_bot/netEaseMusic/modules/song_url"
	"github.com/nautilis/netease_music_bot/netEaseMusic/modules/user_record"
	"github.com/nautilis/netease_music_bot/setting"
)

type BussCode int

const (
	UserNotOpenPlayRecord BussCode = 403
)

func StartBot() {
	_bot, err := tgbotapi.NewBotAPI(setting.AppSetting.TelegramToke)
	if err != nil {
		panic("fail to init bot " + err.Error())
	}
	bot = _bot
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		defer func() {
			err := recover()
			if err != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				fmt.Printf("panic ===> %s\n", string(buf[:n]))
			}
		}()
		if update.Message != nil { // If we got a message
			msg, _ := json.Marshal(update)
			log.Printf("[%s] %s %s", update.Message.From.UserName, update.Message.Text, msg)
			if update.Message.Text == "/start" {
				handleStart(update)
			} else if strings.HasPrefix(update.Message.Text, "/subscribe") {
				handleSubscribe(update)
			} else if strings.HasPrefix(update.Message.Text, "/week_history") {
				handleWeekHistory(update)
			} else {
				handleSearch(update)
			}
		} else if update.CallbackQuery != nil {
			if strings.HasPrefix(update.CallbackQuery.Data, "/song") {
				handleSongPull(update)
			}
		}
	}
}

func handleStart(update tgbotapi.Update) {
	msg := update.Message
	newMsg := tgbotapi.NewMessage(msg.Chat.ID, "this is a bot for netease music\nsend /subscribe/{netease_uid} to subscribe a user\nsend song name to search song\n")
	//newMsg.ReplyToMessageID = update.Message.MessageID
	bot.Send(newMsg)
}

func handleSearch(update tgbotapi.Update) {
	keywords := strings.TrimSpace(update.Message.Text)

	searchResp := search.Query(&search.Data{
		S:      keywords,
		Type:   1,
		Limit:  10,
		Offset: 0,
	}, NetEaseCookie)

	if searchResp.Code != 200 {
		logging.Errorf("fail to search music, keywords: %s, resp:%v", keywords, searchResp)
		return
	}

	var textList []string
	var songIds []int64
	songs := searchResp.Result.Songs
	for idx, s := range songs {
		textList = append(textList, fmt.Sprintf("%d. %s - %s", idx+1, s.Name, s.Ar[0].Name))
		songIds = append(songIds, s.ID)
		if idx == 9 { //
			break
		}
	}
	text := strings.Join(textList, "\n")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	rows := len(songIds) / 5
	if len(songIds)%5 != 0 {
		rows += 1
	}
	var inlineKeyBoardButtons = make([][]tgbotapi.InlineKeyboardButton, rows)
	for i := 0; i < rows; i++ {
		inlineKeyBoardButtons[i] = []tgbotapi.InlineKeyboardButton{}
		for j := i * 5; j < len(songIds) && j < (i+1)*5; j++ {
			cbData := "/song/" + strconv.FormatInt(songIds[j], 10)
			button := tgbotapi.InlineKeyboardButton{
				Text:         strconv.Itoa(j + 1),
				CallbackData: &cbData,
			}
			inlineKeyBoardButtons[i] = append(inlineKeyBoardButtons[i], button)
		}
	}
	msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: inlineKeyBoardButtons}
	bot.Send(msg)

}

func handleSubscribe(update tgbotapi.Update) {
	uid := update.Message.From.ID
	subid := strings.Split(update.Message.Text, "/")[2]
	botData.AddSubUid(strconv.FormatInt(uid, 10), subid)
	logging.Infof("new subscribe => uid %d subscribe %s", uid, subid)
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "now you can tab /week_history access his/her music record"))
}

func handleWeekHistory(update tgbotapi.Update) {
	uid := update.Message.From.ID
	subUids := botData.GetSubUid(strconv.FormatInt(uid, 10))
	if len(subUids) == 0 {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "could not found any netease you had subscribe, you can tap /subscribe/{netease_uid} to subscribe firstly."))
		return
	}
	for _, subid := range subUids {
		songs, code, err := GetMusicRecord(subid)
		if err != nil {
			logging.Errorf("fail to get music record, %v", err.Error())
			return
		}
		if code != 0 {
			return
		}

		var textList []string
		var songIds []int64
		for idx, s := range songs {
			textList = append(textList, fmt.Sprintf("%d. %s - %s", idx+1, s.Name, s.Ar))
			songIds = append(songIds, s.Id)
			if idx == 29 { // 只返归前20 首
				break
			}
		}
		text := strings.Join(textList, "\n")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

		rows := len(songIds) / 5
		if len(songIds)%5 != 0 {
			rows += 1
		}
		var inlineKeyBoardButtons = make([][]tgbotapi.InlineKeyboardButton, rows)
		for i := 0; i < rows; i++ {
			inlineKeyBoardButtons[i] = []tgbotapi.InlineKeyboardButton{}
			for j := i * 5; j < len(songIds) && j < (i+1)*5; j++ {
				cbData := "/song/" + strconv.FormatInt(songIds[j], 10)
				button := tgbotapi.InlineKeyboardButton{
					Text:         strconv.Itoa(j + 1),
					CallbackData: &cbData,
				}
				inlineKeyBoardButtons[i] = append(inlineKeyBoardButtons[i], button)
			}
		}
		msg.ReplyMarkup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: inlineKeyBoardButtons}
		bot.Send(msg)
	}

}

func handleSongPull(update tgbotapi.Update) {
	data := update.CallbackQuery.Data
	songId := strings.Split(data, "/")[2]
	detailResp := song_detail.Query(songId, nil)
	if detailResp.Code != 200 {
		logging.Errorf("song detail access fail")
		return
	}
	song := detailResp.Songs[0]
	songid := []string{songId}
	chatId := update.CallbackQuery.Message.Chat.ID
	time.Sleep(1 * time.Second)
	songResp := song_url.Query(songid, NetEaseCookie)
	if songResp.Code == 200 {
		if len(songResp.Data) > 0 {
			musicUrl := songResp.Data[0].URL
			duration := songResp.Data[0].Time
			if musicUrl == "" {
				re, _ := json.Marshal(songResp)
				logging.Errorf("fail to query song url: %s", string(re))
				musicUrl = fmt.Sprintf("https://music.163.com/song/media/outer/url?id=%s.mp3", songId)
			}
			resp, err := http.Get(musicUrl)
			if err != nil {
				logging.Errorf("fail to download music , err: %s, url: %s", err.Error(), musicUrl)
				return
			}
			defer resp.Body.Close()
			b, _ := ioutil.ReadAll(resp.Body)
			audio := tgbotapi.NewAudio(chatId, tgbotapi.FileBytes{
				Name:  song.Name,
				Bytes: b,
			})
			presp, perr := http.Get(song.Al.PicURL + "?param=200y200")
			if perr != nil {
				logging.Errorf("fail to download pic of music, err: %s", err.Error())
				return
			}
			defer presp.Body.Close()
			pb, _ := ioutil.ReadAll(presp.Body)
			audio.Thumb = tgbotapi.FileBytes{
				Name:  song.Name,
				Bytes: pb,
			}
			audio.Performer = song.Ar[0].Name
			audio.Title = song.Name
			audio.Caption = "by @neteast_music_bot"
			audio.Duration = int(duration / 1000)
			_, err = bot.Send(audio)
			if err != nil {
				logging.Errorf("fail to send to bot %s", err.Error())
			}

		}
	}
}

type song struct {
	Ar   string `json:"ar"`
	Name string `json:"name"`
	Id   int64  `json:"id"`
	Pic  string `json:"pic"`
}

func GetMusicRecord(uid string) ([]song, BussCode, error) {
	recordResp := user_record.Query(&user_record.Data{
		Uid:  uid,
		Type: "1",
	})
	if recordResp.Code != 200 {
		return nil, UserNotOpenPlayRecord, nil
	}

	var songs []song
	for _, p := range recordResp.WeekData {
		songs = append(songs, song{
			Ar:   p.Song.Ar[0].Name,
			Name: p.Song.Name,
			Id:   p.Song.ID,
			Pic:  p.Song.Al.PicURL,
		})
	}
	return songs, 0, nil
}
