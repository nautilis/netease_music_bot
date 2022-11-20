package song_url

import "github.com/nautilis/netease_music_bot/netEaseMusic/util"

type SongUrlResp struct {
	Code int64 `json:"code"`
	Data []struct {
		Br                     int64       `json:"br"`
		CanExtend              bool        `json:"canExtend"`
		Code                   int64       `json:"code"`
		EffectTypes            interface{} `json:"effectTypes"`
		EncodeType             string      `json:"encodeType"`
		Expi                   int64       `json:"expi"`
		Fee                    int64       `json:"fee"`
		Flag                   int64       `json:"flag"`
		FreeTimeTrialPrivilege struct {
			RemainTime     int64 `json:"remainTime"`
			ResConsumable  bool  `json:"resConsumable"`
			Type           int64 `json:"type"`
			UserConsumable bool  `json:"userConsumable"`
		} `json:"freeTimeTrialPrivilege"`
		FreeTrialInfo      interface{} `json:"freeTrialInfo"`
		FreeTrialPrivilege struct {
			ListenType     interface{} `json:"listenType"`
			ResConsumable  bool        `json:"resConsumable"`
			UserConsumable bool        `json:"userConsumable"`
		} `json:"freeTrialPrivilege"`
		Gain        float64     `json:"gain"`
		ID          int64       `json:"id"`
		Level       string      `json:"level"`
		Md5         string      `json:"md5"`
		Payed       int64       `json:"payed"`
		Peak        int64       `json:"peak"`
		PodcastCtrp interface{} `json:"podcastCtrp"`
		RightSource int64       `json:"rightSource"`
		Size        int64       `json:"size"`
		Time        int64       `json:"time"`
		Type        string      `json:"type"`
		Uf          interface{} `json:"uf"`
		URL         string      `json:"url"`
		URLSource   int64       `json:"urlSource"`
	} `json:"data"`
}

type Data struct {
	Ids    []string          `json:"ids"`
	Br     int               `json:"br"`
	Header map[string]string `json:"header"`
}

func (d *Data) SetHeader(h map[string]string) {
	d.Header = h
}

func Query(ids []string, cookie map[string]string) *SongUrlResp {
	resp := &SongUrlResp{}
	util.RequestEapi("https://interface3.music.163.com/eapi/song/enhance/player/url", &Data{
		Ids: ids,
		Br:  999000,
	}, cookie, map[string]string{
		"url":    "/api/song/enhance/player/url",
		"realIP": "27.46.131.60",
	}, resp)
	return resp
}
