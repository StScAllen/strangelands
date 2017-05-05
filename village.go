// village.go
// Village

package main

import "fmt"
import "time"
import "strings"
import "strconv"

type Village struct {
	name            string
	distanceToKeep  int
	size            int
	politicalFavor  int
	villageIndex 	int
	missions		[]Mission
	shopWeapons     []Item
	shopArmor       []Item
	shopProvisions  []Item
	shopApothecary  []Item
	shopCuriosities []Item
	mapX, mapY      int
}

func buildVillages() {
	villages = make([]Village, 8, 8)

	var crowley Village
	crowley.name = "Crowley"
	crowley.distanceToKeep = 1
	crowley.size = 1
	crowley.mapX, crowley.mapY = 20, 10
	crowley.missions = make([]Mission, 0, 0)
	crowley.villageIndex = 0
	// starting village, add a mission to bulletin board.
	crowley.missions = append(crowley.missions, genNewMission(0))
	crowley.missions = append(crowley.missions, genNewMission(0))
	crowley.missions = append(crowley.missions, genNewMission(0))
	
	villages[0] = crowley

	var bristal Village
	bristal.name = "Bristal"
	bristal.distanceToKeep = 2
	bristal.size = 3
	bristal.mapX, bristal.mapY = 17, 3
	bristal.missions = make([]Mission, 0, 0)
	bristal.villageIndex = 1
	villages[1] = bristal

	var faust Village
	faust.name = "Faust"
	faust.distanceToKeep = 2
	faust.size = 3
	faust.mapX, faust.mapY = 37, 2
	faust.missions = make([]Mission, 0, 0)
	faust.villageIndex = 2
	villages[2] = faust

	var gould Village
	gould.name = "Gould"
	gould.distanceToKeep = 2
	gould.size = 2
	gould.mapX, gould.mapY = 35, 15
	gould.missions = make([]Mission, 0, 0)
	gould.villageIndex = 3
	villages[3] = gould

	var elise Village
	elise.name = "Elise"
	elise.distanceToKeep = 3
	elise.size = 2
	elise.mapX, elise.mapY = 56, 15
	elise.missions = make([]Mission, 0, 0)
	elise.villageIndex = 4
	villages[4] = elise

	var autumn Village
	autumn.name = "Autumn"
	autumn.distanceToKeep = 4
	autumn.size = 4
	autumn.mapX, autumn.mapY = 57, 3
	autumn.villageIndex = 5	
	autumn.missions = make([]Mission, 0, 0)
	villages[5] = autumn

	var hollow Village
	hollow.name = "Hollow"
	hollow.distanceToKeep = 13
	hollow.size = 2
	hollow.mapX, hollow.mapY = 2, 13
	hollow.villageIndex = 6
	villages[6] = hollow

	var caustus Village
	caustus.name = "Caustus"
	caustus.distanceToKeep = 1
	caustus.size = 6
	caustus.mapX, caustus.mapY = 35, 11
	caustus.missions = make([]Mission, 0, 0)
	caustus.villageIndex = 7	
	villages[7] = caustus
}

func (village *Village) research() {

}

func (village *Village) viewBulletinBoard() {
	exitFlag := false
	rsp := ""
	
	blank := "│" + packSpaceString("", 72) + " │"
	
	for !exitFlag {
		clearConsole()
		fmt.Println("╔────────────────── " + packSpaceStringCenter(":: " + village.name + " Bulletin Board ::", 36) + " ──────────────────╗")
		fmt.Println(blank)		
		fmt.Println(blank)
		if len(village.missions) < 1 {
			fmt.Println("│ " + packSpaceString("No jobs available.", 72) + " │")
		} else {
			for k := 0; k < len(village.missions); k++ {
				job := fmt.Sprintf("│ %v. ", k)
				fmt.Println(job + packSpaceString(village.missions[k].getDisplayString(1), 69) + " │")
				fmt.Println(blank)			
			}
		}
		fmt.Println(blank)
		fmt.Println("╚──────────────────────────────────────────────────────────────────────────╝")
		fmt.Println("b. Back")
		fmt.Printf("Choose a job to view: ")
		fmt.Scanln(&rsp)
		
		if rsp == "b" {
			exitFlag = true
		} else {
			num,err := strconv.Atoi(rsp)
			
			if err == nil {
				if num < len(village.missions) {
					accepted := village.missions[num].viewAcceptDialog()					
					if accepted {
						exitFlag = true
						
						mission = village.missions[num]
						mission.status = STATUS_ACTIVE
					
						if len(village.missions) > 1 {
							village.missions = append(village.missions[:num], village.missions[num+1:]...)					
						} else {
							village.missions = make([]Mission, 0, 0)
						}
					}
				}
			}
		}
	}
}

func (village *Village) visitTavern() {		// 				
	// rest		 (end day, gain hp)						  
	// get gossip (apprentice tips, politickal gains, ) 
	// gamble  (randomly gain or lose money)			
	// drink  (raise soul by 1 - cost)  		   	 

	exitFlag := false
	rsp := ""
	
	for !exitFlag {
		clearConsole()
		
		fmt.Println("+++ Tavern +++")
		fmt.Println("------------")
		fmt.Println("1. Buy Drink  		(1 Crown)")
		fmt.Println("2. Rent Room   	(3 Crown)")
		fmt.Println("3. Carrouse  ")
		fmt.Println("4. View Bulletin Board")	
		fmt.Println("")	
		fmt.Println("x. Exit")
		fmt.Println("------------")
		fmt.Println("What do you wish to do?")		
		
		fmt.Scanln(&rsp)
		
		if rsp == "x" {
			exitFlag = true
		} else if rsp == "1" {
			showPause("You take in drink and companionship. Your spirits are lifted.")
			character.giveSoul(2)
			character.crowns -= 1
			endDay()
			save()			
		} else if rsp == "2" {	
			showPause("You rest and regather your strength for the long road ahead.")
			character.crowns -= 3
			endDay()
			save()
		} else if rsp == "3" {
			
		} else if rsp == "4" {
			village.viewBulletinBoard()
		}
	}

}

func (village *Village) visitChirurgeon() {
	// set wounds
	// sutures
	// increase healing time
	// costs

}

func (village *Village) politicks() {

}

func doBattle() (string){
	chooseAdventure()
	result := adventure()
	
	if result == DIED {
		showPause("Character died! Game over Man, game over!")
		return "q"
	}
	
	return ""
}

func doMissionPhase() {
	var die Die
	
	if mission.phases[mission.currentPhase - 1].id == PHASE_PUZZLE {
		val := 0
		for k := 0; k < character.skills[0]; k++ {
			val += die.rollxdx(1, 6)
		}
		mission.phases[mission.currentPhase - 1].puzzlePips -= val
		clearConsole()
		showPause(fmt.Sprintf("You solved %v pips of the puzzle!", val))
		if mission.phases[mission.currentPhase - 1].puzzlePips < 1 {
			mission.phases[mission.currentPhase - 1].puzzlePips = 0
			showPause("Puzzle solved!")
			mission.currentPhase++
			mission.viewMissionStatus()
		}
			
	} else if mission.phases[mission.currentPhase - 1].id == PHASE_RESEARCH {
		val := 0
		for k := 0; k < character.skills[2]; k++ {
			val += die.rollxdx(1, 6)
		}
		mission.phases[mission.currentPhase - 1].researchPips -= val
		clearConsole()
		showPause(fmt.Sprintf("You researched %v pips!", val))
		if mission.phases[mission.currentPhase - 1].researchPips < 1 {
			mission.phases[mission.currentPhase - 1].researchPips = 0
			showPause("Research Complete!")
			mission.currentPhase++
			mission.viewMissionStatus()
		}	
	} 
	
	endDay()
}

func (village *Village) visitVillage() string {
	canInvestigate := false
	canBattle := true
	clearConsole()

	fmt.Println("+++ Village of " + village.name + " +++")
	fmt.Println("------------")
	fmt.Println("1. Shop")
	fmt.Println("2. Research Quest")
	fmt.Println("3. Visit Tavern")
	fmt.Println("4. Visit Chirurgeon")
	fmt.Println("5. Politicks - Curry Favor / Influence")
	fmt.Println("6. Travel")
	if mission.typeId != -1 && mission.phases[mission.currentPhase - 1] != PHASE_FIGHT && mission.phases[mission.currentPhase - 1].locationIndex == village.villageIndex {
		quip := ""
		if mission.phases[mission.currentPhase - 1].id == PHASE_PUZZLE {
			quip = "Solve Puzzle"
		} else if mission.phases[mission.currentPhase - 1].id == PHASE_RESEARCH {
			quip = "Research"
		} 
		fmt.Println("r. [Mission: " + quip + "]")
		
		canInvestigate = true
	} 
	if mission.typeId != -1 && mission.currentPhase >= mission.minimumPhases && mission.missionBaseLocation == village.villageIndex {
		fmt.Println("f. [Mission: Battle Monster]")	
		canBattle = true
	} else {
		fmt.Println("")
	}
	
	fmt.Println("q. Quit")
	fmt.Println("")
	fmt.Println(BASE_ACTIONS)
	fmt.Println("")	

	fmt.Println("Select an Option:  ")

	rsp := ""
	rsp2 := ""
	fmt.Scanln(&rsp, &rsp2)

	if rsp == "1" {
		village.shopMenu()
	} else if rsp == "2" {
		village.research()
	} else if rsp == "3" {
		village.visitTavern()
	} else if rsp == "4" {
		village.visitChirurgeon()
	} else if rsp == "5" {
		village.politicks()
	} else if rsp == "r" && canInvestigate {
		doMissionPhase()	
	} else if rsp == "f" && canBattle {
		rsp = doBattle()			
	} else if rsp == "6" {
		showTravelMenu()
	} else if rsp == "s" {
		character.showStatus()
		character.printCharacter(1)
	} else if rsp == "m" {	
		mission.viewMissionStatus()
	} else if rsp == "i" {	
		character.showInventory()
	} else if rsp == "w" {	
		drawWorldMap()			
	} else if rsp == "h" {
		openTownMinutiae()
		
	} else if strings.Contains(rsp, "%give") && strings.Contains(rsp2, "money"){
		showPause("Money received!")
		character.crowns += 200
	} 

	return rsp
}

func showTravelMenu() string {

	showVillages()

	validSelection := false
	dist := 0
	destination := ""
	rsp := ""
	prevIndex := character.villageIndex

	for validSelection == false {
		clearConsole()

		fmt.Println("Travel Menu")
		fmt.Printf("Day: %v \n", game.gameDay)
		fmt.Println("------------")
		dist = getVillageDistance(0)
		fmt.Println(packSpaceString("1. Crowley", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(1)
		fmt.Println(packSpaceString("2. Bristal", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(2)
		fmt.Println(packSpaceString("3. Faust", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(3)
		fmt.Println(packSpaceString("4. Gould", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(4)
		fmt.Println(packSpaceString("5. Elise", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(5)
		fmt.Println(packSpaceString("6. Autumn", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(6)
		fmt.Println(packSpaceString("7. Hollow", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(7)
		fmt.Println(packSpaceString("8. Caustus", 20) + fmt.Sprintf("(%v days travel)", dist))
		fmt.Println("")
		dist = getVillageDistance(99)
		fmt.Println(packSpaceString("k. Return to Keep", 20) + fmt.Sprintf("(%v days travel)", dist))

		fmt.Println("")
		fmt.Println(BASE_ACTIONS)
		fmt.Println("")	
		
		fmt.Println("x. Back")
		fmt.Println("    ----    ")
		fmt.Printf("Where do you wish to travel? ")

		fmt.Scanln(&rsp)

		switch rsp {
		case "1":
			dist = getVillageDistance(0)
			character.villageIndex = 0
			destination = villages[0].name
			validSelection = true
		case "2":
			dist = getVillageDistance(1)
			character.villageIndex = 1
			destination = villages[1].name
			validSelection = true
		case "3":
			dist = getVillageDistance(2)
			character.villageIndex = 2
			destination = villages[2].name
			validSelection = true
		case "4":
			dist = getVillageDistance(3)
			character.villageIndex = 3
			destination = villages[3].name
			validSelection = true
		case "5":
			dist = getVillageDistance(4)
			character.villageIndex = 4
			destination = villages[4].name
			validSelection = true
		case "6":
			dist = getVillageDistance(5)
			character.villageIndex = 5
			destination = villages[5].name
			validSelection = true
		case "7":
			dist = getVillageDistance(6)
			character.villageIndex = 6
			destination = villages[6].name
			validSelection = true
		case "8":
			dist = getVillageDistance(7)
			character.villageIndex = 7
			destination = villages[7].name
			validSelection = true
		case "k":
			dist = getVillageDistance(99)
			character.villageIndex = 99
			destination = "Keep"
			validSelection = true
			
		case "w":
			drawWorldMap()
			validSelection = false
			
		case "s":
			character.showStatus()
			character.printCharacter(1)		
			
		case "h":
			showTravelMinutiae()
		
		case "i":
			character.showInventory()	
			
		case "x":
			validSelection = true
			dist = 0
		}
	}

	if dist > 0 {

		confirm := ""
		fmt.Printf("\nAre you sure you wish to travel to " + destination + "? ")
		fmt.Scanln(&confirm)

		if confirm == "y" {
			showTimePassageScreen(dist)
			fmt.Println("Arrived in "+destination+". After ", dist, " days of travel.")
			fmt.Println("Press any key to continue.")
			tgt := ""
			fmt.Scanln(&tgt)
		} else {
			character.villageIndex = prevIndex
		}

	}

	return rsp
}

func showTimePassageScreen(lapse int) {
	start := 0
	tick := "["
	
	for start < lapse {
		diff := lapse - start
		clearConsole()
		fmt.Println("Day: ", game.gameDay)
		fmt.Println("Time Passes...")
		endDay()
		tick += " █ █ █"
		block := ""
		for k := 0; k < diff-1; k++ {
			block += "      "
		}
		block += " ]"
		
		fmt.Println(tick + block)
		time.Sleep(900 * time.Millisecond)
		start++
	}
}

func (village *Village) goShop() {

	var result string = ""

	for result != "x" {
		result = village.visitVillage()

		if result == "1" {
			village.buyWeaponScreen()
		} else if result == "2" {
			village.buyArmorScreen()
		} else if result == "3" {
			village.buyProvisions()
		} else if result == "4" {
			village.buyCuriosities()
		}
	}

}
