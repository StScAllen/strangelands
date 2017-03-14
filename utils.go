// utils.go

package main

import "fmt"
import "strings"

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


func makeDialogBox(actorName, msg string) ([]string){
	width := 60
//	height := 12
	
	elements := make([]string, 2)
	mid := strings.Repeat("─", width-(3+len(actorName)))
	elements[0] = "╔─" + actorName + mid + "╗"
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