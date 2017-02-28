// test.go

package main
import "fmt"

/*
	Test code to remove for deployment
*/

func (bg * BattleGrid) drawTestGrid(steps [100]AIStep){
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
			if (bg.charGridId == grid.id && i == bg.charYLoc && t == bg.charXLoc){
				row += "C"
				continue
			} else if (bg.hasApprentice && bg.appGridId == grid.id && i == bg.appYLoc && t == bg.appXLoc){
				row += "a"
				continue					
			} else if ((bg.monsterGridId == grid.id) && (i == bg.monsterYLoc) && (t == bg.monsterXLoc)){
				row += "M"
				continue
			} else {
				for k := 0; k < len(steps); k++ {
					if (steps[k].x == t && steps[k].y == i){
						row += "⌂"
						skip = true
						break
					}
				}
			}
			if (skip == false){
				row += grid.grid[i][t]
			}
		}
		
		cgid := id
		
		// draw status rows
		if i == 0 {
			row += "  -PATHS-"
		} else if i == 1 {
			if bg.directionValid(xloc, yloc, 7, cgid){
				row += "  NW"
			} else {
				row += "    "
			}
			if bg.directionValid(xloc, yloc, 0, cgid){
				row += " N "
			} else {
				row += "   "
			}		
			if bg.directionValid(xloc, yloc, 1, cgid){
				row += "NE"
			} else {
				row += "  "
			}			
		} else if i == 2{
			if bg.directionValid(xloc, yloc, 6, cgid){
				row += "  W  "
			} else {
				row += "     "
			}
			if bg.directionValid(xloc, yloc, 2, cgid){
				row += "   E"
			} else {
				row += "    "
			}			
		} else if i == 3{
			if bg.directionValid(xloc, yloc, 5, cgid){
				row += "  SW "
			} else {
				row += "     "
			}
			if bg.directionValid(xloc, yloc, 4, cgid){
				row += "S "
			} else {
				row += "  "
			}		
			if bg.directionValid(xloc, yloc, 3, cgid){
				row += "SE"
			} else {
				row += "  "
			}		
		} else if i == 5 {
			row += "  " + character.name + " Health: ["
			for hlth := 0; hlth <= character.maxhp; hlth++ {
				if (hlth > character.hp){
					row += "-"
				} else {
					row += "*"
				}
			}
			row += "]"
			
		} else if i == 6 {
			if (bg.hasApprentice){
				row += "  " + apprentice.name + " Health: ["
				for hlth := 0; hlth <= apprentice.maxhp; hlth++ {
					if (hlth > apprentice.hp){
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
				if (hlth > bg.monster.hp){
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

func (bg * BattleGrid) drawTestGrid2(testx, testy int, gd Grid){
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
			if (bg.charGridId == grid.id && i == bg.charYLoc && t == bg.charXLoc){
				row += "C"
				continue
			} else if (bg.hasApprentice && bg.appGridId == grid.id && i == bg.appYLoc && t == bg.appXLoc){
				row += "a"
				continue					
			} else if ((bg.monsterGridId == grid.id) && (i == bg.monsterYLoc) && (t == bg.monsterXLoc)){
				row += "M"
				continue
			} else if (testx == t && testy == i){
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
			if bg.directionValid(xloc, yloc, 7, cgid){
				row += "  NW"
			} else {
				row += "    "
			}
			if bg.directionValid(xloc, yloc, 0, cgid){
				row += " N "
			} else {
				row += "   "
			}		
			if bg.directionValid(xloc, yloc, 1, cgid){
				row += "NE"
			} else {
				row += "  "
			}			
		} else if i == 2{
			if bg.directionValid(xloc, yloc, 6, cgid){
				row += "  W  "
			} else {
				row += "     "
			}
			if bg.directionValid(xloc, yloc, 2, cgid){
				row += "   E"
			} else {
				row += "    "
			}			
		} else if i == 3{
			if bg.directionValid(xloc, yloc, 5, cgid){
				row += "  SW "
			} else {
				row += "     "
			}
			if bg.directionValid(xloc, yloc, 4, cgid){
				row += "S "
			} else {
				row += "  "
			}		
			if bg.directionValid(xloc, yloc, 3, cgid){
				row += "SE"
			} else {
				row += "  "
			}		
		} else if i == 5 {
			row += "  " + character.name + " Health: ["
			for hlth := 0; hlth <= character.maxhp; hlth++ {
				if (hlth > character.hp){
					row += "-"
				} else {
					row += "*"
				}
			}
			row += "]"
			
		} else if i == 6 {
			if (bg.hasApprentice){
				row += "  " + apprentice.name + " Health: ["
				for hlth := 0; hlth <= apprentice.maxhp; hlth++ {
					if (hlth > apprentice.hp){
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
				if (hlth > bg.monster.hp){
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