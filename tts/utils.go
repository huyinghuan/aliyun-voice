package tts

import "strings"

//长文本处理
func processLogText(plain string)[]string{
  plain = strings.Replace(plain,"、", "、|", -1)
  plain = strings.Replace(plain,"，", "，|", -1)
  plain = strings.Replace(plain,"。", "。|", -1)
  plain = strings.Replace(plain,"；", "；|", -1)
  plain = strings.Replace(plain,"？", "？|", -1)
  plain = strings.Replace(plain,"！", "！|", -1)
  plain = strings.Replace(plain,",", ",|", -1)
  plain = strings.Replace(plain,";", ";|", -1)
  plain = strings.Replace(plain,"\\?", "?|", -1)
  plain = strings.Replace(plain,"!", "!|", -1)
  arr:= strings.Split(plain, "|")
  textArr:=[]string{}
  i := 0
  length:=len(arr)
  for i < length{
    item := strings.TrimSpace(arr[i])
    if len(item) > 100{
      textArr = append(textArr, item)
      i = i + 1
      continue
    }
    for len(item) < 100 && i < length - 1{
      if len(item + arr[i+1]) > 200{
        break
      }
      item = item + arr[i+1]
      i = i + 1
    }
    textArr = append(textArr, item)
    i = i + 1
  }
  return textArr
}

//处理合成音频字节数组
func getMergeBytesByType(audioType string, voiceBodyList [][]byte)(voiceAll []byte){
  // 不是wav不用处理
  if audioType != "wav"{
    for i:=0; i < len(voiceBodyList); i++{
      voiceAll = append(voiceAll, voiceBodyList[i]...)
    }
    return
  }
  // wav 合并文件的时候需要去掉头部的44字节文件信息， 然后重新填充头部关于文件信息的数据
  for i:=0; i < len(voiceBodyList); i++{
    if i==0{
      voiceAll = append(voiceAll, voiceBodyList[i]...)
    }else{
      voiceAll = append(voiceAll, voiceBodyList[i][44:]...)
    }
  }
  voiceAll[4] = (byte)((len(voiceAll) - 8)  & 0xff)
  voiceAll[5] = (byte)(((len(voiceAll) - 8) >> 8) & 0xff)
  voiceAll[5] = (byte)(((len(voiceAll) - 16) >> 8) & 0xff)
  voiceAll[5] = (byte)(((len(voiceAll) - 24) >> 8) & 0xff)
  voiceAll[40] = (byte)((len(voiceAll) - 44)  & 0xff)
  voiceAll[41] = (byte)(((len(voiceAll) - 44) >> 8) & 0xff)
  voiceAll[42] = (byte)(((len(voiceAll) - 44) >> 8) & 0xff)
  voiceAll[43] = (byte)(((len(voiceAll) - 44) >> 8) & 0xff)
  return
}
