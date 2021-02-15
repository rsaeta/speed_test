package main

// $ time ./main
//   13.317016
//   7.393607837526988
//   ./main  2.72s user 0.11s system 107% cpu 2.633 total
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type deck []int

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d deck) mean() float64 {
	sum := 0.0
	for _, e := range d {
		sum += float64(e)
	}

	return sum / float64(len(d))
}

func (d deck) stdev() float64 {
	avg := d.mean()
	sum := 0.0
	for _, e := range d {
		sum += math.Pow(float64(e)-avg, 2)
	}

	return math.Sqrt(sum / float64(len(d)))
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var results [1_000_000]int
	for i := 0; i < 1_000_000; i++ {
		lh, _ := playItr()
		results[i] = lh
	}

	var resDeck deck = results[:]

	fmt.Println(resDeck.mean())
	fmt.Println(resDeck.stdev())
}

// playItr plays a round of the game and returns the
// size of the hand leftover and the hand itself.
func playItr() (int, []int) {
	d := getDeck()
	d.shuffle()

	hand := []int{}
	for _, c := range d {
		hand = append(hand, c)
		hand = tryReduce(hand)
	}

	return len(hand), hand
}

func tryReduce(hand []int) []int {
	if len(hand) >= 4 {
		handSize := len(hand)
		lastCard := hand[handSize-1]
		fourth := hand[handSize-4]
		if sameRank(lastCard, fourth) {
			return tryReduce(hand[0 : handSize-4])
		}
		if sameSuit(lastCard, fourth) {
			handTmp := hand[0 : handSize-3]
			handTmp = append(handTmp, hand[handSize-1])
			return tryReduce(handTmp)
		}
	}

	return hand
}

func getDeck() deck {
	var d []int
	for i := 0; i < 52; i++ {
		d = append(d, i)
	}
	return d
}

func sameRank(a, b int) bool {
	return a%13 == b%13
}

func sameSuit(a, b int) bool {
	return a/13 == b/13
}

// func getCardRepr(a int) (rankStr, suitStr string) {
// 	var suits [4]string = [4]string{"C", "D", "H", "S"}
// 	suit := a / 13
// 	rank := (a % 13) + 1
// 	if rank == 1 {
// 		rankStr = "A"
// 	} else if rank == 11 {
// 		rankStr = "J"
// 	} else if rank == 12 {
// 		rankStr = "Q"
// 	} else if rank == 13 {
// 		rankStr = "K"
// 	} else {
// 		rankStr = strconv.Itoa(rank)
// 	}

// 	suitStr = suits[suit]
// 	return
// }

// func printDeck(d deck) {
// 	var deckRepr [][]string
// 	for _, card := range d {
// 		rankStr, suitStr := getCardRepr(card)
// 		deckRepr = append(deckRepr, []string{rankStr, suitStr})
// 	}

// 	fmt.Println(deckRepr)
// }
