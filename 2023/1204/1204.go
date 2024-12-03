package main

import (
	"fmt"
	"os"
	"strings"
)

// ----- Scratchcard -----
type Scratchcard struct {
	id          string
	winningNums []string
	cardNums    []string
	numCards    int
}

func (card *Scratchcard) CalculateNumberOfWins() int {
	winners := 0

	for i := 0; i < len(card.winningNums); i++ {
		winningNum := card.winningNums[i]

		if winningNum != "" {
			for x := 0; x < len(card.cardNums); x++ {
				if card.cardNums[x] != "" {
					if card.cardNums[x] == winningNum {
						winners++
						break
					}
				}
			}
		}
	}
