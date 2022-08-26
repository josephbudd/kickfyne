package shuffle

const (
	shuffleFileName = "Shuffle.go"
)

var shuffleTemplate = `package shuffle

import (
	"crypto/rand"
	"math/big"
)

// RandomIndex returns a random number >= 0 && < length.
func RandomIndex(length int) (ri int) {
	if length < 1 {
		return
	}
	// max := big.NewInt(int64(i + 1))
	max := big.NewInt(int64(length))
	var bigI *big.Int
	var err error
	if bigI, err = rand.Int(rand.Reader, max); err != nil {
		return
	}
	ri = int(bigI.Int64())
	return
}

// Indexes returns a shuffled slice of each index for a slice of length sliceLen.
func Indexes(sliceLen int) (shuffled []int) {
	if sliceLen == 0 {
		return
	}
	last := int(sliceLen - 1)
	temp := make(map[int]int, sliceLen)
	var fromIndex, toIndex, randomI int
	for fromIndex = 0; fromIndex < sliceLen; fromIndex++ {
		var alreadyUsed bool
		randomI = RandomIndex(sliceLen)
		for toIndex = randomI; toIndex <= last; toIndex++ {
			if _, alreadyUsed = temp[toIndex]; !alreadyUsed {
				temp[toIndex] = fromIndex
				break
			}
		}
		if !alreadyUsed {
			// This new fromIndex was not previously used. It was just added.
			continue
		}
		for toIndex = 0; toIndex < randomI; toIndex++ {
			if _, alreadyUsed = temp[toIndex]; !alreadyUsed {
				temp[toIndex] = fromIndex
				break
			}
		}
	}
	// build shuffled
	shuffled = make([]int, sliceLen)
	for toIndex, fromIndex = range temp {
		shuffled[toIndex] = fromIndex
	}
	return
}

`
