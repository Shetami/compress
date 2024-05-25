package haffman

import (
	"container/heap"
	"fmt"
	"strings"
	"os"
  )
  
  // Huffman Tree Node
  type Node2 struct {
	char   rune
	freq   int
	left   *Node2
	right  *Node2
  }
  
  // A Min-Heap (priority queue) for the nodes
  type PriorityQueue2 []*Node2
  
  func (pq PriorityQueue2) Len() int { return len(pq) }
  
  func (pq PriorityQueue2) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
  }
  
  func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
  }
  
  func (pq *PriorityQueue2) Push(x interface{}) {
	*pq = append(*pq, x.(*Node2))
  }
  
  func (pq *PriorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
  }
  
  // Build a frequency map from the input string
  func buildFrequencyMap2(text string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, char := range text {
	  freqMap[char]++
	}
	return freqMap
  }
  
  // Build the Huffman Tree
  func buildHuffmanTree2(freqMap map[rune]int) *Node2 {
	pq := make(PriorityQueue2, 0)
	heap.Init(&pq)
  
	for char, freq := range freqMap {
	  heap.Push(&pq, &Node2{char: char, freq: freq})
	}
  
	for pq.Len() > 1 {
	  left := heap.Pop(&pq).(*Node2)
	  right := heap.Pop(&pq).(*Node2)
  
	  merged := &Node2{
		char: 0,
		freq: left.freq + right.freq,
		left: left,
		right: right,
	  }
  
	  heap.Push(&pq, merged)
	}
  
	return heap.Pop(&pq).(*Node2)
  }
  
  // Build the Huffman Codes from the Huffman Tree
  func buildHuffmanCodes(node *Node2, code string, codes map[rune]string) {
	if node == nil {
	  return
	}
  
	if node.left == nil && node.right == nil {
	  codes[node.char] = code
	}
  
	buildHuffmanCodes(node.left, code+"0", codes)
	buildHuffmanCodes(node.right, code+"1", codes)
  }
  

// Encode the input string using the Huffman Codes
func encode_2(text string, codes map[rune]string) string {
	encoded := ""
	for _, char := range text {
	  encoded += codes[char]
	}
	return encoded
  }
  
  // Decode the encoded string using the Huffman Tree
  func decode_2(encoded string, root *Node2) string {
	decoded := ""
	current := root
	for _, bit := range encoded {
	  if bit == '0' {
		current = current.left
	  } else {
		current = current.right
	  }
  
	  if current.left == nil && current.right == nil {
		decoded += string(current.char)
		current = root
	  }
	}
	return decoded
}

func bitStringToByteArray(bitString string) ([]byte, error) {
    length := len(bitString)
    byteArray := make([]byte, ((length+7)/8)+1)  // вычисляем необходимый размер массива байтов
	var postfix_byte byte
	usles_bits := (8*((length+7)/8)) - length
	postfix_byte = byte(usles_bits)

    for i := 0; i < length; i++ {
        if bitString[i] == '1' {
            byteArray[i/8] |= 1 << (7 - uint(i%8))
        } else if bitString[i] != '0' {
            return nil, fmt.Errorf("invalid character %c at position %d", bitString[i], i)
        }
    }
	byteArray[len(byteArray)-1] = postfix_byte

    return byteArray, nil
}

func byteArrayToBitString(byteArray []byte) string {
	var bitString strings.Builder
	last_byte := byteArray[len(byteArray)-1]
	byteArray = byteArray[:len(byteArray)-1]
	usles_bits := int(last_byte)
	for b := 0; b <= len(byteArray) - 1; b++{
		for i := 7; i >= 0; i-- {
			if byteArray[b]&(1<<i) != 0 {
				bitString.WriteByte('1')
			} else {
				bitString.WriteByte('0')
			}
		}
	}
	bitStringResult := bitString.String()
	bitStringResult = bitStringResult[:len(bitStringResult) - usles_bits]
	return bitStringResult
  }

func Haffman2() {
	file := read_file("micro.txt")
	fileInfo, _ := os.Stat("micro.txt")
	fmt.Printf("Size: %d Kb", fileInfo.Size())
	freqMap := buildFrequencyMap2(string(file))
	huffmanTree := buildHuffmanTree2(freqMap)
  
	codes := make(map[rune]string)
	buildHuffmanCodes(huffmanTree, "", codes)
  
	encoded := encode_2(string(file), codes)
	byteArray, err := bitStringToByteArray(encoded)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	write_bin("compressed.ha", byteArray)

	dec := byteArrayToBitString(byteArray)
  
	decoded := decode_2(dec, huffmanTree)

	write_decode("decompressed.txt", []byte(decoded))
  }
  
func HaffmanEncode2(s string) ([]byte, *Node2){
    // Count character frequencies
    freqMap := buildFrequencyMap2(s)

    // Build Huffman tree
    huffmanTree := buildHuffmanTree2(freqMap)

    // Build code table
    codes := make(map[rune]string)
	buildHuffmanCodes(huffmanTree, "", codes)
  
    // Encode data
    encoded := encode_2(s, codes)
	byteArray, err := bitStringToByteArray(encoded)
    if err != nil {
        fmt.Println("Error:", err)
    }

    return byteArray, huffmanTree

}

func HaffmanDecode2(byteArray []byte, n *Node2) string{
    
    // Decode data
    dec := byteArrayToBitString(byteArray)
  
	decoded := decode_2(dec, n)

    return decoded

}