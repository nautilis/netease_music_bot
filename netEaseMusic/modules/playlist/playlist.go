package playlist

import (
	"log"

	"github.com/nautilis/netease_music_bot/netEaseMusic/util"
)

type PlaylistDetailResp struct {
	Code          int         `json:"code"`
	RelatedVideos interface{} `json:"relatedVideos"`
	Playlist      struct {
		Subscribers []struct {
			DefaultAvatar      bool        `json:"defaultAvatar"`
			Province           int         `json:"province"`
			AuthStatus         int         `json:"authStatus"`
			Followed           bool        `json:"followed"`
			AvatarURL          string      `json:"avatarUrl"`
			AccountStatus      int         `json:"accountStatus"`
			Gender             int         `json:"gender"`
			City               int         `json:"city"`
			Birthday           int64       `json:"birthday"`
			UserID             int         `json:"userId"`
			UserType           int         `json:"userType"`
			Nickname           string      `json:"nickname"`
			Signature          string      `json:"signature"`
			Description        string      `json:"description"`
			DetailDescription  string      `json:"detailDescription"`
			AvatarImgID        int64       `json:"avatarImgId"`
			BackgroundImgID    int64       `json:"backgroundImgId"`
			BackgroundURL      string      `json:"backgroundUrl"`
			Authority          int         `json:"authority"`
			Mutual             bool        `json:"mutual"`
			ExpertTags         interface{} `json:"expertTags"`
			Experts            interface{} `json:"experts"`
			DjStatus           int         `json:"djStatus"`
			VipType            int         `json:"vipType"`
			RemarkName         interface{} `json:"remarkName"`
			AvatarImgIDStr     string      `json:"avatarImgIdStr"`
			BackgroundImgIDStr string      `json:"backgroundImgIdStr"`
		} `json:"subscribers"`
		Subscribed bool `json:"subscribed"`
		Creator    struct {
			DefaultAvatar      bool        `json:"defaultAvatar"`
			Province           int         `json:"province"`
			AuthStatus         int         `json:"authStatus"`
			Followed           bool        `json:"followed"`
			AvatarURL          string      `json:"avatarUrl"`
			AccountStatus      int         `json:"accountStatus"`
			Gender             int         `json:"gender"`
			City               int         `json:"city"`
			Birthday           int64       `json:"birthday"`
			UserID             int         `json:"userId"`
			UserType           int         `json:"userType"`
			Nickname           string      `json:"nickname"`
			Signature          string      `json:"signature"`
			Description        string      `json:"description"`
			DetailDescription  string      `json:"detailDescription"`
			AvatarImgID        int64       `json:"avatarImgId"`
			BackgroundImgID    int64       `json:"backgroundImgId"`
			BackgroundURL      string      `json:"backgroundUrl"`
			Authority          int         `json:"authority"`
			Mutual             bool        `json:"mutual"`
			ExpertTags         interface{} `json:"expertTags"`
			Experts            interface{} `json:"experts"`
			DjStatus           int         `json:"djStatus"`
			VipType            int         `json:"vipType"`
			RemarkName         interface{} `json:"remarkName"`
			AvatarImgIDStr     string      `json:"avatarImgIdStr"`
			BackgroundImgIDStr string      `json:"backgroundImgIdStr"`
		} `json:"creator"`
		Tracks []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
			Pst  int    `json:"pst"`
			T    int    `json:"t"`
			Ar   []struct {
				ID    int           `json:"id"`
				Name  string        `json:"name"`
				Tns   []interface{} `json:"tns"`
				Alias []interface{} `json:"alias"`
			} `json:"ar"`
			Alia []interface{} `json:"alia"`
			Pop  float64       `json:"pop"`
			St   int           `json:"st"`
			Rt   interface{}   `json:"rt"`
			Fee  int           `json:"fee"`
			V    int           `json:"v"`
			Crbt interface{}   `json:"crbt"`
			Cf   string        `json:"cf"`
			Dt   int           `json:"dt"`
			H    struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"h"`
			M struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"m"`
			L struct {
				Br   int     `json:"br"`
				Fid  int     `json:"fid"`
				Size int     `json:"size"`
				Vd   float64 `json:"vd"`
			} `json:"l"`
			A           interface{}   `json:"a"`
			Cd          string        `json:"cd"`
			No          int           `json:"no"`
			RtURL       interface{}   `json:"rtUrl"`
			Ftype       int           `json:"ftype"`
			RtUrls      []interface{} `json:"rtUrls"`
			DjID        int           `json:"djId"`
			Copyright   int           `json:"copyright"`
			SID         int           `json:"s_id"`
			Mark        int           `json:"mark"`
			Rtype       int           `json:"rtype"`
			Rurl        interface{}   `json:"rurl"`
			Mst         int           `json:"mst"`
			Cp          int           `json:"cp"`
			Mv          int           `json:"mv"`
			PublishTime int64         `json:"publishTime"`
			Al          struct {
				ID     int           `json:"id"`
				Name   string        `json:"name"`
				PicURL string        `json:"picUrl"`
				Tns    []interface{} `json:"tns"`
				Pic    int64         `json:"pic"`
			} `json:"al,omitempty"`
		} `json:"tracks"`
		TrackIds []struct {
			ID  int         `json:"id"`
			V   int         `json:"v"`
			Alg interface{} `json:"alg"`
		} `json:"trackIds"`
		UpdateFrequency       interface{}   `json:"updateFrequency"`
		BackgroundCoverID     int           `json:"backgroundCoverId"`
		BackgroundCoverURL    interface{}   `json:"backgroundCoverUrl"`
		TitleImage            int           `json:"titleImage"`
		TitleImageURL         interface{}   `json:"titleImageUrl"`
		EnglishTitle          interface{}   `json:"englishTitle"`
		OpRecommend           bool          `json:"opRecommend"`
		Status                int           `json:"status"`
		UserID                int           `json:"userId"`
		SubscribedCount       int           `json:"subscribedCount"`
		CloudTrackCount       int           `json:"cloudTrackCount"`
		Ordered               bool          `json:"ordered"`
		Tags                  []interface{} `json:"tags"`
		Description           interface{}   `json:"description"`
		HighQuality           bool          `json:"highQuality"`
		CreateTime            int64         `json:"createTime"`
		Privacy               int           `json:"privacy"`
		TrackUpdateTime       int64         `json:"trackUpdateTime"`
		TrackNumberUpdateTime int64         `json:"trackNumberUpdateTime"`
		CoverImgID            int64         `json:"coverImgId"`
		NewImported           bool          `json:"newImported"`
		UpdateTime            int64         `json:"updateTime"`
		SpecialType           int           `json:"specialType"`
		CoverImgURL           string        `json:"coverImgUrl"`
		TrackCount            int           `json:"trackCount"`
		CommentThreadID       string        `json:"commentThreadId"`
		AdType                int           `json:"adType"`
		PlayCount             int           `json:"playCount"`
		Name                  string        `json:"name"`
		ID                    int           `json:"id"`
		ShareCount            int           `json:"shareCount"`
		CoverImgIDStr         string        `json:"coverImgId_str"`
		CommentCount          int           `json:"commentCount"`
	} `json:"playlist"`
	Urls       interface{} `json:"urls"`
	Privileges []struct {
		ID      int  `json:"id"`
		Fee     int  `json:"fee"`
		Payed   int  `json:"payed"`
		St      int  `json:"st"`
		Pl      int  `json:"pl"`
		Dl      int  `json:"dl"`
		Sp      int  `json:"sp"`
		Cp      int  `json:"cp"`
		Subp    int  `json:"subp"`
		Cs      bool `json:"cs"`
		Maxbr   int  `json:"maxbr"`
		Fl      int  `json:"fl"`
		Toast   bool `json:"toast"`
		Flag    int  `json:"flag"`
		PreSell bool `json:"preSell"`
	} `json:"privileges"`
}

type Data struct {
	Id string `json:"id"`
	N  string `json:"n"`
	S  string `json:"s"`
}

type Playlist struct {
}

func (playlist *Playlist) Detail(data *Data, cookie map[string]string) *PlaylistDetailResp {
	url := `https://music.163.com/weapi/v3/playlist/detail`
	resp := &PlaylistDetailResp{}
	_, err := util.Request(url, data, cookie, resp)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return resp
}
