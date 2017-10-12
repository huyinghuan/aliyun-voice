package tts

import (
	"os/user"
	"path/filepath"
	"testing"
)

//ALIYUNACCESSID 阿里云 access_id
var ALIYUNACCESSID = ""

//ALIYUNACCESSKEY 阿里云 access_key
var ALIYUNACCESSKEY = ""

func TestGetVoice(t *testing.T) {
	auth := &Auth{AccessID: ALIYUNACCESSID, AccessKey: ALIYUNACCESSKEY}
	if _, err := auth.GetVoice("你好,明天见"); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSaveVoice(t *testing.T) {
	myself, err := user.Current()
	if err != nil {
		panic(err)
	}
	voiceFile := filepath.Join(myself.HomeDir, "Desktop", "a.mp3")
	auth := &Auth{AccessID: ALIYUNACCESSID, AccessKey: ALIYUNACCESSKEY}
	if err := auth.SaveVoice("你好,明天见吧", voiceFile); err != nil {
		t.Error(err)
		t.Fail()
	}
}
