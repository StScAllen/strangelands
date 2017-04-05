// monster.go
package main

import "time"
import "fmt"

// behaviour constants

const STEP_MOVE = 0
const STEP_WAIT = 1
const STEP_ATTACK = 2

const DAMAGE_PHYSICIAL = 0
const DAMAGE_SOUL = 1

var person_bits = []string {"Head", "Arm", "Arm", "Chest", "Chest", "Leg", "Leg",}
var orb_bits = []string {"Body", "Body", "Body", "Body", "Body", "Body", "Body", "Body", "Body", "Body",}
var quad_bits = []string {"Head", "Torso", "Torso", "Torso", "Leg", "Leg", "Leg", "Leg",}

type Monster struct {
	hp, maxhp                     int
	moves                         int
	name                          string
	agi, str, per, intl, cha, gui int
	attacks						  []MonsterAttack					
	bits						  []string
	plan                          AIPlan
	targets						  []int
	body						  []string
	resistance					  []int
	disturbance1                  string
	disturbance2                  string
}

type MonsterAttack struct {
	id 									int
	wRange								int
	dmgType								int 
	atkTurns							int
	accuracy							int
	paddedMod, leatherMod, chainMod 	int
}

var SOUL_SUCK = MonsterAttack{1, 2, DAMAGE_SOUL, 3, 2, 0, 0, 0}
var CHARGE = MonsterAttack{2, 1, DAMAGE_PHYSICIAL, 2, 1, -1, 0, 1}

type AIStep struct {
	action string
	id     int
	x, y   int
	target int
}

type AIPlan struct {
	steps     [100]AIStep  // any "plan" with more than 100 steps is for fools.
	stepCount int
	maneuver  string
	nextStep  int
	target    int // turn id MONST_TURN/CHAR_TURN/APP_TURN constants
	interrupt int //flags set by character or scene actions that will cause monster to recalculate its plan
	charMoved bool
	appMoved  bool
	charDied  bool
	appDied   bool
}

func (mon *Monster) isAlive() bool {
	if mon.hp < 0 {
		return false
	}
	return true
}

func (mon *Monster) getMonsterMoves() int {
	return mon.agi
}

func (mon *Monster) getMonsterVision() int {
	return mon.per
}

func (mon *Monster) getTotalAttackAdjustment() int {
	return 0
}

func (mon *Monster) getTotalDefenseAdjustment() int {
	return 0
}

func (mon *Monster) getMonsterStealthModifier() int {
	stealth := 0

	if mon.agi < 3 {
		stealth -= 2
	} else if mon.agi < 5 {
		stealth -= 1
	} else if mon.agi < 7 {
		stealth += 0
	} else if mon.agi < 9 {
		stealth += 1
	} else {
		stealth += 2
	}

	return stealth
}

func createMonster(id int) Monster {
	var monster Monster

	if id == 1 {
		monster.name = "Will-O-Wisp"
		monster.str = 4
		monster.agi = 8
		monster.per = 8
		monster.cha = 5
		monster.gui = 4
		monster.intl = 5
	
		monster.bits = orb_bits
		
		monster.disturbance1 = "You see a faint glow to the %v"
		monster.disturbance2 = "A sense of despair washes over you. Something is not right here."
		
		monster.targets = ORB_TARGETS
		monster.body = ORB_STRING
		monster.resistance = []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
		monster.attacks = []MonsterAttack{CHARGE, SOUL_SUCK}
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

func (grid *BattleGrid) placeMonster() {

	var dice Die
	monsterNotPlaced := true
	grid.monsterGridId = dice.rollxdx(1, grid.numGrids-1)

	entityGrid := grid.getEntityGrid(grid.monsterGridId)

	for monsterNotPlaced == true {

		grid.monsterXLoc = dice.rollxdx(1, 30)
		grid.monsterYLoc = dice.rollxdx(1, 14)

		if entityGrid.grid[grid.monsterYLoc][grid.monsterXLoc] == " " {
			monsterNotPlaced = false
			log.addInfo("Monster Placed")

		} else {
			log.addInfo("Cannot place monster at " + entityGrid.grid[grid.monsterYLoc][grid.monsterXLoc])
		}

	}
}

func getStepFromTile(tile Tile) AIStep {

	var step AIStep

	step.id = STEP_MOVE
	step.action = "move"
	step.x = tile.x
	step.y = tile.y

	fmt.Println("Step Act ", step.action, " ", step.x, " ", step.y)

	return step
}

func (bg *BattleGrid) getAttack() (int) {
	
	
	return -1
}

func (bg *BattleGrid) isStepValid(step AIStep) bool {

	if bg.monster.plan.nextStep >= bg.monster.plan.stepCount {
		fmt.Println(step)
		fmt.Println("Pathing:  Destination reached?")
		return false
	}

	if step.id == STEP_MOVE {
		fmt.Println(step)
		if bg.isTileOpen(step.x, step.y, bg.monsterGridId, MONST_TURN) {
			return true
		} else {
			return false
		}
	} else if step.id == STEP_ATTACK {
		return true

	} else if step.id == STEP_WAIT {
		return true
	}

	return false
}

func (bg *BattleGrid) doMonsterActivity() int {

	if (!bg.monster.isAlive()){
		return 0
	}

	oldX := bg.monsterXLoc
	oldY := bg.monsterYLoc
	oldG := bg.monsterGridId

	bg.monster.moves = bg.monster.getMonsterMoves()

	// plan exists?
	if bg.monster.plan.stepCount == -1 {
		// create plan
		bg.monster.plan = bg.createMonsterPlan()
	}

	// handle ai interrupts
	if bg.monster.plan.interrupt != 0 {
		if bg.monster.plan.maneuver == "Attack" {
			if bg.monster.plan.target == CHAR_TURN {
				if bg.monster.plan.charMoved {
					log.addAi("Plan interrupted by character move.")
					bg.monster.plan = bg.createMonsterPlan()
				}
			} else if bg.monster.plan.target == APP_TURN {
				if bg.monster.plan.appMoved {
					log.addAi("Plan interrupted by apprentice move.")
					bg.monster.plan = bg.createMonsterPlan()
				}
			}
		}
	}

	for bg.monster.moves > 0 {
		// get the next step from the monster plan
		step := bg.monster.plan.steps[bg.monster.plan.nextStep]

		if bg.isStepValid(step) {
			bg.monster.moves -= 1

			if step.id == STEP_MOVE {
				bg.moveMonsterXY(step.x, step.y)
			} else if step.id == STEP_ATTACK {
				log.addAi("Monster attacks!")
				showPause("Monster attacks!")
			} else if step.id == STEP_WAIT {
				log.addAi("Monster waits")
			}

			if bg.isMonsterVisible() {
				bg.drawGrid()
				time.Sleep(300 * time.Millisecond)
			}

			bg.monster.plan.nextStep += 1

			if bg.monster.plan.nextStep > bg.monster.plan.stepCount {
				bg.monster.plan = bg.createMonsterPlan()
			}

		} else {
			showPause("Step not valid!")
			bg.monster.plan = bg.createMonsterPlan()
		}

		bg.updateActorVisibility()
	}

	log.addAi(fmt.Sprintf("%s move: From: %v : %v (%v) To: %v : %v (%v)", bg.monster.name, oldX, oldY, oldG, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId))
	fmt.Printf("End %s move: %v : %v : %v", bg.monster.name, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId)
	showPause("")

	return 0 // ok to continue (character didn't die or anything)

}
