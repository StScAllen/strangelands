// test.go

package main

import "fmt"


func showDialogBox(diag []string) {
	clearConsole()
	
	for k := 0; k < len(diag); k++{
		fmt.Println(diag[k])
	}
	
	rsp := ""
	fmt.Printf("Press any key to continue.")
	fmt.Scanln(&rsp)
}

func showDialogBoxRight(diag []string) {
	clearConsole()
	
	justBuffer := ""
	// have to cast the string because escape sequences will count as double
	iBuffer := 78 - len([]rune(diag[0]))	
	
	for k := 0; k < iBuffer; k++ {
		justBuffer += " "
	}
	
	for k := 0; k < len(diag); k++{
		fmt.Println(justBuffer + diag[k])
	}
	
	rsp := ""
	fmt.Printf("Press any key to continue.")
	fmt.Scanln(&rsp)
}

func printArrayString(arr []string) {

	for k := 0; k < len(arr); k++{
		fmt.Println(len(arr[k]))
		fmt.Println(arr[k])
	}

	rsp := ""
	fmt.Printf("Press any key to continue.")
	fmt.Scanln(&rsp)
}

func showVillages() {

	clearConsole()

	fmt.Println("Villages")
	fmt.Println("--------------")	
	
	fmt.Printf("Currently in %v \n", character.villageIndex)
	
	for i := range villages {
		fmt.Printf("%s   \t%v : %v  \t%v\n", packSpaceString(villages[i].name, 12), villages[i].mapX, villages[i].mapY, getVillageDistance(i))
	}
	
	fmt.Printf("%s   \t%v : %v  \t%v\n", packSpaceString("Keep", 12), keep.mapX, keep.mapY, getVillageDistance(99))

	rsp := ""
	fmt.Printf("Press any key to continue.")
	fmt.Scanln(&rsp)
}

func drawWorldMap(){
	clearConsole()
		
	charX, charY := 0, 0

	if character.villageIndex == 99 {
		charX = keep.mapX
		charY = keep.mapY
	} else {
		charX = villages[character.villageIndex].mapX
		charY = villages[character.villageIndex].mapY
	}
		
	for k := 0; k < len(worldmap2); k++{
		row := fmt.Sprintf("%s", worldmap2[k])
		if charY == k {
			row = fmt.Sprintf("%s", replaceAtIndex2(row, rune('C'), charX-1));
		}
		if k == 1 {
			row += fmt.Sprintf("  Day: %v", gameDay)
		}
		fmt.Println(row) 
	}
	
	rsp := ""
	fmt.Printf("Press any key to continue.")
	fmt.Scanln(&rsp)
}

/*
	Test code to remove for deployment
*/

func (bg *BattleGrid) drawGridPattern(pattern [8][8]int) {
	clearConsole()

	for k := 0; k < len(pattern); k++ {
		row := ""
		for t := 0; t < len(pattern[k]); t++ {
			if pattern[k][t] == 0 {
				row += "⌂" //fmt.Sprintf("%v", pattern[k][t])
			} else if pattern[k][t] > -1 {
				row += "■" //fmt.Sprintf("%v", pattern[k][t])
			} else {
				row += " " //"."
			}
		}

		fmt.Println(row)
	}

	rsp := ""

	fmt.Scanln(&rsp)
}

func (bg *BattleGrid) drawTestGrid(steps [100]AIStep) {
	clearConsole()
	var grid Grid
	var id, xloc, yloc int

	id = bg.monsterGridId
	xloc = bg.monsterXLoc
	yloc = bg.monsterYLoc

	grid = bg.getEntityGrid(id)

	fmt.Println("Map - " + bg.locationName + " - " + grid.gridName + " - " + bg.monster.name + " - " + times[bg.time] + " - " + weather[bg.weather])
	fmt.Println("------------------------------------------")

	row := ""
	skip := false
	for i := 0; i < len(grid.grid); i++ {
		for t := 0; t < len(grid.grid[i]); t++ {
			skip = false
			if bg.charGridId == grid.id && i == bg.charYLoc && t == bg.charXLoc {
				row += "C"
				continue
			} else if bg.hasApprentice && bg.appGridId == grid.id && i == bg.appYLoc && t == bg.appXLoc {
				row += "a"
				continue
			} else if (bg.monsterGridId == grid.id) && (i == bg.monsterYLoc) && (t == bg.monsterXLoc) {
				row += "M"
				continue
			} else {
				for k := 0; k < len(steps); k++ {
					if steps[k].x == t && steps[k].y == i {
						row += "⌂"
						skip = true
						break
					}
				}
			}
			if skip == false {
				row += grid.grid[i][t]
			}
		}

		cgid := id

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

func (bg *BattleGrid) drawTestGrid2(testx, testy int, gd Grid) {
	clearConsole()
	var grid Grid
	var id, xloc, yloc int

	id = bg.monsterGridId
	xloc = bg.monsterXLoc
	yloc = bg.monsterYLoc

	grid = gd

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
				row += "M"
				continue
			} else if testx == t && testy == i {
				row += "⌂"
				continue
			}

			row += grid.grid[i][t]
		}

		cgid := id

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
