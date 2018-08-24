// behaviours.go

package main

import "fmt"

/*
*	Ai behaviours
 */

const ACTOR_SPOTTED = 1
const ACTOR_KILLED = 2
const MONST_CHANGED_GRID = 3 

/* 
	*  For this behavior, the monster is adjacent to the character go full out on attacks.
*/ 
func (bg *BattleGrid) getDirectAttackBehavior(target int) (int, []Tile, []AIStep) {
	var endSteps []AIStep
	var tiles []Tile
	var count int

	log.addAi("Adding a direct attack behavior")

	monsterMoves := bg.monster.getMonsterMoves()
	count = monsterMoves
	endSteps = make([]AIStep, monsterMoves)

	for k := 0; k < monsterMoves; k++ {
		var step AIStep
		step.action = "attack"
		step.id = STEP_ATTACK
		step.target = target
		step.x = bg.monsterXLoc
		step.y = bg.monsterYLoc

		endSteps[k] = step
	}

	count = 0

	return count, tiles, endSteps
}

/* 
	*  For this behavior, the monster is going to move toward the character until it is close enough to attack.
*/
func (bg *BattleGrid) getMoveAttackBehavior(target int) (int, []Tile, []AIStep) {

	var endSteps []AIStep
	var tiles []Tile
	var count int
	var targetX, targetY int

	tiles = make([]Tile, 0, 0)
	
	log.addAi("(+) Adding a move attack behavior")

	monsterMoves := bg.monster.getMonsterMoves()

	if target == CHAR_TURN {
		targetX = bg.charXLoc
		targetY = bg.charYLoc
	}

	bestX := bg.monsterXLoc - targetX
	bestY := bg.monsterYLoc - targetY

	if bestX > 0 {
		bestX = targetX + 1
	} else if bestX < 0 {
		bestX = targetX - 1
	} else {
		bestX = targetX
	}

	if bestY > 0 {
		bestY = targetY + 1
	} else if bestY < 0 {
		bestY = targetY - 1
	} else {
		bestY = targetY
	}

	count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, bestX, bestY, bg.monsterGridId)

	if count != -1 && count < monsterMoves {
		// monster has enough moves to move to character and attack at least once
		diff := monsterMoves - count
		endSteps = make([]AIStep, diff)
		for i := 0; i < diff; i++ {
			var aiStep AIStep
			aiStep.action = "attack"
			aiStep.id = STEP_ATTACK
			aiStep.x = bestX
			aiStep.y = bestY
			aiStep.target = target
			endSteps[i] = aiStep
		}
	} else if count == -1 {
		count, tiles, endSteps = bg.getPatrolBehavior()
	}

	return count, tiles, endSteps
}

/* 
	*  For this behavior, the monster is going to find a gate and then move to it. It will change grids when it arrives at it.
*/
func (bg *BattleGrid) getChangeGridBehavior() (int, []Tile, []AIStep) {
	var tiles []Tile
	var count int
	var endSteps []AIStep
	var die Die
	
	gates := bg.getGatesForGrid(bg.monsterGridId)
	monsterMoves := bg.monster.getMonsterMoves()
	
	tiles = make([]Tile, 0, 0)
	
	if len(gates) > 0 {
		log.addAi("(+) Adding a change gate behavior")
		roll := die.rollxdx(1, len(gates)) - 1
	
		destGate := gates[roll]
		ex, ey := 0, 0
		
		if destGate.gridid1 == bg.monsterGridId {
			ex = destGate.g1x
			ey = destGate.g1y
		} else {
			ex = destGate.g2x
			ey = destGate.g2y		
		}
		
		count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, ex, ey, bg.monsterGridId)
	
		if count != -1 && count < monsterMoves {
			diff := monsterMoves - count
			endSteps = make([]AIStep, diff)
			for i := 0; i < diff; i++ {
				var aiStep AIStep
				aiStep.action = "wait"
				aiStep.id = STEP_WAIT
				aiStep.x = 0
				aiStep.y = 0
				endSteps[i] = aiStep
			}
		} else if count == -1 {
			count, tiles, endSteps = bg.getPatrolBehavior()
		}
	
	} else {
		return -1, tiles, endSteps
	}
	
	return count, tiles, endSteps
}

/* 
	*  For this behavior, the monster is going to seek a non-visible tile and then move to it. It will wait the rest of its turn when arriving.
*/
func (bg *BattleGrid) getHideBehavior() (int, []Tile, []AIStep) {
	var count int
	var die Die
	var endSteps []AIStep
	var tiles []Tile

	log.addAi("(+) Adding a hiding behavior")

	grid := bg.getEntityGrid(bg.monsterGridId)
	
	monsterMoves := bg.monster.getMonsterMoves()
	tiles = make([]Tile, 0, 0)
	
	// lets look at surrounding tiles in this grid and find one that is obscured...
	
	nearestX, nearestY, nearestDist := 0,0,9999
	for i := 0; i < len(grid.grid); i++ {
		for t := 0; t < len(grid.grid[i]); t++ {
			if bg.isTileObscured(t, i, bg.monsterGridId){
				distX := iAbsDiff(t, bg.monsterXLoc)
				distY := iAbsDiff(i, bg.monsterYLoc)
				thisDist := 0
				
				if distX == distY {
					thisDist = distX
				} else {
					if distX > distY {
						thisDist = distX
					} else {
						thisDist = distY
					}
				}
				
				if nearestDist > thisDist && bg.isTileOpen(t, i, bg.monsterGridId, MONST_TURN) {
					nearestDist = thisDist
					nearestX = t
					nearestY = i
				} else if nearestDist == thisDist && bg.isTileOpen(t, i, bg.monsterGridId, MONST_TURN) {
					if die.rollxdx(1, 5) > 3 {
						nearestDist = thisDist
						nearestX = t
						nearestY = i
					}
				}			
			}
		}
	}
	
	count = -1
	if nearestX+nearestY > 0 {
		count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, nearestX, nearestY, bg.monsterGridId)	
	}

	if count != -1 && count < monsterMoves {
		diff := monsterMoves - count
		endSteps = make([]AIStep, diff)
		for i := 0; i < diff; i++ {
			var aiStep AIStep
			aiStep.action = "wait"
			aiStep.id = STEP_WAIT
			aiStep.x = nearestX
			aiStep.y = nearestY
			endSteps[i] = aiStep
		}
	}

	return count, tiles, endSteps
}

/* 
	*  For this behavior, the monster is going to move to a random location in the current grid.
*/
func (bg *BattleGrid) getRandomMoveBehavior() (int, []Tile, []AIStep) {
	var count int
	var die Die
	var endSteps []AIStep
	var tiles []Tile
	
	tiles = make([]Tile, 0, 0)
	log.addAi("(+) Adding a random Move behavior")

	monsterMoves := bg.monster.getMonsterMoves()

	pathx, pathy := 0, 0
	
	foundFlag := false
	
	for !foundFlag {
		pathx = die.rollxdx(2, 28)
		pathy = die.rollxdx(2, 12)	
		
		if bg.isTileOpen(pathx, pathy, bg.monsterGridId, MONST_TURN) {
			foundFlag = true
		}
	}
	
	count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx, pathy, bg.monsterGridId)

	if count != -1 && count < monsterMoves {
		endSteps = make([]AIStep, 1, 1)

		var aiStep AIStep
		aiStep.action = "defend"
		aiStep.id = STEP_DEFEND
		aiStep.x = pathx
		aiStep.y = pathy
		endSteps[0] = aiStep

	} 

	return count, tiles, endSteps
}


/* 
	*  For this behavior, the monster is going to Patrol to a random corner of the current grid.
*/
func (bg *BattleGrid) getPatrolBehavior() (int, []Tile, []AIStep) {
	var count int
	var die Die
	var endSteps []AIStep
	var tiles []Tile
	
	tiles = make([]Tile, 0, 0)
	log.addAi("(+) Adding a patrolling behavior")

	monsterMoves := bg.monster.getMonsterMoves()

	rslt := die.rollxdx(1, 4)
	pathx, pathy := 16, 12
	if rslt == 1 { //tl
		if bg.isTileOpen(2, 2, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 2, 2
		} else if bg.isTileOpen(3, 3, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 3, 3
		}
	} else if rslt == 2 { //tr
		if bg.isTileOpen(30, 2, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 30, 2
		} else if bg.isTileOpen(30, 3, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 30, 3
		}
	} else if rslt == 3 { //bl
		if bg.isTileOpen(2, 14, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 2, 14
		} else if bg.isTileOpen(3, 14, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 3, 14
		}
	} else if rslt == 4 {
		if bg.isTileOpen(29, 14, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 29, 14
		} else if bg.isTileOpen(30, 14, bg.monsterGridId, MONST_TURN) {
			pathx, pathy = 30, 14
		}
	}

	count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx, pathy, bg.monsterGridId)

	if count != -1 && count < monsterMoves {
		diff := monsterMoves - count
		endSteps = make([]AIStep, diff)
		for i := 0; i < diff; i++ {
			var aiStep AIStep
			aiStep.action = "wait"
			aiStep.id = STEP_WAIT
			aiStep.x = pathx
			aiStep.y = pathy
			endSteps[i] = aiStep
		}
	} else if count == -1 {
		// try a different patrol point
		count, tiles, endSteps = bg.getPatrolBehavior()
	}

	return count, tiles, endSteps
}

/* 
	*  For this behavior, the monster is going to take a couple steps away from the character and then go into 
	*  a defensive shell.
*/
func (bg *BattleGrid) getCowerBehavior() (int, []Tile, []AIStep) {
	var count int
	var endSteps []AIStep
	var tiles []Tile
	
	tiles = make([]Tile, 0, 0)
	log.addAi("(+) Adding a Cowering behavior")

	monsterMoves := bg.monster.getMonsterMoves()
	
	oppX, oppY := 0,0
	
	threat := bg.getBiggestThreat()
	
	count = -1
	
	if threat == -1 {
		// no threat, return do over
		return -1, tiles, endSteps
	} else if threat == CHAR_TURN {
		oppX = bg.monsterXLoc - bg.charXLoc
		oppY = bg.monsterYLoc - bg.charYLoc
	} else {
		oppX = bg.monsterXLoc - bg.appXLoc
		oppY = bg.monsterYLoc - bg.appYLoc	
	}
	
	steps := monsterMoves - 2
	pathx, pathy := bg.monsterXLoc, bg.monsterYLoc
	
	if iAbsVal(oppX) < iAbsVal(oppY) {
		if (oppX < 0) {
			// character is mostly to the east of monster
			if bg.canPathAndMoveHere(pathx-steps, pathy, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-steps, pathy, bg.monsterGridId)			
			} else if bg.canPathAndMoveHere(pathx-(steps-1), pathy+1, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-(steps-1), pathy+1, bg.monsterGridId)				
			} else if bg.canPathAndMoveHere(pathx-(steps-1), pathy-1, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-(steps-1), pathy-1, bg.monsterGridId)				
			}
			
		} else {
			// character is mostly to the west of monster			
			if bg.canPathAndMoveHere(pathx+steps, pathy, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx+steps, pathy, bg.monsterGridId)			
			} else if bg.canPathAndMoveHere(pathx+(steps-1), pathy+1, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx+(steps-1), pathy+1, bg.monsterGridId)				
			} else if bg.canPathAndMoveHere(pathx+(steps-1), pathy-1, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx+(steps-1), pathy-1, bg.monsterGridId)				
			}
		}
	} else {
		if (oppY < 0) {
			// character is mostly to the south of monster
			if bg.canPathAndMoveHere(pathx, pathy-steps, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-steps, pathy, bg.monsterGridId)			
			} else if bg.canPathAndMoveHere(pathx+1, pathy-(steps-1), monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx+1, pathy-(steps-1), bg.monsterGridId)				
			} else if bg.canPathAndMoveHere(pathx-1, pathy-(steps-1), monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-1, pathy-(steps-1), bg.monsterGridId)				
			}			
		} else {
			// character is mostly to the north of monster			
			if bg.canPathAndMoveHere(pathx, pathy+steps, monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-steps, pathy, bg.monsterGridId)			
			} else if bg.canPathAndMoveHere(pathx+1, pathy+(steps-1), monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx+1, pathy+(steps-1), bg.monsterGridId)				
			} else if bg.canPathAndMoveHere(pathx-1, pathy+(steps-1), monsterMoves) {
				count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx-1, pathy+(steps-1), bg.monsterGridId)				
			}				
		}	
	}
	
	if count == -1 {
		// no good grid open, figure something else out
		return count, tiles, endSteps
	}

	if count != -1 && count < monsterMoves {
		endSteps = make([]AIStep, 1, 1)

		var aiStep AIStep
		aiStep.action = "defend"
		aiStep.id = STEP_DEFEND
		aiStep.x = pathx
		aiStep.y = pathy
		endSteps[0] = aiStep

	} 

	return count, tiles, endSteps
}

// checks to see if destination path is reachable based upon tile type and move available
// obstacles could make a direct path unviable within the given moves.
func (bg * BattleGrid) canPathAndMoveHere(pathx, pathy, movesAvail int) (bool) {
	if bg.isTileOpen(pathx, pathy, bg.monsterGridId, MONST_TURN) {
		count,_ := bg.findPath(bg.monsterXLoc, bg.monsterYLoc, pathx, pathy, bg.monsterGridId)
	
		if count <= movesAvail {
			return true
		}
	}
	
	return false
}

// Looks at both character and apprentice and determines who is the biggest threat to us.
// Considers location, grid, and powerlevel.  If neither is on the same grid it will return -1
// Will default to character on draws.
func (bg *BattleGrid) getBiggestThreat() (turn int){
	turn = -1
	power := float32(0.0)
	
	if bg.charGridId == bg.monsterGridId {
		turn = CHAR_TURN
		power = character.getPowerBalance()	
	}
	
	if (bg.hasApprentice && bg.appGridId == bg.monsterGridId){
		if power < apprentice.getPowerBalance() {
			turn = APP_TURN
		}
	}

	return turn
}

// powerbalance is a scale between -X to +X 
// negative balance implies monster has advantage, positive implies character has advantage
// a balance of 0 indicates an even split.  
func (bg * BattleGrid) calcPowerBalance() float32 {
	var balance float32
	balance = 0.0

	balance += character.getPowerBalance()
	if bg.hasApprentice {
		balance += apprentice.getPowerBalance()
	}
	
	balance -= bg.monster.getPowerBalance()

	// monster gets to cheat a bit so give them a balance adjustment
	balance -= 2
	
	return balance
}

func (bg *BattleGrid) createMonsterPlan() AIPlan {
	var plan AIPlan
	var tiles []Tile
	var count int
	var endSteps []AIStep
	var die Die
	//	monsterSeen := bg.isMonsterVisible()
	apprenticeSeen := bg.isApprenticeVisible()
	characterSeen := bg.isCharacterVisible()
	
	powerBalance := bg.calcPowerBalance()
	
	if characterSeen || apprenticeSeen {
		log.addAi("Character Visible: Trying to get path...")

		var charAdj = bg.isActorAdjacent(MONST_TURN, CHAR_TURN)
		var appAdj = bg.isActorAdjacent(MONST_TURN, APP_TURN)

		if !bg.hasApprentice {
			appAdj = false
		}
		
		log.addAi(fmt.Sprintf("Char/App Adj: %s %s", charAdj, appAdj))
		if (charAdj || appAdj){	
			if (powerBalance > float32(die.rollxdx(1, 10))) || (bg.monster.hp == 1 && character.hp > 2) {
				count, tiles, endSteps = bg.getCowerBehavior()	
				plan.maneuver = "Cower"					
			} else {
				if charAdj && !appAdj {
					count, tiles, endSteps = bg.getDirectAttackBehavior(CHAR_TURN)
				} else if appAdj && !charAdj {
					count, tiles, endSteps = bg.getDirectAttackBehavior(APP_TURN)
				} else if charAdj && appAdj {
					if die.rollxdx(1, 2) == 2 {
						count, tiles, endSteps = bg.getDirectAttackBehavior(CHAR_TURN)
					} else {
						count, tiles, endSteps = bg.getDirectAttackBehavior(APP_TURN)
					}
				}			
			}		
		} else {
			// if neither are adjacent then move to attack, or hide
			
			if powerBalance > 3 && die.rollxdx(1, 10) > 6 {
				count, tiles, endSteps = bg.getHideBehavior()	
				plan.maneuver = "Stalk"						
			} else {
				if characterSeen && !apprenticeSeen {
					count, tiles, endSteps = bg.getMoveAttackBehavior(CHAR_TURN)
				} else if !characterSeen && apprenticeSeen {
					count, tiles, endSteps = bg.getMoveAttackBehavior(APP_TURN)
				} else {
					// both seen, choose randomly
					if die.rollxdx(1, 2) == 2 {
						count, tiles, endSteps = bg.getMoveAttackBehavior(CHAR_TURN)
					} else {
						count, tiles, endSteps = bg.getMoveAttackBehavior(APP_TURN)
					}
				}		
				plan.maneuver = "Attack"				
			}
		}

	} else {
		if bg.turnCounter > 5 && bg.monster.gridChangeCoolDown < 1 {
			roll := die.rollxdx(1, 10)
			if roll < 4 {
				// Change grids
				count, tiles, endSteps = bg.getChangeGridBehavior()
				plan.maneuver = "ChangeGrid"				
			
			} else if roll < 7 {
				count, tiles, endSteps = bg.getRandomMoveBehavior()	
				plan.maneuver = "MoveRandom"
				
			} else {
				// patrol to random corner
				count, tiles, endSteps = bg.getPatrolBehavior()
				plan.maneuver = "Patrol"			
			}
		} else {
			// patrol or move random
			roll := die.rollxdx(1, 4)
			if roll < 2 {
				count, tiles, endSteps = bg.getRandomMoveBehavior()	
				plan.maneuver = "MoveRandom"			
			} else {
				count, tiles, endSteps = bg.getPatrolBehavior()
				plan.maneuver = "Patrol"			
			}			
		}
	}

	if count == -1 {
		// whatever we tried to do was unsuccessful, try a random loc move 
		count, tiles, endSteps = bg.getRandomMoveBehavior()	
	}
	
	fmt.Println("Path found, building plan- ", " Steps:  ", count)

	plan.stepCount = count
	plan.nextStep = 0
	plan.charMoved = false
	plan.appMoved = false
	plan.interrupt = 0

	countUp := 0
	if count > 0 {
		for k := count - 1; k >= 0; k-- {
			plan.steps[countUp] = getStepFromTile(tiles[k])
			countUp++
		}
	}
	if len(endSteps) > 0 {
		for k := 0; k < len(endSteps); k++ {
			plan.steps[countUp] = endSteps[k]
			plan.stepCount += 1
			countUp++
		}
	}

	if DEBUG_ON {
		for k := 0; k < countUp; k++ {
			fmt.Println(plan.steps[k])
		}
		showPause("")	
		
		bg.drawTestGrid(plan.steps)
		showPause("")
	}

	return plan
}
