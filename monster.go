// monster.go
package main

import "time"
import "fmt"

// behaviour constants

const STEP_MOVE = 0
const STEP_WAIT = 1
const STEP_ATTACK = 2

type Monster struct {
	hp, maxhp int
	moves int
	name string
	agi, str, per, intl, cha, gui int
	plan AIPlan
}

type AIStep struct {
	action string
	id int
	x, y int
}

type AIPlan struct {
	steps [100]AIStep
	stepCount int
	maneuver string
	nextStep int
}

func (mon * Monster) getMonsterMoves() (int) {
	return mon.agi
}

func (mon * Monster) getMonsterVision() (int) {
	return mon.per
}

func createMonster(id int) (Monster){
	var monster Monster

	if (id == 1){
		monster.name = "Will-O-Wisp"
		monster.str = 4
		monster.agi = 8
		monster.per = 8
		monster.cha = 5
		monster.gui = 4
		monster.intl = 5
	}
	
	monster.moves = monster.agi
	monster.hp = monster.str
	monster.maxhp = monster.str
	
	// create a void plan
	var initPlan AIPlan
	initPlan.stepCount = -1
	initPlan.maneuver = "init"
	initPlan.nextStep = -1
	
	monster.plan = initPlan
	
	return monster
}

func getStepFromTile(tile Tile) (AIStep){

	var step AIStep
	
	step.id = STEP_MOVE
	step.action = "move"
	step.x = tile.x
	step.y = tile.y
	
	fmt.Println("Step Act ", step.action, " ", step.x, " ",  step.y)
	
	return step
}

func (bg *BattleGrid) isStepValid(step AIStep) (bool) {

	if bg.monster.plan.nextStep >= bg.monster.plan.stepCount {
		fmt.Println(step)
		fmt.Println("Pathing:  Destination reached?")
		return false
	}

	if (step.id == STEP_MOVE){
		fmt.Println(step)
		if (bg.isTileOpen(step.x, step.y, bg.monsterGridId, MONST_TURN)){
			return true
		} else {
			return false
		}
	}

	return false
}

func (bg * BattleGrid) createMonsterPlan()(AIPlan){
	var plan AIPlan
	var tiles [MAX_PF_TILES]Tile
	var count int
	var die Die
	
//	monsterSeen := bg.isMonsterVisible()
//	apprenticeSeen := bg.isApprenticeVisible()
	characterSeen := bg.isCharacterVisible()
	
	monsterMoves := bg.monster.getMonsterMoves()
	
	
	if characterSeen {
		log.addAi("Character Visible: Trying to get path...")
		count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, bg.charXLoc, bg.charYLoc, bg.monsterGridId)	
	
		if (count < monsterMoves) {
			// monster has enough moves to move to character and attack at least once
		}
	} else {
		// patrol to random corner
		rslt := die.rollxdx(1, 4)
		
		if rslt == 1 {  //tl
			if (bg.isTileOpen(2, 2, bg.monsterGridId, MONST_TURN)) {
			
			} else if (bg.isTileOpen(3, 3, bg.monsterGridId, MONST_TURN)){
			
			}
		} else if rslt == 2 { //tr
			if (bg.isTileOpen(30, 2, bg.monsterGridId, MONST_TURN)) {
			
			} else if (bg.isTileOpen(30, 3, bg.monsterGridId, MONST_TURN)){
			
			}
		} else if rslt == 3 {	//bl
			if (bg.isTileOpen(2, 14, bg.monsterGridId, MONST_TURN)) {
			
			} else if (bg.isTileOpen(3, 14, bg.monsterGridId, MONST_TURN)){
			
			}
		} else if rslt == 4 {
			if (bg.isTileOpen(29, 14, bg.monsterGridId, MONST_TURN)) {
			
			} else if (bg.isTileOpen(30, 14, bg.monsterGridId, MONST_TURN)){
			
			}
		}
		
	}
	
	fmt.Println("Path found, building plan- ", " Steps:  ", count)

	//plan.steps = make([]AIStep, 100)
	plan.stepCount = count
	plan.nextStep = 0
	countUp := 0
	if (count > 0){
		for k:= count-1; k >= 0; k-- {
			plan.steps[countUp] = getStepFromTile(tiles[k])
			countUp++
		}
	}
	
	// showPause("")
	bg.drawTestGrid(plan.steps)
	showPause("")
	
	return plan
}

func (bg *BattleGrid) doMonsterActivity() (int){
	
	oldX := bg.monsterXLoc
	oldY := bg.monsterYLoc
	oldG := bg.monsterGridId
	
	bg.monster.moves = bg.monster.getMonsterMoves()
	
	// plan exists?
	if bg.monster.plan.stepCount == -1 {
		// create plan
		bg.monster.plan = bg.createMonsterPlan()
	}
	
	for ; bg.monster.moves > 0; {		
		// get the next step from the monster plan
		step := bg.monster.plan.steps[bg.monster.plan.nextStep]
		
		if (bg.isStepValid(step)){
			bg.monster.moves -= 1
		
			if (step.id == STEP_MOVE){
				bg.moveMonsterXY(step.x, step.y)
			}
			
			if (bg.isMonsterVisible()){
				bg.drawGrid()
				time.Sleep(300 * time.Millisecond)
			}	
			
			bg.monster.plan.nextStep += 1
			
			if (bg.monster.plan.nextStep > bg.monster.plan.stepCount){
				bg.monster.plan = bg.createMonsterPlan()
			}
			
		} else {
			showPause("Step not valid!")
			bg.monster.plan = bg.createMonsterPlan()
		}
				
		bg.updateVisibility()	
	}
	
	log.addAi(fmt.Sprintf("%s move: From: %v : %v (%v) To: %v : %v (%v)", bg.monster.name, oldX, oldY, oldG, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId))
	fmt.Printf("End %s move: %v : %v : %v", bg.monster.name, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId)
	showPause("")
	
	return 0	// ok to continue (character didn't die or anything)

}