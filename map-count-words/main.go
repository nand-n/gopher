package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	//Open a file
	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer f.Close()
	//The frequency of words in the file
	words, err := freq(f)
	if err != nil {
		log.Fatalf("Error form freq in main : %s", err)
	}

	//Display  the word frequencies
	//For word , freq: range words {
	// fmt.Printf("%s \t\t %d\n" , word, freq)
	// }

	//sort the words frequencies
	pairs := sortWordFrequency(words)
	//Print the sorted results
	for _, pair := range pairs {
		fmt.Printf("%s \t\t %d\n", pair.key, pair.Value)
	}
	//WOrd with greatest frequency and it's frequency
	w, n, err := maxWord(words)
	if err != nil {
		log.Fatalf("Error with maxWord: %s\n ", err)
	}
	fmt.Printf("%#v has frequency %d", w, n)

}

func freq(r io.Reader) (map[string]int, error) {
	// Create  a map  to store word frequencies
	wordFreq := make(map[string]int)
	// create scanner to read the file
	s := bufio.NewScanner(r)
	s.Split((bufio.ScanWords))
	// Read each word and update the word frequency
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordFreq[word]++
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return wordFreq, nil
}

type Pair struct {
	key   string
	Value int
}

//implement the Len, Less , and swap  methods on PairList
// to satisfy the sort.Interface interface.

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortWordFrequency(m map[string]int) PairList {
	//convert  the map to pair list
	pairs := make(PairList, len(m))
	i := 0
	for key, value := range m {
		pairs[i] = Pair{key, value}
		i++
	}
	//sort the pair list
	sort.Sort(pairs)
	return pairs

}

func maxWord(m map[string]int) (string, int, error) {
	if len(m) == 0 {
		return "", 0, fmt.Errorf("no words found")
	}
	var maxKey string
	var maxValue int
	for key, value := range m {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}
	return maxKey, maxValue, nil
}
