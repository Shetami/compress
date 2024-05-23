package haffman

import (
	"container/heap"
	"fmt"
	"os"
	//"strconv"
)

var Usles int

type Node struct {
    char  byte
    freq  int
    left  *Node
    right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Node)
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func buildHuffmanTree(freqMap map[byte]int) *Node {
    pq := make(PriorityQueue, len(freqMap))
    i := 0
    for char, freq := range freqMap {
        pq[i] = &Node{char: char, freq: freq}
        i++
    }
    heap.Init(&pq)

    for pq.Len() > 1 {
        left := heap.Pop(&pq).(*Node)
        right := heap.Pop(&pq).(*Node)
        parent := &Node{
            freq:  left.freq + right.freq,
            left:  left,
            right: right,
        }
        heap.Push(&pq, parent)
    }
    return heap.Pop(&pq).(*Node)
}

func buildCodeTable(root *Node, code string, codeTable map[byte]string) {
    if root == nil {
        return
    }

    if root.left == nil && root.right == nil {
        codeTable[root.char] = code
    }

    buildCodeTable(root.left, code+"0", codeTable)
    buildCodeTable(root.right, code+"1", codeTable)
}


func encode(data []byte, codeTable map[byte]string) []byte {
    var encoded []byte
    var buffer uint8 // используем 8-битный буфер для хранения битов
    var length uint8 // текущая длина буфера
    for _, char := range data {
        code := codeTable[char]
        for _, bit := range code {
            //fmt.Println(bit)
            if bit == '1' {
                buffer |= 1 << (7 - length) // устанавливаем бит в буфере
            }
            length++
            if length == 8 { // если буфер заполнен, добавляем его к закодированным данным
                encoded = append(encoded, buffer)
                buffer = 0 // сбрасываем буфер
                length = 0 // сбрасываем длину
            }
        }
    }
    if length > 0 {
        buffer <<= (7 - length) // сдвигаем буфер на оставшиеся позиции
        encoded = append(encoded, buffer)
    }
    //fmt.Println(codeTable)
    // fmt.Println(encoded)
    return encoded
}



func decode(encoded []byte, root *Node) []byte {
    var decoded []byte
    node := root
    var buffer uint8 // используем 8-битный буфер для чтения битов из закодированных данных
    
    for b := 0; b < len(encoded); b++ {
        buffer = uint8(encoded[b]) // загружаем новый байт в буфер
        for i := 7; i >= 0; i-- {
            bit := (buffer >> uint(i)) & 1 // извлекаем бит из буфера
            if bit == 0 {
                node = node.left
            } else {
                node = node.right
            }
            if node.left == nil && node.right == nil {
                decoded = append(decoded, node.char)
                node = root
            }
        }
    }
    
    
    return decoded
}

func frequencies(data []byte) map[byte]int {
	freqMap := make(map[byte]int)
    for _, char := range data {
        freqMap[char]++
    }
	return freqMap
}

func read_file(file string) []byte{
	inputFile := file
    data, err := os.ReadFile(inputFile)
    if err != nil {
        fmt.Println("Error reading input file:", err)
    }

	return data
}

func write_bin(file string, encodedData []byte) {
	encodedFile := file
    err := os.WriteFile(encodedFile, encodedData, 0644)
    if err != nil {
        fmt.Println("Error writing encoded file:", err)
        return
    }
}

func write_decode(file string, decodedData []byte) {
	err := os.WriteFile(file, decodedData, 0644)
    if err != nil {
        fmt.Println("Error writing decoded file:", err)
        return
    }
}


func HaffmanCompress() {
    // Read input file
    file := read_file("test2.txt")

    // Count character frequencies
    freqMap := frequencies(file)

    // Build Huffman tree
    root := buildHuffmanTree(freqMap)

    // Build code table
    codeTable := make(map[byte]string)
    buildCodeTable(root, "", codeTable)

    // Encode data
    encodedData := encode(file, codeTable)

    // Write encoded data to file
    write_bin("encode.txt", encodedData)

    // Decode data
    decodedData := decode(encodedData, root)

    // Write decoded data to file
    write_decode("decode.txt", decodedData)

}

func HaffmanEncode(b []byte) ([]byte, *Node){
    // Count character frequencies
    freqMap := frequencies(b)

    // Build Huffman tree
    root := buildHuffmanTree(freqMap)

    // Build code table
    codeTable := make(map[byte]string)
    buildCodeTable(root, "", codeTable)

    // Encode data
    encodedData := encode(b, codeTable)

    return encodedData, root

}

func HaffmanDecode(b []byte, n *Node) []byte{
    
    // Decode data
    decodedData := decode(b, n)

    return decodedData

}