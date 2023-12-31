package day7

import (
	"fmt"
	"robertbrignull/adventofcode2023/shared"
	"sort"
	"strconv"
)

type Card int

const (
	Ace Card = iota
	King
	Queen
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
	Five
	Four
	Three
	Two
	Joker
)

func (x Card) String() string {
	switch x {
	case Ace:
		return "Ace"
	case King:
		return "King"
	case Queen:
		return "Queen"
	case Jack:
		return "Jack"
	case Ten:
		return "Ten"
	case Nine:
		return "Nine"
	case Eight:
		return "Eight"
	case Seven:
		return "Seven"
	case Six:
		return "Six"
	case Five:
		return "Five"
	case Four:
		return "Four"
	case Three:
		return "Three"
	case Two:
		return "Two"
	case Joker:
		return "Joker"
	default:
		return fmt.Sprintf("%d", int(x))
	}
}

type Hand struct {
	cards    [5]Card
	bid      int
	handType HandType
}

func readCard(c byte, useJokers bool) (Card, error) {
	switch c {
	case 'A':
		return Ace, nil
	case 'K':
		return King, nil
	case 'Q':
		return Queen, nil
	case 'J':
		if useJokers {
			return Joker, nil
		} else {
			return Jack, nil
		}
	case 'T':
		return Ten, nil
	case '9':
		return Nine, nil
	case '8':
		return Eight, nil
	case '7':
		return Seven, nil
	case '6':
		return Six, nil
	case '5':
		return Five, nil
	case '4':
		return Four, nil
	case '3':
		return Three, nil
	case '2':
		return Two, nil
	default:
		return 0, fmt.Errorf("unrecognised card: %c", c)
	}
}

func readHand(line string, useJokers bool) (Hand, error) {
	hand := Hand{}

	for i := 0; i < 5; i++ {
		card, err := readCard(line[i], useJokers)
		if err != nil {
			return Hand{}, err
		}
		hand.cards[i] = card
	}

	bid, err := strconv.Atoi(line[6:])
	if err != nil {
		return Hand{}, err
	}
	hand.bid = bid

	hand.handType = getHandType(hand.cards, useJokers)

	return hand, nil
}

func readHands(lines []string, useJokers bool) ([]Hand, error) {
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		hand, err := readHand(line, useJokers)
		if err != nil {
			return nil, err
		}
		hands[i] = hand
	}
	return hands, nil
}

type HandType int

const (
	FiveOfAKind HandType = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func (x HandType) String() string {
	switch x {
	case FiveOfAKind:
		return "FiveOfAKind"
	case FourOfAKind:
		return "FourOfAKind"
	case FullHouse:
		return "FullHouse"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case TwoPair:
		return "TwoPair"
	case OnePair:
		return "OnePair"
	case HighCard:
		return "HighCard"
	default:
		return fmt.Sprintf("%d", int(x))
	}
}

func getHandTypeWithoutJokers(handCards [5]Card) HandType {
	cards := [5]Card{}
	copy(cards[:], handCards[:])

	sort.SliceStable(cards[:], func(i, j int) bool {
		return cards[i] < cards[j]
	})

	if cards[0] == cards[4] {
		return FiveOfAKind
	}

	if cards[0] == cards[3] || cards[1] == cards[4] {
		return FourOfAKind
	}

	if (cards[0] == cards[2] && cards[3] == cards[4]) || (cards[0] == cards[1] && cards[2] == cards[4]) {
		return FullHouse
	}

	if cards[0] == cards[2] || cards[1] == cards[3] || cards[2] == cards[4] {
		return ThreeOfAKind
	}

	if (cards[0] == cards[1] && cards[2] == cards[3]) || (cards[1] == cards[2] && cards[3] == cards[4]) || (cards[0] == cards[1] && cards[3] == cards[4]) {
		return TwoPair
	}

	if cards[0] == cards[1] || cards[1] == cards[2] || cards[2] == cards[3] || cards[3] == cards[4] {
		return OnePair
	}

	return HighCard
}

func getHandType(cards [5]Card, useJokers bool) HandType {
	handType := getHandTypeWithoutJokers(cards)

	if !useJokers {
		return handType
	}

	numJokers := 0
	for _, card := range cards {
		if card == Joker {
			numJokers++
		}
	}

	if numJokers == 0 {
		return handType
	}

	switch handType {
	case FiveOfAKind:
		return FiveOfAKind

	case FourOfAKind:
		return FiveOfAKind

	case FullHouse:
		return FiveOfAKind

	case ThreeOfAKind:
		return FourOfAKind

	case TwoPair:
		if numJokers == 1 {
			return FullHouse
		} else {
			return FourOfAKind
		}

	case OnePair:
		return ThreeOfAKind
	}

	// case HighCard
	return OnePair
}

func compareHands(handA Hand, handB Hand) bool {
	if handA.handType != handB.handType {
		return handA.handType > handB.handType
	}

	for i := 0; i < 5; i++ {
		if handA.cards[i] != handB.cards[i] {
			return handA.cards[i] > handB.cards[i]
		}
	}

	return true
}

// Time taken: 51 minutes
func Part1() (string, error) {
	lines, err := shared.ReadFileLines("days/day7/input.txt")
	if err != nil {
		return "", err
	}

	hands, err := readHands(lines, false)
	if err != nil {
		return "", err
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	winnings := 0
	for i, hand := range hands {
		winnings += hand.bid * (i + 1)
	}

	return strconv.Itoa(winnings), nil
}

// Time taken: 31 minutes
func Part2() (string, error) {
	lines, err := shared.ReadFileLines("days/day7/input.txt")
	if err != nil {
		return "", err
	}

	hands, err := readHands(lines, true)
	if err != nil {
		return "", err
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	winnings := 0
	for i, hand := range hands {
		winnings += hand.bid * (i + 1)
	}

	return strconv.Itoa(winnings), nil
}
