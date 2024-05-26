package bwtrle

import (
	"fmt"
	"os"

	"github.com/Shetami/compress/bwt"
	"github.com/Shetami/compress/rle"
)

func BwtRle(f string){
	

    data, err := os.ReadFile(f)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    // Сжатие текста
	transformed:= bwt.BWT(string(data))
    compressedText := rle.Compress(string(transformed))

	// Запись сжатого текста в файл
    outputFile, err := os.Create("compressed.br")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer outputFile.Close()

    _, err = outputFile.WriteString(string(compressedText))
    if err != nil {
        fmt.Println("Ошибка записи в файл:", err)
        return
    }

    fmt.Println("Текст успешно сжат и записан в файл compressed.txt")

	decompressedText := rle.DecompressRLE(compressedText)
	original := bwt.Invert(decompressedText)

	outputFileDecompressed, err := os.Create("decompressed.txt")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer outputFileDecompressed.Close()

    _, err = outputFileDecompressed.WriteString(string(original))
    if err != nil {
        fmt.Println("Ошибка записи в файл:", err)
        return
    }

    fmt.Println("Сжатый текст успешно распакован и записан в файл decompressed.txt")

}