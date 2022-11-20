package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Headers = []string{
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
}

func getUserAgent() string {
	return Headers[rand.Intn(len(Headers)-1)]
}

func mergeCookie(cookieMap map[string]string) string {
	cookieList := []string{}
	for k, v := range cookieMap {
		cookieList = append(cookieList, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(cookieList, ";")
}

type Data interface {
	SetHeader(map[string]string)
}

func RequestEapi(url string, data Data, cookie map[string]string, optional map[string]string, targetResp interface{}) (map[string]string, error) {
	client := http.Client{}

	header := map[string]string{
		"osver":       cookie["osver"],
		"deviceId":    cookie["deviceId"],
		"appver":      "8.7.01",
		"versioncode": "140",
		"mobilename":  cookie["mobilename"],
		"buildver":    strconv.FormatInt(time.Now().Unix(), 10),
		"resolution":  "1920x1080",
		"__csrf":      cookie["__csrf"],
		"os":          "android",
		"channel":     cookie["channel"],
		"requestId":   fmt.Sprintf("%d_0%d", time.Now().UnixMilli(), rand.Intn(1000)),
		"MUSIC_U":     cookie["MUSIC_U"],
		//"MUSIC_A":     "bf8bfeabb1aa84f9c8c3906c04a04fb864322804c83f5d607e91a04eae463c9436bd1a17ec353cf780b396507a3f7464e8a60f4bbc019437993166e004087dd32d1490298caf655c2353e58daa0bc13cc7d5c198250968580b12c1b8817e3f5c807e650dd04abd3fb8130b7ae43fcc5b",
	}
	data.SetHeader(header)
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("before => %s\n", string(dataByte))
	body := EapiEncrypt(optional["url"], dataByte)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if ip := optional["realIP"]; ip != "" {
		req.Header.Set("X-Real-IP", ip)
		req.Header.Set("X-Forwarded-For", ip)
	}
	userAgent := getUserAgent()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	if strings.Contains(url, "music.163.com") {
		req.Header.Set("Referer", "https://music.163.com")
	}
	req.Header.Set("Cookie", mergeCookie(header))

	resp, err := client.Do(req)
	//fmt.Println(resp.Cookies())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, targetResp)
	cookieMap := map[string]string{}
	for _, c := range resp.Cookies() {
		tmp := strings.Split(c.String(), ";")
		for _, m := range tmp {
			//t.Log(m)
			tmp1 := strings.Split(m, "=")
			cookieMap[strings.TrimSpace(tmp1[0])] = strings.TrimSpace(tmp1[1])
		}
	}
	return cookieMap, err

}

func Request(url string, data interface{}, cookie map[string]string, targetResp interface{}) (map[string]string, error) {
	client := http.Client{}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := weapiEncrypt(dataByte)
	//body := `params=w8kouBju%2BmqQ2O9av9cBttY1I9gAdnvDGyYXBhd0hCCXJ87FpwEEX2yQ8tOLj10%2BJltTf0sff2D3OEfrwU52ohFXYGYYwByHzC1xX8hq4KduuOpCINJ75FdewC9RWQl%2FLEXbkDI%2FwSnw1qzOrU6pvyOqcInZ7rTqDypP5ApaYsZcnBr3r28lc6mKgFLz23XGmIc79WNA%2FDQ7pdWy8SPSdozVnXL2mXkvGBS9LSwz%2FTg%3D&encSecKey=3d14dddf249ba20a7ccf5024d63fb6b7eb60c8ff73b9f77eb7d8344238dfd6d35f96ed937bbcf21fa2f18e0b6befde8d43dfae412b039e999f43bac863fedb6a49e6065819dc702251a20ff888c8edbd95400c5cffd0ce6c7f9fa3a5612b2c7a61ed6178a1d028b3dbe1470e40660dfc2c3dcd81e4f90d92d88375759566c066`
	//fmt.Println(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	userAgent := getUserAgent()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	if strings.Contains(url, "music.163.com") {
		req.Header.Set("Referer", "https://music.163.com")
	}
	req.Header.Set("Cookie", mergeCookie(cookie))
	//req.Header.Set("Set-Cookie", cookie2)

	resp, err := client.Do(req)
	//fmt.Println(resp.Cookies())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, targetResp)
	if err != nil {
		return nil, err
	}
	cookieMap := map[string]string{}
	for _, c := range resp.Cookies() {
		tmp := strings.Split(c.String(), ";")
		for _, m := range tmp {
			//t.Log(m)
			tmp1 := strings.Split(m, "=")
			cookieMap[strings.TrimSpace(tmp1[0])] = strings.TrimSpace(tmp1[1])
		}
	}
	return cookieMap, err
}
