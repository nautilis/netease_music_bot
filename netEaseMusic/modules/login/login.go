package login

import (
	"github.com/nautilis/netease_music_bot/netEaseMusic/util"
)

type LoginResp struct {
	LoginType int `json:"loginType"`
	Code      int `json:"code"`
	Account   struct {
		ID                 int    `json:"id"`
		UserName           string `json:"userName"`
		Type               int    `json:"type"`
		Status             int    `json:"status"`
		WhitelistAuthority int    `json:"whitelistAuthority"`
		CreateTime         int64  `json:"createTime"`
		Salt               string `json:"salt"`
		TokenVersion       int    `json:"tokenVersion"`
		Ban                int    `json:"ban"`
		BaoyueVersion      int    `json:"baoyueVersion"`
		DonateVersion      int    `json:"donateVersion"`
		VipType            int    `json:"vipType"`
		ViptypeVersion     int64  `json:"viptypeVersion"`
		AnonimousUser      bool   `json:"anonimousUser"`
	} `json:"account"`
	Token   string `json:"token"`
	Profile struct {
		UserID        int    `json:"userId"`
		VipType       int    `json:"vipType"`
		Gender        int    `json:"gender"`
		AccountStatus int    `json:"accountStatus"`
		AvatarImgID   int64  `json:"avatarImgId"`
		DefaultAvatar bool   `json:"defaultAvatar"`
		AvatarURL     string `json:"avatarUrl"`
		Nickname      string `json:"nickname"`
		Birthday      int64  `json:"birthday"`
		City          int    `json:"city"`
		Experts       struct {
		} `json:"experts"`
		Mutual                    bool        `json:"mutual"`
		RemarkName                interface{} `json:"remarkName"`
		ExpertTags                interface{} `json:"expertTags"`
		AuthStatus                int         `json:"authStatus"`
		BackgroundImgID           int64       `json:"backgroundImgId"`
		UserType                  int         `json:"userType"`
		Province                  int         `json:"province"`
		DjStatus                  int         `json:"djStatus"`
		Description               string      `json:"description"`
		AvatarImgIDStr            string      `json:"avatarImgIdStr"`
		BackgroundImgIDStr        string      `json:"backgroundImgIdStr"`
		DetailDescription         string      `json:"detailDescription"`
		Followed                  bool        `json:"followed"`
		BackgroundURL             string      `json:"backgroundUrl"`
		Signature                 string      `json:"signature"`
		Authority                 int         `json:"authority"`
		Followeds                 int         `json:"followeds"`
		Follows                   int         `json:"follows"`
		EventCount                int         `json:"eventCount"`
		PlaylistCount             int         `json:"playlistCount"`
		PlaylistBeSubscribedCount int         `json:"playlistBeSubscribedCount"`
	} `json:"profile"`
	Bindings []struct {
		Expired      bool   `json:"expired"`
		URL          string `json:"url"`
		UserID       int    `json:"userId"`
		TokenJSONStr string `json:"tokenJsonStr"`
		ExpiresIn    int64  `json:"expiresIn"`
		BindingTime  int64  `json:"bindingTime"`
		RefreshTime  int    `json:"refreshTime"`
		ID           int    `json:"id"`
		Type         int    `json:"type"`
	} `json:"bindings"`
}

type LoginData struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	RememberLogin string `json:"rememberLogin"`
	CsrfToken     string `json:"csrf_token"`
}

func Query(data *LoginData, cookie map[string]string) (map[string]string, *LoginResp, error) {
	url := "https://music.163.com/weapi/login/"
	if cookie == nil {
		cookie = map[string]string{}
	}
	cookie["os"] = "pc"
	cookie["appver"] = "2.9.7"
	resp := &LoginResp{}
	cookies, err := util.Request(url, data, cookie, resp)
	if err != nil {
		return nil, nil, err
	}
	return cookies, resp, nil
}
