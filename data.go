package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/nautilis/netease_music_bot/logging"
	"github.com/nautilis/netease_music_bot/setting"
)

type BotData struct {
	uid2SubUid map[string][]string
	subLock    sync.Mutex
}

func NewBotData() *BotData {
	inst := &BotData{
		uid2SubUid: map[string][]string{},
	}
	inst.InitData()
	return inst
}

func (d *BotData) GetSubUid(uid string) (subUids []string) {
	d.subLock.Lock()
	defer d.subLock.Unlock()
	subUids = d.uid2SubUid[uid]
	return
}

func (d *BotData) getUid2SubUidFile() string {
	return fmt.Sprintf(setting.AppSetting.RuntimeRootPath + "uid-2-subuid.gob")
}

func (d *BotData) AddSubUid(uid string, subUid string) {
	d.subLock.Lock()
	defer d.subLock.Unlock()
	subUids := d.uid2SubUid[uid]
	for _, id := range subUids {
		if id == subUid {
			return
		}
	}
	subUids = append(subUids, subUid)
	d.uid2SubUid[uid] = subUids
}

func (d *BotData) DumpData() {
	var data bytes.Buffer
	enc := gob.NewEncoder(&data)
	err := enc.Encode(d.uid2SubUid)
	if err != nil {
		logging.Error("fail to encode uid 2 subuid")
		return
	}
	if err := ioutil.WriteFile(d.getUid2SubUidFile(), data.Bytes(), 0655); err != nil {
		logging.Error("fail to write uid-to-subuid.gob, err: %s", err.Error())
	}
}

func (d *BotData) InitData() {
	tmp := map[string][]string{}
	b, err := ioutil.ReadFile(d.getUid2SubUidFile())
	data := bytes.NewBuffer(b)
	if err != nil {
		logging.Errorf("fail to read uid-to-subuid.gob, err: %v", err.Error())
		return
	}
	decode := gob.NewDecoder(data)
	err = decode.Decode(&tmp)
	if err != nil {
		logging.Errorf("fail to decode uid-to-subuid.gob, err: %v", err.Error())
		return
	}
	d.uid2SubUid = tmp
}
