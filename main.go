package main

import (
	"fmt"
	"log"
)

var (
	vies          = 10
	GuessedLetter []rune
)

type GameState struct {
	Vies          int    `json:"vies"`
	GuessedLetter []rune `json:"letters_devinees"`
	display       []rune `json:"display"`
	Mot           string `json:"mot"`
}

func main() {
	log.SetFlags(0)

	mot, err := RandomWord("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	currentDisplay := Display(mot)

	for {
		fmt.Printf("La partie va commencer. Si vous souhaiter interrompre la partie tappez 'stop' et la partie s'arrêtera.\nSi vous souhaitez reprendre là où vous l'avez interrompu tappez 'go'. Bonne chance à vous !\n")
		fmt.Printf("Vies restantes : %d\n", vies)
		fmt.Printf("lettres déjà saisies : %s\n", string(GuessedLetter))
		PositionJose()

		AsciiArt(currentDisplay, "standard.txt")

		fmt.Println("Mot actuel :", string(currentDisplay))
		guess, err := RequestLetter()
		if err != nil {
			log.Fatal(err)
		}

		if guess == "stop" {
			SaveGame(mot, currentDisplay)
			fmt.Println("Jeu sauvegardé. À bientôt !")
			break
		} else if guess == "go" {
			loadGame(&mot, &currentDisplay)
			continue
		}

		if len(guess) > 1 {
			if guess == mot {
				fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
				break
			} else {
				vies -= 2
				fmt.Printf("Mauvais mot ! Vous avez perdu 2 vies. Vies restantes : %d\n", vies)
			}
		} else {
			letter := rune(guess[0])
			if !Contains(GuessedLetter, letter) {
				GuessedLetter = append(GuessedLetter, letter)
				if !GoodLetter(mot, letter) {
					vies--
					fmt.Printf("La letter '%c' n'est pas dans le mot.\n", letter)
				} else {
					UpdateDisplay(mot, currentDisplay, letter)
				}
			} else {
				fmt.Printf("Vous avez déjà deviné la letter '%c'.\n", letter)
			}
		}

		AsciiArt(currentDisplay, "standard.txt")

		if string(currentDisplay) == mot {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", mot)
			break
		}

		if vies <= 0 {
			fmt.Println("Désolé, vous avez perdu. Le mot était :", mot)
			fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========
			`)
			break
		}
	}
}
