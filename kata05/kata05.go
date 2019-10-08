package kata05

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"

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

func getHash(word string) [md5.Size]byte {
	data := []byte(word)
	return md5.Sum(data)
}

func splitHash(hash string, bitsPerSection int) []uint16 {
	if hash == "" {
		return []uint16{}
	}

	decoded, err := hex.DecodeString(hash)
	if err != nil {
		log.Fatal(err)
	}
	inMap := bitmap.TSFromData(decoded, true)
	outMap := bitmap.NewTS(bitsPerSection)

	outputLength := int(math.Ceil(float64(len(hash)*4) / float64(bitsPerSection)))
	fmt.Printf("Guessing that outputLength is %v\n", outputLength)
	output := make([]uint16, outputLength)
	inLoc, outLoc := 0, 0
	for {
		outMap.Set(outLoc, inMap.Get(inLoc))

		if outLoc++; outLoc >= bitsPerSection {
			outLoc = 0
			setOutput(&output, outMap, inLoc, bitsPerSection)
			fmt.Printf("resetting outmap\n")
			outMap = bitmap.NewTS(bitsPerSection)
		}

		if inLoc++; inLoc >= len(decoded)*8 {
			fmt.Printf("at end of input, quitting FOR loop\n")
			if outLoc > 0 {
				fmt.Printf("Outputing LEFTOVER data:\n")
				setOutput(&output, outMap, inLoc, bitsPerSection)
			}
			break
		}
	}

	return output
}

func setOutput(output *[]uint16, outMap *bitmap.Threadsafe, inLoc int, bitsPerSection int) {
	fmt.Printf("Trying to get outData from outMap\n")
	outData := outMap.Data(false)
	fmt.Printf("Initial outData: %v\n", outData)
	for len(outData) < 2 {
		outData = append(outData, 0)
	}
	fmt.Printf("Trying to get outValue from outData: %v\n", outData)
	outValue := binary.LittleEndian.Uint16(outData)
	fmt.Printf("Trying to set outValue: %v to output array at location: %v\n", outValue, inLoc/bitsPerSection)
	(*output)[inLoc/bitsPerSection] = outValue
}
