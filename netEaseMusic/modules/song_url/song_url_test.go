package song_url

import "testing"

func TestSongUrl(t *testing.T) {
	resp := Query([]string{"85786"}, nil)
	t.Log(resp)
	t.Log(resp.Data[0].URL)
}
