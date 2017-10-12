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
var EncodeTypeDefault = "pcm"

var VoiceNameList = []string{"xiaogang", "xiaoyun"}
var VoiceNameDefault = "xiaoyun"

var VolumeDefault = 50

var SampleRateDefault = 16000

func (tts TTS) GetEncodeType() string {
	if tts.EncodeType == "" {
		return EncodeTypeDefault
	}
	for _, v := range EncodeTypeList {
		if v == tts.EncodeType {
			return v
		}
	}
	return EncodeTypeDefault
	log.Println(fmt.Sprintf("TTS params EncodeType: %s unsupport! auto use %s", tts.EncodeType, EncodeTypeDefault))
	return EncodeTypeDefault
}

func (tts TTS) GetVoiceName() string {
	if tts.VoiceName == "" {
		return VoiceNameDefault
	}
	for _, v := range VoiceNameList {
		if v == tts.VoiceName {
			return v
		}
	}
	log.Println(fmt.Sprintf("TTS params VoiceName: %s unsupport! auto use %s", tts.VoiceName, VoiceNameDefault))
	return VoiceNameDefault
}

func (tts TTS) GetVolume() string {
	if tts.Volume == 0 {
		return strconv.Itoa(VolumeDefault)
	}
	if tts.Volume > 100 || tts.Volume <= 0 {
		log.Println(fmt.Sprintf("TTS params Volume: %d unsupport! auto use %d", tts.Volume, VolumeDefault))
		return strconv.Itoa(VolumeDefault)
	}
	return strconv.Itoa(tts.Volume)
}

func (tts TTS) GetSampleRate() string {
	if tts.SampleRate == 0 {
		return strconv.Itoa(SampleRateDefault)
	}
	if tts.SampleRate != 8000 || tts.SampleRate != 16000 {
		log.Println(fmt.Sprintf("TTS params SampleRate: %d unsupport! auto use %d", tts.SampleRate, SampleRateDefault))
		return strconv.Itoa(SampleRateDefault)
	}
	return strconv.Itoa(tts.SampleRate)
}

func (tts TTS) GetSpeechRate() string {
	if tts.SpeechRate == 0 {
		return strconv.Itoa(0)
	}
	if tts.SpeechRate > 500 || tts.SpeechRate <= -500 {
		log.Println(fmt.Sprintf("TTS params SpeechRate: %d unsupport! auto use %d", tts.SpeechRate, 0))
		return strconv.Itoa(0)
	}
	return strconv.Itoa(tts.SpeechRate)
}

func (tts TTS) GetPitchRate() string {
	if tts.PitchRate == 0 {
		return strconv.Itoa(0)
	}
	if tts.PitchRate > 500 || tts.PitchRate <= -500 {
		log.Println(fmt.Sprintf("TTS params PitchRate: %d unsupport! auto use %d", tts.PitchRate, 0))
		return strconv.Itoa(0)
	}
	return strconv.Itoa(tts.PitchRate)
}
