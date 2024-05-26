package rle

import (
    "fmt"
    "os"
    "strconv"
	"strings"
)

func Compress(input string) string {
    if input == "" {
        return ""
    }

    compressed := ""
    count := 1
    for i := 1; i < len(input); i++ {
        if input[i] != input[i-1] {
			if count == 1{
				compressed += string(input[i-1])
			} else {
				compressed += strconv.Itoa(count) + string(input[i-1])
			}
            
            count = 1
        } else {
            count++
        }
    }
    compressed += strconv.Itoa(count) + string(input[len(input)-1])
    return compressed
}

func DecompressRLE(input string) string {
	var result strings.Builder
	countStr := ""
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			countStr += string(input[i])
		} else {
			count, _ := strconv.Atoi(countStr)
			if count == 0 {
				count = 1 // Если отсутствует префикс, устанавливаем количество повторений в 1
			}
			for j := 0; j < count; j++ {
				result.WriteByte(input[i])
			}
			countStr = ""
		}
	}
	return result.String()
}

func Rle(f string) {

    data, err := os.ReadFile(f)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Сжатие текста
    compressedText := Compress(string(data))

    // Запись сжатого текста в файл
    outputFile, err := os.Create("compressed.txt")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer outputFile.Close()

    _, err = outputFile.WriteString(compressedText)
    if err != nil {
        fmt.Println("Ошибка записи в файл:", err)
        return
    }


    // Распаковка сжатого текста
    decompressedText := DecompressRLE(compressedText)

    // Запись распакованного текста в файл
    outputFileDecompressed, err := os.Create("decompressed.txt")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer outputFileDecompressed.Close()

    _, err = outputFileDecompressed.WriteString(decompressedText)
    if err != nil {
        fmt.Println("Ошибка записи в файл:", err)
        return
    }

}
