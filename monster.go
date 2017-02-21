// monster.go
package main

import "time"
import "fmt"

type Monster struct {
	hp, maxhp int
	moves int
	name string
	agi, str, per, intl, cha, gui int
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
	
	return monster
}

func (bg *BattleGrid) isMonsterVisible() (bool) {

	var sameCharGrid bool = false
	var sameAppGrid bool = false
	
	if (bg.monsterGridId != bg.charGridId && bg.monsterGridId != bg.appGridId){
		return false
	}
	
	
	if (bg.monsterGridId == bg.charGridId){
		sameCharGrid = true
	}
	if (bg.monsterGridId == bg.appGridId){
		sameAppGrid = true
	}

	if (sameCharGrid){
		if (bg.inViewRange(bg.monsterXLoc, bg.monsterYLoc, bg.charXLoc, bg.charYLoc, character.per)){
			return true
		}
	}
	if (sameAppGrid){
		if bg.inViewRange(bg.monsterXLoc, bg.monsterYLoc, bg.appXLoc, bg.appYLoc, apprentice.per) {
			return true
		}
	}	

	return false
}

func (bg *BattleGrid) doMonsterActivity() (int){
	var die Die	
	
	oldX := bg.monsterXLoc
	oldY := bg.monsterYLoc
	oldG := bg.monsterGridId
	
	bg.monster.moves = bg.monster.getMonsterMoves()
	die.rollxdx(1,4)
	for ; bg.monster.moves > 0; {
		bg.monster.moves -= 1
		count, options := bg.getMoveOptions(bg.monsterGridId, bg.monsterXLoc, bg.monsterYLoc)
		if (count > 0 && len(options) == count){
			var move = options[die.rollxdx(0, count-1)]
			bg.moveMonster(move)
			
			if (bg.isGate(MONST_TURN)){
				if (selectedGate.gridid1 == bg.monsterGridId){
					bg.monsterGridId = selectedGate.gridid2
					bg.monsterXLoc = selectedGate.g2x
					bg.monsterYLoc = selectedGate.g2y
				} else {
					bg.monsterGridId = selectedGate.gridid1
					bg.monsterXLoc = selectedGate.g1x
					bg.monsterYLoc = selectedGate.g1y
				} 
			}
			
			if (bg.isMonsterVisible()){
				bg.drawGrid()
				time.Sleep(300 * time.Millisecond)
			}
		}
	}
	
	log.addAi(fmt.Sprintf("%s move: From: %v : %v (%v) To: %v : %v (%v)", bg.monster.name, oldX, oldY, oldG, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId))
	fmt.Printf("End %s move: %v : %v : %v", bg.monster.name, bg.monsterXLoc, bg.monsterYLoc, bg.monsterGridId)
	rsp := ""
	fmt.Scanln(&rsp)
	
	return 0	// ok to continue (character didn't die or anything)

}