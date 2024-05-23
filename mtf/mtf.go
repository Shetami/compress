package mtf

// encodeMTF кодирует входную строку с помощью метода MTF
func MTFEncode(input string) []byte {
	alphabet := make([]byte, 256)
	for i := range alphabet {
		alphabet[i] = byte(i)
	}

	var output []byte
	for _, char := range input {
		var index byte
		for i, val := range alphabet {
			if val == byte(char) {
				index = byte(i)
				break
			}
		}
		output = append(output, index)
		copy(alphabet[1:index+1], alphabet[0:index])
		alphabet[0] = byte(char)
	}

	return output
}

// decodeMTF декодирует входной список с помощью метода MTF
func MTFDncode(input []byte) string {
	alphabet := make([]byte, 256)
	for i := range alphabet {
		alphabet[i] = byte(i)
	}

	var output []rune
	for _, index := range input {
		char := rune(alphabet[index])
		output = append(output, char)
		copy(alphabet[1:index+1], alphabet[0:index])
		alphabet[0] = byte(char)
	}

	return string(output)
}