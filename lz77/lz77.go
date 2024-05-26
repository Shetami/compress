package lz77

import (
	"bytes"
	"fmt"
	"os"
)

const (
	WindowSize = 4096
	lookaheadBufferSize = 18
)

type token struct {
	offset int
	length int
	next   byte
}

func Compress(input []byte) []token {
	var tokens []token
	inputLen := len(input)
	for i := 0; i < inputLen; {
		matchOffset, matchLength := 0, 0

		for j := 1; j <= WindowSize && i-j >= 0; j++ {
			k := 0
			for k < lookaheadBufferSize && i+k < inputLen && input[i-j+k] == input[i+k] {
				k++
			}
			if k > matchLength {
				matchLength = k
				matchOffset = j
			}
		}

		next := byte(0)
		if i+matchLength < inputLen {
			next = input[i+matchLength]
		}

		tokens = append(tokens, token{matchOffset, matchLength, next})
		i += matchLength + 1
	}

	return tokens
}

func writeCompressed(tokens []token, filename string) {
	var buffer bytes.Buffer
	for _, tok := range tokens {
		buffer.WriteByte(byte(tok.offset >> 8))
		buffer.WriteByte(byte(tok.offset & 0xff))
		buffer.WriteByte(byte(tok.length))
		buffer.WriteByte(tok.next)
	}
	err := os.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error writing compressed file:", err)
	}
}

func Decompress(tokens []token) []byte {
	var buffer bytes.Buffer

	for _, tok := range tokens {
		start := buffer.Len() - tok.offset
		for i := 0; i < tok.length; i++ {
			buffer.WriteByte(buffer.Bytes()[start+i])
		}
		buffer.WriteByte(tok.next)
	}

	return buffer.Bytes()
}

func readCompressed(filename string) []token {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading compressed file:", err)
		return nil
	}

	var tokens []token
	for i := 0; i < len(data); i += 4 {
		offset := int(data[i])<<8 | int(data[i+1])
		length := int(data[i+2])
		next := data[i+3]
		tokens = append(tokens, token{offset, length, next})
	}

	return tokens
}

func Convert(tokens []token) string {
	var buffer bytes.Buffer
  
	for _, token := range tokens {
	  start := buffer.Len() - token.offset
	  for i := 0; i < token.length; i++ {
		buffer.WriteByte(buffer.Bytes()[start+i])
	  }
	  if token.next != 0 {
		buffer.WriteByte(token.next)
	  }
	}
  
	return buffer.String()
  }

  func InverseConvert(input string, windowSize int) []token {
	var tokens []token
	inputLength := len(input)
  
	for i := 0; i < inputLength; {
	  var offset, length int
	  var next byte
  
	  // Устанавливаем границы окна
	  start := i - windowSize
	  if start < 0 {
		start = 0
	  }
	  end := i
  
	  // Поиск самого длинного совпадения
	  for j := start; j < end; j++ {
		k := 0
		for k < end-j && i+k < inputLength && input[j+k] == input[i+k] {
		  k++
		}
		if k > length {
		  offset = end - j
		  length = k
		}
	  }
  
	  // Определяем следующий символ
	  if i+length < inputLength {
		next = input[i+length]
	  }
  
	  // Добавляем токен в список
	  tokens = append(tokens, token{offset: offset, length: length, next: next})
	  i += length + 1
	}
  
	return tokens
  }
  

func LZ77(f string) {

	outputFile := "compressed.lz77"
	decompressFile := "decompressed.txt"

	input, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	tokens := Compress(input)
	writeCompressed(tokens, outputFile)

	tokens = readCompressed(outputFile)
	if tokens == nil {
		return
	}

	output := Decompress(tokens)
	err = os.WriteFile(decompressFile, output, 0644)
	if err != nil {
		fmt.Println("Error writing decompressed file:", err)
	}
}
