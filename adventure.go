// aventure.go
package main

import "fmt"
import "strings"

const FLED_BATTLE = 0
const FINISHED_MISSION = 1
const DIED = -1

var mission Mission

func (bg *BattleGrid) canCharacterCast(char Character, currTurns int) bool {
	return true
}

func (bg *BattleGrid) doPlayerAttack(turn int, hand int) {
	var die Die

	if turn == CHAR_TURN {
		adj := character.getTotalAttackAdjustment(hand)
		atkRoll := die.rollxdx(1, 20)
		atkTotal := adj + atkRoll

		def := bg.monster.getTotalDefenseAdjustment()
		defRoll := die.rollxdx(1, 20)
		defTotal := def + defRoll

		fmt.Println(fmt.Sprintf("Character rolls %v + %v = [%v]", atkRoll, adj, atkTotal))
		fmt.Println(fmt.Sprintf("Monster rolls %v + %v = [%v]", defRoll, def, defTotal))

		if atkTotal > defTotal {
			fmt.Println(character.name + " hits!")
		} else {
			showPause(character.name + " misses!")
			return
		}

		character.handSlots[hand].durability -= 1
		
		diff := atkTotal - defTotal
		tBonus := 0
		for ; diff >= 5; diff -= 5 {
			tBonus++
		}

		fmt.Println(fmt.Sprintf("Bonus is %v", getSigned(tBonus)))
		targetRoll := die.rollxdx(1, 10)
		totalTarget := targetRoll + tBonus
		if totalTarget > 10 {
			totalTarget = 10
		}

		fmt.Println(fmt.Sprintf("Target is %v + %v = %v", targetRoll, tBonus, totalTarget))
		crits := ""
		hits := 1
		if totalTarget == 10 {
			hits++
			for totalTarget == 10 {
				crits += "[Crit!]"
				targetRoll = die.rollxdx(1, 10)
				totalTarget = targetRoll + tBonus
				if totalTarget >= 10 {
					totalTarget = 10
					hits++
				}
			}
			fmt.Println(crits)
		}
		fmt.Println("Hit on " + bg.monster.body[totalTarget-1])

		penetrationBonus := 0
		diff = atkTotal - defTotal
		for ; diff >= 2; diff -= 2 {
			penetrationBonus++
		}
		fmt.Println(fmt.Sprintf("Penetration bonus is %v", penetrationBonus))
		penetrationRoll := die.rollxdx(1, 20)
		totalPenetration := penetrationBonus + penetrationRoll

		fmt.Println(fmt.Sprintf("Penetration Roll: %v + %v = [%v]", penetrationRoll, penetrationBonus, totalPenetration))
		fmt.Println(fmt.Sprintf("Resistance is %v", bg.monster.resistance[targetRoll-1]))

		if totalPenetration > bg.monster.resistance[targetRoll-1] {
			showPause(fmt.Sprintf("Attack penetrates! Monster takes %v hits!", hits))
			bg.monster.hp -= hits
		} else {
			showPause("Monster soaks the attack.")
		}
		
		if character.handSlots[hand].isBroken() {
			showPause(character.handSlots[hand].name + " has broken!")
		}
		
	} else {
		adj := apprentice.getTotalAttackAdjustment(hand)
		atkRoll := die.rollxdx(1, 20)
		atkTotal := adj + atkRoll

		def := bg.monster.getTotalDefenseAdjustment()
		defRoll := die.rollxdx(1, 20)
		defTotal := def + defRoll

		fmt.Println(fmt.Sprintf("%s rolls %v + %v = [%v]", apprentice.name, atkRoll, adj, atkTotal))
		fmt.Println(fmt.Sprintf("Monster rolls %v + %v = [%v]", defRoll, def, defTotal))

		if atkTotal > defTotal {
			fmt.Println(apprentice.name + " hits!")
		} else {
			showPause(apprentice.name + " misses!")
			return
		}

		apprentice.handSlots[hand].durability -= 1
		
		diff := atkTotal - defTotal
		tBonus := 0
		for ; diff >= 5; diff -= 5 {
			tBonus++
		}

		fmt.Println(fmt.Sprintf("Bonus is %v", getSigned(tBonus)))
		targetRoll := die.rollxdx(1, 10)
		totalTarget := targetRoll + tBonus
		if totalTarget > 10 {
			totalTarget = 10
		}

		fmt.Println(fmt.Sprintf("Target is %v + %v = %v", targetRoll, tBonus, totalTarget))
		crits := ""
		hits := 1
		if totalTarget == 10 {
			hits++
			for totalTarget == 10 {
				crits += "[Crit!]"
				targetRoll = die.rollxdx(1, 10)
				totalTarget = targetRoll + tBonus
				if totalTarget >= 10 {
					totalTarget = 10
					hits++
				}
			}
			fmt.Println(crits)
		}
		fmt.Println("Hit on " + bg.monster.body[totalTarget-1])

		penetrationBonus := 0
		diff = atkTotal - defTotal
		for ; diff >= 2; diff -= 2 {
			penetrationBonus++
		}
		fmt.Println(fmt.Sprintf("Penetration bonus is %v", penetrationBonus))
		penetrationRoll := die.rollxdx(1, 20)
		totalPenetration := penetrationBonus + penetrationRoll

		fmt.Println(fmt.Sprintf("Penetration Roll: %v + %v = [%v]", penetrationRoll, penetrationBonus, totalPenetration))
		fmt.Println(fmt.Sprintf("Resistance is %v", bg.monster.resistance[targetRoll-1]))

		if totalPenetration > bg.monster.resistance[targetRoll-1] {
			showPause(fmt.Sprintf("Attack penetrates! Monster takes %v hits!", hits))
			bg.monster.hp -= hits
		} else {
			showPause("Monster soaks the attack.")
		}
		
		if apprentice.handSlots[hand].isBroken() {
			showPause(apprentice.handSlots[hand].name + " has broken!")
		}
	}
}

// return - hand value
func (bg *BattleGrid) canCharacterAttack(char Character, currTurns int) int {
	if bg.isMonsterVisible() {
		if bg.isMonsterInAttackRange(bg.turn) && bg.isAttackPathClear(bg.turn) {
			if char.handSlots[LEFT].typeCode == ITEM_TYPE_WEAPON && char.handSlots[LEFT].atkTurns <= currTurns && !char.handSlots[LEFT].isBroken(){
				return LEFT
			} else if char.handSlots[RIGHT].typeCode == ITEM_TYPE_WEAPON && char.handSlots[RIGHT].atkTurns <= currTurns && !char.handSlots[RIGHT].isBroken(){
				return RIGHT
			} else {
				fmt.Println("no attackable weapons " + string(currTurns))
			}
		} else {
			fmt.Println("monster not in range")
		}
	} else {
		fmt.Println("monster not vis")
	}

	return -1
}

func (bg *BattleGrid) showConfimExit() (bool) {
	clearConsole()

	if character.hp < 1 {
		fmt.Println("You have been defeated by the monster. Game Over.")
		fmt.Println()
		showPause("  Press any key to exit.")	
		return true
		
	} else if bg.monster.isAlive() {
		fmt.Println("Exiting without defeating the monster will cause you to fail this mission.")
		fmt.Println()
		fmt.Println("  Are you sure you wish to exit?")

	} else if !bg.monster.isAlive() {
		fmt.Println("Monster Defeated!")
		fmt.Println("If you exit any uncollected loot will be lost.")
		fmt.Println()
		fmt.Println("  Are you sure you wish to exit?")	
	}
	
	rsp := ""
	fmt.Scanln(&rsp)
	
	if rsp == "y" || rsp == "Y" || rsp == "yes"{
		return true
	} 

	return false
}

func (bg *BattleGrid) showFoundLoot(idx int) string {
	loop := true
	rsp := ""
	var loot Loot

	log.addInfo(fmt.Sprintf("Showing loot for idx %v", idx))
	
	for loop {
		clearConsole()

		loot = bg.getEntityGrid(bg.currGrid).loot[idx]

		fmt.Println("╔═════════ Found ══════════╗")
		fmt.Println("║" + packSpaceString(loot.container, 26) + "║")
		fmt.Println("║                          ║")
		fmt.Println("║  Crowns: " + packSpace(loot.crowns, 4) + "            ║")
		fmt.Println("║                          ║")

		if loot.empty {
			fmt.Println("║ " + packSpaceString("Empty", 26) + "║")
		} else {
			for k := 0; k < len(loot.items); k++ {
				fmt.Println("║ " + packSpace(k, 1) + ". " + packSpaceString(loot.items[k].name, 22) + "║")
				fmt.Println("║                          ║")
			}
		}

		fmt.Println("║                          ║")
		fmt.Println("╚══════════════════════════╝")
		fmt.Println("")
		fmt.Println("ta : Take All")
		fmt.Println("tc : Take Crowns")
		fmt.Println("t# : Take #")
		fmt.Println("l : Leave")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		fmt.Scanln(&rsp)

		if rsp == "l" || rsp == "L" {
			loop = false
		} else if rsp == "tc" {
			character.crowns += loot.crowns
			bg.getEntityGrid(bg.currGrid).loot[idx].crowns = 0

			if len(bg.getEntityGrid(bg.currGrid).loot[idx].items) < 1 {
				loot.empty = true
				loop = false
			}
			
		} else if rsp == "ta" {
			loop = false
			if !loot.empty {
				character.crowns += loot.crowns
				bg.getEntityGrid(bg.currGrid).loot[idx].crowns = 0
				loot.empty = true
				bg.getEntityGrid(bg.currGrid).loot[idx].empty = true

				if bg.turn == CHAR_TURN {				
					for k := 0; len(bg.getEntityGrid(bg.currGrid).loot[idx].items) > 0; {
						addOK := character.giveCharacterItem(loot.items[k])
						if addOK {
							if len(bg.getEntityGrid(bg.currGrid).loot[idx].items) > 1 {
								bg.getEntityGrid(bg.currGrid).loot[idx].items = append(bg.getEntityGrid(bg.currGrid).loot[idx].items[:k], bg.getEntityGrid(bg.currGrid).loot[idx].items[k+1:]...)							
							} else {
								bg.getEntityGrid(bg.currGrid).loot[idx].items = make([]Item, 0, 0)
							}
						}
					}
				} else {
					for k := 0; len(bg.getEntityGrid(bg.currGrid).loot[idx].items) > 0; {
						addOK := apprentice.giveCharacterItem(loot.items[k])
						if addOK {
							if len(bg.getEntityGrid(bg.currGrid).loot[idx].items) > 1 {
								bg.getEntityGrid(bg.currGrid).loot[idx].items = append(bg.getEntityGrid(bg.currGrid).loot[idx].items[:k], bg.getEntityGrid(bg.currGrid).loot[idx].items[k+1:]...)							
							} else {
								bg.getEntityGrid(bg.currGrid).loot[idx].items = make([]Item, 0, 0)
							}						
						}
					}
				}
			}
		} else if rsp == "t0" && len(loot.items) > 0 {
			k := 0

			if bg.turn == CHAR_TURN {
				addOK := character.giveCharacterItem(loot.items[k])
				if addOK {
					bg.getEntityGrid(bg.currGrid).loot[idx].items = append(bg.getEntityGrid(bg.currGrid).loot[idx].items[:k], bg.getEntityGrid(bg.currGrid).loot[idx].items[k+1:]...)
				} else {
					showPause("Character cannot take this item because they will be over encumbered.")
				}
			} else {
				addOK := apprentice.giveCharacterItem(loot.items[k])
				if addOK {
					bg.getEntityGrid(bg.currGrid).loot[idx].items = append(bg.getEntityGrid(bg.currGrid).loot[idx].items[:k], bg.getEntityGrid(bg.currGrid).loot[idx].items[k+1:]...)
				} else {
					showPause("Character cannot take this item because they will be over encumbered.")
				}
			}

		} else if rsp == "t1" && len(loot.items) > 1 {
			k := 1
			if bg.turn == CHAR_TURN {
				addOK := character.giveCharacterItem(loot.items[k])
				if addOK {
					bg.getEntityGrid(bg.currGrid).loot[idx].items = append(bg.getEntityGrid(bg.currGrid).loot[idx].items[:k], bg.getEntityGrid(bg.currGrid).loot[idx].items[k+1:]...)
				} else {
					showPause("Character cannot take this item because they will be over encumbered.")
				}
			} else {
				addOK := apprentice.giveCharacterItem(loot.items[k])
				if addOK {
					bg.getEntityGrid(bg.currGrid).loot[idx].items = append(bg.getEntityGrid(bg.currGrid).loot[idx].items[:k], bg.getEntityGrid(bg.currGrid).loot[idx].items[k+1:]...)
				} else {
					showPause("Character cannot take this item because they will be over encumbered.")
				}
			}
		}
	}

	if len(bg.getEntityGrid(bg.currGrid).loot[idx].items) < 1 && bg.getEntityGrid(bg.currGrid).loot[idx].crowns < 1 {
		bg.getEntityGrid(bg.currGrid).loot[idx].empty = true
	}

	return rsp
}

func (bg *BattleGrid) searchLocation(turn int) {

	if turn == CHAR_TURN {
		if bg.allGrids[bg.charGridId].isLootAtLoc(bg.charXLoc, bg.charYLoc) {
			indx := bg.allGrids[bg.charGridId].getLootAtLoc(bg.charXLoc, bg.charYLoc)

			if indx > -1 {
				bg.showFoundLoot(indx)
			} else {
				showPause("Nothing found at this location.")
			}

		} else {
			showPause("Nothing found at this location.")
		}

	} else if turn == APP_TURN {
		if bg.allGrids[bg.appGridId].isLootAtLoc(bg.appXLoc, bg.appYLoc) {
			indx := bg.allGrids[bg.appGridId].getLootAtLoc(bg.appXLoc, bg.appYLoc)

			if indx > -1 {
				bg.showFoundLoot(indx)
			} else {
				showPause("Nothing found at this location.")
			}
		} else {
			showPause("Nothing found at this location.")
		}
	}
}

func (bg *BattleGrid) getAvailableActions(char Character, currTurns int) string {

	actions := ""

	if currTurns < 1 {
		actions = "(Inventory) (Status) (End Turn) (Help) (Exit) \n\nAction: "
	} else {
		actions = "(Move ++)"

		if bg.canCharacterAttack(char, currTurns) > -1 {
			actions += " (Attack)"
		}

		if bg.canCharacterCast(char, currTurns) {
			actions += " (Cast)"
		}

		if currTurns > 0 {
			actions += " (Defend) (Search) (Wait)"
			
			npcs := bg.getAdjacentNPCs()
			
			if len(npcs) > 0 {
				actions += " (Talk)"
			}
			
		}

		actions += "\n(Inventory) (Status) (End Turn) (Help) (Exit) \n\nAction: "
	}

	return actions
}


func chooseAdventure(random bool) (int) {
	rslt := 0
	
	if !random {
		showPause("Mission title is: " + mission.title)	
		rslt = adventure(mission.typeId)
	} else {
		showPause("Mission title is: " + "Accosted on the roadside.")	
		rslt = adventure(-1)		
	}
	
	return rslt
}

func adventure(mid int) (result int) {
	var bg = buildBattleGrid(mid)
	rsp := ""
	rsp2 := ""
	rsp3 := ""

	currTurns := 0
	maxTurns := 0
	// mission start setup
	currTurns = character.getCharacterMoves()
	maxTurns = character.getCharacterMoves()

	result = FLED_BATTLE
	
	playFlag := false
	
	for !playFlag {
		bg.drawGrid()
		fmt.Printf("Turns: %v / %v  (%v : %v : %v) \n", currTurns, maxTurns, bg.charXLoc, bg.charYLoc, bg.charGridId)
		descrip := ""
		if bg.turn == CHAR_TURN {
			descrip = bg.getAvailableActions(character, currTurns)
		} else if bg.turn == APP_TURN {
			descrip = bg.getAvailableActions(apprentice, currTurns)
		}

		fmt.Printf(descrip)
		fmt.Scanln(&rsp, &rsp2)

		if result == DIED && rsp != "exit" {
			showPause("You and your apprentice have died.")
			continue;
		}
		
		if strings.Contains(rsp, "move") && len(rsp2) > 0 {
			if currTurns < 1 {
				fmt.Println("No turns remain. End your turn.")
				fmt.Scanln(&rsp3)
			} else {
				direct := convertCardinalStringToInt(rsp2)

				if bg.turn == CHAR_TURN {
					if bg.directionValid(bg.charXLoc, bg.charYLoc, direct, bg.charGridId) {
						bg.moveCharacter(direct)
						currTurns -= 1
						bg.monster.plan.interrupt = ACTOR_SPOTTED
						bg.monster.plan.charMoved = true

						if bg.isGate(CHAR_TURN) {
							if selectedGate.gridid1 == bg.charGridId {
								bg.charGridId = selectedGate.gridid2
								bg.charXLoc = selectedGate.g2x
								bg.charYLoc = selectedGate.g2y
							} else {
								bg.charGridId = selectedGate.gridid1
								bg.charXLoc = selectedGate.g1x
								bg.charYLoc = selectedGate.g1y
							}
						}

					} else {
						fmt.Println("Path is blocked in this direction!")
						fmt.Scanln(&rsp3)
					}
				} else if bg.turn == APP_TURN {
					if bg.directionValid(bg.appXLoc, bg.appYLoc, direct, bg.appGridId) {
						bg.moveCharacter(direct)
						currTurns -= 1
						bg.monster.plan.interrupt = 1
						bg.monster.plan.appMoved = true
						
						if bg.isGate(APP_TURN) {
							if selectedGate.gridid1 == bg.appGridId {
								bg.appGridId = selectedGate.gridid2
								bg.appXLoc = selectedGate.g2x
								bg.appYLoc = selectedGate.g2y
							} else {
								bg.appGridId = selectedGate.gridid1
								bg.appXLoc = selectedGate.g1x
								bg.appYLoc = selectedGate.g1y
							}
						}
					} else {
						fmt.Println("Path is blocked in this direction!")
						fmt.Scanln(&rsp3)
					}
				}

				bg.updateActorVisibility()
			}
		} else if strings.Contains(rsp, "wait") {
			showPause("You cower in the darkness...")
			currTurns -= 1
			
		} else if strings.Contains(rsp, "talk") {
			showPause("TODO: Talk to NPCs if they exist, otherwise talk to yourself. Shut up.")
			currTurns -= 1
			
		} else if strings.Contains(rsp, "defend") {
			if currTurns < 1 {
				fmt.Println("No turns remain. You need at least one turn available to defend. End your turn.")
				fmt.Scanln(&rsp3)
			} else {
				if bg.turn == CHAR_TURN {			
					showPause("Turn Defense: " + getSigned(currTurns))	
					character.turnDefense = currTurns
					currTurns = 0
				} else if bg.turn == APP_TURN {
					showPause("Turn Defense: " + getSigned(currTurns))	
					apprentice.turnDefense = currTurns	
					currTurns = 0
				}			
			}			
		} else if strings.Contains(rsp, "end") && strings.Contains(rsp2, "turn") {
			// TODO:  THIS SHOULD BE 0!
			if currTurns > 99 {
				fmt.Println("You still have turns remaining... the darkness patiently waits.")
				fmt.Scanln(&rsp3)
			} else {
				if bg.turn == CHAR_TURN {
					if bg.hasApprentice && apprentice.isMotile() {
						bg.turn = APP_TURN
						bg.currGrid = bg.appGridId
						currTurns = apprentice.getCharacterMoves()
						maxTurns = apprentice.getCharacterMoves()
					} else {
						bg.turn = MONST_TURN
					}
				} else if bg.turn == APP_TURN {
					bg.turn = MONST_TURN
				}

				if bg.turn == MONST_TURN {
					bg.monster.turnDefense = 0
					bg.currGrid = bg.monsterGridId
					rslt := bg.doMonsterActivity()
					bg.turnCounter++
					character.turnDefense = 0
					apprentice.turnDefense = 0
					
					if rslt == 0 {
						currTurns = character.getCharacterMoves()
						maxTurns = character.getCharacterMoves()
						bg.turn = CHAR_TURN
						bg.currGrid = bg.charGridId
						bg.monster.plan.interrupt = 0 // clear any interrupts
					} else if rslt == CHARACTER_KILLED {
						if bg.hasApprentice && apprentice.isAlive() {
							showPause("You have died! Your apprentice assumes your banner!")
							var loot Loot
							loot.container = character.name + " (Dead)"
							loot.crowns = character.crowns
							loot.locX = bg.charXLoc
							loot.locY = bg.charYLoc
							loot.items = character.getListOfPossessions()
							bg.allGrids[bg.charGridId].loot = append(bg.allGrids[bg.charGridId].loot, loot)
							
							character = apprentice
							bg.hasApprentice = false
							var blankApprentice Character
							apprentice = blankApprentice
						} else {
							rsp = "exit"
							result = DIED						
						}
					} else if rslt ==  APPRENTICE_KILLED {
						// apprentice has died!
						var loot Loot
						loot.container = apprentice.name + " (Dead)"
						loot.crowns = 0
						loot.locX = bg.appXLoc
						loot.locY = bg.appYLoc
						loot.items = apprentice.getListOfPossessions()
						bg.allGrids[bg.appGridId].loot = append(bg.allGrids[bg.appGridId].loot, loot)
						
						bg.hasApprentice = false
						var blankApprentice Character
						apprentice = blankApprentice
						
						// TODO: show a death notice
					}
				}
			}
		} else if strings.Contains(rsp, "status") {
			if bg.turn == CHAR_TURN {
				character.showStatus()
				character.printCharacter(1)
			} else {
				apprentice.showStatus()
				apprentice.printCharacter(1)
			}
		} else if strings.Contains(rsp, "view") {
			if strings.Contains(rsp2, "-log") {
				log.displayLog()
			} else if strings.Contains(rsp2, "-pattern") {
				bg.drawGridPattern(bg.gridPattern)
			} else if strings.Contains(rsp2, "-balance") {
				showPause(fmt.Sprintf("Power Balance: %s",  getSignedFloat32(bg.calcPowerBalance())))
			}
		} else if strings.Contains(rsp, "inventory") {
			if bg.turn == CHAR_TURN {
				character.showInventory()
			} else {
				apprentice.showInventory()
			}
		} else if strings.Contains(rsp, "attack") {
			if bg.turn == CHAR_TURN {
				hand := bg.canCharacterAttack(character, currTurns)
				if hand > -1 {
					showPause("Character Attacks with " + character.handSlots[hand].name + "!")
					currTurns -= character.handSlots[hand].atkTurns
					bg.doPlayerAttack(CHAR_TURN, hand)
				} else {
					showPause("Unable to attack: not enough turns or no usable weapon equipped!")					
				}
			} else {
				hand := bg.canCharacterAttack(apprentice, currTurns)
				if hand > -1 {
					showPause("Apprentice Attacks with " + apprentice.handSlots[hand].name + "!")
					currTurns -= apprentice.handSlots[hand].atkTurns
					bg.doPlayerAttack(APP_TURN, hand)
				} else {
					showPause("Unable to attack: not enough turns or no usable weapon equipped!")					
				}
			}
			
			if !bg.monster.isAlive() {
				// killed monster, flag mission as success
				result = FINISHED_MISSION
			}
			
		} else if strings.Contains(rsp, "search") {
			if currTurns < 1 {
				fmt.Println("No turns remain. End your turn.")
				fmt.Scanln(&rsp3)
			} else {
				currTurns -= 1
				bg.searchLocation(bg.turn)			
			}			
		} else if strings.Contains(rsp, "exit") {
			playFlag = bg.showConfimExit()
			
		} else if strings.Contains(rsp, "die") {
			character.hp = 0
			apprentice.hp = 0
			currTurns = 0
			result = DIED
		} else if strings.Contains(rsp, "kill") {
			if strings.Contains(rsp2, "-monst") {
				bg.monster.hp = 0
				result = FINISHED_MISSION
			}
		}
	}
	
	return result
}
