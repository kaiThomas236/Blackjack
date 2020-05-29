package main

import (
  "fmt"
  "strconv"
  //"math"
  "math/rand"
  "time"
  "os"
)

func getDeck() ([52][2]string) {
  var ranks [13]string
  for i := 2; i < 11; i++ {
    ranks[i-1] = strconv.Itoa(i)
  }
  ranks[0] = "Ace"
  ranks[10] = "Jack"
  ranks[11] = "Queen"
  ranks[12] = "King"

  suits := [4]string{"Spades", "Hearts", "Clubs", "Diamonds"}

  var deck [52][2]string
  i := 0

  for j := 0; j < len(ranks); j++ {
    for k := 0; k < len(suits); k++ {
      deck[i][0] = ranks[j]
      deck[i][1] = suits[k]
      i++
    }
  }

  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  r1.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

  return deck
}

func updateDeck(deck [][2]string) ([][2]string) {
  return deck[0:len(deck)-1]
}


func startGame(deck [][2]string) ([][2]string, [][2]string, [][2]string){

  var dHand [][2]string
  dHand = deck[len(deck)-2:len(deck)]
  deck = updateDeck(deck)
  deck = updateDeck(deck)

  var pHand [][2]string
  pHand = deck[len(deck)-2:len(deck)]
  deck = updateDeck(deck)
  deck = updateDeck(deck)
  //x = append(x,c...)
  //fmt.Println(dHand, pHand)

  return dHand, pHand, deck
}

func getValue (card [2]string, dealer bool, prev int) (int) {
  //fmt.Println(card)
  if card[0] == "Jack" || card[0] == "Queen" || card[0] == "King" {
    return 10
  }else if card[0] == "Ace" {
    //this is a default for now but ideally I want
    //if dealer:
      //if prev hand val + 11 <=21:
        //return 11
      //else return 1
    //if player:
      //ask the player what they want, print their prev hand val
    return 11
  }else {
    rank := card[0]
    val, err := strconv.Atoi(rank)
    if err != nil {
      fmt.Println("garbage")
    }
    return val
  }
}





func main() {
  //user input "Would you like to play a game of blackjack?"
  //if yes: runGame()

  temp := getDeck()
  var deck [][2]string = temp[0:len(temp)]
  //fmt.Println(len(deck))

  dHand, pHand, deck := startGame(deck)
  //fmt.Println(len(deck))
  var dHandVal int
  for i := 0; i < len(dHand); i++ {
    dHandVal += getValue(dHand[i], false, 0)
  }

  var pHandVal int
  for i := 0; i < len(pHand); i++ {
    pHandVal += getValue(pHand[i], false, 0)
  }

  fmt.Println("Your starting hand is", pHand, "and the value of your hand is", pHandVal)

  for pHandVal <= 21 && dHandVal <= 21 {
    if pHandVal == 21 {
      fmt.Println("You win! Your winning hand is", pHand, "and the value is 21.")
      os.Exit(1)
    }
    if dHandVal == 21 {
      fmt.Println("The dealer wins. The winning hand is", dHand, "and the value is 21. Try again!")
      os.Exit(1)
    }
    //ask the player if they'd like to hit
    pPrevVal := pHandVal
    dPrevVal := dHandVal
    pHit := true
    var dHit bool
    if dHandVal <= 18 {
      dHit = true
    }

    if pHit == true {
      fmt.Println("You drew", deck[len(deck)-1])
      pHand = append(pHand, deck[len(deck)-1])//x = append(x,c...)
      deck = updateDeck(deck)
      pHandVal = 0
      for i := 0; i < len(pHand); i++ {
        pHandVal += getValue(pHand[i], false, pPrevVal)
      }
    }
    if dHit == true {
      dHand = append(dHand, deck[len(deck)-1])//x = append(x,c...)
      deck = updateDeck(deck)
      dHandVal = 0
      for i := 0; i < len(dHand); i++ {
        dHandVal += getValue(dHand[i], false, dPrevVal)
      }
    }

    fmt.Println("Your current hand is", pHand, "and the value of your hand is", pHandVal)

    if dHandVal > 21 || pHandVal > 21 {
      if dHandVal > 21 && pHandVal > 21 {
        fmt.Println("You both tie. Nobody wins.")
        os.Exit(1)
      }else if pHandVal < 21 {
        fmt.Println("The dealer goes bust and you win. The losing hand is", dHand, "and the value is", dHandVal)
        os.Exit(1)
      }else if dHandVal < 21 {
      fmt.Println("You go bust and the dealer wins. The losing hand is", pHand, "and the value is", pHandVal)
      os.Exit(1)
      }
    }



  }

}
