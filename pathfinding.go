// pathfinding.go
// A* implementation, other

package main

import "fmt"
import "time"

const MAX_PF_TILES = 200

var passableTiles = []string{" ", ".", "/", "\\"}
var seeThroughTiles = []string{" ", ".", "░"}
var pathGrid PathfindingGrid
var path []Tile
var pathCount int

type Tile struct {
	id               int
	hVal, fVal, gVal int
	x, y             int
	closed           bool
	parentId         int
}

type PathfindingGrid struct {
	tiles             [][]Tile
	destX, destY      int
	startX, startY    int
	shortestPathCount int
}

// returns Tile with id of sid and failure flag
func (pfGrid *PathfindingGrid) getTileById(sid int) (int, Tile) {
	var fTile Tile

	for i := range pfGrid.tiles {
		for k := range pfGrid.tiles[i] {
			if pfGrid.tiles[i][k].id == sid {
				return 0, pfGrid.tiles[i][k]
			}
		}
	}

	showPause(fmt.Sprintf("Cannot find tile in lookup: %v", sid))

	return -1, fTile
}

func (pfGrid *PathfindingGrid) getTileByCardinal(sid int) (int, Tile) {
	var fTile Tile

	for i := range pfGrid.tiles {
		for k := range pfGrid.tiles[i] {
			if pfGrid.tiles[i][k].id == sid {
				return 0, pfGrid.tiles[i][k]
			}
		}
	}

	return -1, fTile
}

func (bg *BattleGrid) createPathfindingGrid(xsiz, ysiz, ex, ey int, gd Grid) PathfindingGrid {

	pfGrid := make([][]Tile, ysiz)
	for i := range pfGrid {
		pfGrid[i] = make([]Tile, xsiz)
	}

	counter := 0
	for i := range pfGrid {
		for k := range pfGrid[i] {
			pfGrid[i][k].id = counter
			pfGrid[i][k].fVal = 0
			pfGrid[i][k].gVal = 0
			pfGrid[i][k].parentId = -1
			pfGrid[i][k].closed = !bg.isPassable(gd.grid[i][k])
			pfGrid[i][k].x = k
			pfGrid[i][k].y = i
			pfGrid[i][k].hVal = iAbsDiff(k, ex) + iAbsDiff(i, ey) // "city block" dist to dest

			counter++
		}
	}

	var pGrid PathfindingGrid
	pGrid.tiles = pfGrid
	pGrid.destX = ex
	pGrid.destY = ey

	return pGrid
}

func (bg *BattleGrid) isPassable(tile string) bool {
	for i := 0; i < len(passableTiles); i++ {
		if tile == passableTiles[i] {
			return true
		}
	}

	return false
}

func (bg *BattleGrid) isSeeThrough(tile string) bool {
	for i := 0; i < len(seeThroughTiles); i++ {
		if tile == seeThroughTiles[i] {
			return true
		}
	}

	return false
}

func (bg *BattleGrid) getAvailableTiles(tx, ty, pid int, playGrid Grid) (tiles []Tile, count int) {

	tmpTiles := make([]Tile, 8)
	count = 0

	for dir := 0; dir < 8; dir++ { // loop through cardinals
		//fmt.Println("Original XY: ", tx, ty, dir)
		x, y := getXYFromCardinal(tx, ty, dir)
		//fmt.Println("After XY: ", x, y, " - ", playGrid.maxY, playGrid.maxX)
		if x >= 0 && y >= 0 && x <= playGrid.maxY && y <= playGrid.maxX {
			//fmt.Println(fmt.Sprintf("range- from: %v, %v  to: %v, %v ", len(pathGrid.tiles), len(pathGrid.tiles[0]), y, x))
			tile := pathGrid.tiles[y][x]
			if tile.closed == false && bg.directionValid(tx, ty, dir, playGrid.id) {
				//showPause(fmt.Sprintf("range off: tmptiles- %v, %v ", len(tmpTiles), count))
				tmpTiles[count] = tile
				count++
			} else {
				pathGrid.tiles[y][x].closed = true // not available, so close tile
			}
		}
	}

	if count > 0 {
		tiles = make([]Tile, count)
		for k := 0; k < count; k++ {
			tiles[k] = tmpTiles[k]
			tiles[k].parentId = pid
			//pathGrid.tiles[tiles[k].y][tiles[k].x].parentId = pid
		}
		//fmt.Println(tiles)
	} else {
		tiles = make([]Tile, 1)
	}

	//showPause(fmt.Sprintf("Found %v tiles!", count))

	return tiles, count
}

func (bg *BattleGrid) findPath(sx int, sy int, ex int, ey int, gid int) (int, []Tile) {

	debugPause(fmt.Sprintf("Finding path from %v, %v to %v, %v on grid %v", sx, sy, ex, ey, gid))

	if sx == ex && sy == ey {
		return -1, path // can't move to/from same square
	}

	path = make([]Tile, 0, 0)
	
	openList := make([]Tile, 0)
	closedList := make([]Tile, 0)

	playGrid := bg.getEntityGrid(gid)
	pathGrid = bg.createPathfindingGrid(len(playGrid.grid[0]), len(playGrid.grid), ex, ey, playGrid)
	pathGrid.startX = sx
	pathGrid.startY = sy

	tile := pathGrid.tiles[sy][sx]
	tile.fVal = 0

	openList = append(openList, tile)

	var endFlag int = -1

	ticks := time.Now().Unix()
	var newTicks int64 = 0
	
	for len(openList) > 0 {
		// find q tile
		highFIndex := -1
		highF := 999
		for i := 0; i < len(openList); i++ {
			if openList[i].fVal <= highF {
				highFIndex = i
				highF = openList[i].fVal
			}
		}

		// pop q tile
		q := openList[highFIndex]
		openList = append(openList[:highFIndex], openList[highFIndex+1:]...)

		// get adjacent tiles
		availTiles, tCount := bg.getAvailableTiles(q.x, q.y, q.id, playGrid)
		// fmt.Println("*** Caught: ")
		// fmt.Println(availTiles)
		// showPause("")
		for k := 0; k < tCount; k++ {
			kTile := availTiles[k]
			if kTile.x == ex && kTile.y == ey {
				// destination tile, end search
				pathGrid.tiles[kTile.y][kTile.x].parentId = q.id
				fmt.Println(kTile)
				showPause("End tile reached")
				endFlag = 1
				break
			}

			pathGrid.tiles[kTile.y][kTile.x].gVal = q.gVal + 1
			// hVal is calculated when the grid is built
			pathGrid.tiles[kTile.y][kTile.x].fVal = pathGrid.tiles[kTile.y][kTile.x].gVal + pathGrid.tiles[kTile.y][kTile.x].hVal

			//			fmt.Println("pathGrid: %v, %v, %v   kTile: %v, %v, %v", pathGrid.tiles[kTile.y][kTile.x].gVal, pathGrid.tiles[kTile.y][kTile.x].hVal, pathGrid.tiles[kTile.y][kTile.x].fVal, kTile.gVal, kTile.hVal, kTile.fVal)

			// Check for existing tile in openList
			skip := false
			for i := range openList {
				iTile := openList[i]
				if iTile.id == kTile.id {
					if iTile.fVal < pathGrid.tiles[kTile.y][kTile.x].fVal {
						skip = true
						break
					}
				}
			}

			// check for existing tile in closeList
			for i := range closedList {
				iTile := closedList[i]
				if iTile.id == kTile.id {
					skip = true
					break
				}
			}

			if !skip {
				pathGrid.tiles[kTile.y][kTile.x].parentId = q.id
				openList = append(openList, pathGrid.tiles[kTile.y][kTile.x])
			}
		}

		closedList = append(closedList, q)

		if endFlag == 1 {
			break
		}
		
		newTicks = time.Now().Unix()
		
		if (newTicks - ticks) > 4 {
			// time out
			log.addWarn("AI timed out finding path.")
			endFlag = -2
			break
		}
	}

	if endFlag == 1 {
		endTile := pathGrid.tiles[ey][ex]
		// fmt.Println("End Step:")
		// fmt.Println(endTile)
		pathCount = 1
		path = append(path, endTile)
		chk, stepTile := pathGrid.getTileById(endTile.parentId)
		if chk == -1 {
			return 1, path
		}
		// fmt.Println("First Step:")
		// fmt.Println(stepTile)
		// fmt.Println(fmt.Sprintf("End coords: %v, %v", sx, sy)) // start position because we are moving backwards
		newX, newY := stepTile.x, stepTile.y
		id := stepTile.parentId
		path = append(path, stepTile)
		pathCount++

		chk, stepTile = pathGrid.getTileById(stepTile.parentId)
		if chk == -1 {
			return 2, path
		}
		
		for true {
			_, checkTile := pathGrid.getTileById(id)
			path = append(path, checkTile)
			// fmt.Println(pathCount)
			// fmt.Println(path[pathCount])
			// showPause("")
			pathCount++
			newX = checkTile.x
			newY = checkTile.y
			id = checkTile.parentId
			if newX == sx && newY == sy {
				break
			}
		}		

	} else {
		if endFlag == -2 {
			fmt.Println("AI Pathfinding timed out!")
		}
		debugPause("**************************** No Path Found!")
		pathCount = -1
	}

	debugPause("**************************** COMPLETE PATH FOUND!")

	return pathCount, path
}
