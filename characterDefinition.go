// CharacterDefine.go
package main

import "fmt"
import "strings"

const VERSION = ".02a"

var skills = []string {"Puzzles", "Alchemy", "Haggle", "Instruction", "Spellcraft", "Research", "Politicking"}
var weaponSkills = []string {"Knife", "Sword", "Crossbow", "Polearm", "Axe", "Mace"}
					  					  
const NUM_SKILLS = 9

// new char attributes - Perception (view distance, searching)
//					   - Agility (actions per turn)
//					   - Strength
// 					   - Intellect
//					   - Charm
//					   - Guile	 

/* type Character2 struct {
	str, dex, con, intl, wis, cha int
	class string
	alignment string
	name string
	hp, maxhp int
	mana, maxmana int
	weight, maxweight int
	skillLevels [NUM_SKILLS]int
	gold int
	lvl int
	turns int
	numWeapons, numArmor, numItems int
	weapons [3]Weapon
	armors [3]Armor
	items [10]Item
} */

type Character struct {
	agi, str, per, intl, cha, gui int
	name string
	hp, maxhp int
	weight, maxweight int
	gold int
	lvl int
	exp int
	turns int
	spellbook Spellbook
}

// can have special items to increase moves
func (char * Character) getCharacterMoves() (int) {
	return char.agi
}

func getName() (string) {

	clearConsole()
	var flag bool = true
	rsp := ""
	
	for flag {
		fmt.Println("--- Choose a Character Name ---")
		fmt.Println("A name is nothing more than a tool. Don't forget that.")
		fmt.Println("")
		fmt.Println("Enter a name: ")
		
		fmt.Scanln(&rsp)

		if (len(strings.Trim(rsp, " ")) > 0){
			rsp2 := ""
			fmt.Println("")
			fmt.Println("(Y/N) Do you wish to use " + rsp + "?")
			fmt.Scanln(&rsp2)

			if (rsp2 == "y" || rsp2 == "Y"){
				flag = false
			}
		}
	}
	
	return rsp	  
}

func (c *Character) purchaseStats() {	
	var flag bool = true
	var points int = 8
	
	for flag {
		clearConsole()
		fmt.Println("--- Purchase Attributes ---")
		fmt.Println("Its not what you have done, but what you will do.")
		fmt.Println("")
		fmt.Printf("1. Perception: %v  (Vision, awareness, aim)\n", c.per)
		fmt.Printf("2. Strength:   %v  (Damage, encumberance, health)\n", c.str)
		fmt.Printf("3. Agility:    %v  (Movement, attacking, defense)\n", c.agi)
		fmt.Printf("4. Intellect:  %v  (Spells, skillcraft)\n", c.intl)		
		fmt.Printf("5. Charm:      %v  (Bartering, apprentice building)\n",  c.cha)
		fmt.Printf("6. Guile:      %v  (Experience, skillcraft)\n",  c.gui)		
		fmt.Println("")
		fmt.Println("7. Reset")		
		fmt.Println("8. Finished")			
		fmt.Println("--------------------")		
		fmt.Printf("Points remaining: %v \n", points)
		fmt.Println("Choose an attribute to add a point: ")		
		rsp := ""
		fmt.Scanln(&rsp)
		
		if (rsp == "8"){
			flag = false
		} else if (rsp == "7"){
			c.per = 3
			c.str = 3
			c.intl = 3
			c.agi = 3
			c.cha = 3
			c.gui = 3
			points = 8
		} else {		
			if points < 1 {
				rsp2 := ""
				fmt.Println("No points remain. Press enter to continue.")
				fmt.Scanln(&rsp2)
			} else {
			
				switch(rsp) {				
					case "1": 	
						c.per += 1
						points -= 1
					case "2": 	
						c.str += 1
						points -= 1
					case "3": 	
						c.agi += 1
						points -= 1
					case "4": 	
						c.intl += 1
						points -= 1
					case "5": 	
						c.cha += 1
						points -= 1
					case "6": 	
						c.gui += 1
						points -= 1											
				}
			
			}		
		}
	}
}

func (c *Character) getTotalStats() (int){	
	return c.str + c.agi + c.intl + c.gui + c.cha + c.per
}


func createCharacter() (Character) {
	var character Character
	
	character.name = getName()
	character.str = 3
	character.agi = 3
	character.intl = 3	
	character.gui = 3
	character.cha = 3
	character.per = 3
	
	character.purchaseStats()
	
	character.lvl = 1
	
	diff := 72 - character.getTotalStats()
	
	character.gold = 20 + diff
	
	character.hp = character.str
	character.maxhp = character.hp
	character.hp -= 1
	
	character.maxweight = character.str + 1
	character.weight = 0
	
	character.exp = 0

	return character
}

func (character *Character) printCharacter(pause int) {
	
	clearConsole()
	
	fmt.Printf("Name: %s    ", character.name)
	fmt.Printf("Level: %v    ", character.lvl)
	fmt.Printf("Exp: %v    ", character.exp)

	fmt.Println()
	fmt.Printf("Hp: %v / %v  ", character.hp, character.maxhp)
	fmt.Printf("Encumb: %v / %v  st", character.weight, character.maxweight)

	fmt.Println()	
	fmt.Printf("Per: %v \n", character.per)	
	fmt.Printf("Str: %v \n", character.str)
	fmt.Printf("Agi: %v \n", character.agi)
	fmt.Printf("Int: %v \n", character.intl)
	fmt.Printf("Cha: %v \n", character.cha)
	fmt.Printf("Gui: %v \n", character.gui)	
	
	fmt.Println()
	
	fmt.Printf("\nGold: %v", character.gold)
	
	if (pause > 0) {
		rsp := "n"
		fmt.Println("\nPress enter to continue.")
		fmt.Scanln(&rsp)
	}
}