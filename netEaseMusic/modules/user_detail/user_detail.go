package user_detail

import (
	"log"

	"github.com/nautilis/netease_music_bot/netEaseMusic/util"
)

type UserDetail struct{}

type UserDetailResp struct {
	AdValid  bool `json:"adValid"`
	Bindings []struct {
		BindingTime  int64       `json:"bindingTime"`
		Expired      bool        `json:"expired"`
		ExpiresIn    int64       `json:"expiresIn"`
		ID           int64       `json:"id"`
		RefreshTime  int64       `json:"refreshTime"`
		TokenJSONStr interface{} `json:"tokenJsonStr"`
		Type         int64       `json:"type"`
		URL          string      `json:"url"`
		UserID       int64       `json:"userId"`
	} `json:"bindings"`
	Code                     int64 `json:"code"`
	CreateDays               int64 `json:"createDays"`
	CreateTime               int64 `json:"createTime"`
	Level                    int64 `json:"level"`
	ListenSongs              int64 `json:"listenSongs"`
	MobileSign               bool  `json:"mobileSign"`
	PcSign                   bool  `json:"pcSign"`
	PeopleCanSeeMyPlayRecord bool  `json:"peopleCanSeeMyPlayRecord"`
	Profile                  struct {
		AccountStatus             int64         `json:"accountStatus"`
		AllSubscribedCount        int64         `json:"allSubscribedCount"`
		ArtistIdentity            []interface{} `json:"artistIdentity"`
		AuthStatus                int64         `json:"authStatus"`
		Authority                 int64         `json:"authority"`
		AvatarDetail              interface{}   `json:"avatarDetail"`
		AvatarImgID               int64         `json:"avatarImgId"`
		AvatarImgIDStr            string        `json:"avatarImgId_str"`
		AvatarURL                 string        `json:"avatarUrl"`
		BackgroundImgID           int64         `json:"backgroundImgId"`
		BackgroundImgIDStr        string        `json:"backgroundImgIdStr"`
		BackgroundURL             string        `json:"backgroundUrl"`
		Birthday                  int64         `json:"birthday"`
		Blacklist                 bool          `json:"blacklist"`
		CCount                    int64         `json:"cCount"`
		City                      int64         `json:"city"`
		CreateTime                int64         `json:"createTime"`
		DefaultAvatar             bool          `json:"defaultAvatar"`
		Description               string        `json:"description"`
		DetailDescription         string        `json:"detailDescription"`
		DjStatus                  int64         `json:"djStatus"`
		EventCount                int64         `json:"eventCount"`
		ExpertTags                interface{}   `json:"expertTags"`
		Experts                   struct{}      `json:"experts"`
		FollowMe                  bool          `json:"followMe"`
		FollowTime                interface{}   `json:"followTime"`
		Followed                  bool          `json:"followed"`
		Followeds                 int64         `json:"followeds"`
		Follows                   int64         `json:"follows"`
		Gender                    int64         `json:"gender"`
		Mutual                    bool          `json:"mutual"`
		NewFollows                int64         `json:"newFollows"`
		Nickname                  string        `json:"nickname"`
		PlaylistBeSubscribedCount int64         `json:"playlistBeSubscribedCount"`
		PlaylistCount             int64         `json:"playlistCount"`
		Province                  int64         `json:"province"`
		RemarkName                interface{}   `json:"remarkName"`
		SCount                    int64         `json:"sCount"`
		SDJPCount                 int64         `json:"sDJPCount"`
		Signature                 string        `json:"signature"`
		UserID                    int64         `json:"userId"`
		UserType                  int64         `json:"userType"`
		VipType                   int64         `json:"vipType"`
	} `json:"profile"`
	UserPoint struct {
		Balance      int64 `json:"balance"`
		BlockBalance int64 `json:"blockBalance"`
		Status       int64 `json:"status"`
		UpdateTime   int64 `json:"updateTime"`
		UserID       int64 `json:"userId"`
		Version      int64 `json:"version"`
	} `json:"userPoint"`
}

func (slf *UserDetail) UserDetail(uid string) *UserDetailResp {
	url := `https://music.163.com/weapi/v1/user/detail/` + uid
	resp := &UserDetailResp{}
	_, err := util.Request(url, nil, nil, resp)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return resp
}
