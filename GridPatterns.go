// GridPatterns.go

package main

import "fmt"

var empty = [8][8]int{
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, 0, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
	{-1, -1, -1, -1, -1, -1, -1, -1},
}

func getEmptyPatternCopy() [8][8]int {

	var pattern [8][8]int

	for k := 0; k < len(empty); k++ {
		for t := 0; t < len(empty[k]); t++ {
			pattern[k][t] = empty[k][t]
		}
	}

	return pattern
}

func getAvailableDirections(pattern [8][8]int, x, y int) []int {

	dirs := make([]int, 0)

	if y > 0 && pattern[y-1][x] == -1 { // check add NORTH
		dirs = append(dirs, 0)
	}

	if y < 7 && pattern[y+1][x] == -1 { // check add SOUTH
		dirs = append(dirs, 2)
	}

	if x > 0 && pattern[y][x-1] == -1 { // check add WEST
		dirs = append(dirs, 3)
	}

	if x < 7 && pattern[y][x+1] == -1 { // check add WEST
		dirs = append(dirs, 1)
	}

	return dirs
}

func getGridPatternXY(pattern [8][8]int, gridid int) (x, y int) {
	x = -1
	y = -1

	for k := 0; k < len(empty); k++ {
		for t := 0; t < len(empty[0]); t++ {
			if pattern[k][t] == gridid {
				return t, k
			}
		}
	}

	return x, y
}

func haveGateFor(gates []Gate, id1, id2 int) bool {
	for k := 0; k < len(gates); k++ {
		if gates[k].gridid1 == id1 && gates[k].gridid2 == id2 {
			return true
		}
		if gates[k].gridid1 == id2 && gates[k].gridid2 == id1 {
			return true
		}
	}

	return false
}

func getNeighbors(x, y int, pattern [8][8]int) [][4]int {

	var neighbors = make([][4]int, 0)

	if x > 0 {
		if pattern[y][x-1] != -1 {
			neighbors = append(neighbors, [4]int{x - 1, y, pattern[y][x-1], 0})
		}
	}

	if x < 7 {
		if pattern[y][x+1] != -1 {
			neighbors = append(neighbors, [4]int{x + 1, y, pattern[y][x+1], 2})
		}
	}

	if y > 0 {
		if pattern[y-1][x] != -1 {
			neighbors = append(neighbors, [4]int{x, y - 1, pattern[y-1][x], 3})
		}
	}

	if y < 7 {
		if pattern[y+1][x] != -1 {
			neighbors = append(neighbors, [4]int{x, y + 1, pattern[y+1][x], 1})
		}
	}

	return neighbors
}

func (bg *BattleGrid) findGateLoc(pattern [8][8]int, gate Gate) (x1, y1, x2, y2 int) {

	var die Die
	cardFlag, primId := 0, 0

	for k := 0; k < len(pattern); k++ {
		for t := 0; t < len(pattern[k]); t++ {
			if pattern[k][t] == gate.gridid1 {
				yadj := 0
				yadj = k + 1
				xadj := 0
				xadj = t + 1
				if yadj < 7 && pattern[k+1][t] == gate.gridid2 {
					// north/south relationship gridid1 on top
					cardFlag = NORTH
					primId = gate.gridid1
				} else if (xadj < len(pattern[k])-1) && (pattern[k][t+1] == gate.gridid2) {
					// east/west relationship gridid1 west
					cardFlag = WEST
					primId = gate.gridid1
				}

			} else if pattern[k][t] == gate.gridid2 {
				yadj := 0
				yadj = k + 1
				xadj := 0
				xadj = t + 1
				if yadj < 7 && pattern[yadj][t] == gate.gridid1 {
					// north/south relationship gridid2 on top
					cardFlag = NORTH
					primId = gate.gridid2
				} else if (yadj < len(pattern[k])-1) && xadj < 7 && (pattern[k][t+1] == gate.gridid1) {
					// east/west relationship gridid2 west
					cardFlag = WEST
					primId = gate.gridid2
				}
			}
		}
	}

	searchFlag := true

	for searchFlag {
		if cardFlag == NORTH {
			x := die.rollxdx(2, 30)
			if primId == gate.gridid1 {
				if bg.allGrids[gate.gridid1].grid[15][x] == "─" {
					if bg.allGrids[gate.gridid2].grid[0][x] == "─" {
						// safe match, add gatelocs
						x1 = x
						x2 = x
						y1 = 15
						y2 = 0
						break
					}
				}
			} else {
				if bg.allGrids[gate.gridid2].grid[15][x] == "─" {
					if bg.allGrids[gate.gridid1].grid[0][x] == "─" {
						// safe match, add gatelocs
						x1 = x
						x2 = x
						y1 = 0
						y2 = 15
						break
					}
				}
			}
		} else if cardFlag == WEST {
			y := die.rollxdx(2, 14)
			fmt.Println("Newgate is ", gate.gridid1, " ", gate.gridid2, " Y: ", y)

			if primId == gate.gridid1 {
				if bg.allGrids[gate.gridid1].grid[y][31] == "│" {
					if bg.allGrids[gate.gridid2].grid[y][0] == "│" {
						// safe match, add gatelocs
						x1 = 31
						x2 = 0
						y1 = y
						y2 = y
						break
					}
				}
			} else {
				if bg.allGrids[gate.gridid2].grid[y][31] == "│" {
					if bg.allGrids[gate.gridid1].grid[y][0] == "│" {
						// safe match, add gatelocs
						x1 = 0
						x2 = 31
						y1 = y
						y2 = y
						break
					}
				}
			}

		}
	}

	return x1, y1, x2, y2
}

func (bg *BattleGrid) createGatesForGrid(pattern [8][8]int) {

	var gates []Gate
	counter := 0

	fmt.Println("Creating gates for BattleGrid...")

	for k := 0; k < len(pattern); k++ {
		for t := 0; t < len(pattern[k]); t++ {
			if pattern[k][t] > -1 {
				neighbors := getNeighbors(t, k, pattern)

				for i := 0; i < len(neighbors); i++ {
					if !haveGateFor(gates, pattern[k][t], neighbors[i][2]) {
						var newGate Gate
						newGate.gridid1 = pattern[k][t]
						newGate.gridid2 = neighbors[i][2]
						newGate.g1x, newGate.g1y, newGate.g2x, newGate.g2y = bg.findGateLoc(pattern, newGate)
						bg.gates[counter] = newGate
						gates = append(gates, newGate)
						bg.allGrids[newGate.gridid1].grid[newGate.g1y][newGate.g1x] = GATE1
						bg.allGrids[newGate.gridid2].grid[newGate.g2y][newGate.g2x] = GATE1
						counter++
					}
				}
			}
		}
	}
}

func (bg *BattleGrid) createGridPattern() {

	var die Die
	dirs := []int{0, 0, 0, 0}
	nextGrid := 1
	numGrids := bg.numGrids
	// first, lets define our grid. grid 0 is always 3,2:

	gridPattern := getEmptyPatternCopy()

	//how many gates attached to grid 0?
	gridsPlaced := make([]int, 1)
	gridsPlaced[0] = 0

	fmt.Println("Building Patterns for ", numGrids, " grids.")

	for num := 0; num < numGrids; num++ {
		gridid := gridsPlaced[die.rollxdx(1, len(gridsPlaced))-1]
		x, y := getGridPatternXY(gridPattern, gridid)

		if x > -1 && y > -1 {
			fmt.Println("Getting Directions for ", gridid, " num is ", num, " numGrids is ", numGrids)
			fmt.Println("Grid xy is ", x, " ", y)
			dirs = getAvailableDirections(gridPattern, x, y)

			if len(dirs) > 0 {
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
				fmt.Println("Placed grid ", nextGrid)
				nextGrid += 1
				if nextGrid >= numGrids {
					break
				}
			} else {
				// failed to get a direction with this grid, do over
				num -= 1
			}
		}
	}

	rsp := ""
	fmt.Scanln(&rsp)
	bg.drawGridPattern(gridPattern)
	bg.createGatesForGrid(gridPattern)

	bg.gridPattern = gridPattern
}
