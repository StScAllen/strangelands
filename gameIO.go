// gameio.go

package main

import "io/ioutil"
import "strings"
import "strconv"
import "fmt"
import "os"


func (c *Character) save() {

	var filename string
	
	filename = "save.txt"
	
	file, err := os.Create(filename)

	var saveString string 
	
	saveString += c.name + ","
	saveString += fmt.Sprintf("%v,", c.lvl)
	
	saveString += fmt.Sprintf("%v,", c.str)
	saveString += fmt.Sprintf("%v,", c.per)
	saveString += fmt.Sprintf("%v,", c.agi)
	saveString += fmt.Sprintf("%v,", c.intl)
	saveString += fmt.Sprintf("%v,", c.cha)
	saveString += fmt.Sprintf("%v,", c.gui)
	
	saveString += fmt.Sprintf("%v,", c.hp)
	saveString += fmt.Sprintf("%v,", c.maxhp)
	
	saveString += fmt.Sprintf("%v,", c.weight)
	saveString += fmt.Sprintf("%v,", c.maxweight)
	
	saveString += fmt.Sprintf("%v,", c.gold)
	
	if (err == nil){
		defer file.Close()
		
		file.WriteString(saveString)	
		fmt.Println("Game Saved!")
	}
}

func loadGame() (Character){	
	var char Character
		
	data, err := ioutil.ReadFile("save.txt")
	if (err == nil){
		charData := fmt.Sprintf("%s", data)
		
		if (len(charData) > 0){
		
			var bits []string = strings.Split(charData, ",")
			
			fmt.Println(bits)
			
			char.name = bits[0]
			char.lvl, err = strconv.Atoi(bits[1])
			
			char.str, err = strconv.Atoi(bits[2])
			char.per, err = strconv.Atoi(bits[3])
			char.agi, err = strconv.Atoi(bits[4])
			char.intl, err = strconv.Atoi(bits[5])
			char.cha, err = strconv.Atoi(bits[6])
			char.gui, err = strconv.Atoi(bits[7])
			
			char.hp, err = strconv.Atoi(bits[8])
			char.maxhp, err = strconv.Atoi(bits[9])
			
			char.weight, err = strconv.Atoi(bits[10])
			char.maxweight, err = strconv.Atoi(bits[11])
			
			char.gold, err = strconv.Atoi(bits[12])

		}
		
		fmt.Println("-----")

		fmt.Printf("\n%s", data)
		fmt.Println("Game Loaded! ")
		
	}
	
	return char
}