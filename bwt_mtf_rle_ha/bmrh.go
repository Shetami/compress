package bwtmtfrleha

import(
	"fmt"
	"os"

	"github.com/Shetami/compress/bwt"
	"github.com/Shetami/compress/haffman"
	"github.com/Shetami/compress/mtf"
	"github.com/Shetami/compress/rle"

)

func write_bin(file string, encodedData []byte) {
	encodedFile := file
    err := os.WriteFile(encodedFile, encodedData, 0644)
    if err != nil {
        fmt.Println("Error writing encoded file:", err)
        return
    }
}

func BMRH(f string){

	
	data, err := os.ReadFile(f)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    // Сжатие текста
	transformed:= bwt.BWT(string(data))
	mtf_transformed := mtf.MTFEncode(transformed)
	rle_transform := rle.RLECompress2(mtf_transformed)
	ha_transformed, root := haffman.HaffmanEncode2(string(rle_transform))

	write_bin("compressed.bmrh", []byte(ha_transformed))

	// Разжатие текста
	ha_decode := haffman.HaffmanDecode2(ha_transformed, root)
	rle_decode := rle.RLEDecompress2([]byte(ha_decode))
	mtf_decode := mtf.MTFDncode([]byte(rle_decode))
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