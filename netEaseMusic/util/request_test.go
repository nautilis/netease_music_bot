package util

import "testing"

func TestEapi(t *testing.T) {
	after := EapiEncrypt("/api/song/enhance/player/url", []byte(`{"ids":"[33894312]","br":999000,"header":{"appver":"8.7.01","versioncode":"140","buildver":"1668842016","resolution":"1920x1080","__csrf":"","os":"pc","requestId":"1668842016582_0788","MUSIC_A":"bf8bfeabb1aa84f9c8c3906c04a04fb864322804c83f5d607e91a04eae463c9436bd1a17ec353cf780b396507a3f7464e8a60f4bbc019437993166e004087dd32d1490298caf655c2353e58daa0bc13cc7d5c198250968580b12c1b8817e3f5c807e650dd04abd3fb8130b7ae43fcc5b"}}`))
	t.Log(after)
}
