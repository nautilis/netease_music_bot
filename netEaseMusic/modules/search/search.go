package search

import (
	"fmt"

	"github.com/nautilis/netease_music_bot/netEaseMusic/util"
)

type SearchResp struct {
	Code   int64 `json:"code"`
	Result struct {
		SearchQcReminder interface{} `json:"searchQcReminder"`
		SongCount        int64       `json:"songCount"`
		Songs            []struct {
			A  interface{} `json:"a"`
			Al struct {
				ID     int64         `json:"id"`
				Name   string        `json:"name"`
				Pic    int64         `json:"pic"`
				PicURL string        `json:"picUrl"`
				PicStr string        `json:"pic_str"`
				Tns    []interface{} `json:"tns"`
			} `json:"al"`
			Alia []string `json:"alia"`
			Ar   []struct {
				Alia  []string      `json:"alia"`
				Alias []string      `json:"alias"`
				ID    int64         `json:"id"`
				Name  string        `json:"name"`
				Tns   []interface{} `json:"tns"`
			} `json:"ar"`
			Cd                string      `json:"cd"`
			Cf                string      `json:"cf"`
			Copyright         int64       `json:"copyright"`
			Cp                int64       `json:"cp"`
			Crbt              interface{} `json:"crbt"`
			DjID              int64       `json:"djId"`
			Dt                int64       `json:"dt"`
			EntertainmentTags interface{} `json:"entertainmentTags"`
			Fee               int64       `json:"fee"`
			Ftype             int64       `json:"ftype"`
			H                 struct {
				Br   int64 `json:"br"`
				Fid  int64 `json:"fid"`
				Size int64 `json:"size"`
				Sr   int64 `json:"sr"`
				Vd   int64 `json:"vd"`
			} `json:"h"`
			Hr struct {
				Br   int64 `json:"br"`
				Fid  int64 `json:"fid"`
				Size int64 `json:"size"`
				Sr   int64 `json:"sr"`
				Vd   int64 `json:"vd"`
			} `json:"hr"`
			ID int64 `json:"id"`
			L  struct {
				Br   int64 `json:"br"`
				Fid  int64 `json:"fid"`
				Size int64 `json:"size"`
				Sr   int64 `json:"sr"`
				Vd   int64 `json:"vd"`
			} `json:"l"`
			M struct {
				Br   int64 `json:"br"`
				Fid  int64 `json:"fid"`
				Size int64 `json:"size"`
				Sr   int64 `json:"sr"`
				Vd   int64 `json:"vd"`
			} `json:"m"`
			Mark                 int64       `json:"mark"`
			Mst                  int64       `json:"mst"`
			Mv                   int64       `json:"mv"`
			Name                 string      `json:"name"`
			No                   int64       `json:"no"`
			NoCopyrightRcmd      interface{} `json:"noCopyrightRcmd"`
			OriginCoverType      int64       `json:"originCoverType"`
			OriginSongSimpleData struct {
				AlbumMeta struct {
					ID   int64  `json:"id"`
					Name string `json:"name"`
				} `json:"albumMeta"`
				Artists []struct {
					ID   int64  `json:"id"`
					Name string `json:"name"`
				} `json:"artists"`
				Name   string `json:"name"`
				SongID int64  `json:"songId"`
			} `json:"originSongSimpleData"`
			Pop       float64 `json:"pop"`
			Privilege struct {
				ChargeInfoList []struct {
					ChargeMessage interface{} `json:"chargeMessage"`
					ChargeType    int64       `json:"chargeType"`
					ChargeURL     interface{} `json:"chargeUrl"`
					Rate          int64       `json:"rate"`
				} `json:"chargeInfoList"`
				Cp                 int64  `json:"cp"`
				Cs                 bool   `json:"cs"`
				Dl                 int64  `json:"dl"`
				DlLevel            string `json:"dlLevel"`
				DownloadMaxBrLevel string `json:"downloadMaxBrLevel"`
				DownloadMaxbr      int64  `json:"downloadMaxbr"`
				Fee                int64  `json:"fee"`
				Fl                 int64  `json:"fl"`
				FlLevel            string `json:"flLevel"`
				Flag               int64  `json:"flag"`
				FreeTrialPrivilege struct {
					ListenType     interface{} `json:"listenType"`
					ResConsumable  bool        `json:"resConsumable"`
					UserConsumable bool        `json:"userConsumable"`
				} `json:"freeTrialPrivilege"`
				ID             int64       `json:"id"`
				MaxBrLevel     string      `json:"maxBrLevel"`
				Maxbr          int64       `json:"maxbr"`
				Payed          int64       `json:"payed"`
				Pl             int64       `json:"pl"`
				PlLevel        string      `json:"plLevel"`
				PlayMaxBrLevel string      `json:"playMaxBrLevel"`
				PlayMaxbr      int64       `json:"playMaxbr"`
				PreSell        bool        `json:"preSell"`
				Rscl           interface{} `json:"rscl"`
				Sp             int64       `json:"sp"`
				St             int64       `json:"st"`
				Subp           int64       `json:"subp"`
				Toast          bool        `json:"toast"`
			} `json:"privilege"`
			Pst           int64         `json:"pst"`
			PublishTime   int64         `json:"publishTime"`
			ResourceState bool          `json:"resourceState"`
			Rt            string        `json:"rt"`
			RtURL         interface{}   `json:"rtUrl"`
			RtUrls        []interface{} `json:"rtUrls"`
			Rtype         int64         `json:"rtype"`
			Rurl          interface{}   `json:"rurl"`
			SID           int64         `json:"s_id"`
			Single        int64         `json:"single"`
			SongJumpInfo  interface{}   `json:"songJumpInfo"`
			Sq            struct {
				Br   int64 `json:"br"`
				Fid  int64 `json:"fid"`
				Size int64 `json:"size"`
				Sr   int64 `json:"sr"`
				Vd   int64 `json:"vd"`
			} `json:"sq"`
			St         int64       `json:"st"`
			T          int64       `json:"t"`
			TagPicList interface{} `json:"tagPicList"`
			Tns        []string    `json:"tns"`
			V          int64       `json:"v"`
			Version    int64       `json:"version"`
		} `json:"songs"`
	} `json:"result"`
}

type Data struct {
	S      string `json:"s"`
	Type   int    `json:"type"` // // 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func Query(d *Data, cookie map[string]string) *SearchResp {
	d.Type = 1
	resp := &SearchResp{}
	_, err := util.Request("https://music.163.com/weapi/cloudsearch/get/web", d, cookie, resp)
	if err != nil {
		fmt.Printf("Error fail to request search %s\n", err.Error())
	}
	return resp
}
