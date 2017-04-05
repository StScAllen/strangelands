package main

import "fmt"
import "os"
import "os/exec"

var character Character
var apprentice Character
var keep Keep
var gameDay, dayCounter int

var villages []Village

var log Log

const VERSION = ".04a"

// this needs to be command prompt generic
func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showPause(messge string) {
	fmt.Println(messge)
	rsp := ""
	fmt.Scanln(&rsp)
}

func finalExit() {
	clearConsole()
	log.displayLog()
	log.writeToFile()
}

func endDay() {
	character.hp += 1

	if character.hp > character.maxhp {
		character.hp = character.maxhp
	}

	gameDay++
	dayCounter++

	if dayCounter == 21 {
		dayCounter = 0
		updateShops()
	}
}

func showGameMenu() string {

	clearConsole()

	fmt.Println("Strange Lands   (v" + VERSION + ")")
	fmt.Println("------------")
	fmt.Println("1. New Game")
	fmt.Println("2. Load Game")
	fmt.Println("3. Quit")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	if rsp == "" {
		rsp = showGameMenu()
	}

	return rsp
}

func init() {
	gameDay = 1
	log = openLog()
	dieInit()
}

func main() {
	rsp := showGameMenu()

	defer finalExit()

	if rsp == "1" { // new game, make a character
		rsp = "n"
		for rsp != "y" && rsp != "Y" {
			character = createCharacter()
			character.printCharacter(0)

			fmt.Print("\nUse this character? ")
			fmt.Scanln(&rsp)
			clearConsole()
		}
		
		keep = createKeep()
		buildVillages()
		updateShops()
		
	} else if rsp == "2" {
		character, keep = loadGame()
		buildVillages()
		updateShops()
		character.printCharacter(1)
		character.gold = 800
	} else if rsp == "3" {
		return
	}

	clearConsole()

	rsp = ""
	gameFlag := true
	for gameFlag {
		if (character.villageIndex == 99){
			rsp = keep.visitKeep()
		} else if character.villageIndex < 9 {
			rsp = villages[character.villageIndex].visitVillage()		
		}
		
		if (rsp == "q"){
			gameFlag = false
		}
	}

}
