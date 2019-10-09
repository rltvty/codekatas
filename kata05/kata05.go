package kata05

import (
	"bufio"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/boljen/go-bitmap"
)

/*
So, this kata is fairly straightforward. Implement a Bloom filter based spell checker.
You’ll need some kind of bitmap, some hash functions, and a simple way of reading in
the dictionary and then the words to check. For the hash function, remember that you
can always use something that generates a fairly long hash (such as MD5) and then take
your smaller hash values by extracting sequences of bits from the result.

Play with using different numbers of hashes, and with different bitmap sizes.
*/

const bitmapLength = 256

func populateBloomFilter(bitsPerSection int) []bool {
	file, err := os.Open("./wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	output := make([]bool, int(math.Exp2(float64(bitsPerSection))))

	fileReader := bufio.NewReader(file)
	keepReading := true

	for keepReading {
		line, err := fileReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				keepReading = false
			} else {
				log.Fatal(err)
			}
		}

		flags := splitHash(getHash(line), bitsPerSection)
		for _, v := range flags {
			output[v] = true
		}
	}

	return output
}

func spellCheck(bloomFilter []bool, word string, bitsPerSection int) bool {
	flags := splitHash(getHash(word), bitsPerSection)
	for _, v := range flags {
		if bloomFilter[v] == false {
			return false
		}
	}
	return true
}

func getHash(word string) string {
	data := []byte(word)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func splitHash(hash string, bitsPerSection int) []uint16 {
	if hash == "" {
		return []uint16{}
	}
	bitsInHash := len(hash) * 4
	outputLength := int(math.Ceil(float64(bitsInHash) / float64(bitsPerSection)))

	inMap := getBitmap(hash)
	outMap := bitmap.NewTS(bitsPerSection)

	output := make([]uint16, outputLength)
	inLoc, outLoc := 0, 0
	for {
		outMap.Set(outLoc, inMap[bitsInHash-inLoc-1])

		if outLoc++; outLoc >= bitsPerSection {
			outLoc = 0
			setOutput(&output, outMap, inLoc, bitsPerSection)
			outMap = bitmap.NewTS(bitsPerSection)
		}

		if inLoc++; inLoc >= bitsInHash {
			if outLoc > 0 {
				setOutput(&output, outMap, inLoc, bitsPerSection)
			}
			break
		}
	}
	return output
}

func setOutput(output *[]uint16, outMap *bitmap.Threadsafe, inLoc int, bitsPerSection int) {
	outData := outMap.Data(false)
	for len(outData) < 2 {
		outData = append(outData, 0)
	}
	outValue := binary.LittleEndian.Uint16(outData)
	(*output)[inLoc/bitsPerSection] = outValue
}

func reverseBytes(s []byte) []byte {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
	return s
}

func getBitmap(hash string) []bool {
	if hash == "" {
		return []bool{}
	}
	bitsInHash := len(hash) * 4

	if len(hash)%2 == 1 {
		hash = fmt.Sprintf("0%s", hash)
	}

	decoded, err := hex.DecodeString(hash)
	if err != nil {
		log.Fatal(err)
	}
	decoded = reverseBytes(decoded)
	inMap := bitmap.TSFromData(decoded, true)

	output := make([]bool, bitsInHash)
	for i := 0; i < bitsInHash; i++ {
		output[i] = inMap.Get(bitsInHash - i - 1)
	}

	return output
}
