package login

import (
	"encoding/json"
	"testing"
	"time"
)

//jsonStr := `{"username":"bevanpf@163.com","password":"a6979486d8684a9721bf6f939855803a","rememberLogin":"true","csrf_token":""}`

func TestJson(t *testing.T) {
	data, _ := json.Marshal(&LoginData{
		Username:      "bevanpf@163.com",
		Password:      "a6979486d8684a9721bf6f939855803a",
		RememberLogin: "true",
		CsrfToken:     "",
	})
	cookies, resp, err := Query(&LoginData{
		Username:      "bevanpf@163.com",
		Password:      "a6979486d8684a9721bf6f939855803a",
		RememberLogin: "true",
		CsrfToken:     "",
	}, map[string]string{
		"os":     "pc",
		"appver": "2.9.7",
	})
	t.Log(string(data))
	t.Log(cookies)

	if err != nil {
		t.Error(err.Error())
	}
	expireAt, err := time.Parse(time.RFC1123, cookies["Expires"])
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(expireAt.Unix())
	t.Log(resp)

}
