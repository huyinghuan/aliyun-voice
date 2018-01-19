# 阿里云 智能语音识别 SDK

  GO 语言版本

----------------------------

Install

```
go get https://github.com/huyinghuan/aliyun-voice
```

## TTS 【语音合成服务】

package:
```
"github.com/huyinghuan/aliyun-voice/tts"
```

### tts.GetAuth(accessID string, accessKey string):*Auth

阿里云认证。
```
  auth := tts.GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
```

### auth.GetVoice(text string):([]byte, error)

获取语音文件字节数组，或抛出错误。
`注意！！！  有时候 语音文件会返回空字节数组，这时是由于aliyun api不稳定造成。
由于aliyun api不会报错，因此这里也不会作为错误抛出，需要自己捕获空字节数组进行处理`

###  auth.SaveVoice(text string, dist string):error

存储语音文件到指定目录【dist】，或抛出错误。

### auth.GetLongVoice(text string):([]byte, []error)

获取长文本文件字节数组，或抛出错误

### auth.SaveLongVoice(text string, dist string):[]error

存储语音文件到指定目录【dist】，或抛出错误。

### auth.TTSConfig

设置语音文件属性,参考：https://help.aliyun.com/document_detail/52793.html?spm=5176.doc30422.6.587.Z6Muvv

```
auth.TTSConfig = &auth.TTS{
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

or 

```
auth.TTSConfig.EncodeType = xxxxx 
auth.TTSConfig.Volume = 90
...
```

### Test

填写 `tts/index_test.go`的 `ALIYUNACCESSID` 和`ALIYUNACCESSKEY`

```
go test aliyun-voice/tts -v
```

### Example

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

## ASR 【语音识别服务】

package:
```
"github.com/huyinghuan/aliyun-voice/asr"
```


### ars.GetAuth(accessID string, accessKey string):*Auth

阿里云认证。
```
  auth := ars.GetAuth(ALIYUNACCESSID, ALIYUNACCESSKEY)
```

### ars.ARSConfig 音频文件属性设置
```
type ARSConfig struct{
      Model string //语音模型
      ContentType string //编码  可选: pcm, wav, opu, opus, speex
      SampleRate int //采样率
}
```

### ars.GetOneWord(voice []byte)(result map[string]string, err error)

获取语音识别结果。 返回字段见 [识别结果返回](https://help.aliyun.com/document_detail/52787.html)

注意，这里的默认音频文件格式为 `wav`, 频率为 `16000`, 语音模型为: `chat` 更多格式 见[RestfulAPI] (https://help.aliyun.com/document_detail/52787.html)


### 测试用例

见 `asr/index_test`  里面的音频文件是由```tts```通过文字生成提供的。