// pathfinding.go
// A* implementation, other 

package main
import "container/list"
import "math"

var passableTiles = []string {" ", ".", "/", "\\"}	

type Tile struct {
	id int
	hVal, fVal, gVal float32
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

func createPathfindingGrid(xsiz int, ysiz int) (PathfindingGrid){

	pfGrid := make([][]Tile, ysiz)
	for i := range pfGrid {
		pfGrid[i] = make([]Tile, xsiz)
	}

	counter := 0
	for i := range pfGrid {
		for k := range pfGrid[i] {
			pfGrid[i][k].id = counter
			pfGrid[i][k].hVal = 0
			pfGrid[i][k].fVal = 0
			pfGrid[i][k].gVal = 0
			pfGrid[i][k].parentId = 0
			pfGrid[i][k].closed = false
			
			counter++
		}
	}
	
	var pathGrid PathfindingGrid
	pathGrid.tiles = pfGrid
	
	return pathGrid
}

func (bg *BattleGrid) isPassable(tile string) (bool){
	for i := 0; i < len(passableTiles); i++ {	
		if (tile == passableTiles[i]){
			return true
		}
	}
	return false
}

func (bg *BattleGrid) checkRouteTiles(sx, sy, ex, ey int, pathGrid PathfindingGrid, playGrid Grid){
	
	// check north tile
	var x,y = bg.getXYFromCardinal(sx, sy, NORTH)
	if (x >= 0 && y >= 0 && x <= playGrid.maxX && y <= playGrid.maxY){
		tile := pathGrid.tiles[y][x]
		if tile.closed == false && bg.directionValid(sx, sy, NORTH, playGrid.id){
			tile.hVal = iAbsDiff(sx, ex) + iAbsDiff(sy, ey)	// calc city walk distance
			tile.gVal = 1
			tile.fVal = tile.hVal + tile.gVal
			tile.closed = false;
			tile.parentId = pathGrid.tiles[y][x].id
		} else {
			tile.closed = true
		}	
	}
	
	x,y = bg.getXYFromCardinal(sx, sy, NORTHEAST)
	if (x >= 0 && y >= 0 && x <= playGrid.maxX && y <= playGrid.maxY){
		tile := pathGrid.tiles[y][x]
		if tile.closed == false && bg.directionValid(sx, sy, NORTHEAST, playGrid.id){
			tile.hVal = iAbsDiff(sx, ex) + iAbsDiff(sy, ey)	// calc city walk distance
			tile.gVal = 1
			tile.fVal = tile.hVal + tile.gVal
			tile.closed = false;
			tile.parentId = pathGrid.tiles[y][x].id
		} else {
			tile.closed = true
		}	
	}

	x,y = bg.getXYFromCardinal(sx, sy, EAST)
	if (x >= 0 && y >= 0 && x <= playGrid.maxX && y <= playGrid.maxY){
		tile := pathGrid.tiles[y][x]
		if tile.closed == false && bg.directionValid(sx, sy, EAST, playGrid.id){
			tile.hVal = iAbsDiff(sx, ex) + iAbsDiff(sy, ey)	// calc city walk distance
			tile.gVal = 1
			tile.fVal = tile.hVal + tile.gVal
			tile.closed = false;
			tile.parentId = pathGrid.tiles[y][x].id
		} else {
			tile.closed = true
		}	
	}	
	
}

func (bg *BattleGrid) findPath(sx int, sy int, ex int, ey int, gid int){
	
	playGrid := bg.getEntityGrid(gid)	
	pathGrid := createPathfindingGrid(playGrid.maxX, playGrid.maxY)

	clList := list.New()
	opList := list.New()
	
	
	
	

}