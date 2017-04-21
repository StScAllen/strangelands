// aventure.go
package main

import "fmt"
import "strings"

const FLED_BATTLE = 0
const FINISHED_MISSION = 1
const DIED = -1

func chooseAdventure() {

}

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
			showPause("Character hits!")
		} else {
			showPause("Character misses!")
			return
		}

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
		showPause("Hit on " + bg.monster.body[totalTarget-1])

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

	} else {

	}
}

// return - hand value
func (bg *BattleGrid) canCharacterAttack(char Character, currTurns int) int {
	if bg.isMonsterVisible() {
		if bg.isMonsterInAttackRange(bg.turn) && bg.isAttackPathClear(bg.turn) {
			if char.handSlots[LEFT].typeCode == ITEM_TYPE_WEAPON && char.handSlots[LEFT].atkTurns <= currTurns {
				return LEFT
			} else if char.handSlots[RIGHT].typeCode == ITEM_TYPE_WEAPON && char.handSlots[RIGHT].atkTurns <= currTurns {
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
		}

		actions += "\n(Inventory) (Status) (End Turn) (Help) (Exit) \n\nAction: "
	}

	return actions
}

func adventure() (result int) {

	var bg = buildBattleGrid(1)
	rsp := ""
	rsp2 := ""
	rsp3 := ""

	currTurns := 0
	maxTurns := 0
	// mission start setup
	currTurns = character.getCharacterMoves()
	maxTurns = character.getCharacterMoves()

	result = FLED_BATTLE
	
	for rsp != "exit" && rsp != "Exit" {
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
		} else if strings.Contains(rsp, "end") && strings.Contains(rsp2, "turn") {
			// TODO:  THIS SHOULD BE 0!
			if currTurns > 99 {
				fmt.Println("You still have turns remaining... the darkness patiently waits.")
				fmt.Scanln(&rsp3)
			} else {
				if bg.turn == CHAR_TURN {
					if bg.hasApprentice {
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
					bg.currGrid = bg.monsterGridId
					rslt := bg.doMonsterActivity()
					bg.turnCounter++
					
					if rslt == 0 {
						currTurns = character.getCharacterMoves()
						maxTurns = character.getCharacterMoves()
						bg.turn = CHAR_TURN
						bg.currGrid = bg.charGridId
						bg.monster.plan.interrupt = 0 // clear any interrupts
					} else if rslt == CHARACTER_KILLED {
						rsp = "exit"
						result = DIED
						// character has died!
						// if character dies, and they have an apprentice, the apprentice becomes the new character
						// assuming the apprentice lives
					}  else if rslt ==  APPRENTICE_KILLED {
						// apprentice has died!
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
				showPause("Power Balance: " + getSignedFloat32(bg.calcPowerBalance()))
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
					showPause("Character Attacks!")
					currTurns -= character.handSlots[hand].atkTurns
					bg.doPlayerAttack(CHAR_TURN, hand)
				}
			} else {
				hand := bg.canCharacterAttack(apprentice, currTurns)
				if hand > -1 {
					showPause("Apprentice Attacks!")
					currTurns -= apprentice.handSlots[hand].atkTurns
					bg.doPlayerAttack(APP_TURN, hand)
				}
			}
			
			if !bg.monster.isAlive() {
				// killed monster, flag mission as success
				result = FINISHED_MISSION
			}
			
		} else if strings.Contains(rsp, "search") {
			currTurns -= 1
			bg.searchLocation(bg.turn)
		}
	}
	
	return result
}
