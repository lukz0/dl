package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (wv *weightVals) completeRandomizer() {
	for i := range *wv {
		for j := range (*wv)[i] {
			for k := range (*wv)[i][j] {
				(*wv)[i][j][k] = randomAxonWeight()
			}
		}
	}
}

/*
 * [][][][][]
 * nma = length
 * nmi = 3
 * actI = 3
 *
 * [][][][x][]
 * nma = length-1
 * nmi = 0
 * actI = 0
 *
 * [x][][][x][]
 * nma = length-2
 * nmi = 0
 * actI = 1
 *
 * [x][x][][x][]
 * nma = 1
 * nmi = 1
 * actI = 4
 */

func (wv *weightVals) mutate(amount int) {
	var (
		length           int
		isMutatedArray   []bool
		notMutatedAmount int
	)
	for i := range *wv {
		for j := range (*wv)[i] {
			length += len((*wv)[i][j])
		}
	}

	notMutatedAmount = length

	isMutatedArray = make([]bool, length)

	for i := 0; i < amount; i++ {
		if notMutatedAmount <= 0 {
			panic("Amount of requested mutated axons bigger than the amount of axons in network")
		}
		notMutatedIndex := rand.Int63n(int64(notMutatedAmount))
		actualIndex := 0
		var indexCountdown int64 = notMutatedIndex

		for {
			if isMutatedArray[actualIndex] {
				actualIndex++
			} else {
				if indexCountdown > 0 {
					indexCountdown--
					actualIndex++
				} else {
					break
				}
			}
		}
		isMutatedArray[actualIndex] = true

		notMutatedAmount--
	}

	var globalIndex int
	for i := range *wv {
		for j := range (*wv)[i] {
			for k := range (*wv)[i][j] {
				if isMutatedArray[globalIndex] {
					(*wv)[i][j][k] = randomAxonWeight()
				}
				globalIndex++
			}
		}
	}
}

func randomAxonWeight() float64 {
	return (rand.Float64() - 0.5) * 2
}
