package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	bwtmtfha "github.com/Shetami/compress/bwt_mtf_ha"
	bwtmtfrleha "github.com/Shetami/compress/bwt_mtf_rle_ha"
	bwtrle "github.com/Shetami/compress/bwt_rle"
	"github.com/Shetami/compress/haffman"
	"github.com/Shetami/compress/lz77"
	"github.com/Shetami/compress/lz77ha"
	"github.com/Shetami/compress/rle"
)
func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Use 'compress -h' for more information")
	}


	help := `
		Usage: compress [param] [filename]

		-ha: compress Haffman method

		-r: compress Run Length Encoding method

		-bmh: compress Burrows Wheeler Transform -> Move To Front -> Haffman method

		-br: compress Burrows Wheeler Transform -> Run Length Encoding method

		-bmrh: compress Burrows Wheeler Transform -> Move To Front -> Run Length Encoding -> Haffman method

		-l: compress Lempel Ziv Welch(LZ77) method

		-lh: compress Lempel Ziv Welch(LZ77) -> Haffman method
	`
	if args[0] == "-h"{
		fmt.Println(help)
		return
	}

	inputFile := args[1]

	switch args[0]{
		case "-ha":
			haffman.Haffman2(inputFile)
		case "-r":
			rle.Rle(inputFile)
		case "-bmh":
			bwtmtfha.BWT_MTF_HA(inputFile)
		case "-br":
			bwtrle.BwtRle(inputFile)
		case "-bmrh":
			bwtmtfrleha.BMRH(inputFile)
		case "-l":
			lz77.LZ77(inputFile)
		case "-lh":
			lz77ha.LZ77HA(inputFile)
	}

	statInputInfo, _ := os.Stat(inputFile)
	
	// Путь к каталогу, в котором вы хотите искать файлы
    directory := "."

    // Регулярное выражение для сопоставления файлов по шаблону
    pattern := regexp.MustCompile(`^compressed\.[a-z]+$`)

	var fileCompressInfo fs.FileInfo
    // Функция для обработки каждого файла в каталоге
    err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Проверяем, соответствует ли имя файла регулярному выражению
        if pattern.MatchString(info.Name()) {
            // Если да, то выводим информацию о файле
            fileInfo, err := os.Stat(path)
            if err != nil {
                return err
            }
			fileCompressInfo = fileInfo
        }
        return nil
    })

    if err != nil {
        fmt.Println("Ошибка:", err)
    }

	fmt.Printf("Size: %d Kb -> %d Kb\n", statInputInfo.Size(), fileCompressInfo.Size())
}