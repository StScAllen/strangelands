// battlegrid.go
package main

import "fmt"
import "strings"

// cardinal constants
const NORTH = 0
const NORTHEAST = 1
const EAST = 2
const SOUTHEAST = 3
const SOUTH = 4
const SOUTHWEST = 5
const WEST = 6
const NORTHWEST = 7
const STAY = -1

// cardinal constants end

const EMPTY_TILE = " "
const HIDDEN_TILE = "▒"
const FOG_TILE = "░"
const WATER_TILE = "░"

const GATE1 = "/"
const GATE2 = "\\"

const DAY = 0
const NIGHT = 1
const DUSK = 2
const DAWN = 3

const CLEAR = 0 // no penalties
const FOGGY = 1 // vision reduced by 1
const RAIN = 2  // movement reduced by 1
const SNOW = 3  // movement & vision reduced by 1

const CHAR_TURN = 0
const APP_TURN = 1
const MONST_TURN = 2

var selectedGate Gate

var weather = []string{"Sunny", "Foggy", "Rain", "Snow"}
var times = []string{"Day", "Night", "Dusk", "Dawn"}

type Gate struct {
	gridid1  int
	gridid2  int
	g1x, g1y int
	g2x, g2y int
}

type Grid struct {
	grid       [][]string
	id         int
	gridName   string
	maxX, maxY int
}

type BattleGrid struct {
	allGrids                             [32]Grid
	gates                                [200]Gate
	gridPattern                          [8][8]int
	numGrids                             int
	monster                              Monster
	locationName                         string
	charXLoc, charYLoc                   int // character locs
	appXLoc, appYLoc                     int // apprentice locs
	monsterXLoc, monsterYLoc             int // monster1 locs
	charGridId, monsterGridId, appGridId int
	currGrid                             int
	time                                 int
	weather                              int
	turn                                 int // char=0, app=1, monst=2
	hasApprentice                        bool
	monsterSpotted                       bool
	characterSpotted                     bool
	apprenticeSpotted                    bool
}

func (gd *Grid) addCemetaryDecorations() {
	var die Die

	for i := 0; i < die.rollxdx(4, 14); i++ {
		x := die.rollxdx(2, 14)
		y := die.rollxdx(2, 14)

		if die.rollxdx(1, 5) > 2 {
			gd.grid[x][y] = " "
		} else {
			gd.grid[x][y] = "∩"
		}
	}
}

func (bg *BattleGrid) isActorAdjacent(whoFlag, targetFlag int) bool {

	if whoFlag == MONST_TURN {
		if targetFlag == CHAR_TURN {
			if iAbsDiff(bg.monsterXLoc, bg.charXLoc) < 2 && iAbsDiff(bg.monsterYLoc, bg.charYLoc) < 2 {
				// character is adjacent to monster
				return true
			}
		} else if targetFlag == APP_TURN {
			if iAbsDiff(bg.monsterXLoc, bg.appXLoc) < 2 && iAbsDiff(bg.monsterYLoc, bg.appYLoc) < 2 {
				// character is adjacent to monster
				return true
			}
		}
	} else if whoFlag == CHAR_TURN {
		if targetFlag == MONST_TURN {
			if iAbsDiff(bg.monsterXLoc, bg.charXLoc) < 2 && iAbsDiff(bg.monsterYLoc, bg.charYLoc) < 2 {
				// monster is adjacent to character
				return true
			}
		} else if targetFlag == APP_TURN {
			if iAbsDiff(bg.charXLoc, bg.appXLoc) < 2 && iAbsDiff(bg.charYLoc, bg.appYLoc) < 2 {
				// apprentice is adjacent to character
				return true
			}
		}
	}

	return false
}

func (bg *BattleGrid) isCharacterVisible() bool {

	if bg.monsterGridId != bg.charGridId {
		return false
	}

	if bg.inViewRange(bg.charXLoc, bg.charYLoc, bg.monsterXLoc, bg.monsterYLoc, bg.monster.getMonsterVision()) {
		if !bg.characterSpotted {
			log.addAi("Character Spotted!")
		}
		bg.characterSpotted = true
		return true
	}

	return false
}

func (bg *BattleGrid) isApprenticeVisible() bool {

	if bg.hasApprentice == false {
		return false
	}

	if bg.monsterGridId != bg.appGridId {
		return false
	}

	if bg.inViewRange(bg.appXLoc, bg.appYLoc, bg.monsterXLoc, bg.monsterYLoc, bg.monster.getMonsterVision()) {
		if !bg.characterSpotted {
			log.addAi("Apprentice Spotted!")
		}
		bg.characterSpotted = true
		return true
	}

	return false
}

func (bg * BattleGrid) isAttackPathClear() (bool){

}

func (bg * BattleGrid) isMonsterInAttackRange(turn int) (bool){
	if bg.isMonsterVisible() == false {
		return false
	}
	
	weaponRange := 0
	actorX, actorY := 0, 0
	if turn == CHAR_TURN {
		weaponRange = character.getWeaponRange()	
		actorX = bg.charXLoc
		actorY = bg.charYLoc
	} else {
		weaponRange = apprentice.getWeaponRange()	
		actorX = bg.appXLoc
		actorY = bg.appYLoc		
	}

	if weaponRange < 1 {
		return false
	}

	actorDistance := getCrowDistance(actorX, actorY, bg.monsterXLoc, bg.monsterYLoc)

	fmt.Println("actor dist is", actorDistance)
	
	if actorDistance <= weaponRange {
		return true
	}
	
	return false
}

func (bg *BattleGrid) isMonsterVisible() bool {

	var sameCharGrid bool = false
	var sameAppGrid bool = false

	if bg.monsterGridId != bg.charGridId && bg.monsterGridId != bg.appGridId {
		return false
	}

	if bg.monsterGridId == bg.charGridId {
		sameCharGrid = true
	}
	if bg.monsterGridId == bg.appGridId {
		sameAppGrid = true
	}

	if sameCharGrid {
		if bg.inViewRange(bg.monsterXLoc, bg.monsterYLoc, bg.charXLoc, bg.charYLoc, character.per) {
			if !bg.monsterSpotted {
				log.addAi("Monster Spotted!")
			}
			bg.monsterSpotted = true
			return true
		}
	}
	if sameAppGrid {
		if bg.inViewRange(bg.monsterXLoc, bg.monsterYLoc, bg.appXLoc, bg.appYLoc, apprentice.per) {
			return true
		}
	}

	return false
}

func (bg *BattleGrid) isTileObscured(x, y, gridid int) bool {
	if gridid != bg.charGridId && gridid != bg.appGridId {
		return true
	}

	//check := fmt.Sprintf("is tile obscured? %v, %v", x, y, gridid)
	//fmt.Println(check)
	//log.addInfo(check)

	checkGrid := bg.getEntityGrid(gridid)
	hidden := false

	if y == 0 || y == 15 {
		return false
	}
	if x == 0 || x == 31 {
		return false
	}

	if gridid == bg.charGridId {
		// get direction perspective of character to tile
		xdiff := x - bg.charXLoc
		ydiff := y - bg.charYLoc

		if xdiff > 0 {
			if ydiff > 0 {
				if !bg.isSeeThrough(checkGrid.grid[y-1][x-1]) {
					hidden = true
				}
			} else if ydiff < 0 {
				if !bg.isSeeThrough(checkGrid.grid[y+1][x-1]) {
					hidden = true
				}
			} else {
				if !bg.isSeeThrough(checkGrid.grid[y][x-1]) {
					hidden = true
				}
			}
		} else if xdiff < 0 {
			if ydiff > 0 {
				if !bg.isSeeThrough(checkGrid.grid[y-1][x+1]) {
					hidden = true
				}
			} else if ydiff < 0 {
				if !bg.isSeeThrough(checkGrid.grid[y+1][x+1]) {
					hidden = true
				}
			} else {
				if !bg.isSeeThrough(checkGrid.grid[y][x+1]) {
					hidden = true
				}
			}
		} else {
			if ydiff > 0 {
				if !bg.isSeeThrough(checkGrid.grid[y-1][x]) {
					hidden = true
				}
			} else if ydiff < 0 {
				if !bg.isSeeThrough(checkGrid.grid[y+1][x]) {
					hidden = true
				}
			}
		}
	}

	if bg.hasApprentice && gridid == bg.appGridId {
		xdiff := bg.appXLoc - x
		ydiff := bg.appYLoc - y

		if xdiff > 0 {
			if ydiff > 0 {
				if !bg.isSeeThrough(checkGrid.grid[y-1][x-1]) {
					hidden = true
				} else {
					hidden = false
				}
			} else if ydiff < 0 {
				if !bg.isSeeThrough(checkGrid.grid[y+1][x-1]) {
					hidden = true
				} else {
					hidden = false
				}
			} else {
				if !bg.isSeeThrough(checkGrid.grid[y][x-1]) {
					hidden = true
				} else {
					hidden = false
				}
			}
		} else if xdiff < 0 {
			if ydiff > 0 {
				if !bg.isSeeThrough(checkGrid.grid[y-1][x+1]) {
					hidden = true
				} else {
					hidden = false
				}
			} else if ydiff < 0 {
				if !bg.isSeeThrough(checkGrid.grid[y+1][x+1]) {
					hidden = true
				} else {
					hidden = false
				}
			} else {
				if !bg.isSeeThrough(checkGrid.grid[y][x+1]) {
					hidden = true
				} else {
					hidden = false
				}
			}
		} else {
			if ydiff > 0 {
				if !bg.isSeeThrough(checkGrid.grid[y-1][x]) {
					hidden = true
				} else {
					hidden = false
				}
			} else if ydiff < 0 {
				if !bg.isSeeThrough(checkGrid.grid[y+1][x]) {
					hidden = true
				} else {
					hidden = false
				}
			}
		}
	}

	return hidden
}

func (bg *BattleGrid) updateActorVisibility() {
	bg.isMonsterVisible()
	bg.isCharacterVisible()
	bg.isApprenticeVisible()
}

func (bg *BattleGrid) isGate(turn int) bool {

	var xloc, yloc, gridId int

	if turn == CHAR_TURN {
		xloc = bg.charXLoc
		yloc = bg.charYLoc
		gridId = bg.charGridId

	} else if turn == APP_TURN {
		xloc = bg.appXLoc
		yloc = bg.appYLoc
		gridId = bg.charGridId
	} else if turn == MONST_TURN {
		xloc = bg.monsterXLoc
		yloc = bg.monsterYLoc
		gridId = bg.monsterGridId
	}

	for _, g := range bg.gates {
		if g.gridid1 == gridId {
			if g.g1x == xloc && g.g1y == yloc {
				selectedGate = g
				return true
			}
		} else if g.gridid2 == gridId {
			if g.g2x == xloc && g.g2y == yloc {
				selectedGate = g
				return true
			}
		}
	}

	return false
}

func (bg *BattleGrid) getEntityGrid(id int) Grid {
	for _, gd := range bg.allGrids {
		if gd.id == id {
			return gd
		}
	}

	var nilGrid Grid
	nilGrid.id = -1

	return nilGrid
}

func getXYFromCardinal(locX, locY, cardinal int) (x, y int) {

	x, y = locX, locY

	switch cardinal { // switch always breaks unless you use fallthrough

	case NORTH:
		y = locY - 1
	case NORTHEAST:
		y = locY - 1
		x = locX + 1
	case EAST:
		x = locX + 1
	case SOUTHEAST:
		y = locY + 1
		x = locX + 1
	case SOUTH:
		y = locY + 1
	case SOUTHWEST:
		y = locY + 1
		x = locX - 1
	case WEST:
		x = locX - 1
	case NORTHWEST:
		x = locX - 1
		y = locY - 1
	case STAY:
		x = locX
		y = locY
	}

	return x, y
}

func (grid *BattleGrid) moveCharacter(cardinal int) {
	var newX, newY int

	if grid.turn == CHAR_TURN {
		newX = grid.charXLoc
		newY = grid.charYLoc
	} else { // app turn
		newX = grid.appXLoc
		newY = grid.appYLoc
	}

	newX, newY = getXYFromCardinal(newX, newY, cardinal)

	if grid.turn == CHAR_TURN {
		grid.charXLoc = newX
		grid.charYLoc = newY
	} else { // app turn
		grid.appXLoc = newX
		grid.appYLoc = newY
	}
}

func (grid *BattleGrid) moveMonster(cardinal int) {
	var newX, newY int

	newX = grid.monsterXLoc
	newY = grid.monsterYLoc

	newX, newY = getXYFromCardinal(newX, newY, cardinal)

	grid.monsterXLoc = newX
	grid.monsterYLoc = newY
}

func (grid *BattleGrid) moveMonsterXY(x, y int) {
	grid.monsterXLoc = x
	grid.monsterYLoc = y
}

func (bg *BattleGrid) getMoveOptions(gridId int, xloc int, yloc int) (int, []int) {
	count := 0

	for i := 0; i < 8; i++ {
		if bg.directionValid(xloc, yloc, i, gridId) {
			count += 1
		}
	}

	var moves = make([]int, count)

	counter := 0
	for i := 0; i < 8; i++ {
		if bg.directionValid(xloc, yloc, i, gridId) {
			moves[counter] = i
			counter += 1
		}
	}

	return count, moves
}

// checks whether a move in the cardinal direction from current location is valid
// passing in STAY constant will check current block for collision
func (grid *BattleGrid) directionValid(locX int, locY int, cardinal int, gridId int) bool {
	var newX, newY int

	//fmt.Printf("loc is %v  %v  %s \n", locX, locY, grid.grid[locX][locY])
	//fmt.Printf("card is %v  ", cardinal)

	newX = locX
	newY = locY

	newX, newY = getXYFromCardinal(newX, newY, cardinal)

	if newX < 0 || newY < 0 {
		return false
	}

	var tgrid Grid = grid.getEntityGrid(gridId)

	if newX > tgrid.maxY || newY > tgrid.maxX {
		return false
	}

	// check for character collisions
	if cardinal != STAY {
		if gridId == grid.charGridId && newX == grid.charXLoc && newY == grid.charYLoc {
			return false
		} else if grid.hasApprentice && gridId == grid.appGridId && newX == grid.appXLoc && newY == grid.appYLoc {
			return false
		} else if gridId == grid.monsterGridId && newX == grid.monsterXLoc && newY == grid.monsterYLoc {
			return false
		}
	}

	if grid.isPassable(tgrid.grid[newY][newX]) {
		return true
	}

	return false
}

func (grid *BattleGrid) isTileOpen(tx, ty, gridId, turn int) bool {
	var tgrid Grid = grid.getEntityGrid(gridId)

	if gridId == grid.charGridId && tx == grid.charXLoc && ty == grid.charYLoc {
		return false
	} else if grid.hasApprentice && gridId == grid.appGridId && tx == grid.appXLoc && ty == grid.appYLoc {
		return false
	} else if gridId == grid.monsterGridId && tx == grid.monsterXLoc && ty == grid.monsterYLoc && turn != MONST_TURN {
		return false
	}

	if grid.isPassable(tgrid.grid[ty][tx]) {
		return true
	}

	return false
}

func convertCardinalStringToInt(cardinal string) int {

	cardinal = strings.ToUpper(cardinal)

	switch cardinal { // switch always breaks unless you use fallthrough

	case "N":
		return 0
	case "NE":
		return 1
	case "E":
		return 2
	case "SE":
		return 3
	case "S":
		return 4
	case "SW":
		return 5
	case "W":
		return 6
	case "NW":
		return 7
	default:
		return -1
	}
}

func (grid *BattleGrid) inViewRange(x int, y int, charX int, charY int, charPer int) bool {

	var vRange int = 0

	if grid.time == DAY {
		return true
	}

	if charPer < 3 {
		vRange = 1
	} else if charPer < 5 {
		vRange = 2
	} else if charPer < 7 {
		vRange = 3
	} else if charPer < 9 {
		vRange = 4
	} else {
		vRange = 5
	}

	if grid.time == DUSK || grid.time == DAWN {
		vRange += 1
	}

	if grid.weather == FOGGY || grid.weather == SNOW {
		vRange -= 1
	}

	distX := 0
	distY := 0
	if x > charX {
		distX = x - charX
	} else {
		distX = charX - x
	}

	if y > charY {
		distY = y - charY
	} else {
		distY = charY - y
	}

	if distX > vRange || distY > vRange {
		return false
	}
	if ((distX + distY) / 2) > vRange {
		return false
	}

	return true
}

func (bg *BattleGrid) drawGrid() {

	clearConsole()
	var grid Grid
	var id, xloc, yloc int

	if bg.turn == CHAR_TURN {
		id = bg.charGridId
		xloc = bg.charXLoc
		yloc = bg.charYLoc
	} else if bg.turn == APP_TURN {
		id = bg.appGridId
		xloc = bg.appXLoc
		yloc = bg.appYLoc
	} else {
		id = bg.monsterGridId
		xloc = bg.monsterXLoc
		yloc = bg.monsterYLoc
	}

	grid = bg.getEntityGrid(id)

	fmt.Println("Map - " + bg.locationName + " - " + grid.gridName + " - " + bg.monster.name + " - " + times[bg.time] + " - " + weather[bg.weather])
	fmt.Println("------------------------------------------")

	row := ""
	for i := 0; i < len(grid.grid); i++ {
		for t := 0; t < len(grid.grid[i]); t++ {

			if bg.charGridId == grid.id && i == bg.charYLoc && t == bg.charXLoc {
				row += "C"
				continue
			} else if bg.hasApprentice && bg.appGridId == grid.id && i == bg.appYLoc && t == bg.appXLoc {
				row += "a"
				continue
			} else if (bg.monsterGridId == grid.id) && (i == bg.monsterYLoc) && (t == bg.monsterXLoc) {
				if !bg.isPassable(grid.grid[i][t]) {
					log.addAi("Monster is stuck! (" + grid.grid[i][t] + ")")
				}
				if bg.isMonsterVisible() {
					row += "M"
				} else {
					row += HIDDEN_TILE
				}

				continue
			}

			if bg.inViewRange(t, i, bg.charXLoc, bg.charYLoc, character.per) || (bg.hasApprentice && bg.inViewRange(t, i, bg.appXLoc, bg.appYLoc, apprentice.per)) {
				if bg.isTileObscured(t, i, grid.id) {
					row += HIDDEN_TILE
				} else {
					row += grid.grid[i][t]
				}

			} else {
				row += HIDDEN_TILE
			}
		}

		cgid := 0
		if bg.turn == CHAR_TURN {
			cgid = bg.charGridId
		} else if bg.turn == APP_TURN {
			cgid = bg.appGridId
		}

		// draw status rows
		if i == 0 {
			row += "  -PATHS-"
		} else if i == 1 {
			if bg.directionValid(xloc, yloc, 7, cgid) {
				row += "  NW"
			} else {
				row += "    "
			}
			if bg.directionValid(xloc, yloc, 0, cgid) {
				row += " N "
			} else {
				row += "   "
			}
			if bg.directionValid(xloc, yloc, 1, cgid) {
				row += "NE"
			} else {
				row += "  "
			}
		} else if i == 2 {
			if bg.directionValid(xloc, yloc, 6, cgid) {
				row += "  W  "
			} else {
				row += "     "
			}
			if bg.directionValid(xloc, yloc, 2, cgid) {
				row += "   E"
			} else {
				row += "    "
			}
		} else if i == 3 {
			if bg.directionValid(xloc, yloc, 5, cgid) {
				row += "  SW "
			} else {
				row += "     "
			}
			if bg.directionValid(xloc, yloc, 4, cgid) {
				row += "S "
			} else {
				row += "  "
			}
			if bg.directionValid(xloc, yloc, 3, cgid) {
				row += "SE"
			} else {
				row += "  "
			}
		} else if i == 5 {
			row += "  " + character.name + " Health: ["
			for hlth := 0; hlth <= character.maxhp; hlth++ {
				if hlth > character.hp {
					row += "-"
				} else {
					row += "*"
				}
			}
			row += "]"

		} else if i == 6 {
			if bg.hasApprentice {
				row += "  " + apprentice.name + " Health: ["
				for hlth := 0; hlth <= apprentice.maxhp; hlth++ {
					if hlth > apprentice.hp {
						row += "-"
					} else {
						row += "*"
					}
				}
				row += "]"
			}
		} else if i == 7 {
			row += "  " + bg.monster.name + " Health: ["
			for hlth := 0; hlth <= bg.monster.maxhp; hlth++ {
				if hlth > bg.monster.hp {
					row += "-"
				} else {
					row += "*"
				}
			}
			row += "]"
		}

		fmt.Println(row)
		row = ""
	}
}

func (bg *BattleGrid) addGates() {

	var tGate Gate
	var die Die

	// gate 1
	tGate.gridid1 = 0
	tGate.g1x = bg.allGrids[0].maxY
	tGate.g1y = die.rollxdx(2, 12)
	bg.allGrids[0].grid[tGate.g1y][tGate.g1x] = GATE1

	tGate.gridid2 = 1
	tGate.g2x = 0
	tGate.g2y = tGate.g1y
	bg.allGrids[1].grid[tGate.g2y][tGate.g2x] = GATE1

	bg.gates[0] = tGate

	// gate 2
	tGate.gridid1 = 1
	tGate.g1x = bg.allGrids[1].maxY
	tGate.g1y = die.rollxdx(2, 12)
	bg.allGrids[1].grid[tGate.g1y][tGate.g1x] = GATE1

	tGate.gridid2 = 2
	tGate.g2x = 0
	tGate.g2y = tGate.g1y
	bg.allGrids[2].grid[tGate.g2y][tGate.g2x] = GATE1

	bg.gates[1] = tGate

	// gate 3  n/s gate
	tGate.gridid1 = 2
	tGate.g1x = die.rollxdx(2, 12)
	tGate.g1y = bg.allGrids[2].maxX
	bg.allGrids[2].grid[tGate.g1y][tGate.g1x] = GATE1

	tGate.gridid2 = 3
	tGate.g2x = tGate.g1x
	tGate.g2y = 0
	bg.allGrids[3].grid[tGate.g2y][tGate.g2x] = GATE1

	bg.gates[2] = tGate
}

func createSquareGrid(height int, width int) Grid {

	newGrid := make([][]string, height)
	for i := range newGrid {
		newGrid[i] = make([]string, width)
	}

	for k := range newGrid {
		for t := range newGrid[k] {
			newGrid[k][t] = EMPTY_TILE
		}
	}

	// create walled structure
	// set walls
	for t := 1; t < width; t++ {
		newGrid[0][t] = "─"
		newGrid[height-1][t] = "─"
	}
	for t := 1; t < height; t++ {
		newGrid[t][width-1] = "│"
		newGrid[t][0] = "│"
	}

	// set corners...
	newGrid[0][0] = "┌"
	newGrid[height-1][width-1] = "┘"
	newGrid[height-1][0] = "└"
	newGrid[0][width-1] = "┐"

	var retGrid Grid
	retGrid.maxX = height - 1
	retGrid.maxY = width - 1

	retGrid.grid = newGrid

	return retGrid
}

func buildBattleGrid(id int) BattleGrid {

	var grid BattleGrid
	var monster Monster

	fmt.Printf("Building Grid: %v   \n ", id)

	grid.currGrid = 0     // default
	grid.time = DAY       // default
	grid.weather = CLEAR  // default
	grid.turn = CHAR_TURN // default
	grid.numGrids = 4     // default
	grid.hasApprentice = false

	if id == 1 { // cemetary
		grid.numGrids = 4
		// for random, done after number of grids is assigned!

		for k := 0; k < grid.numGrids; k++ {
			g1 := createSquareGrid(16, 32)
			g1.addCemetaryDecorations()
			g1.id = k
			g1.gridName = fmt.Sprintf("%v", k)
			grid.allGrids[k] = g1
			grid.setRandomStamp(g1.maxX, g1.maxY, k)
		}

		grid.createGridPattern()

		grid.currGrid = 0
		monster = createMonster(1)

		grid.monster = monster
		grid.locationName = "Cemetary"

		grid.charXLoc = 1
		grid.charYLoc = 1
		grid.charGridId = 0

		grid.appXLoc = 2
		grid.appYLoc = 1
		grid.appGridId = 0

		grid.characterSpotted = false
		grid.monsterSpotted = false
		grid.apprenticeSpotted = false

		//grid.addGates()
		grid.placeMonster()

	} else {
		/* 		grid.grid = SMALL_GRID
		   		monster.name = "Manticore"
		   		grid.monster = monster */

		grid.charXLoc = 1
		grid.charYLoc = 1

		grid.monsterXLoc = 6
		grid.monsterYLoc = 6
	}

	return grid
}
