package tts

import (
	"fmt"
	"log"
	"strconv"
)

type TTS struct {
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

var EncodeTypeList = []string{"pcm", "wav", "mp3", "alaw"}

var VoiceNameList = []string{"xiaogang", "xiaoyun"}

func (tts TTS) GetEncodeType() string {
	for _, v := range EncodeTypeList {
		if v == tts.EncodeType {
			return v
		}
	}
	log.Println(fmt.Sprintf("TTS params EncodeType: %s unsupport! auto use %s", tts.EncodeType, "mp3"))
	return "mp3"
}

func (tts TTS) GetVoiceName() string {
	for _, v := range VoiceNameList {
		if v == tts.VoiceName {
			return v
		}
	}
	log.Println(fmt.Sprintf("TTS params VoiceName: %s unsupport! auto use %s", tts.VoiceName, "xiaoyun"))
	return "xiaoyun"
}

func (tts TTS) GetVolume() string {
	if tts.Volume > 100 || tts.Volume <= 0 {
		log.Println(fmt.Sprintf("TTS params Volume: %d unsupport! auto use %d", tts.Volume, 50))
		return strconv.Itoa(50)
	}
	return strconv.Itoa(tts.Volume)
}

func (tts TTS) GetSampleRate() string {
	if tts.SampleRate != 8000 && tts.SampleRate != 16000 {
		log.Println(fmt.Sprintf("TTS params SampleRate: %d unsupport! auto use %d", tts.SampleRate, 16000))
		return strconv.Itoa(16000)
	}
	return strconv.Itoa(tts.SampleRate)
}

func (tts TTS) GetSpeechRate() string {
	if tts.SpeechRate > 500 || tts.SpeechRate <= -500 {
		log.Println(fmt.Sprintf("TTS params SpeechRate: %d unsupport! auto use %d", tts.SpeechRate, 0))
		return strconv.Itoa(0)
	}
	return strconv.Itoa(tts.SpeechRate)
}

func (tts TTS) GetPitchRate() string {
	if tts.PitchRate > 500 || tts.PitchRate <= -500 {
		log.Println(fmt.Sprintf("TTS params PitchRate: %d unsupport! auto use %d", tts.PitchRate, 0))
		return strconv.Itoa(0)
	}
	return strconv.Itoa(tts.PitchRate)
}

func (tts TTS) GetTTSnus() string {
	if tts.TssNus != 0 && tts.TssNus != 1 {
		log.Println(fmt.Sprintf("TTS params TTSnus: %d unsupport! auto use %d", tts.TssNus, 1))
		return strconv.Itoa(1)
	}
	return strconv.Itoa(tts.TssNus)
}

func (tts TTS) GetBackgroundMusicID() string {
	if tts.BackgroundMusicID == -1 {
		return strconv.Itoa(-1)
	}
	if tts.BackgroundMusicID != 0 && tts.BackgroundMusicID != 1 {
		log.Println(fmt.Sprintf("TTS params BackgroundMusicID: %d unsupport! auto ignore", tts.BackgroundMusicID))
		return strconv.Itoa(-1)
	}
	return strconv.Itoa(tts.BackgroundMusicID)
}

func (tts TTS) GetBackgroundMusicOffset() string {
	return strconv.Itoa(tts.BackgroundMusicOffset)
}

func (tts TTS) GetBackgroundMusicVolume() string {
	return strconv.Itoa(tts.BackgroundMusicVolume)
}
