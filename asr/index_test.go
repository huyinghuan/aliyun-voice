package asr

import (
  "path/filepath"
  "os/user"
  "io/ioutil"
  "testing"
  "log"
)

//ALIYUNACCESSID 阿里云 access_id
var ALIYUNACCESSID = ""

//ALIYUNACCESSKEY 阿里云 access_key
var ALIYUNACCESSKEY = ""

func TestGetOneWord(t *testing.T){
  auth := GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
  myself, _ := user.Current()
  voiceFile := filepath.Join(myself.HomeDir, "Desktop", "test.wav")
  voice, err:=ioutil.ReadFile(voiceFile)
  if err!=nil{
    t.Fail()
  }

  result, e:= auth.GetOneWord(voice)
  if e!=nil{
    log.Println(e)
    t.Fail()
  }
  log.Println(result)
}