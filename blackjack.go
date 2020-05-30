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
    //fmt.Println(card)
    if dealer == true {
      if prev + 11 == 21 || prev + 11 < 19 {
        return 11
      } else {
        return 1
      }
    }else {
      var pChoice int
      fmt.Printf("You have pulled an Ace. The value of your hand before this Ace is %v. Would you like this Ace to have a value of 11 (11) or 1 (1)?", prev)
      fmt.Scan(&pChoice)
      if pChoice == 11 {
        return 11
      }else if pChoice == 1 {
        return 1
      }else {
        fmt.Println("Not a valid input. Value is 11.")
        return 11
      }
    }

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
    dHandVal += getValue(dHand[i], true, 0)
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
    if dHandVal > 21 || pHandVal > 21 {
      if dHandVal > 21 && pHandVal > 21 {
        fmt.Println("You both bust. Nobody wins.")
        os.Exit(1)
      }else if pHandVal < 21 {
        fmt.Println("The dealer goes bust and you win. The losing hand is", dHand, "and the value is", dHandVal)
        os.Exit(1)
      }else if dHandVal < 21 {
        fmt.Println("You go bust and the dealer wins. The losing hand is", pHand, "and the value is", pHandVal)
        os.Exit(1)
      }
    }

    //ask the player if they'd like to hit
    var pInput string
    fmt.Println("Would you like to hit (hit/h) or pass (pass/p)?")
    fmt.Scan(&pInput)
    var pHit bool

    if pInput == "hit" || pInput == "h" {
      pHit = true
    }else if pInput == "pass" || pInput == "p" {
      pHit = false
    }else {
      fmt.Println("Not a valid input. Passing.")
      pHit = false
    }

    var dHit bool
    if dHandVal <= 18 {
      dHit = true
    }

    if dHit == false && pHit == false {
      if pHandVal > dHandVal {
        fmt.Println("You both decided to pass and you win! The value of your hand is", pHandVal, "and the value of the dealer's hand is", dHandVal)
        os.Exit(1)
      }else if pHandVal == dHandVal {
        fmt.Println("You both decided to pass and you tied, nobody wins. The value of your hand is", pHandVal, "and the value of the dealer's hand is", dHandVal)
        os.Exit(1)
      }else {
        fmt.Println("You both decided to pass and the dealer wins. The value of your hand is", pHandVal, "and the value of the dealer's hand is", dHandVal)
        os.Exit(1)
      }

    }

    pPrevVal := pHandVal
    dPrevVal := dHandVal


    if pHit == true {
      fmt.Println("You drew", deck[len(deck)-1])
      pHand = append(pHand, deck[len(deck)-1])//x = append(x,c...)
      deck = updateDeck(deck)
      pHandVal = pPrevVal + getValue(pHand[len(pHand)-1], false, pPrevVal)
    }
    if dHit == true {
      dHand = append(dHand, deck[len(deck)-1])//x = append(x,c...)
      deck = updateDeck(deck)
      dHandVal = dPrevVal + getValue(dHand[len(dHand)-1], true, dPrevVal)
    }

    fmt.Println("Your current hand is", pHand, "and the value of your hand is", pHandVal)


    if dHandVal > 21 || pHandVal > 21 {
      if dHandVal > 21 && pHandVal > 21 {
        fmt.Println("You both bust. Nobody wins.")
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
  if dHandVal > 21 || pHandVal > 21 {
    if dHandVal > 21 && pHandVal > 21 {
      fmt.Println("You both bust. Nobody wins.")
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
