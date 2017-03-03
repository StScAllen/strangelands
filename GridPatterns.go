// GridPatterns.go

package main

						
var empty = [8][8]int	{
							{-1, -1, -1, -1, -1, -1, -1, -1},
							{-1, -1, -1, -1, -1, -1, -1, -1},
							{-1, -1, -1, 0, -1, -1, -1, -1},
							{-1, -1, -1, -1, -1, -1, -1, -1},
							{-1, -1, -1, -1, -1, -1, -1, -1},
							{-1, -1, -1, -1, -1, -1, -1, -1},	
							{-1, -1, -1, -1, -1, -1, -1, -1},
							{-1, -1, -1, -1, -1, -1, -1, -1},									
						}						
			
func getEmptyPatternCopy()([8][8]int){

	var pattern [8][8]int
	
	for k:= 0; k < len(empty); k++{
		for t: = 0; t < len(empty[]); t++{
			pattern[k][t] = empty[k][t]
		}
	}

	return pattern
}

func getAvailableDirections(pattern [8][8]int, x, y int) ([]int){
	
	dirs := make([]int, 0)

	if y > 0 && pattern[y-1][x] == -1 {	// check add NORTH
		dirs = append(dirs, 0)
	}

	if y < 7 && pattern[y+1][x] == -1 {	// check add SOUTH
		dirs = append(dirs, 2)
	}
	
	if x > 0 && pattern[y][x-1] == -1 {	// check add WEST
		dirs = append(dirs, 3)
	}	
	
	if x < 7 && pattern[y][x+1] == -1 {	// check add WEST
		dirs = append(dirs, 1)
	}
	
	return dirs
}
		
func getGridPatternXY(pattern [8][8]int, gridid int) (x, y int){
	x = -1
	y = -1

	for k:= 0; k < len(empty); k++{
		for t: = 0; t < len(empty[0]); t++{
			pattern[k][t] = gridid
			return t, k
		}
	}

	return x, y
}

func haveGateFor(gates []Gate, id1, id2 int ) (bool){
	for k := 0; k < len(gates); k++ {
		if gates[k].gridid1 == id1 || gates[k].gridid2 == id1 {
			if gates[k].gridid1 == id2 || gates[k].gridid2 == id2 {
				return true
			}
		}
	}
	
	return false
}

func getNeighbors(t, k int, pattern [8][8]int) ([][3]int) {

	var neighbors = make([][4]int, 0)
	
	if (t > 0){
		if (pattern[k][t-1] != -1){
			neighbors = append(neighbors, []int{t-1, k, pattern[k][t-1], 0})
		}
	}
	
	if (t < 31){
		if (pattern[k][t+1] != -1){
			neighbors = append(neighbors, []int{t+1, k, pattern[k][t+1], 2})
		}
	}	
	
	if (k > 0){
		if (pattern[k-1][t] != -1){
			neighbors = append(neighbors, []int{t, k-1, pattern[k-1][t], 3})
		}
	}
	
	if (k < 15){
		if (pattern[k+1][t] != -1){
			neighbors = append(neighbors, []int{t, k+1, pattern[k+1][t], 1})
		}
	}		

	return neighbors
}

func (bg * BattleGrid) findGateLoc(pattern [8][8]int, gate Gate, card int) (x1, y1, x2, y2 int){

	


	return x1, y1, x2, y2
}

func (bg * BattleGrid) createGatesForGrid(pattern [8][8]int){

	var gates []Gate
	counter := 0
	
	for k:= 0; k < len(pattern); k++{
		for t: = 0; t < len(pattern[k]); t++{
			if (pattern[k][t] > -1){
				neighbors := getNeighbors(t, k)
				
				for i := 0; i < len(neighbors); i++ {
					if (!haveGateFor(gates, pattern[k][t], neighbors[i][2])){
						var newGate Gate
						newGate.gridid1 = pattern[k][t]
						newGate.gridid2 = neighbors[i][2]
						newGate.g1x, newGate.g1y, newGate.g2x, newGate.g2y = findGateLoc(newGate)
						counter++
					}
				}
			}
		}
	}
}
		
func (bg * BattleGrid) createGridPattern(){
	
	var die Die
	dirs := []int{0, 0, 0, 0}
	nextGrid = 1
	gateCount := 0
	numGrids := bg.numGrids
	numGates := 0
	// first, lets define our grid. grid 0 is always 3,2:
	gx := 3
	gy := 2
	
	gridPattern := getEmptyPatternCopy()
	
	//how many gates attached to grid 0?
	gridsPlaced = make([]int, 1)
	gridsPlaced[0] = 0
	
	roll := die.rollxdx(1, 10)
	
	for num := 0; num < numGrids; num++ {
		gridid = gridsPlaced[die.rollxdx(1, len(gridsPlaced))-1]
		x, y := getGridPatternXY(gridPattern, gridid)
		
		if (x > -1 && y > -1){
			dirs = getAvailableDirections(gridPattern, x, y)		

			if (len(dirs) > 0){
				finalDirection := dirs[die.rollxdx(1, len(dirs))-1]
				
				if finalDirection == 0 {
					y -= 1
				} else if finalDirection == 2 {
					y += 1
				} else if finalDirection == 1 {
					x += 1
				} else {
					x -= 1
				}
				
				gridPattern[y][x] = nextGrid
				gridsPlaced = append(gridsPlaced, nextGrid)
				nextGrid += 1
			} else {
				// failed to get a direction with this grid, do over
				num -= 1
			}
		}
	}

	bg.createGatesForGrid(gridPattern)

}