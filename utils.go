// utils.go

package main

import "fmt"
import "strings"

const DIALOG_RIGHT = 0
const DIALOG_LEFT = 1

func replaceAtIndex(str string, replacement string, index int) string {
    return str[:index] + replacement + str[index+1:]
}

func replaceAtIndex2(str string, replacement rune, index int) string {
    out := []rune(str)
    out[index] = replacement
    return string(out)
}

func getSigned(val int) (string) {
	if val > 0 {
		return fmt.Sprintf("+%v", val)
	} 

	return fmt.Sprintf("%v", val)
}

func convertPoundsToStone(lbs int) (string) {

	if lbs < 14 {
		return fmt.Sprintf("0 stone %v", lbs)
	}
	
	stone := 0
	for ; lbs > 14; lbs -= 14 {
		stone++
	}
	
	return fmt.Sprintf("%v stone, %v", stone, lbs)
}

func convertStoneToPounds(stone int) (int) {
	return stone * 14
}

func getVillageDistance(idx int) (int){
	currX, currY := 0,0
	destX, destY := 0,0
	
	if idx == 99 {
		destX, destY = keep.mapX, keep.mapY
	} else {
		destX, destY = villages[idx].mapX, villages[idx].mapY	
	}
	
	if character.villageIndex == 99 {
		currX, currY = keep.mapX, keep.mapY
	} else {
		currX, currY = villages[character.villageIndex].mapX, villages[character.villageIndex].mapY
	}

	distX := iAbsDiff(currX, destX)
	distY := iAbsDiff(currY, destY)
	
	if distX == 0 && distY == 0 {
		return 0
	}
	
	daysTravel := int((distX+distY) / 6)
	
	if daysTravel < 1 {
		daysTravel = 1
	}
	
	return daysTravel
}

func packSpace(num int, digits int) string {
	ret := fmt.Sprintf("%v", num)

	for len(ret) < digits {
		ret += " "
	}

	return ret
}

func packSpaceString(str string, digits int) string {
	for len(str) < digits {
		str += " "
	}

	return str
}

func makeDialogString(str string) (string) {
	str = "\"" + str + "\""
	
	return str
}


func makeDialogBox(actorName, msg string, side int) ([]string){
	width := 60
//	height := 12
	
	elements := make([]string, 2)
	mid := strings.Repeat("─", width-(3+len(actorName)))
	if side == DIALOG_LEFT{
		elements[0] = "╔─" + actorName + mid + "╗"	
	} else {
		elements[0] = "╔" + mid + actorName + "-╗"
	}

	elements[1] = "│" + packSpaceString(" ", 58) + "│"		
	if len(msg) < width {
		row := packSpaceString(msg, 58)
		elements = append(elements, "│ " + row + " │")
	} else {
		charProcessed := 0
		bits := strings.Split(msg, " ")
		fmt.Println("bits are: ", bits)
		lastbit := 0
		for charProcessed < len(msg) {
			row := "│ "
			
			for k:= lastbit; k < len(bits); k++ {
				if (len(bits[k]) + len(row)) < (width-2) {
					row = row + bits[k] + " "
					fmt.Println("Adding bit: ", bits[k])
					fmt.Println("Now: " + row)
					fmt.Println("Total Chars: ", len(row))
					if (k == (len(bits)-1)){
						fmt.Println("End found!")
						row = packSpaceString(row, 60)
						row += " │"
						elements = append(elements, row)
						charProcessed += len(row)
						break						
					}
				} else {
					row = packSpaceString(row, 60)				
					charProcessed += (len(row) - 2)
					row += " │"
					elements = append(elements, row)
					lastbit = k
					break
				}
				lastbit++
			}
		}
	}
	
	endcap := "╚" + strings.Repeat("─", 58) + "╝"	
	elements = append(elements, endcap)
	
	rsp := ""
	fmt.Printf("\nPress any key to continue.")
	fmt.Scanln(&rsp)
	
	return elements
}