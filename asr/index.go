package asr

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
  "strconv"
  "encoding/json"
)
//API 阿里云地址
var API = "https://nlsapi.aliyun.com/recognize"

//Auth 认证信息配置
type authenticate struct {
	AccessID  string
	AccessKey string
  ASRConfig *ASRConfig
}

func GetAuth(accessID string, accessKey string) *authenticate {
	auth := &authenticate{
		AccessID:  accessID,
		AccessKey: accessKey,
    ASRConfig: &ASRConfig{
		  Model: "chat",
		  ContentType: "wav",
      SampleRate: 16000,
    },
	}
	return auth
}

func MD5BASE64(body[]byte)string{
  hasher := md5.New()
  hasher.Write(body)
  return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
//getAuth 获取认证字符串
func (auth authenticate) getAuth(voice []byte, date string) string {
  bodyMD5 := MD5BASE64([]byte(MD5BASE64(voice)))
  mac := hmac.New(sha1.New, []byte(auth.AccessKey))
  feature := fmt.Sprintf("%s\napplication/json\n%s\n%s\n%s",
    "POST",
      bodyMD5,
      fmt.Sprintf("audio/%s;samplerate=%d", auth.ASRConfig.ContentType, auth.ASRConfig.SampleRate),
      date)

  mac.Write([]byte(feature))
  macHash := mac.Sum(nil)
  return base64.StdEncoding.EncodeToString(macHash)
}

func (auth *authenticate) GetOneWord(voice []byte)(result map[string]string,err error){
  client := &http.Client{}
  date := time.Now().Local().Format("Mon, 02 Jan 2006 15:04:05 MST")
  apiURL := fmt.Sprintf("%s?model=%s", API, auth.ASRConfig.Model)
  req, err := http.NewRequest("POST", apiURL, bytes.NewReader(voice))
  if err != nil {
    return nil, err
  }
  req.Header.Add("Authorization", fmt.Sprintf("Dataplus %s:%s", auth.AccessID, auth.getAuth(voice, date)))
  req.Header.Add("Content-Type", fmt.Sprintf("audio/%s;samplerate=%d", auth.ASRConfig.ContentType, auth.ASRConfig.SampleRate))
  req.Header.Add("accept", "application/json")
  req.Header.Add("Content-length", strconv.Itoa(len(voice)))
  req.Header.Add("date", date)
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }
  json.Unmarshal(respBody, &result)
  return result, nil
}
