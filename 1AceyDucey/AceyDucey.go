package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
//map containing card number and name
var cards map[int]string = map[int]string{
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
	10: "10",
	11: "JACK",
	12: "QUEEN",
	13: "KING",
	14: "ACE",
}

//struct for the game
type AceyDuceyGame struct {
	TotalAmount int
}

//main function
func main() {
	var KeepPlaying bool

	//Show game info
	ShowGameInfo()

	//main gameplay loop
	for {
		//initialize game
		game := AceyDuceyGame{
			TotalAmount : 100,
		}
		//start game loop
		game.Play()

		//ask player, do they want to play again?
		KeepPlaying = AskPlayAgain()

		//exit if player does not want to play again
		if !KeepPlaying {
			fmt.Println("OK HOPE YOU HAD FUN")
			break
		}
	}
}

func (g AceyDuceyGame) Play() {
	//current amount, initialized to given amount
	var currAmount int = g.TotalAmount
	//main game loop
	for {
		fmt.Println()
		//show how much does player have
		ShowPlayerAmount(currAmount)

		//draw cards
		card1, card2 := DrawCards()
		//show cards
		ShowCards(card1, card2)

		//ask player to bet
		betAmount := AskPlayerToBet(currAmount)

		//draw player card, between 2 and 14 and show
		playerCard := DrawPlayerCard()

		//show player's card
		ShowPlayerCard(playerCard)

		//determine if player won
		playerWon := DidPlayerWin(card1, card2, playerCard)

		//calculate current player amount based on bet and winnings and show message
		currAmount = CalculateWinnings(currAmount, betAmount, playerWon)

		//game ends when current amount reaches 0
		if currAmount <= 0 {
			break
		}
	}
}

//show player card
func ShowPlayerCard(playerCard int) {
	fmt.Printf("YOUR CARD : %s\n", cards[playerCard])
}

//function to calculate winnings
func CalculateWinnings(currAmount, betAmount int, playerWon bool) int {
	if !playerWon {
		fmt.Println("SORRY, YOU LOSE")
		return currAmount - betAmount
	}

	fmt.Println("YOU WIN!!!")
	return currAmount + betAmount
}


//Did player win
func DidPlayerWin(card1, card2, playerCard int) bool {
	return playerCard > card1 && playerCard < card2
}

//Draw the card for the player
func DrawPlayerCard() int {
	return DrawCard(2, 14)
}

//function to ask player to bet, and react to the bet
func AskPlayerToBet(currAmount int) int {
	var bet int
	var validBet bool = false
	for {
		//ask and read bet
		fmt.Println()
		fmt.Print("WHAT IS YOUR BET : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		parsedBet, err := strconv.Atoi(scanner.Text())

		//is bet valid, or did it error out
		if err != nil {
			fmt.Println("That is an invalid bet.")
		} else if parsedBet > currAmount {
			fmt.Println("SORRY, MY FRIEND BUT YOU BET TOO MUCH")
			fmt.Printf("YOU HAVE ONLY %d DOLLARS TO BET\n", currAmount)
		} else if parsedBet >= 0 && parsedBet <= currAmount {
			validBet = true
			bet = parsedBet
		} else {
			fmt.Println("That is an invalid bet.")
		}

		//break logic
		if validBet {
			break
		}
	}
	//react to bet
	if bet <= 0 {
		bet = 0
		fmt.Println("CHICKEN!!!")
	}
	return bet
}

//function to show cards to player
func ShowCards(card1, card2 int) {
	fmt.Println("HERE ARE YOUR NEXT TWO CARDS ")
	fmt.Printf("%s %s", cards[card1], cards[card2])
	fmt.Println()
}

//function to draw cards between 2 and 14 (included)
func DrawCards() (int, int) {
	var card1, card2 int = 0, 0

	//do this till we satisfy the following criteria
	//card1 and card2 between 2 and 14 included
	//card1 < card2
	for {
		//draw card1, random between 2 and 14 (included)
		//card1
		card1 = DrawCard(2, 14)
		//card2 needs to be drawn till it's not equal to card1
		for {
			card2 = DrawCard(2, 14)
			if card1 != card2 {
				break
			}
		}
		//check card1 and card2
		if card1 < card2 {
			break
		}
	}
	return card1, card2
}

//Draw a number between two numbers included
func DrawCard(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end - start + 1) + start
}

//function to show player the amount they are left with
func ShowPlayerAmount(amount int) {
	fmt.Printf("YOU NOW HAVE : %d DOLLARS\n", amount)
}

//Ask player if they want to play again
func AskPlayAgain() bool {
	fmt.Print("TRY AGAIN (YES OR NO) : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	val := strings.ToUpper(scanner.Text())
	return  val == "YES" || val == "Y"
}

func ShowGameInfo() {
	fmt.Println(`ACEY DUCEY CARD GAME
	CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY

	ACEY-DUCEY IS PLAYED IN THE FOLLOWING MANNER
	THE DEALER (COMPUTER) DEALS TWO CARDS FACE UP
	YOU HAVE AN OPTION TO BET OR NOT BET DEPENDING
	ON WHETHER OR NOT YOU FEEL THE CARD WILL HAVE
	A VALUE BETWEEN THE FIRST TWO.
	IF YOU DO NOT WANT TO BET, INPUT A 0`)
}
