package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "king", "queen", "jack", "ten":
		value = 10
	case "nine":
		value = 9
	case "eight":
		value = 8
	case "seven":
		value = 7
	case "six":
		value = 6
	case "five":
		value = 5
	case "four":
		value = 4
	case "three":
		value = 3
	case "two":
		value = 2
	default:
		value = 0
	}
	return value
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackJack{
		return "W"
	}
	else if isBlackJack && dealerScore>20{
		return "P"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {

}
