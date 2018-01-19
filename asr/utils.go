package asr

import (
  "encoding/binary"
)

func reverse(s []byte) {
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}

func getInt(data []byte)int{
  reverse(data)
  a := binary.BigEndian.Uint32(data)
  return int(a)
}