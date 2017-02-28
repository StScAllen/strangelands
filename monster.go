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
	// lets just create a simple plan to test pathfinding for now
	var plan AIPlan
	//var die Die
	var tiles [MAX_PF_TILES]Tile
	var count int
	// lets try to go to loc 4, 4
	
	if (bg.isTileOpen(4, 4, bg.monsterGridId, MONST_TURN) && bg.monsterXLoc != 4 && bg.monsterYLoc != 4) {
		log.addAi("MONST AI: Tile 4, 4 is available, trying to get path...")
		count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, 4, 4, bg.monsterGridId)			
		if (count < 1){
			fmt.Println("cannot find path")
			log.addAi("... can't find path!")
		} else {
			fmt.Println("PATH FOUND!")
			log.addAi("PATH FOUND!")
			for k := 0; k < count; k++ {
				log.addAi(fmt.Sprintf("Step: %v   to %v : %v", k, tiles[0].x, tiles[0].y))
			}
		}
		
	} else if (bg.isTileOpen(6, 6, bg.monsterGridId, MONST_TURN) && bg.monsterXLoc != 6 && bg.monsterYLoc != 6){
		count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, 6, 6, bg.monsterGridId)
		if (count < 1){
			fmt.Println("cannot find path")
			log.addAi("... can't find path!")
		} else {
			fmt.Println("PATH FOUND!")
			log.addAi("PATH FOUND!")
			for k := 0; k < count; k++ {
				log.addAi(fmt.Sprintf("Step: %v   to %v : %v", k, tiles[0].x, tiles[0].y))
			}
		}	
	} else if (bg.isTileOpen(8, 8, bg.monsterGridId, MONST_TURN) && bg.monsterXLoc != 8 && bg.monsterYLoc != 8) {
		count, tiles = bg.findPath(bg.monsterXLoc, bg.monsterYLoc, 8, 8, bg.monsterGridId)
		if (count < 1){
			fmt.Println("cannot find path")
			log.addAi("... can't find path!")
		} else {
			fmt.Println("PATH FOUND!")		
			log.addAi("PATH FOUND!")
			for k := 0; k < count; k++ {
				log.addAi(fmt.Sprintf("Step: %v   to %v : %v", k, tiles[0].x, tiles[0].y))
			}
		}	
	}
	fmt.Println("Path found, building plan- ", " Steps:  ", count)
	showPause("")
	//plan.steps = make([]AIStep, 100)
	plan.stepCount = count
	plan.nextStep = 0
	if (count > 0){
		for k:= count-1; k >= 0; k-- {
			plan.steps[k] = getStepFromTile(tiles[k])
		}
	}
	
	bg.drawTestGrid(plan.steps)
	
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
		
		// count, options := bg.getMoveOptions(bg.monsterGridId, bg.monsterXLoc, bg.monsterYLoc)
		// if (count > 0 && len(options) == count){
			// var move = options[die.rollxdx(0, count-1)]
			// bg.moveMonster(move)
			
			// if (bg.isGate(MONST_TURN)){
				// if (selectedGate.gridid1 == bg.monsterGridId){
					// bg.monsterGridId = selectedGate.gridid2
					// bg.monsterXLoc = selectedGate.g2x
					// bg.monsterYLoc = selectedGate.g2y
				// } else {
					// bg.monsterGridId = selectedGate.gridid1
					// bg.monsterXLoc = selectedGate.g1x
					// bg.monsterYLoc = selectedGate.g1y
				// } 
			// }
			
			// if (bg.isMonsterVisible()){
				// bg.drawGrid()
				// time.Sleep(300 * time.Millisecond)
			// }
		// }
	}
	
	log.addAi(fmt.Sprintf("%s move: From: %v : %v (%v) To: %v : %v (%v)", bg.monster.name, oldX, oldY, oldG, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId))
	fmt.Printf("End %s move: %v : %v : %v", bg.monster.name, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId)
	rsp := ""
	fmt.Scanln(&rsp)
	
	return 0	// ok to continue (character didn't die or anything)

}