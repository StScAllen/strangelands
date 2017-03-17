// aventure.go
package main

import "fmt"
import "strings"

func chooseAdventure() {

}

func (bg * BattleGrid) canCharacterCast(char Character, currTurns int) (bool){
	return true
}

func (bg * BattleGrid) canCharacterAttack(char Character, currTurns int) (bool){
	if bg.isMonsterVisible() {
		if bg.isMonsterInAttackRange(bg.turn) && bg.isAttackPathClear(bt.turn){
			if (char.handSlots[0].typeCode == ITEM_TYPE_WEAPON && char.handSlots[0].atkTurns <= currTurns){
				return true
			} else if (char.handSlots[1].typeCode == ITEM_TYPE_WEAPON && char.handSlots[1].atkTurns <= currTurns){
				return true
			} else {
				fmt.Println("no attackable weapons")
			}
		} else {
			fmt.Println("monster not in range")
		}
	} else {
		fmt.Println("monster not vis")
	}
	
	return false
}

func (bg *BattleGrid) getAvailableActions(char Character, currTurns int) (string){

	actions := ""

	if currTurns < 1 {
		actions = "(Inventory) (Status) (End Turn) (Help) (Exit) \n\nAction: "
	} else {
		actions = "(Move ++)"
		
		if bg.canCharacterAttack(char, currTurns){
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

func adventure() {

	var bg = buildBattleGrid(1)
	rsp := ""
	rsp2 := ""
	rsp3 := ""

	currTurns := 0
	maxTurns := 0
	// mission start setup
	currTurns = character.getCharacterMoves()
	maxTurns = character.getCharacterMoves()

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
						bg.monster.plan.interrupt = 1
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
			if currTurns > 0 {
				fmt.Println("You still have turns remaining... the darkness patiently waits.")
				fmt.Scanln(&rsp3)
			} else {
				if bg.turn == CHAR_TURN {
					if bg.hasApprentice {
						bg.turn = APP_TURN
						currTurns = apprentice.getCharacterMoves()
						maxTurns = apprentice.getCharacterMoves()
					} else {
						bg.turn = MONST_TURN
					}
				} else if bg.turn == APP_TURN {
					bg.turn = MONST_TURN
				}

				if bg.turn == MONST_TURN {
					rslt := bg.doMonsterActivity()

					if rslt == 0 {
						currTurns = character.getCharacterMoves()
						maxTurns = character.getCharacterMoves()
						bg.turn = CHAR_TURN
						bg.monster.plan.interrupt = 0 // clear any interrupts
					} else if rslt == 1 {
						// character has died!
						// if character dies, and they have an apprentice, the apprentice becomes the new character
						// assuming the apprentice lives
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
			}
		} else if strings.Contains(rsp, "inventory") {
			if bg.turn == CHAR_TURN {
				character.showInventory()
			} else {
				apprentice.showInventory()
			}
		}
	}
}
