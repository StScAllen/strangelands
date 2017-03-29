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

const VERSION = ".03a"

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

	if dayCounter == 6 {
		dayCounter = 0
		updateShops()
	}
}

func showTopMenu() string {

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
		rsp = showTopMenu()
	}

	return rsp
}

func init() {
	gameDay = 1
	log = openLog()
	dieInit()
}

func main() {
	rsp := showTopMenu()

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
		rsp = villages[character.villageIndex].visitVillage()

		if rsp == "1" {
			villages[character.villageIndex].goShop()
		} else if rsp == "2" {
			keep.goKeep()
			character.save()
		} else if rsp == "3" {
			txt := "It's cold and dark here. Shadows from my waning fire dance across the vacant expanse. Sometimes the end of the world looks as bleak and sorrowful as its beginning. But only if just."
			showDialogBoxRight(makeDialogBox("Joe Durden", txt, DIALOG_RIGHT))
		} else if rsp == "4" {
			chooseAdventure()
			adventure()
		} else if rsp == "5" {		
			drawWorldMap()
		} else if rsp == "6" {
			openMinutiae()
		} else if rsp == "7" {
			gameFlag = false
		} else if rsp == "9" {
			showTravelMenu()
		}
	}

}
