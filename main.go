package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type info struct {
	key   byte
	total int
	lhs   int
	rhs   int
}

type kv struct {
	key   string
	value int
}

func lhs(ch byte) bool {
	return (strings.Contains("zxqjkhryeaiou", string(ch)))
}

func main() {
	b, err := ioutil.ReadFile("words.txt")
	if err != nil {
		os.Exit(-1)
	}

	digrams := make(map[string]int)
	freqs := make(map[byte]info)
	words := strings.Fields(string(b))

	for _, w := range words {
		var chk, chv byte
		var key string
		var stats info

		for i := 1; i < len(w); i++ {
			key = w[i-1 : i+1]
			digrams[key]++

			chk = w[i]
			chv = w[i-1]
			for i := 0; i < 2 && chk != chv; i++ {
				stats = freqs[chk]
				if lhs(chv) {
					stats.lhs++
				} else {
					stats.rhs++
				}
				freqs[chk] = stats
				chk, chv = chv, chk
			}
		}

		for i := 0; i < len(w); i++ {
			chk = w[i]
			stats = freqs[chk]
			stats.total++
			freqs[chk] = stats
		}

	}

	sortDigs := make([]kv, 0)
	for k, v := range digrams {
		sortDigs = append(sortDigs,
			kv{key: k,
				value: v})
	}

	sort.Slice(sortDigs, func(i, j int) bool {
		return (sortDigs[i].value > sortDigs[j].value || (sortDigs[i].value == sortDigs[j].value && sortDigs[i].key < sortDigs[j].key))
	})

	for i, v := range sortDigs {
		fmt.Printf("%v(%v)\t", v.key, v.value)
		if i%14 == 13 {
			fmt.Println()
		}
	}

	fmt.Println("")

	sortFreqs := make([]info, 0)
	for k, v := range freqs {
		sortFreqs = append(sortFreqs,
			info{key: k,
				total: v.total,
				lhs:   v.lhs,
				rhs:   v.rhs,
			})
	}

	sort.Slice(sortFreqs, func(i, j int) bool {
		return (sortFreqs[i].total > sortFreqs[j].total)
	})

	lhsTot := 0
	grandTot := 0
	for _, v := range sortFreqs {
		if lhs(v.key) {
			lhsTot += v.total
		}
		grandTot += v.total
		fmt.Printf("%v(%v:%v:%v) ", string(v.key), v.total, v.lhs, v.rhs)
	}

	fmt.Printf("\nLHS = %v out of %v", lhsTot, grandTot)

}
