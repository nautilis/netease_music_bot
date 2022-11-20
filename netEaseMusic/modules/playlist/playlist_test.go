package playlist

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPlaylist_Detail(t *testing.T) {
	data := &Data{
		Id: "69724682",
		N:  "10000",
		S:  "8",
	}
	cookie := map[string]string{
		"os":            "pc",
		"__csrf":        "fcdc7115c74430cad7ec148b2713ed70",
		"__remember_me": "true",
		"NMTID":         "00OgJ6CkxzKn7ktTUVkpGdzf3CHw_MAAAF7NUniww",
		"MUSIC_U":       "da9df33ae4acd5cd047360ea3165104e5e9b5233b736633da8c4feed9de595c36ea7bb836ffc05cb6c0b02903764249b",
	}

	playlist := &Playlist{}
	resp := playlist.Detail(data, cookie)
	b, _ := json.Marshal(resp)
	t.Log(string(b))
	if resp == nil {
		t.Error()
	}
	link := fmt.Sprintf(`https://music.163.com/#/song?id=%d`, resp.Playlist.Tracks[0].ID)
	t.Log(resp.Playlist.Tracks[0].Name, resp.Playlist.Tracks[0].Ar[0].Name, link)
}
