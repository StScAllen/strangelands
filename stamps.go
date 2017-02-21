// stamps.go
// architectural stamps for game
package main

//import "fmt"		
//import "strings"

var PALLETE string = "▓ ▒ ░ ■ ║ ╣ ║ ╝ ╚ ╩ ╠ ╬ ═ ╦ ╔ ╗ │ └ ┘ ┌ ┐ ─ ┴ ├ ┤ ┬ ┼ ╞ ╫ ┴ ╨ ╥ ╖ ╒ ╓"
var PALLETE2 string ="↨ ↔ ▀ █ ▐ ▲ ► ▼ ◄ « ˄ « ˄ ∞ ⌂ ☼ ♥ ♪ ♫ ± Σ Φ ∩ †"		

var well = [][]string{
							{" ", "-", " "},
							{"(", "O", ")"},
							{" ", "-", " "},
						   }

var open_crypt = [][]string{
							{"╔", "╦", "╗"},
							{"║", " ", "║"},
							{"╚", "\\", "╝"},
						   }	
						 
var closed_crypt = [][]string{
							{"╔", "╦", "╗"},
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
							{"/", "|",  "\\"},
							{"/", "|", "\\"},
							{"/", "|", "\\"},	
						}

func (bg * BattleGrid) setRandomStamp(maxx int, maxy int, gidx int){
	
	var die Die 
		
	roll := die.rollxdx(1, 4)
	
	var stamp [][]string
	
	if (roll == 1){
		stamp = open_crypt
	} else if (roll == 2){
		stamp = closed_crypt
	} else if (roll == 3){
		stamp = wagon
	} else {
		stamp = well
	}
	
	stamp = tree

	xs := die.rollxdx(1, (maxx-len(stamp))-1)
	ys := die.rollxdx(1, (maxy-len(stamp[0]))-1)	
	
	bg.setStamp(stamp, xs, ys, gidx)
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