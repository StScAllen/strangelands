package main

import "fmt"

var character Character
var apprentice Character
var keep Keep

var game Game
var villages []Village
var actors []Character	// important npcs to track.

var log Log

const VERSION = ".18a"

const DEBUG_ON = true

const BASE_ACTIONS = "[s. status   i. inventory   m. mission   w. world map   h. minutiae]"

type Game struct {
	gameDay        		int
	dayCounter     		int
	weekCounter    		int
	monthCounter   		int
	itemInstanceId 		int
	missionInstanceId	int
	charInstanceId 		int
	darkness			int		// value of total darkness, impacts encounters, missions, etc.
	disposition			int    // 0 in village, 1 in battlegrid, 2 intro/exit menus
	battleGrid 			*BattleGrid
}

func init() {
	game.gameDay = 1
	game.itemInstanceId = 1
	game.charInstanceId = 2
	game.darkness = 5
	log = openLog()
	dieInit()
}

func finalExit() {
	clearConsole()
	log.displayLog()
	log.writeToFile()
}

func showEndDayDetails(details string) {
	clearConsole()
	
	fmt.Println(details)
	
	showPause("So ends another day. Your toils continue...");
}

func endDay(restQuality int, showDetails bool) {
	var die Die
	details := ""
	
	details += "The consuming darkness marks an end to your day's labors...\n"
	game.darkness++

	if character.hp == character.maxhp {
		details += "... You are healthy.\n"
	} else {
		roll := die.rollxdx(1, 4)
		if roll <= restQuality {
			details += "... Your wounds are healing nicely.\n"
			character.hp += 1
		} else {
			details += "... Your wounds show little improvement.\n"
		}
	}
	if character.hp > character.maxhp {
		character.hp = character.maxhp
	}
	
	if apprentice.exists() {
		if apprentice.hp == apprentice.maxhp {
			details += "... Your apprentice is healthy.\n"
		} else {
			roll := die.rollxdx(1, 4)
			if roll <= restQuality {
				details += "... Your apprentice's wounds are healing nicely.\n"
				apprentice.hp += 1
			} else {
				details += "... Your apprentice's wounds show little improvement.\n"
			}
		}
		if apprentice.hp > apprentice.maxhp {
			apprentice.hp = apprentice.maxhp
		}
	}
		
	game.gameDay++
	game.dayCounter++

	// remove any missions that expired
	removeExpiredMissions()
	// small chance a mission will be added to a random villages bulletin board.
	maybeAddMission()
	
	details += keep.endDay()  // do keep maintenance
	
	// update villages
	for k := 0; k < len(villages); k++ {
		villages[k].endDay()
	}
	
	// do end of week stuff
	if game.dayCounter == 7 {
		game.weekCounter++
		game.dayCounter = 0
		//		showWeekEnd()

		// do end of month stuff
		if game.weekCounter == 4 {
			game.monthCounter++
			updateShops()
			game.weekCounter = 0
			//			showMonthEnd()
		}
	} else {
		//		showDayEnd()
	}
	
	if showDetails {
		showEndDayDetails(details)
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
			character = createPlayerCharacter()
			character.printCharacter(0)

			fmt.Print("\nUse this character? ")
			fmt.Scanln(&rsp)
			clearConsole()
		}

		keep = createKeep()
		buildVillages()
		updateShops()
		mission = getBlankMission()
		buildOrphanage()
		actors = make([]Character, 0, 0)	// build empty actors holder

		save()

	} else if rsp == "2" {
		// this needs to be done, before loading saves.
		buildVillages()
		err = loadGame()
		if err == -1 {
			showPause("Game File is missing or corrupted!")
			return
		}
		character.printCharacter(1)
		if apprentice.exists() {
			apprentice.printCharacter(1)
		}

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
