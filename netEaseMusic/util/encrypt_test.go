package util

import (
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	//jsonStr := `{"username":"bevanpf@163.com","password":"a6979486d8684a9721bf6f939855803a","rememberLogin":"true","csrf_token":""}`
	jsonStr := `{"username":"bevanpf@163.com","password":"a6979486d8684a9721bf6f939855803a","rememberLogin":"true","csrf_token":""}`
	b := weapiEncrypt([]byte(jsonStr))
	t.Log(string(b))
}
