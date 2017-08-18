// utils.go

package main

import "fmt"
import "strings"
import "os"
import "os/exec"
import "math"

const DIALOG_RIGHT = 0
const DIALOG_LEFT = 1

// this needs to be command prompt generic
func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showPause(messge string) {
	fmt.Println(messge)
	rsp := ""
	fmt.Scanln(&rsp)
}

func getResponse(question string) string{
	fmt.Println(question)
	rsp := ""
	fmt.Scanln(&rsp)
	return rsp
}

func debugPause(messge string) {
	if DEBUG_ON {
		fmt.Println(messge)
		rsp := ""
		fmt.Scanln(&rsp)	
	}
}

func replaceAtIndex(str string, replacement string, index int) string {
	return str[:index] + replacement + str[index+1:]
}

func replaceAtIndex2(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

func getSigned(val int) string {
	if val > 0 {
		return fmt.Sprintf("+%v", val)
	}

	return fmt.Sprintf("%v", val)
}

func getSignedFloat32(val float32) string {
	if val > 0 {
		return fmt.Sprintf("+%f", val)
	}

	return fmt.Sprintf("%f", val)
}

func getCityBlockDistance(sx, sy, ex, ey int) int {
	return iAbsDiff(sx, ex) + iAbsDiff(sy, ey)
}

func getCrowDistance(sx, sy, ex, ey int) int {

	if sx == ex {
		return iAbsDiff(sy, ey)
	} else if sy == ey {
		return iAbsDiff(sx, ex)
	}

	xDist := iAbsDiff(sx, ex)
	yDist := iAbsDiff(sy, ey)

	if xDist == yDist {
		return xDist
	}

	return (iAbsDiff(xDist, yDist) + xDist)
}

// gives absolute value between int values
func iAbsDiff(x1, x2 int) int {
	if x1 > x2 {
		return x1 - x2
	} else {
		return x2 - x1
	}
}

func iAbsVal(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func getCardinalStringFromRelativePosition(relX, relY int, shortCard bool) string {
	retVal := "BAD!"

	if relX == 0 && relY > 0 {
		if shortCard {
			retVal = "N"
		} else {
			retVal = "North"
		}
	} else if relX == 0 && relY < 0 {
		if shortCard {
			retVal = "S"
		} else {
			retVal = "South"
		}
	} else if relX == 0 && relY == 0 {
		if shortCard {
			retVal = "B"
		} else {
			retVal = "Beneath"
		}
	} else if relX > 0 && relY > 0 {
		if shortCard {
			retVal = "NW"
		} else {
			retVal = "NorthWest"
		}
	} else if relX > 0 && relY < 0 {
		if shortCard {
			retVal = "SW"
		} else {
			retVal = "SouthWest"
		}
	} else if relX > 0 && relY == 0 {
		if shortCard {
			retVal = "W"
		} else {
			retVal = "West"
		}
	} else if relX < 0 && relY == 0 {
		if shortCard {
			retVal = "E"
		} else {
			retVal = "East"
		}
	} else if relX < 0 && relY > 0 {
		if shortCard {
			retVal = "NE"
		} else {
			retVal = "NorthEast"
		}
	} else if relX < 0 && relY < 0 {
		if shortCard {
			retVal = "SE"
		} else {
			retVal = "SouthEast"
		}
	}

	return retVal
}

func convertPoundsToStone(lbs int) string {

	if lbs < 14 {
		return fmt.Sprintf("0 stone %v", lbs)
	}

	stone := 0
	for ; lbs > 14; lbs -= 14 {
		stone++
	}

	return fmt.Sprintf("%v stone, %v", stone, lbs)
}

func convertStoneToPounds(stone int) int {
	return stone * 14
}

func getVillageDistance(idx int) int {
	currX, currY := 0, 0
	destX, destY := 0, 0

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

	daysTravel := int((distX + distY) / 6)

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

func packSpaceStringCenter(str string, digits int) string {
	padAmount := float64(digits - len(str))
	var half float64 
	var remain float64

	half = math.Floor(padAmount / float64(2))	
	remain = float64(padAmount / float64(2)) - half

	for k := float64(0); k < half; k++ {
		str = " " + str
	}
	
	for k := float64(0); k < half; k++ {
		str += " "
	}
	
	if remain > 0 {
		str += " "
	}
	
	return str
}

func packSpaceStringWithToken(str string, digits int, token string) string {
	for len(str) < digits {
		str += token
	}

	return str
}

func makeDialogString(str string) string {
	str = "\"" + str + "\""

	return str
}

func getOnOffString(flag bool) (string) {
	if flag {
		return "On"
	} 
	
	return "Off"
}

func makeDialogBox(actorName, msg string, side int) []string {
	width := 60
	//	height := 12

	elements := make([]string, 2)
	mid := strings.Repeat("─", width-(3+len(actorName)))
	if side == DIALOG_LEFT {
		elements[0] = "╔─" + actorName + mid + "╗"
	} else {
		elements[0] = "╔" + mid + actorName + "-╗"
	}

	elements[1] = "│" + packSpaceString(" ", 58) + "│"
	if len(msg) < width {
		row := packSpaceString(msg, 58)
		elements = append(elements, "│ "+row+" │")
	} else {
		charProcessed := 0
		bits := strings.Split(msg, " ")
		fmt.Println("bits are: ", bits)
		lastbit := 0
		for charProcessed < len(msg) {
			row := "│ "

			for k := lastbit; k < len(bits); k++ {
				if (len(bits[k]) + len(row)) < (width - 2) {
					row = row + bits[k] + " "
					fmt.Println("Adding bit: ", bits[k])
					fmt.Println("Now: " + row)
					fmt.Println("Total Chars: ", len(row))
					if k == (len(bits) - 1) {
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
