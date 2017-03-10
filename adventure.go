// aventure.go
package main

import "fmt"
import "strings"

func chooseAdventure() {

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
		if currTurns < 1 {
			descrip = "(Inventory) (Status) (End Turn) (Help) (Exit) \n\nAction: "
		} else {
			descrip = "(Move ++) (Attack) (Defend) (Get) (Search) (Cast) (Wait)\n(Inventory) (Status) (End Turn) (Help) (Exit) \n\nAction: "
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
		} else if strings.Contains(rsp, "view"){
			if strings.Contains(rsp2, "-log")  {
				log.displayLog()
			} else if strings.Contains(rsp2, "-pattern")  {
				bg.drawGridPattern(bg.gridPattern)
			}

		}
	}
}
