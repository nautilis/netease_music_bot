package user_playlist

import (
	"encoding/json"
	"testing"
)

func TestUserPlaylist_UserPlaylist(t *testing.T) {
	userPlaylist := &UserPlaylist{}
	cookie := map[string]string{
		"os":            "pc",
		"__csrf":        "fcdc7115c74430cad7ec148b2713ed70",
		"__remember_me": "true",
		"NMTID":         "00OgJ6CkxzKn7ktTUVkpGdzf3CHw_MAAAF7NUniww",
		"MUSIC_U":       "da9df33ae4acd5cd047360ea3165104e5e9b5233b736633da8c4feed9de595c36ea7bb836ffc05cb6c0b02903764249b",
	}
	resp := userPlaylist.Query(&Data{
		Uid:    "32953014",
		Limit:  30,
		Offset: 0,
	}, cookie)
	r, _ := json.Marshal(resp)
	t.Log(r)
}
