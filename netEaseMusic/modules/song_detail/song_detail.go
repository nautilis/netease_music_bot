package song_detail

import (
	"fmt"
	"log"
	"strings"

	"github.com/nautilis/netease_music_bot/netEaseMusic/util"
)

type Data struct {
	C   string `json:"c"`
	Ids string `json:"ids"`
}

type Resp struct {
	Code       int64 `json:"code"`
	Privileges []struct {
		ChargeInfoList []struct {
			ChargeMessage interface{} `json:"chargeMessage"`
			ChargeType    int64       `json:"chargeType"`
			ChargeURL     interface{} `json:"chargeUrl"`
			Rate          int64       `json:"rate"`
		} `json:"chargeInfoList"`
		Cp                 int64 `json:"cp"`
		Cs                 bool  `json:"cs"`
		Dl                 int64 `json:"dl"`
		DownloadMaxbr      int64 `json:"downloadMaxbr"`
		Fee                int64 `json:"fee"`
		Fl                 int64 `json:"fl"`
		Flag               int64 `json:"flag"`
		FreeTrialPrivilege struct {
			ResConsumable  bool `json:"resConsumable"`
			UserConsumable bool `json:"userConsumable"`
		} `json:"freeTrialPrivilege"`
		ID        int64       `json:"id"`
		Maxbr     int64       `json:"maxbr"`
		Payed     int64       `json:"payed"`
		Pl        int64       `json:"pl"`
		PlayMaxbr int64       `json:"playMaxbr"`
		PreSell   bool        `json:"preSell"`
		Rscl      interface{} `json:"rscl"`
		Sp        int64       `json:"sp"`
		St        int64       `json:"st"`
		Subp      int64       `json:"subp"`
		Toast     bool        `json:"toast"`
	} `json:"privileges"`
	Songs []struct {
		A  interface{} `json:"a"`
		Al struct {
			ID     int64         `json:"id"`
			Name   string        `json:"name"`
			Pic    int64         `json:"pic"`
			PicURL string        `json:"picUrl"`
			PicStr string        `json:"pic_str"`
			Tns    []interface{} `json:"tns"`
		} `json:"al"`
		Alia []interface{} `json:"alia"`
		Ar   []struct {
			Alias []interface{} `json:"alias"`
			ID    int64         `json:"id"`
			Name  string        `json:"name"`
			Tns   []interface{} `json:"tns"`
		} `json:"ar"`
		Cd        string      `json:"cd"`
		Cf        string      `json:"cf"`
		Copyright int64       `json:"copyright"`
		Cp        int64       `json:"cp"`
		Crbt      interface{} `json:"crbt"`
		DjID      int64       `json:"djId"`
		Dt        int64       `json:"dt"`
		Fee       int64       `json:"fee"`
		Ftype     int64       `json:"ftype"`
		H         struct {
			Br   int64   `json:"br"`
			Fid  int64   `json:"fid"`
			Size int64   `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"h"`
		ID int64 `json:"id"`
		L  struct {
			Br   int64   `json:"br"`
			Fid  int64   `json:"fid"`
			Size int64   `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"l"`
		M struct {
			Br   int64   `json:"br"`
			Fid  int64   `json:"fid"`
			Size int64   `json:"size"`
			Vd   float64 `json:"vd"`
		} `json:"m"`
		Mark                 int64         `json:"mark"`
		Mst                  int64         `json:"mst"`
		Mv                   int64         `json:"mv"`
		Name                 string        `json:"name"`
		No                   int64         `json:"no"`
		NoCopyrightRcmd      interface{}   `json:"noCopyrightRcmd"`
		OriginCoverType      int64         `json:"originCoverType"`
		OriginSongSimpleData interface{}   `json:"originSongSimpleData"`
		Pop                  float64       `json:"pop"`
		Pst                  int64         `json:"pst"`
		PublishTime          int64         `json:"publishTime"`
		ResourceState        bool          `json:"resourceState"`
		Rt                   string        `json:"rt"`
		RtURL                interface{}   `json:"rtUrl"`
		RtUrls               []interface{} `json:"rtUrls"`
		Rtype                int64         `json:"rtype"`
		Rurl                 interface{}   `json:"rurl"`
		SID                  int64         `json:"s_id"`
		Single               int64         `json:"single"`
		St                   int64         `json:"st"`
		T                    int64         `json:"t"`
		Tns                  []string      `json:"tns"`
		V                    int64         `json:"v"`
		Version              int64         `json:"version"`
	} `json:"songs"`
}

func Query(ids string, cookies map[string]string) *Resp {
	resp := &Resp{}
	idList := strings.Split(ids, ",")
	cL := []string{}
	for _, id := range idList {
		cL = append(cL, fmt.Sprintf("{'id': %s}", id))
	}

	data := &Data{
		C:   fmt.Sprintf("[%s]", strings.Join(cL, ",")),
		Ids: fmt.Sprintf("[%s]", ids),
	}
	_, err := util.Request(`https://music.163.com/weapi/v3/song/detail`, data, cookies, resp)
	if err != nil {
		log.Println("Error", err.Error())
	}
	return resp
}
