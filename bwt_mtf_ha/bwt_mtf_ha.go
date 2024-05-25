package bwtmtfha

import (
	"fmt"
	"os"

	"github.com/Shetami/compress/bwt"
	"github.com/Shetami/compress/haffman"
	"github.com/Shetami/compress/mtf"
)

func write_bin(file string, encodedData []byte) {
	encodedFile := file
    err := os.WriteFile(encodedFile, encodedData, 0644)
    if err != nil {
        fmt.Println("Error writing encoded file:", err)
        return
    }
}

func BWT_MTF_HA(){
	filename := "test.txt"

    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    // Сжатие текста
	transformed:= bwt.BWT(string(data))
	mtf_transformed := mtf.MTFEncode(transformed)
	ha_transformed, root := haffman.HaffmanEncode2(string(mtf_transformed))
    //fmt.Println("HA:", ha_transformed)

	write_bin("compressed.bmh", []byte(ha_transformed))

	// Разжатие текста
	ha_decode := haffman.HaffmanDecode2(ha_transformed, root)
    //fmt.Println("INV:", ha_decode)
	mtf_decode := mtf.MTFDncode([]byte(ha_decode))
	bwt_decode := bwt.Invert(mtf_decode)

	outputFileDecompressed, err := os.Create("decompressed.txt")
    if err != nil {
        fmt.Println("Ошибка создания файла:", err)
        return
    }
    defer outputFileDecompressed.Close()

    _, err = outputFileDecompressed.WriteString(string(bwt_decode))
    if err != nil {
        fmt.Println("Ошибка записи в файл:", err)
        return
    }
}