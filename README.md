## 阿里云 智能语音识别 SDK

  GO 语言版本

----------------------------

Install

```
go get https://github.com/huyinghuan/aliyun-voice
```

### TTS 【语音合成服务】

package:
```
"github.com/huyinghuan/aliyun-voice/tts"
```

#### tts.GetAuth(accessID string, accessKey string)

阿里云认证。
```
  auth := tts.GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
```

##### auth.GetVoice(text string):([]byte, error)

获取语音文件字节数组，或抛出错误。

#####  auth.SaveVoice(text string, dist string):error

存储语音文件到指定目录【dist】，或抛出错误。

#### auth.tts

设置语音文件属性,参考：https://help.aliyun.com/document_detail/52793.html?spm=5176.doc30422.6.587.Z6Muvv

```
auth.tts = &auth.TTS{
  EncodeType            string
	VoiceName             string
	Volume                int //0-100
	SampleRate            int //8000 or 16000
	SpeechRate            int //-500-500
	PitchRate             int //-500-500
	TssNus                int // 0 or 1
	BackgroundMusicID     int // 0 or 1
	BackgroundMusicOffset int // default 0,
	BackgroundMusicVolume int //default 50, 0-100
}
```

#### Example

```
package main

import (
	"log"
	"os/user"
	"path/filepath"

	"github.com/huyinghuan/aliyun-voice/tts"
)

func main() {

	//ALIYUNACCESSID 阿里云 access_id
	var ALIYUNACCESSID = "Your aliyun access_id"

	//ALIYUNACCESSKEY 阿里云 access_key
	var ALIYUNACCESSKEY = "Your aliyun access_key"
	myself, _ := user.Current()
	voiceFile := filepath.Join(myself.HomeDir, "Desktop", "b.mp3")
	auth := tts.GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
	if err := auth.SaveVoice("窗前明月光，地上鞋一双。", voiceFile); err != nil {
		log.Fatalln(err)
	}
}
```