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
  "log"
  "sync"
)

//Auth 认证信息配置
type authenticate struct {
	AccessID  string
	AccessKey string
	TTSConfig *TTS
}

func GetAuth(accessID string, accessKey string) *authenticate {
	auth := &authenticate{
		AccessID:  accessID,
		AccessKey: accessKey,
		TTSConfig: &TTS{
			EncodeType:            "mp3",
			VoiceName:             "xiaoyun",
			Volume:                50,
			SampleRate:            16000,
			SpeechRate:            0,
			PitchRate:             0,
			TssNus:                1,
			BackgroundMusicID:     -1,
			BackgroundMusicOffset: 0,
			BackgroundMusicVolume: 50,
		},
	}
	return auth
}

//API 阿里云地址
var API = "https://nlsapi.aliyun.com/speak"

//getAuth 获取认证字符串
func (auth authenticate) getAuth(text string, date string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	bodyMD5 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	mac := hmac.New(sha1.New, []byte(auth.AccessKey))
	feature := fmt.Sprintf("%s\n%s\n%s\n%s\n%s", "POST", fmt.Sprintf("audio/%s,application/json", auth.TTSConfig.GetEncodeType()), bodyMD5, "text/plain", date)
	mac.Write([]byte(feature))
	macHash := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macHash)
}
func (auth authenticate) GetUrlParams() string {
	v := &url.Values{}
	v.Add("encode_type", auth.TTSConfig.GetEncodeType())
	v.Add("voice_name", auth.TTSConfig.GetVoiceName())
	v.Add("volume", auth.TTSConfig.GetVolume())
	v.Add("sample_rate", auth.TTSConfig.GetSampleRate())
	v.Add("speech_rate", auth.TTSConfig.GetSpeechRate())
	v.Add("pitch_rate", auth.TTSConfig.GetPitchRate())
	v.Add("tts_nus", auth.TTSConfig.GetTTSnus())
	backgroundMusicID := auth.TTSConfig.GetBackgroundMusicID()
	if backgroundMusicID != "-1" {
		v.Add("background_music_id", backgroundMusicID)
		v.Add("background_music_offset", auth.TTSConfig.GetBackgroundMusicOffset())
		v.Add("background_music_volume", auth.TTSConfig.GetBackgroundMusicVolume())
	}

	return v.Encode()
}

//GetLongVoice 获取长文本声音。
func (auth authenticate) GetLongVoice(text string)(voice[]byte, err []error){
  textList := processLogText(text)
  voiceBodyList := make([][]byte, len(textList))
  var wg sync.WaitGroup
  wg.Add(len(textList))
  emptyVoice:=map[int]string{}
  for index, text := range textList{
    go func(index int, text string){
      defer wg.Done()
      body,e:=auth.GetVoice(text)
      if e!=nil{
        log.Println(e)
        emptyVoice[index] = text
        return
      }
      if len(body) == 0{
        emptyVoice[index] = text
      }else{
        voiceBodyList[index] = body
      }
    }(index, text)
  }
  wg.Wait()

  //错误情况重试最大
  tryMAXCount := 5
  tryIndex := 0
  errorMap := map[int]error{}
  for tryIndex < tryMAXCount && len(emptyVoice) != 0{
    for key, value := range emptyVoice{
      body,e:=auth.GetVoice(value)
      if e!=nil{
        errorMap[key]= e
        log.Println(e)
        return
      }
      if len(body) == 0{
        emptyVoice[key] = value
      }else{
        voiceBodyList[key] = body
        delete(emptyVoice, key)
        delete(errorMap, key)
      }
    }
    tryIndex++
  }
  if len(emptyVoice)!= 0{
    for key, value := range emptyVoice{
      errorMap[key] = fmt.Errorf("索引:%d,语音内容获取失败:%s\n",key, value)
    }
  }
  if len(errorMap)!=0{
    errList := make([]error, len(errorMap))
    for _,value:=range errorMap{
      errList = append(errList, value)
    }
    return nil, errList
  }
  return getMergeBytesByType(auth.TTSConfig.EncodeType, voiceBodyList),nil
}

//SaveLongVoice 保存长文本声音
func (auth authenticate) SaveLongVoice(text string, dist string)[]error{
  file, err := os.Create(dist)
  if err!=nil{
    return []error{err}
  }
  body, errList:=auth.GetLongVoice(text)
  if errList != nil{
    return errList
  }
  defer file.Close()
  if _, err = file.Write(body); err!=nil{
    return []error{err}
  }
  return nil
}

//GetVoice 根据文本获取声音
//Params: text: 合成声音的文本
//Return 声音[]byte和error
func (auth authenticate) GetVoice(text string) ([]byte, error) {
  text = strings.TrimSpace(text)
	client := &http.Client{}
	date := time.Now().Local().Format("Mon, 02 Jan 2006 15:04:05 MST")
	apiURL := fmt.Sprintf("%s?%s", API, auth.GetUrlParams())
	req, err := http.NewRequest("POST", apiURL, bytes.NewReader([]byte(text)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Dataplus %s:%s", auth.AccessID, auth.getAuth(text, date)))
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("accept", fmt.Sprintf("audio/%s,application/json", auth.TTSConfig.GetEncodeType()))
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
	if len(text)!= 0 && len(respBody)==0{
	  log.Printf("⚠️警告阿里云API响应内容为空，准备重试:%s\n", text,)
	  return respBody, nil
  }
	if strings.Contains(contentType, "json") {
		return nil, fmt.Errorf(string(respBody))
	}
	return respBody, nil
}

//SaveVoice 存储声音文件
func (auth authenticate) SaveVoice(text string, dist string) error {
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
	file.Close()
	return nil
}
