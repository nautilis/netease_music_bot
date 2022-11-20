package search

import "testing"

func TestQuery(t *testing.T) {
	resp := Query(&Data{
		S:      "当",
		Limit:  10,
		Offset: 0,
	}, nil)
	t.Log(resp)
	t.Log(resp.Code)
	t.Log(resp.Result.Songs[0].Name, resp.Result.Songs[0].Ar[0].Name)
}
