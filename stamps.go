// stamps.go
// architectural stamps for game
package main

//import "fmt"		
//import "strings"

var PALLETE string = "▓ ▒ ░ ■ ║ ╣ ║ ╝ ╚ ╩ ╠ ╬ ═ ╦ ╔ ╗ │ └ ┘ ┌ ┐ ─ ┴ ├ ┤ ┬ ┼ ╞ ╫ ┴ ╨ ╥ ╖ ╒ ╓"
var PALLETE2 string ="↨ ↔ ▀ █ ▐ ▲ ► ▼ ◄ « ˄ « ˄ ∞ ⌂ ☼ ♥ ♪ ♫ ± Σ Φ ∩ † "		

var well = [][]string{
							{"┌", "─", "┐"},
							{"│", "O", "│"},
							{"└", "─", "┘"},
						   }

var open_crypt = [][]string{
							{"╔", "╦", "╗"},
							{"║", " ", "║"},
							{"║", " ", "║"},
							{"╚", "\\", "╝"},
						   }	
						 
var closed_crypt = [][]string{
							{"╔", "╦", "╗"},
							{"║", " ", "║"},
							{"║", " ", "║"},							
							{"╚", "═", "╝"},
						   }	
						   
var wagon = [][]string	{	   
							{" ", "┌", "┬", "═", "┬", "┐", " "},
							{"O", "┤", "│", " ", "│", "├", "O"},
							{" ", "│", "│", " ", "│", "│", " "},
							{"O", "┤", "└", "─", "┘", "├", "O"},
							{" ", "└", "─", "╥", "─", "┘", " "},							
							{" ", " ", " ", "╩", " ", " ", " "},				
						}
						
var house = [][]string 	{
							{"┌", "─","─", "─", "─",  "─", "─",  "─", "─", "─", "┐"},
							{"│", " "," ", "┌", "─",  "─", "─",  "┐", " ", " ", "│"},
							{"│", " "," ", "│", "\\", " ", "/",  "│", " ", " ", "│"},
							{"│", " "," ", "│", " ",  "■", " ",  "│", " ", " ", "│"},							
							{"│", " "," ", "│", "/",  " ", "\\", "│", " ", " ", "│"},		
							{"│", " "," ", "└", "─",  "─", "─",  "┘", " ", " ", "│"},							
							{"└", "─","─", "─", "─",  "/", "─",  "─", "─", "─", "┘"},
						}						
						
var tree = [][]string 	{
							{"\\", "|",  "/"},
							{"-", "·", "-"},
							{"/", "|", "\\"},	
						}
						
var pond1 = [][]string 	{
							{" ", " ", "░", "░", "░", " ", " "},
							{" ", "░", "░", "░", "░", "░", " "},	
							{"░", "░", "░", "░", "░", "░", "░"},
							{"░", "░", "░", "░", "░", "░", "░"},
							{" ", "░", "░", "░", "░", "░", " "},
							{" ", " ", "░", "░", "░", " ", " "},							
						}												

var pond2 = [][]string 	{
							{" ", " ", " ", "░", " ", " ", " "},
							{" ", " ", "░", "░", "░", " ", " "},	
							{" ", "░", "░", "░", "░", "░", " "},
							{" ", " ", "░", "░", "░", " ", " "},
							{" ", " ", " ", "░", " ", " ", " "},
							
						}						

var pond3 = [][]string 	{
							{" ", " ", " ", " ", "░", "░", "░", " ", " ", " ", " "},
							{" ", " ", "░", "░", "░", "░", "░", "░", "░", " ", " "},	
							{" ", "░", "░", "░", "░", "░", "░", "░", "░", "░", " "},
							{"░", "░", "░", "░", "░", "░", "░", "░", "░", "░", "░"},
							{"░", "░", "░", "░", "░", "░", "░", "░", "░", "░", "░"},
							{" ", "░", "░", "░", "░", "░", "░", "░", "░", "░", " "},							
							{" ", " ", "░", "░", "░", "░", "░", "░", "░", " ", " "},
							{" ", " ", " ", " ", "░", "░", "░", " ", " ", " ", " "},							
						}	
						
func (bg * BattleGrid) addStreamVertical(xloc int, gidx int){
	
	var die Die 
	
	var bridgeY = die.rollxdx(4, 13)
	
	for k := 0; k < len(bg.allGrids[gidx].grid); k++ {
		if (k == bridgeY){
			bg.allGrids[gidx].grid[k][xloc-1] = "┌"		
			bg.allGrids[gidx].grid[k][xloc] = "─"
			bg.allGrids[gidx].grid[k][xloc+1] = "─"	
			bg.allGrids[gidx].grid[k][xloc+2] = "┐"
		} else if (k+2 == bridgeY){
			bg.allGrids[gidx].grid[k][xloc-1] = "└"		
			bg.allGrids[gidx].grid[k][xloc] = "─"
			bg.allGrids[gidx].grid[k][xloc+1] = "─"	
			bg.allGrids[gidx].grid[k][xloc+2] = "┘"
		} else if (k+1 == bridgeY){
			bg.allGrids[gidx].grid[k][xloc] = " "
			bg.allGrids[gidx].grid[k][xloc+1] = " "	
		} else {
			bg.allGrids[gidx].grid[k][xloc] = "░"
			bg.allGrids[gidx].grid[k][xloc+1] = "░"		
		}
	}
}		

func (bg * BattleGrid) addStreamHorizontal(yloc int, gidx int){
	
	var die Die 
	
	var bridgeY = die.rollxdx(4, 13)
	
	for k := 0; k < len(bg.allGrids[gidx].grid[0]); k++ {
		if (k == bridgeY){
			bg.allGrids[gidx].grid[yloc-1][k] = "┌"		
			bg.allGrids[gidx].grid[yloc][k] = "│"
			bg.allGrids[gidx].grid[yloc+1][k] = "│"	
			bg.allGrids[gidx].grid[yloc+2][k] = "└"
		} else if (k+2 == bridgeY){
			bg.allGrids[gidx].grid[yloc-1][k] = "┐"		
			bg.allGrids[gidx].grid[yloc][k] = "│"
			bg.allGrids[gidx].grid[yloc+1][k] = "│"	
			bg.allGrids[gidx].grid[yloc+2][k] = "┘"
		} else if (k+1 == bridgeY){
			bg.allGrids[gidx].grid[yloc][k] = " "
			bg.allGrids[gidx].grid[yloc+1][k] = " "	
		} else {
			bg.allGrids[gidx].grid[yloc][k] = "░"
			bg.allGrids[gidx].grid[yloc+1][k] = "░"		
		}
	}
}					
						
func (bg * BattleGrid) setRandomStamp(maxx int, maxy int, gidx int){
	
	var die Die 
	var skip bool	
	roll := die.rollxdx(1, 8)
	
	var stamp [][]string
	
	skip = false
	if (roll == 1){
		stamp = open_crypt
	} else if (roll == 2){
		stamp = closed_crypt
	} else if (roll == 3){
		stamp = wagon
	} else if (roll == 4){
		stamp = tree	
	} else if (roll == 5){
		stamp = pond1	
	} else if (roll == 6){
		stamp = well	
	} else if (roll == 7){
		bg.addStreamVertical(die.rollxdx(4, 28), gidx)	
		skip = true
	} else {
		bg.addStreamHorizontal(die.rollxdx(4, 12), gidx)
		skip = true
	}
		
//	stamp = pond3

	if (skip == false){
		xs := die.rollxdx(2, (maxx-len(stamp))-1)
		ys := die.rollxdx(2, (maxy-len(stamp[0]))-1)	
		
		bg.setStamp(stamp, xs, ys, gidx)	
	}

}						
						
func (bg * BattleGrid) setStamp(stamp [][]string, x int, y int, gidx int){
	for k := 0; k < len(stamp); k++ {
		for t := 0; t < len(stamp[k]); t++ {
			if (k+x < len(bg.allGrids[gidx].grid)){
				if t+y < len(bg.allGrids[gidx].grid[0]){
					bg.allGrids[gidx].grid[k+x][t+y] = stamp[k][t]
				}
			}
		}
	}
}