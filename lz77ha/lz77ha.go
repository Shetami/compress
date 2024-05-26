package lz77ha

import (
	"fmt"
	"os"

	"github.com/Shetami/compress/haffman"
	"github.com/Shetami/compress/lz77"
)
  
func write_bin(file string, encodedData []byte) {
	encodedFile := file
    err := os.WriteFile(encodedFile, encodedData, 0644)
    if err != nil {
        fmt.Println("Error writing encoded file:", err)
        return
    }
}

func LZ77HA(){
	filename := "test.txt"

    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
	tokens := lz77.Compress(data)
	tokens_str := lz77.Convert(tokens)
	ha_transformed, root := haffman.HaffmanEncode2(tokens_str)
	write_bin("compressed.lzh", []byte(ha_transformed))

	ha_decode := haffman.HaffmanDecode2(ha_transformed, root)
	tokens = lz77.InverseConvert(ha_decode, lz77.WindowSize)

	output := lz77.Decompress(tokens)
	err = os.WriteFile("decompress.txt", output, 0644)
	if err != nil {
		fmt.Println("Error writing decompressed file:", err)
	}

}