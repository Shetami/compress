package rle

import (
	"bytes"
  )

func RLECompress2(data []byte) []byte {
	var result bytes.Buffer
	n := len(data)
	for i := 0; i < n; i++ {
	  count := 1
	  for i+1 < n && data[i] == data[i+1] && count < 255 {
		count++
		i++
	  }
	  result.WriteByte(data[i])
	  result.WriteByte(byte(count))
	}
	return result.Bytes()
  }
  
  // RLEDecompress decompresses a byte array using Run-Length Encoding.
  func RLEDecompress2(data []byte) []byte {
	var result bytes.Buffer
	n := len(data)
	for i := 0; i < n; i += 2 {
	  value := data[i]
	  count := int(data[i+1])
	  for j := 0; j < count; j++ {
		result.WriteByte(value)
	  }
	}
	return result.Bytes()
  }