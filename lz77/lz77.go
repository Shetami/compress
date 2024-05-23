package lz77

import (
	"bytes"
	"fmt"
	"os"
)

const (
	windowSize = 4096
	lookaheadBufferSize = 18
)

type token struct {
	offset int
	length int
	next   byte
}

func compress(input []byte) []token {
	var tokens []token
	inputLen := len(input)
	for i := 0; i < inputLen; {
		matchOffset, matchLength := 0, 0

		for j := 1; j <= windowSize && i-j >= 0; j++ {
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

func decompress(tokens []token) []byte {
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

func LZ77() {

	inputFile := "input2.txt"
	outputFile := "compressed.lz77"
	decompressFile := "decompress.txt"

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	tokens := compress(input)
	writeCompressed(tokens, outputFile)

	tokens = readCompressed(outputFile)
	if tokens == nil {
		return
	}

	output := decompress(tokens)
	err = os.WriteFile(decompressFile, output, 0644)
	if err != nil {
		fmt.Println("Error writing decompressed file:", err)
	}
}
