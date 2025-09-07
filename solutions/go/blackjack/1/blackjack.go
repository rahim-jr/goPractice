package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card{
        case "ace" : 
         	return 11 
        case "two" : 
        	return 2 
        case "three" :	
        	return 3
        case "four" :
        	return 4 
        case "five" : 
        	return 5 
        case "six" :
        	return 6
        case "seven" :
        	return 7 
        case "eight" : 
        	return 8 
        case "nine" : 
        	return 9
        case "ten", "jack", "king" ,"queen": 
        	return 10 
        default: 
        	return 0 
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	player1 := ParseCard(card1)
	player2 := ParseCard(card2)
	dealer := ParseCard(dealerCard)

	sum := player1 + player2

	// Rule 1: Pair of aces
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	// Rule 2: Blackjack
	if sum == 21 {
		if dealer == 11 || dealer == 10 { // dealer has ace or 10-value card
			return "S"
		}
		return "W"
	}

	// Rule 3: Sum 17–20 → Stand
	if sum >= 17 && sum <= 20 {
		return "S"
	}

	// Rule 4: Sum 12–16 → depends on dealer
	if sum >= 12 && sum <= 16 {
		if dealer >= 7 {
			return "H"
		}
		return "S"
	}

	// Rule 5: Sum ≤ 11 → Hit
	return "H"
}
