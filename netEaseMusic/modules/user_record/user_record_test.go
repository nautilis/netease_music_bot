package user_record

import (
	"testing"

	"github.com/unknwon/com"
)

func TestUserRecord(t *testing.T) {
	resp := Query(&Data{
		Uid:  "396444917",
		Type: "1",
	})
	t.Log(com.ToStr(resp))
}
