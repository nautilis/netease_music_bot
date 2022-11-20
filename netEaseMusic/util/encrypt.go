package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"strings"
)

var iv = []byte("0102030405060708")
var base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var publicKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB`
var eapiKey = "e82ckenh8dichen8"

func weapiEncrypt(jsonByte []byte) string {

	presetKey := []byte("0CoJUm6Qyw8W8jud")

	randBytes := make([]byte, 16)
	_, err := rand.Read(randBytes)
	if err != nil {
		log.Fatal(err.Error())
	}
	key := []byte{}
	for _, b := range randBytes {
		key = append(key, base62[b%62])
	}
	firstEncode := AesEncrypt(jsonByte, presetKey, iv)
	params := AesEncrypt([]byte(firstEncode), key, iv)
	encSecKey, _ := RsaEncodeNoPadding([]byte(Reverse(string(key))))
	return fmt.Sprintf(`params=%s&encSecKey=%s`, url.QueryEscape(params), url.QueryEscape(encSecKey))

}

func EapiEncrypt(url string, jsonByte []byte) string {
	message := fmt.Sprintf("nobody%suse%smd5forencrypt", url, string(jsonByte))
	//fmt.Println("message => ", message)
	h := md5.New()
	h.Write([]byte(message))
	digest := hex.EncodeToString(h.Sum(nil))
	//fmt.Println("digest => ", digest)
	data := fmt.Sprintf("%s-36cd479b6b5-%s-36cd479b6b5-%s", url, string(jsonByte), digest)
	//fmt.Println("data => ", data)
	return fmt.Sprintf("params=%s", strings.ToUpper(AesEncryptEcb([]byte(data), []byte(eapiKey), nil)))

}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(data []byte, key []byte, iv []byte) (b64 string) {
	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCEncrypter(block, iv)
	dataAfterPadding := PKCS7Padding(data, aes.BlockSize)
	ciphertext := make([]byte, len(dataAfterPadding))
	mode.CryptBlocks(ciphertext, dataAfterPadding)
	b64 = base64.StdEncoding.EncodeToString(ciphertext)
	return
}

func AesEncryptEcb(data []byte, key []byte, iv []byte) string {
	block, _ := aes.NewCipher(key)
	mode := NewECBEncrypter(block)
	dataAfterPadding := PKCS7Padding(data, aes.BlockSize)
	ciphertext := make([]byte, len(dataAfterPadding))
	mode.CryptBlocks(ciphertext, dataAfterPadding)
	return hex.EncodeToString(ciphertext)
}

func AesDecrypt(encryptedData, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	mode := cipher.NewCBCDecrypter(block, iv)
	originData := make([]byte, len(encryptedData))
	mode.CryptBlocks(originData, encryptedData)
	originData = PKCS7UnPadding(originData)
	return originData
}

func RsaEncodeNoPadding(originData []byte) (string, error) {
	key, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", err
	}

	pk, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	pub := pk.(*rsa.PublicKey)
	encrypted := new(big.Int)
	e := big.NewInt(int64(pub.E))
	payload := new(big.Int).SetBytes(originData)
	encrypted.Exp(payload, e, pub.N)
	return fmt.Sprintf("%x", encrypted.Bytes()), nil
}

func Reverse(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range []rune(s) {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}
