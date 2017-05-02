package main

import "fmt"

var character Character
var apprentice Character
var keep Keep

var game Game
var villages []Village

var log Log

const VERSION = ".08a"

const DEBUG_ON = true

type Game struct {
	gameDay        		int
	dayCounter     		int
	weekCounter    		int
	monthCounter   		int
	itemInstanceId 		int
	missionInstanceId	int
}

func init() {
	game.gameDay = 1
	game.itemInstanceId = 1
	log = openLog()
	dieInit()
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

	game.gameDay++
	game.dayCounter++

	if game.dayCounter == 7 {
		game.weekCounter++
		game.dayCounter = 0
		//		showWeekEnd()

		if game.weekCounter == 4 {
			game.monthCounter++
			updateShops()
			game.weekCounter = 0
			//			showMonthEnd()
		}
	} else {
		//		showDayEnd()
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

func main() {
	debugPause("\n\n< [DEBUG IS ON] >")
	defer finalExit()

	BEGIN:

	rsp := showGameMenu()
	err := 0

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
		mission = genNewMission()
		save()

	} else if rsp == "2" {
		err = loadGame()
		if err == -1 {
			showPause("Game File is missing or corrupted!")
			return
		}

		buildVillages()
		updateShops()
		character.printCharacter(1)
		character.crowns = 800
	} else if rsp == "3" {
		return
	}

	clearConsole()

	rsp = ""
	gameFlag := true
	for gameFlag {
		if character.villageIndex == 99 {
			rsp = keep.visitKeep()
		} else if character.villageIndex < 9 {
			rsp = villages[character.villageIndex].visitVillage()
		}

		if rsp == "q" {
			gameFlag = false
			goto BEGIN
		}
	}

}
