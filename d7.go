package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day 6! \n")

	// Cards - Bid Amount

	// 32T3K 765
	// T55J5 684
	// KK677 28
	// KTJJT 220
	// QQQJA 483

	data, _ := Read("d7.txt")
	ans := splitData(data)
	fmt.Printf("Data: %v \n", data)
	fmt.Printf("Ans: %d  \n", ans)

}

func splitData(data []string) int {
	hands := []string{}
	bids := []int{}

	for _, v := range data {
		splitData := strings.Split(v, " ")
		bidNum, _ := strconv.Atoi(splitData[1])
		bids = append(bids, bidNum)
		hands = append(hands, splitData[0])
	}
	sHands := sortHands(hands, bids)
	ans := 0
	for i, hand := range sHands {
		ans += ((i + 1) * *&hand.Bid) // Deref at addr
	}
	return ans
}

// We will score each hand, then we can sort these based on these metrics witha basic sorting algo
type Hand struct {
	Cards string
	Bid   int
	Score int // Score of the hand,
}

func (h *Hand) p() string {
	s := fmt.Sprintf("Cards %s \n Bid: %d \n Score: %d", h.Cards, h.Bid, h.Score)
	return s
}

var CardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"~": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

// Given a map where ea key is a card number, count the number of cards in the "hand" and determine the rank
// High card, One pair, Three of a kind, Full House, Four of a Kind, Five of a Kind
func sortHands(handStrs []string, bids []int) []Hand {
	var hands []Hand

	for i, handStr := range handStrs {
		cards := handStr
		bid := bids[i]
		// score := getScoreP1(handStr)
		score := getScoreP2(handStr)
		hand := Hand{cards, bid, score}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]

		if a.Score < b.Score {
			return true
		} else if a.Score > b.Score {
			return false
		} else {
			return highestCard(a.Cards, b.Cards) // return true if a is less of a high card
		}

	})

	fmt.Printf("Sorted hands: %v \n", hands)
	return hands
}

func highestCard(aCards string, bCards string) bool {
	for i, ch := range aCards {
		a := string(ch)
		b := string(bCards[i])
		if CardValues[a] < CardValues[b] {
			return true
		} else if CardValues[a] > CardValues[b] {
			return false
		}
	}
	return false
}

func getScoreP1(handStr string) int {
	// Given a string representing cards, determine the hand
	hm := make(map[string]int)
	for _, ch := range handStr {
		addToMap(hm, string(ch))
	}

	// 7 Five of a kind 			map len = 1
	// 6 Four of a kind 			map len = 2 with a 1
	// 5 Full House of a kind 		map len = 2 with a 2
	// 4 Three of a kind 			map len = 3 with 3, 1, 1
	// 3 Two Pair	 				map len = 3 with 2,2,1
	// 2 One Pair	 				map len = 4
	// 1 High Card	 				map len = 5

	// PT 2 now we have a joker which needs to make the best hand possible
	// 7 Five of a kind 			map len = 1 			[! 5]
	// 6 Four of a kind 			map len = 2 with a 1 	[! 4, @ 1]
	// 5 Full House of a kind 		map len = 2 with a 2	[!3, @ 2]
	// 4 Three of a kind 			map len = 3 with 3, 1, 1[!3, @ 1, # 1]
	// 3 Two Pair	 				map len = 3 with 2,2,1  [!2, @ 2, # 1]
	// 2 One Pair	 				map len = 4             [!2, @ 1, # 1, $ 1]
	// 1 High Card	 				map len = 5				[!1, @ 1, # 1, $ 1, % 1]

	ml := len(hm)
	if ml == 1 {
		return 7
	} else if ml == 2 {
		// Check case
		for _, v := range hm {
			if v == 1 {
				return 6
			}
		}
		return 5
	} else if ml == 3 {
		// Check case
		for _, v := range hm {
			if v == 3 {
				return 4
			}
		}
		return 3
	} else if ml == 4 {
		return 2
	} else if ml == 5 {
		return 1
	}
	// fmt.Printf("Map: %v \n", hm)
	return 0
}

func getScoreP2(handStr string) int {
	// Given a string representing cards, determine the hand
	hm := make(map[string]int)
	for _, ch := range handStr {
		addToMap(hm, string(ch))
	}

	fmt.Printf("Map: %v \n", hm)

	// redistribute the jokers to the highest cards
	const joker = "J"
	numJokers, exists := hm[joker]
	if exists {
		// In order of values high to low, we need to increment up to 5 and then to the next to distribute the jokers
		skeys := getSortedkeys(hm)
		for _, k := range skeys {
			// currentVal := hm[k]
			// While numJokers && currentVal < 5
			for hm[k] < 5 && numJokers > 0 {
				hm[k]++
				numJokers--
			}

		}
	}
	fmt.Printf("Dist-Map: %v \n", hm)
	// PT 2 now we have a joker which needs to make the best hand possible
	// 7 Five of a kind 			map len = 1 			[! 5]
	// 6 Four of a kind 			map len = 2 with a 1 	[! 4, @ 1]
	// 5 Full House of a kind 		map len = 2 with a 2	[!3, @ 2]
	// 4 Three of a kind 			map len = 3 with 3, 1, 1[!3, @ 1, # 1]
	// 3 Two Pair	 				map len = 3 with 2,2,1  [!2, @ 2, # 1]
	// 2 One Pair	 				map len = 4             [!2, @ 1, # 1, $ 1]
	// 1 High Card	 				map len = 5				[!1, @ 1, # 1, $ 1, % 1]

	ml := len(hm)
	if ml == 1 {
		return 7
	} else if ml == 2 {
		// Check case
		for _, v := range hm {
			if v == 1 {
				return 6
			}
		}
		return 5
	} else if ml == 3 {
		// Check case
		for _, v := range hm {
			if v == 3 {
				return 4
			}
		}
		return 3
	} else if ml == 4 {
		return 2
	} else if ml == 5 {
		return 1

	}
	return 0
}

func getSortedkeys(m map[string]int) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		if k != "J" {
			keys[i] = k
			i++
		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}

func addToMap(hm map[string]int, ch string) {
	_, exists := hm[ch]
	if !exists {
		hm[ch] = 1
	} else {
		hm[ch]++
	}
}
