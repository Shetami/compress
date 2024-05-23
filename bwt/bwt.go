package bwt

import (
	"sort"
    "strings"
)

type Suffix struct {
	index int
	rank  [2]int
}

func buildSuffixArray(txt string, n int) []int {
	suffixes := make([]Suffix, n)
	for i := 0; i < n; i++ {
		suffixes[i].index = i
		suffixes[i].rank[0] = int(txt[i]) - int('a')
		if i+1 < n {
			suffixes[i].rank[1] = int(txt[i+1]) - int('a')
		} else {
			suffixes[i].rank[1] = -1
		}
	}


	sort.Slice(suffixes, func(i, j int) bool {
		if suffixes[i].rank[0] == suffixes[j].rank[0] {
			return suffixes[i].rank[1] < suffixes[j].rank[1]
		}
		return suffixes[i].rank[0] < suffixes[j].rank[0]
	})

	ind := make([]int, n)
	k := 4
	for k < 2*n {
		rank := 0
		prevRank := suffixes[0].rank[0]
		suffixes[0].rank[0] = rank
		ind[suffixes[0].index] = 0

		for i := 1; i < n; i++ {
			if suffixes[i].rank[0] == prevRank && suffixes[i].rank[1] == suffixes[i-1].rank[1] {
				prevRank = suffixes[i].rank[0]
				suffixes[i].rank[0] = rank
			} else {
				prevRank = suffixes[i].rank[0]
				rank++
				suffixes[i].rank[0] = rank
			}
			ind[suffixes[i].index] = i
		}

		for i := 0; i < n; i++ {
			nextIndex := suffixes[i].index + k/2
			if nextIndex < n {
				suffixes[i].rank[1] = suffixes[ind[nextIndex]].rank[0]
			} else {
				suffixes[i].rank[1] = -1
			}
		}

		sort.Slice(suffixes, func(i, j int) bool {
			if suffixes[i].rank[0] == suffixes[j].rank[0] {
				return suffixes[i].rank[1] < suffixes[j].rank[1]
			}
			return suffixes[i].rank[0] < suffixes[j].rank[0]
		})

		k *= 2
	}

	suffixArr := make([]int, n)
	for i := 0; i < n; i++ {
		suffixArr[i] = suffixes[i].index
	}

	return suffixArr
}

func findLastChar(inputText string, suffixArr []int, n int) string {
	bwtArr := make([]byte, n)
	for i := 0; i < n; i++ {
		j := suffixArr[i] - 1
		if j < 0 {
			j += n
		}
		bwtArr[i] = inputText[j]
	}
	return string(bwtArr)
}

func Invert(bwtArr string) string {
	lenBWT := len(bwtArr)
    i := 0
    for string(bwtArr[i]) != "$"{
        i++
    }
    x := i
	sortedBWT := make([]byte, lenBWT)
	copy(sortedBWT, bwtArr)
	sort.Slice(sortedBWT, func(i, j int) bool { return sortedBWT[i] < sortedBWT[j] })

	lShift := make([]int, lenBWT)
	arr := make([][]int, 128)

	for i := 0; i < lenBWT; i++ {
		arr[bwtArr[i]] = append(arr[bwtArr[i]], i)
	}

	for i := 0; i < 128; i++ {
		sort.Ints(arr[i])
	}

	for i := 0; i < lenBWT; i++ {
		lShift[i] = arr[sortedBWT[i]][0]
		arr[sortedBWT[i]] = arr[sortedBWT[i]][1:]
	}

	decoded := make([]byte, lenBWT)
	for i := 0; i < lenBWT; i++ {
		x = lShift[x]
		decoded[i] = bwtArr[x]
	}

    res := strings.ReplaceAll(string(decoded), "$", "")
	return res
}
func BWT(st string) string{
    st = st + "$"
    n := len(st)

	suffixArr := buildSuffixArray(st, n)
	bwtArr := findLastChar(st, suffixArr, n)

    return bwtArr
}


