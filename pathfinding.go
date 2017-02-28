// pathfinding.go
// A* implementation, other 

package main
import "fmt"

const MAX_PF_TILES = 2500

var passableTiles = []string {" ", ".", "/", "\\"}	
var pathGrid PathfindingGrid
var path [MAX_PF_TILES]Tile
var pathCount int

type Tile struct {
	id int
	hVal, fVal, gVal int
	x, y int
	closed bool
	parentId int
}

type PathfindingGrid struct {
	tiles [][]Tile
	destX, destY int
	startX, startY int
	shortestPathCount int
}

func iAbsDiff(x1, x2 int) (int){
	if (x1 > x2){
		return x1 - x2
	} else {
		return x2 - x1
	}
}

func getCityBlockDistance(sx, sy, ex, ey int) (int){
	return iAbsDiff(sx, ex) + iAbsDiff(sy, ey)
}

func getCrowDistance(sx, sy, ex, ey int) (int){
	xDist := iAbsDiff(sx, ex)
	yDist := iAbsDiff(sy, ey)
	
	if (xDist == yDist){
		return xDist
	} 
	
	return (iAbsDiff(xDist, yDist) + xDist)
}

// returns Tile with id of sid and failure flag
func (pfGrid * PathfindingGrid) getTileById(sid int) (int, Tile){
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

func (pfGrid * PathfindingGrid) getTileByCardinal(sid int) (int, Tile){
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

func (bg *BattleGrid) createPathfindingGrid(xsiz, ysiz, ex, ey int, gd Grid) (PathfindingGrid){

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
			pfGrid[i][k].hVal = iAbsDiff(k, ex) + iAbsDiff(i, ey)	// "city block" dist to dest
			
			counter++
		}
	}
	
	var pGrid PathfindingGrid
	pGrid.tiles = pfGrid
	pGrid.destX = ex
	pGrid.destY = ey
	
	return pGrid
}

func (bg *BattleGrid) isPassable(tile string) (bool){
	for i := 0; i < len(passableTiles); i++ {	
		if (tile == passableTiles[i]){
			return true
		} 
	}
	
	fmt.Println("Tile not passable: " + tile)
	return false
}

func showPause(messge string){
	fmt.Println(messge)
	rsp := ""
	fmt.Scanln(&rsp)	
}

func (bg *BattleGrid) getAvailableTiles(tx, ty, pid int, playGrid Grid) (tiles []Tile, count int){

	tmpTiles := make([]Tile, 8)
	count = 0
	
	for dir := 0; dir < 8; dir++ {	// loop through cardinals
		fmt.Println("Original XY: ", tx, ty, dir)
		x, y := getXYFromCardinal(tx, ty, dir)
		fmt.Println("After XY: ", x, y, " - ", playGrid.maxY, playGrid.maxX)
		if (x >= 0 && y >= 0 && x <= playGrid.maxY && y <= playGrid.maxX){
			fmt.Println(fmt.Sprintf("range- from: %v, %v  to: %v, %v ", len(pathGrid.tiles), len(pathGrid.tiles[0]), y, x))
			tile := pathGrid.tiles[y][x]
			if (tile.closed == false && bg.directionValid(tx, ty, dir, playGrid.id)) {
				//showPause(fmt.Sprintf("range off: tmptiles- %v, %v ", len(tmpTiles), count))			
				tmpTiles[count] = tile
				count++
			} else {
				pathGrid.tiles[y][x].closed = true	// not available, so close tile
			}
		}
	}

	if count > 0 {
		tiles = make([]Tile, count)
		for k := 0; k < count; k++ {
			tiles[k] = tmpTiles[k]
			tiles[k].parentId = pid
			pathGrid.tiles[tiles[k].y][tiles[k].x].parentId = pid
		}
		fmt.Println(tiles)
	} else {
		tiles = make([]Tile, 1)
	}

	//showPause(fmt.Sprintf("Found %v tiles!", count))
	
	return tiles, count
}

/* func (bg *BattleGrid) checkRouteTiles(sx, sy, ex, ey int, playGrid Grid) (int){
	
	availTiles, tCount := bg.getAvailableTiles(sx, sy, playGrid)
	
	var lowestFTile Tile
	lowestF := 99
	skipper := 0
	
	for k := 0; k < tCount; k++ {
		cX := availTiles[k].x
		cY := availTiles[k].y
		
		if (pathGrid.tiles[cY][cX].parentId > -1){	// previously accessed tile
			_, parentTile := pathGrid.getTileById(pathGrid.tiles[cY][cX].parentId)
			newG := parentTile.gVal + 1
			newF := newG + pathGrid.tiles[cY][cX].hVal
			
			if (newF <= pathGrid.tiles[cY][cX].fVal){
				pathGrid.tiles[cY][cX].parentId = pathGrid.tiles[sy][sx].id
				pathGrid.tiles[cY][cX].fVal = newF
				pathGrid.tiles[cY][cX].gVal = newG
			}
		
			if (lowestF >= pathGrid.tiles[cY][cX].fVal){
				lowestF = pathGrid.tiles[cY][cX].fVal
				lowestFTile = pathGrid.tiles[cY][cX]
				
				showPause(fmt.Sprintf("Closed: %v, %v", cX, cY))
				pathGrid.tiles[sy][sx].closed = true
				skipper++
			} 
		
		} else { // untouched tile
			pathGrid.tiles[cY][cX].parentId = pathGrid.tiles[sy][sx].id
			pathGrid.tiles[cY][cX].gVal = pathGrid.tiles[sy][sx].gVal + 1
			pathGrid.tiles[cY][cX].fVal = pathGrid.tiles[cY][cX].gVal + pathGrid.tiles[cY][cX].hVal
			
			if (lowestF >= pathGrid.tiles[cY][cX].fVal){
				lowestF = pathGrid.tiles[cY][cX].fVal
				lowestFTile = pathGrid.tiles[cY][cX]
			} 
			
			if (cX == ex && cY == ey){
				// end tile!
				fmt.Println("+++ Found END TILE! +++")
				lowestF = pathGrid.tiles[cY][cX].fVal
				lowestFTile = pathGrid.tiles[cY][cX]
				kX := cX
				kY := cY
				pathCount = 0
				for (kX != pathGrid.startX && kY != pathGrid.startY){
					path[pathCount] = pathGrid.tiles[kY][kX]
					_, tile := pathGrid.getTileById(pathGrid.tiles[kY][kX].parentId)
					kX = tile.x
					kY = tile.y
					pathCount++
				}
								
				return pathCount
			}
		}
	}
	
	if lowestF != 99 {
		bg.drawTestGrid2(lowestFTile.x, lowestFTile.y, playGrid)
	
		showPause(fmt.Sprintf("Found Tile: %v, %v", lowestFTile.x, lowestFTile.y))
		return bg.checkRouteTiles(lowestFTile.x, lowestFTile.y, ex, ey, playGrid)
	} else {
		pathGrid.tiles[sy][sx].closed = true
	}
	
	return -1	// no route found
} */

func (bg *BattleGrid) findPath(sx int, sy int, ex int, ey int, gid int) (int, [MAX_PF_TILES]Tile){
	
	showPause(fmt.Sprintf("Finding path from %v, %v to %v, %v on grid %v", sx, sy, ex, ey, gid))
	
	openList := make([]Tile, 1)
	closedList := make([]Tile, 0)
	
	playGrid := bg.getEntityGrid(gid)	
	pathGrid = bg.createPathfindingGrid(len(playGrid.grid[0]), len(playGrid.grid), ex, ey, playGrid)
	pathGrid.startX = sx
	pathGrid.startY = sy
		
	tile := pathGrid.tiles[sy][sx]
	tile.fVal = 0
	
	openList = append(openList, tile)

	var endFlag int = -1
	
	for len(openList) > 0 {
		// find q tile
		highFIndex := -1
		highF := 999
		for i := 0; i < len(openList); i++ {
			if (openList[i].fVal <= highF){
				highFIndex = i
				highF = openList[i].fVal
			}
		}
		
		// pop q tile
		q := openList[highFIndex]
		openList = append(openList[:highFIndex], openList[highFIndex+1:]...)
		
		// get adjacent tiles
		availTiles, tCount := bg.getAvailableTiles(q.x, q.y, q.id, playGrid)		
		//fmt.Println("*** Caught: ")
		//fmt.Println(availTiles)
		//showPause("")
		for k := 0; k < tCount; k++ {
			kTile := availTiles[k]
			if (kTile.x == ex && kTile.y == ey){
				// destination tile, end search
				pathGrid.tiles[kTile.y][kTile.x].parentId = q.id
				fmt.Println(kTile)
				showPause("")
				endFlag = 1
				break
			} 
		
			pathGrid.tiles[kTile.y][kTile.x].gVal = q.gVal + 1
			// hVal is calculated when the grid is built
			pathGrid.tiles[kTile.y][kTile.x].fVal = pathGrid.tiles[kTile.y][kTile.x].gVal + pathGrid.tiles[kTile.y][kTile.x].hVal
		
		
			// Check for existing tile in openList
			// check for existing tile in closeList
		
			openList = append(openList, kTile)
		}
		
		closedList = append(closedList, q)
		
		if (endFlag == 1){
			break
		}
	}
	
	if (endFlag == 1){
		endTile := pathGrid.tiles[ey][ex]
		fmt.Println("End Step:")		
		fmt.Println(endTile)
		pathCount = 1
		path[0] = endTile
		_, stepTile := pathGrid.getTileById(endTile.parentId)
		fmt.Println("First Step:")
		fmt.Println(stepTile)
		showPause("**************************** COMPLETE PATH FOUND!")
		for stepTile.x != ex && stepTile.y != ey {
			id := stepTile.parentId
			_, stepTile = pathGrid.getTileById(id)
			path[pathCount] = stepTile
			fmt.Println(pathCount)
			fmt.Println(path[pathCount])
			showPause("")
			pathCount++
		}
	} else {
		showPause("**************************** No Path Found!")
	}
	
	return pathCount, path
}