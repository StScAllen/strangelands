// monster.go
package main

import "time"
import "fmt"

// behaviour constants

const STEP_MOVE = 0
const STEP_WAIT = 1
const STEP_ATTACK = 2

const DAMAGE_PHYSICAL = 0
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
	name								string
	id 									int
	wRange								int
	dmgType								int 
	atkTurns							int
	accuracy							int
	paddedMod, leatherMod, chainMod 	int
}

var SOUL_SUCK = MonsterAttack{"Soul Suck", 1, 2, DAMAGE_SOUL, 3, 2, 0, 0, 0}
var CHARGE = MonsterAttack{"Charge", 2, 1, DAMAGE_PHYSICAL, 2, 1, -1, 0, 1}

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

// Looks at monster turns, and attacks available and picks one, or returns -1 if no attacks are available.
func (bg *BattleGrid) getAttack() (int, int) {
	var die Die
	var attacksAvailable = make([]MonsterAttack, 0, 0)
	
	for k := range bg.monster.attacks {
		// get attacks that can be done withing the amount of available turns
		if bg.monster.attacks[k].atkTurns <= bg.monster.moves {
			// add if there is a target within range			
			if bg.getActorInAttackRange(bg.monster.attacks[k].wRange) > -1 {
				attacksAvailable = append(attacksAvailable, bg.monster.attacks[k])			
			}
		}
	}

	// if we have attacks available and targets in range, pick an attack and target
	if len(attacksAvailable) > 0 {
		attackIndex, targetIndex := -1, -1
		attackIndex = die.rollxdx(1, len(attacksAvailable))
	
		actor := bg.getActorInAttackRange(attacksAvailable[attackIndex-1].wRange)
		
		if (actor == 2) {
			targetIndex = die.rollxdx(1,2) - 1
		} else {
			targetIndex = actor
		}
		
		for k := range bg.monster.attacks {
			if bg.monster.attacks[k].id == attacksAvailable[attackIndex].id {
				attackIndex = k
			}
		}
		
		return attackIndex, targetIndex
	}
	
	return -1, -1
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
		atk,_ := bg.getAttack()
		return (atk > -1) && (bg.monster.moves >= bg.monster.attacks[atk].atkTurns)

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
			if step.id == STEP_MOVE {
				bg.moveMonsterXY(step.x, step.y)
				bg.monster.moves -= 1
			} else if step.id == STEP_ATTACK {
				attackIndex, tgtIndex := bg.getAttack()
				
				if attackIndex > -1 {
					if (tgtIndex == CHAR_TURN) {
						log.addAi("Monster attacks " + character.name + " with " + bg.monster.attacks[attackIndex].name)
						showPause("Monster attacks " + character.name + " with " + bg.monster.attacks[attackIndex].name)				
					} else {
						log.addAi("Monster attacks " + apprentice.name + " with " + bg.monster.attacks[attackIndex].name)
						showPause("Monster attacks " + apprentice.name + " with " + bg.monster.attacks[attackIndex].name)				
					}
					
					bg.monster.moves -= bg.monster.attacks[attackIndex].atkTurns				
				} else {
					log.addAi("Monster wants to attack but no attack available - waits")
					bg.monster.moves -= 1
				}
				
			} else if step.id == STEP_WAIT {
				log.addAi("Monster waits")
				bg.monster.moves -= 1
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
