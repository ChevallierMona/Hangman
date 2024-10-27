package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func SaveGame(mot string, currentDisplay []rune) {
	state := GameState{
		Vies:          vies,
		GuessedLetter: GuessedLetter,
		display:       currentDisplay,
		Mot:           mot,
	}

	file, err := os.Create("save.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(state); err != nil {
		log.Fatal(err)
	}
}

func loadGame(mot *string, currentDisplay *[]rune) {
	file, err := os.Open("save.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var state GameState
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&state); err != nil {
		log.Fatal(err)
	}

	vies = state.Vies
	GuessedLetter = state.GuessedLetter
	*mot = state.Mot
	*currentDisplay = state.display

	fmt.Println("Partie reprise avec succès !")
}

func AsciiArt(currentDisplay []rune, filename string) {
	asciiArtMap := make(map[rune][]string)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < 297; i++ {
		if !scanner.Scan() {
			log.Fatal("Fichier trop court")
		}
	}

	for i := 0; i < 26; i++ {
		var lines []string
		for j := 0; j < 9; j++ {
			if scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
		}
		letter := rune('a' + i)
		asciiArtMap[letter] = lines
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var asciiLines []string
	for i := 0; i < 9; i++ {
		asciiLines = append(asciiLines, "")
	}

	for _, letter := range currentDisplay {
		if lines, exists := asciiArtMap[letter]; exists {
			for i := 0; i < len(lines); i++ {
				asciiLines[i] += lines[i] + "  "
			}
		} else {

			for i := 0; i < 9; i++ {
				asciiLines[i] += "________  "
			}
		}
	}

	for _, line := range asciiLines {
		fmt.Println(line)
	}
}

func RandomWord(filename string) (string, error) {
	var words []string
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	motAleatoire := words[r.Intn(len(words))]

	return motAleatoire, nil
}

func Display(toFind string) []rune {
	display := make([]rune, len(toFind))
	lettersVisibles := (len(toFind) - 1) / 2

	indices := make([]int, len(toFind))
	for i := range indices {
		indices[i] = i
	}
	rand.Shuffle(len(indices), func(i, j int) {
		indices[i], indices[j] = indices[j], indices[i]
	})

	for i := 0; i < lettersVisibles; i++ {
		display[indices[i]] = rune(toFind[indices[i]])
	}

	for i := range display {
		if display[i] == 0 {
			display[i] = '_'
		}
	}

	return display
}

func RequestLetter() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Veuillez saisir une lettre ou un mot : ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if len(input) > 0 {
		return input, nil
	}

	return "", fmt.Errorf("aucune entrée saisie")
}

func UpdateDisplay(toFind string, currentDisplay []rune, letter rune) {
	for i, char := range toFind {
		if char == letter {
			currentDisplay[i] = letter
		}
	}
}

func GoodLetter(mot string, letter rune) bool {
	for _, char := range mot {
		if char == letter {
			return true
		}
	}
	return false
}

func Contains(letters []rune, letter rune) bool {
	for _, l := range letters {
		if l == letter {
			return true
		}
	}
	return false
}

func PositionJose() {
	switch vies {
	case 9:
		fmt.Println("=========")
	case 8:
		fmt.Println(`
      |  
      |  
      |  
      |  
      |  
=========
		`)
	case 7:
		fmt.Println(`
  +---+  
      |  
      |  
      |  
      |  
      |
=========
		`)
	case 6:
		fmt.Println(`
  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========
		`)
	case 5:
		fmt.Println(`
  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========
		`)
	case 4:
		fmt.Println(`
  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========
		`)
	case 3:
		fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========
		`)
	case 2:
		fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========
		`)
	case 1:
		fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========
		`)
	case 0:
		fmt.Println(`
  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========
		`)
	}
}
