// village.go
// Village

package main

import "fmt"
import "time"
import "strings"
import "strconv"

var orphanage []Character

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

func (v *Village) getSaveString() (string) {
	villageBlock := BLOCK_VILLAGE + ","

	villageBlock += fmt.Sprintf("%v,", v.villageIndex)
	villageBlock += fmt.Sprintf("%v,", v.size)
	villageBlock += fmt.Sprintf("%v,", v.politicalFavor)

	villageBlock += fmt.Sprintf("%v,", len(v.shopWeapons))
	villageBlock += fmt.Sprintf("%v,", len(v.shopArmor))
	villageBlock += fmt.Sprintf("%v,", len(v.shopProvisions))
	villageBlock += fmt.Sprintf("%v,", len(v.shopApothecary))
	villageBlock += fmt.Sprintf("%v,", len(v.shopCuriosities))
	
	villageBlock += "◄" // line end, now do equip
	
	for k := 0; k < len(v.shopWeapons); k++ {
		villageBlock += v.shopWeapons[k].getSaveString()
	}
	
	for k := 0; k < len(v.shopArmor); k++ {
		villageBlock += v.shopArmor[k].getSaveString()
	}

	for k := 0; k < len(v.shopProvisions); k++ {
		villageBlock += v.shopProvisions[k].getSaveString()
	}

	for k := 0; k < len(v.shopApothecary); k++ {
		villageBlock += v.shopApothecary[k].getSaveString()
	}

	for k := 0; k < len(v.shopCuriosities); k++ {
		villageBlock += v.shopCuriosities[k].getSaveString()
	}
	
	
	// finally:
	villageBlock += "■"

	return villageBlock
}

func unpackVillageBlock(idx int, block string) (int, Village) {
	var vill Village
	
	lines := strings.Split(block, "◄")

	bits := strings.Split(lines[0], ",")

	if bits[0] == BLOCK_VILLAGE {
		fmt.Println("Loading " + BLOCK_VILLAGE + "...")
	} else {
		log.addError("Cant find VILLAGE block.")
		fmt.Println("VILLAGE Block not found!")
		return -1, vill
	}

	lineCounter := 0

	if (villages[idx].villageIndex != idx) {
		log.addError("Village indexes mismatch! Corrupt Save file?")
		fmt.Println("VILLAGE index mismatch! Cannot load save file.")
		return -1, vill	
	}
	
	vill = villages[idx]
	
	vill.size, _ = strconv.Atoi(bits[2])
	vill.politicalFavor, _ = strconv.Atoi(bits[3])
	
	numWeapons, numArmor, numProv, numApoth, numCur := 0,0,0,0,0
	
	numWeapons, _ = strconv.Atoi(bits[4])
	numArmor, _ = strconv.Atoi(bits[5])
	numProv, _ = strconv.Atoi(bits[6])
	numApoth, _ = strconv.Atoi(bits[7])
	numCur, _ = strconv.Atoi(bits[8])
	
	lineCounter = 1
	
	for k := 0; k < numWeapons; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		vill.shopWeapons = append(vill.shopWeapons, itm)
		lineCounter++
	}
	
	for k := 0; k < numArmor; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		vill.shopArmor = append(vill.shopArmor, itm)
		lineCounter++
	}
	
	for k := 0; k < numProv; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		vill.shopProvisions = append(vill.shopProvisions, itm)
		lineCounter++
	}
	
	for k := 0; k < numApoth; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		vill.shopApothecary = append(vill.shopApothecary, itm)
		lineCounter++
	}	
	
	for k := 0; k < numCur; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		vill.shopCuriosities = append(vill.shopCuriosities, itm)
		lineCounter++
	}	
	
	return 1, vill
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
	gould.mapX, gould.mapY = 32, 18
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

	var autumn Village					// Autumn has the orphanage
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

func buildOrphanage() {
	var die Die
	
	orphanage = make([]Character, 0, 0)
	
	for k := 0; k < die.rollxdx(1, 3) + 1; k++ {
		appr := getRandomApprentice()
		orphanage = append(orphanage, appr)
	}
}

func (village * Village) visitOrphanage() {
	exitFlag := false
	rsp := ""
	counter := 0
	
	for !exitFlag {
		clearConsole()
		
		fmt.Println("+++ Orphanage +++")
		fmt.Println(makeDialogString("The following children are all left handed..."))		
		fmt.Println(makeDialogString("...a strange request to be sure but so long as you have the"))			
		fmt.Println(makeDialogString("adoption fee, which is, of course, the usual 25 crowns..."))			
		fmt.Println("------------")
		if len(orphanage) < 1 {
			fmt.Println("No left handed children available.")
		} else {
			counter = 0
			for k := 0; k < len(orphanage); k++ {
				counter++
				row := fmt.Sprintf("%v. %s", counter, orphanage[k].name)
				fmt.Println(row)
			}		
		}

		fmt.Println("")		
		fmt.Println("x. Exit  ")
		
		fmt.Scanln(&rsp)
		
		if rsp == "x" {
			exitFlag = true
		} else {
			num, err := strconv.Atoi(rsp)
		
			if err == nil {
				if num-1 < len(orphanage) {
					if character.crowns >= 25 {
						orphanage[num-1].printCharacter(0)
						
						adopt := getResponse("\n\nDo you wish to adopt?")

						if adopt == "y" || adopt == "Y" {
							character.crowns -= 25
							indx := num-1
							showPause(fmt.Sprintf("Congrats! %s has been sent to your keep to get settled in.", orphanage[indx].name))
							keep.addNewApprenticeToKeep(orphanage[indx])
							orphanage = append(orphanage[:indx], orphanage[indx+1:]...)
						}
					} else {
						showPause("Not enough crowns to adopt!")		
					}
				} else {
					showPause("Invalid selection.")
				}			
			}
		}
	}
	
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
			endDay(2, true)
			save()			
		} else if rsp == "2" {	
			showPause("You rest and regather your strength for the long road ahead.")
			character.crowns -= 3
			endDay(3, true)
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

func doBattle(random bool) (string){
	result := chooseAdventure(random)
	
	if result == DIED {
		showPause("Character died! Game over Man, game over!")
		return "q"
	} else if random == false && result == FINISHED_MISSION {
		showMissionComplete()
	} else if random && result == FINISHED_MISSION {
		showPause("You defeat the bandits and continue your journey...")
	}
	
	return ""
}

func drawPipMap(dice []int) {

	const PUZZLE_TOKEN = "[Φ]"
	const RESEARCH_TOKEN = "[Σ]"

	if mission.phases[mission.currentPhase - 1].id == PHASE_PUZZLE {
		diceResults := ""
		total := 0
		
		for i := 0; i < len(dice); i++ {
			diceResults += fmt.Sprintf(" [%v] ", dice[i])
			total += dice[i]
		}	
	
		sTotal := fmt.Sprintf("%v", total)
	
	//[" + packSpaceStringCenter(sTotal, 4) + "]
	
		fmt.Println("┌─────────╡ " + packSpaceStringCenter("Puzzle Required", 19) + " ╞─────────┐   ┌───────── Dice ────────┐ ")

		fmt.Println("  " + packSpaceStringCenter(" ", 39) + "     " + diceResults)
		
		pips := "  " 
		counter := 0
		rowcount := 0
		
		for k := 0; k < mission.phases[mission.currentPhase - 1].puzzlePips; k++ {
			counter++
			pips += PUZZLE_TOKEN
			
			if counter > 9 {
				counter = 0
				rowcount++
				if rowcount == 1 {
					pips +=  "     └────────╜" + packSpaceStringCenter(sTotal, 5) + "╙────────┘"			
				}
				
				pips += "\n  "
			// } else if counter == mission.phases[mission.currentPhase - 1].puzzlePips {
				// pips = packSpaceString(pips, 25) + "     └────────╜" + packSpaceStringCenter(sTotal, 5) + "╙────────┘"	
			} else {
				pips += " "
			}
		}
		
		fmt.Println(pips)
	} else if mission.phases[mission.currentPhase - 1].id == PHASE_RESEARCH {
	
		diceResults := ""
		total := 0
		
		for i := 0; i < len(dice); i++ {
			diceResults += fmt.Sprintf(" [%v] ", dice[i])
			total += dice[i]
		}	
		sTotal := fmt.Sprintf("%v", total)
	
		fmt.Println("┌─────────╡ " + packSpaceStringCenter("Research Required", 19) + " ╞─────────┐   ┌───────── Dice ────────┐ ")
		fmt.Println("  " + packSpaceStringCenter(" ", 39) + "     " + diceResults)

		pips := "  " 
		counter := 0
		rowcount := 0
		
		for k := 0; k < mission.phases[mission.currentPhase - 1].researchPips; k++ {
			counter++
			pips += RESEARCH_TOKEN
			
			if counter > 9 {
				counter = 0
				rowcount++
				if rowcount == 1 {
					pips +=  "     └────────╜" + packSpaceStringCenter(sTotal, 5) + "╙────────┘"				
				}
				pips += "\n  "
			} else {
				pips += " "
			}
		}
		
		fmt.Println(pips)
	}
	fmt.Println("")
	fmt.Println("└─────────────────────────────────────────┘")
	fmt.Println("")
}

func doMissionPhase() {
	var die Die
			
	exitFlag := false
	rsp := ""

	dice := make([]int, 0, 0)

	rerolls := 1
	verb := ""
	
	if mission.phases[mission.currentPhase - 1].id == PHASE_PUZZLE {
		for d := 0; d < character.skills[0]; d++{
			dice = append(dice, 0)
		}	
		verb = "Solved!"
	} else if mission.phases[mission.currentPhase - 1].id == PHASE_RESEARCH {
		for d := 0; d < character.skills[2]; d++{
			dice = append(dice, 0)
		}		
		verb = "Researched"
	}

	total := 0	
	for !exitFlag {
		clearConsole()
		drawPipMap(dice)
		
		commands := ""
		if rerolls > 0 {
			commands = "[r. roll/reroll    e. exit]"
		} else {
			commands = "[a. accept]"
		}
		
		//fmt.Println(packSpaceString(diceResults, 40) + fmt.Sprintf("(%v)", total))
		fmt.Println("")
		fmt.Println(commands)
		fmt.Printf("Choose an action: ")
		
		fmt.Scanln(&rsp)
		
		if rsp == "r" && rerolls > 0 {
			total = 0
			for i := 0; i < len(dice); i++ {
				dice[i] = die.rollxdx(1, 6)
				total += dice[i]
			}		
			
			rerolls -= 1
		} else if rsp == "e" {
			exitFlag = true
		} else if rsp == "a" {
			exitFlag = true
			
			if total > 0 {
				endFlag := false
				for k := 0; k < total; k++ {
					if mission.phases[mission.currentPhase - 1].id == PHASE_PUZZLE {
						mission.phases[mission.currentPhase - 1].puzzlePips -= 1
						if mission.phases[mission.currentPhase - 1].puzzlePips < 1 {
							endFlag = true
							mission.phases[mission.currentPhase - 1].complete = 1
						}
					} else if  mission.phases[mission.currentPhase - 1].id == PHASE_RESEARCH {
						mission.phases[mission.currentPhase - 1].researchPips -= 1					
						if mission.phases[mission.currentPhase - 1].researchPips < 1 {
							endFlag = true
							mission.phases[mission.currentPhase - 1].complete = 1
						}
					}
					clearConsole()
					drawPipMap(dice)
					time.Sleep(400 * time.Millisecond)
					if endFlag {
						break
					}
				}
			}
			
			if mission.phases[mission.currentPhase - 1].complete == 0 {
				showPause(fmt.Sprintf("%v pips %s!", total, verb))
			} else {
				showPause(fmt.Sprintf("Phase Complete! (%v experience earned)", mission.phases[mission.currentPhase - 1].rewardExperience))			
				character.exp += mission.phases[mission.currentPhase - 1].rewardExperience
				mission.currentPhase++
			}
		}
	}
	
	endDay(1, true)
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
	
	if village.villageIndex == 5 {
		fmt.Println("7. Visit Orphanage")		
	}
	
	if mission.typeId != -1 && mission.phases[mission.currentPhase - 1].id != PHASE_FIGHT && mission.phases[mission.currentPhase - 1].locationIndex == village.villageIndex {
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
		rsp = doBattle(false)			
	} else if rsp == "6" {
		rsp = showTravelMenu()
	} else if rsp == "7" && village.villageIndex == 5 {
		village.visitOrphanage()		
	} else if rsp == "s" {
		character.printCharacter(1)
		character.showStatus()
		if apprentice.instanceId > 0 {
			apprentice.printCharacter(1)
			apprentice.showStatus()
		}
	} else if rsp == "m" {	
		mission.viewMissionStatus()
	} else if rsp == "i" {	
		character.showInventory()
		if apprentice.instanceId > 0 {
			apprentice.showInventory()
		}
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
			
		case "m":
			mission.viewMissionStatus()
			
		case "s":
			character.printCharacter(1)	
			character.showStatus()	
			if apprentice.instanceId > 0 {
				apprentice.printCharacter(1)
				apprentice.showStatus()
			}
			
		case "h":
			showTravelMinutiae()
		
		case "i":
			character.showInventory()	
			if apprentice.instanceId > 0 {
				apprentice.showInventory()
			}
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
		
			// check to see if we are attacked by bandits on the road...
			var die Die
			attacked := false
			
			for j := 0; j < dist; j++ {
				if die.rollxdx(1, 100) > 93 {
					// attacked by bandits
					showPause("You have been attacked by bandits during your travel!")
					attacked = true
					break;
				}
			}
			
			if attacked {
				atk := doBattle(true)
				
				if atk == "q" {
					// character died, exit
					return "q"
				}
			}
		
			showTimePassageScreen(dist)
			fmt.Println("Arrived in " + destination + ". After ", dist, " days of travel.")
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
		endDay(0, false)
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
