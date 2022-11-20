package user_playlist

import (
	"log"

	"github.com/nautilis/netease_music_bot/netEaseMusic/util"
)

type UserPlaylist struct {
}

type Resp struct {
	Code     int64 `json:"code"`
	More     bool  `json:"more"`
	Playlist []struct {
		AdType             int64       `json:"adType"`
		Anonimous          bool        `json:"anonimous"`
		Artists            interface{} `json:"artists"`
		BackgroundCoverID  int64       `json:"backgroundCoverId"`
		BackgroundCoverURL interface{} `json:"backgroundCoverUrl"`
		CloudTrackCount    int64       `json:"cloudTrackCount"`
		CommentThreadID    string      `json:"commentThreadId"`
		CoverImgID         int64       `json:"coverImgId"`
		CoverImgIDStr      string      `json:"coverImgId_str"`
		CoverImgURL        string      `json:"coverImgUrl"`
		CreateTime         int64       `json:"createTime"`
		Creator            struct {
			AccountStatus       int64       `json:"accountStatus"`
			Anchor              bool        `json:"anchor"`
			AuthStatus          int64       `json:"authStatus"`
			AuthenticationTypes int64       `json:"authenticationTypes"`
			Authority           int64       `json:"authority"`
			AvatarDetail        interface{} `json:"avatarDetail"`
			AvatarImgID         int64       `json:"avatarImgId"`
			AvatarImgIDStr      string      `json:"avatarImgIdStr"`
			AvatarURL           string      `json:"avatarUrl"`
			BackgroundImgID     int64       `json:"backgroundImgId"`
			BackgroundImgIDStr  string      `json:"backgroundImgIdStr"`
			BackgroundURL       string      `json:"backgroundUrl"`
			Birthday            int64       `json:"birthday"`
			City                int64       `json:"city"`
			DefaultAvatar       bool        `json:"defaultAvatar"`
			Description         string      `json:"description"`
			DetailDescription   string      `json:"detailDescription"`
			DjStatus            int64       `json:"djStatus"`
			ExpertTags          []string    `json:"expertTags"`
			Experts             interface{} `json:"experts"`
			Followed            bool        `json:"followed"`
			Gender              int64       `json:"gender"`
			Mutual              bool        `json:"mutual"`
			Nickname            string      `json:"nickname"`
			Province            int64       `json:"province"`
			RemarkName          interface{} `json:"remarkName"`
			Signature           string      `json:"signature"`
			UserID              int64       `json:"userId"`
			UserType            int64       `json:"userType"`
			VipType             int64       `json:"vipType"`
		} `json:"creator"`
		Description           string        `json:"description"`
		EnglishTitle          interface{}   `json:"englishTitle"`
		HighQuality           bool          `json:"highQuality"`
		ID                    int64         `json:"id"`
		Name                  string        `json:"name"`
		NewImported           bool          `json:"newImported"`
		OpRecommend           bool          `json:"opRecommend"`
		Ordered               bool          `json:"ordered"`
		PlayCount             int64         `json:"playCount"`
		Privacy               int64         `json:"privacy"`
		RecommendInfo         interface{}   `json:"recommendInfo"`
		SharedUsers           interface{}   `json:"sharedUsers"`
		SpecialType           int64         `json:"specialType"`
		Status                int64         `json:"status"`
		Subscribed            bool          `json:"subscribed"`
		SubscribedCount       int64         `json:"subscribedCount"`
		Subscribers           []interface{} `json:"subscribers"`
		Tags                  []string      `json:"tags"`
		TitleImage            int64         `json:"titleImage"`
		TitleImageURL         interface{}   `json:"titleImageUrl"`
		TotalDuration         int64         `json:"totalDuration"`
		TrackCount            int64         `json:"trackCount"`
		TrackNumberUpdateTime int64         `json:"trackNumberUpdateTime"`
		TrackUpdateTime       int64         `json:"trackUpdateTime"`
		Tracks                interface{}   `json:"tracks"`
		UpdateFrequency       interface{}   `json:"updateFrequency"`
		UpdateTime            int64         `json:"updateTime"`
		UserID                int64         `json:"userId"`
	} `json:"playlist"`
	Version string `json:"version"`
}

type Data struct {
	Uid    string `json:"uid"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (*UserPlaylist) Query(data *Data, cookie map[string]string) *Resp {
	resp := &Resp{}
	_, err := util.Request(`https://music.163.com/weapi/user/playlist`, data, cookie, resp)
	if err != nil {
		log.Println("Error", err.Error())
		return nil
	}
	return resp

}
