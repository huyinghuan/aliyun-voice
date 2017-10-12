package tts

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//Auth 认证信息配置
type Auth struct {
	AccessID  string
	AccessKey string
	tts       TTS
}

//API 阿里云地址
var API = "https://nlsapi.aliyun.com/speak?encode_type=mp3"

//getAuth 获取认证字符串
func (auth Auth) getAuth(text string, date string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	bodyMD5 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	mac := hmac.New(sha1.New, []byte(auth.AccessKey))
	feature := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", "POST", "audio/mp3,application/json", bodyMD5, "text/plain", date)
	mac.Write([]byte(feature))
	macHash := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macHash)
}
func (auth Auth) GetUrlParams() string {
	v := &url.Values{}
	v.Add("encode_type", auth.tts.GetEncodeType())
	v.Add("voice_name", auth.tts.GetVoiceName())
	v.Add("volume", auth.tts.GetVolume())
	v.Add("sample_rate", auth.tts.GetSampleRate())
	v.Add("speech_rate", auth.tts.GetSpeechRate())
	v.Add("pitch_rate", auth.tts.GetPitchRate())
	return v.Encode()
}

//GetVoice 根据文本获取声音
//Params: text: 合成声音的文本
//Return 声音[]byte和error
func (auth Auth) GetVoice(text string) ([]byte, error) {
	client := &http.Client{}
	date := time.Now().Local().Format("Mon, 02 Jan 2006 15:04:05 MST")
	apiURL := fmt.Sprintf("%s?%s", API, auth.GetUrlParams())
	fmt.Println(apiURL)
	req, err := http.NewRequest("POST", API, bytes.NewReader([]byte(text)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Dataplus %s:%s", auth.AccessID, auth.getAuth(text, date)))
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("accept", "audio/mp3,application/json")
	req.Header.Add("date", date)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("content-type")
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if strings.Contains(contentType, "json") {
		return nil, fmt.Errorf(string(respBody))
	}
	return respBody, nil
}

//SaveVoice 存储声音文件
func (auth Auth) SaveVoice(text string, dist string) error {
	voice, err := auth.GetVoice(text)
	if err != nil {
		return err
	}
	//open a file for writing
	file, err := os.Create(dist)
	if err != nil {
		return err
	}
	file.Write(voice)
	//file.Close()
	return nil
}
