package kata06

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

type wordTree struct {
	words    []string
	branches map[string]*wordTree
}

func newWordTree() *wordTree {
	return &wordTree{[]string{}, make(map[string]*wordTree)}
}

func buildWordTree() *wordTree {
	file, err := os.Open("../kata05/wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := newWordTree()

	fileReader := bufio.NewReader(file)
	keepReading := true

MAINLOOP:
	for keepReading {
		word, err := fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				keepReading = false
			} else {
				log.Fatal(err)
			}
		}
		word = strings.ToLower(strings.Trim(word, "\n"))
		chars := []string{}
		for _, char := range word {
			chars = append(chars, fmt.Sprintf("%c", char))
		}
		sort.Strings(chars)

		currentTree := output
		for _, char := range chars {
			nextTree, ok := currentTree.branches[char]
			if !ok {
				nextTree = newWordTree()
				currentTree.branches[char] = nextTree
			}
			currentTree = nextTree
		}
		for _, w := range currentTree.words {
			if w == word {
				continue MAINLOOP
			}
		}
		currentTree.words = append(currentTree.words, word)
	}

	return output
}

func processBranch(tree *wordTree, c chan []string, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(tree.words) > 1 {
		c <- tree.words
	}
	for _, branch := range tree.branches {
		wg.Add(1)
		go processBranch(branch, c, wg)
	}
}

func waitForComplete(c chan []string, wg *sync.WaitGroup) {
	wg.Wait()
	close(c)
}

func anagrams() {
	var wg sync.WaitGroup
	c := make(chan []string)
	initialTree := buildWordTree()
	wg.Add(1)
	go processBranch(initialTree, c, &wg)
	go waitForComplete(c, &wg)

	output := [][]string{}
	sets := 0
	biggestSet := []string{}
	for i := range c {
		sets++
		output = append(output, i)
		if len(i) > len(biggestSet) {
			biggestSet = i
		}
		for _, word := range i {
			if word == "crepitus" || word == "paste" || word == "punctilio" || word == "sunders" {
				fmt.Printf("Set from readme: \n")
				fmt.Println(i)
			}
		}
	}
	fmt.Printf("Found %v sets of anagrams\n", sets)
	fmt.Printf("Biggest set: \n")
	fmt.Println(biggestSet)
}
