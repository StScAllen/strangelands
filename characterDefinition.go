// CharacterDefine.go
package main

import "fmt"
import "strings"

var skills = []string{"Puzzles", "Alchemy", "Haggle", "Instruction", "Spellcraft", "Research", "Politicking", "Chirurgery"}
var weaponSkills = []string{"Knife", "Sword", "Crossbow", "Polearm", "Axe", "Mace"}

const NUM_SKILLS = 9
const LEFT = 0
const RIGHT = 1

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
	name                          string // max name length for FORMATTING reasons is 23 characters!
	hp, maxhp                     int
	soul, maxsoul                 int // soul is both a spiritual hp and a tool to craft/power artefacts
	weight, maxweight             int
	gold                          int
	lvl                           int
	exp                           int
	turns                         int
	handSlots                     [2]Item
	armorSlots                    [8]Item
	inventory                     []Item
	spellbook                     Spellbook
	villageIndex				  int
}

// can have special items to increase moves
func (char *Character) getCharacterMoves() int {
	return char.agi
}

func (char *Character) getTotalAttackAdjustment(handSlot int) (int) {
	adj := 0

	adj += char.handSlots[handSlot].accuracy

	return adj
}

func (char *Character) getTotalDefenseAdjustment(handSlot int) (int) {
	adj := 0

	adj += char.handSlots[handSlot].defense

	return adj
}

func (char *Character) getWeaponRange() int {
	hand1 := char.handSlots[0]
	hand2 := char.handSlots[1]
	itmRange := -1
	
	if hand1.id != -1 {
		itmRange = hand1.wRange
	} 
	
	if hand2.id != -1 {
		if hand2.wRange > itmRange {
			itmRange = hand2.wRange
		}
	}
	
	fmt.Println("Weapon range is ", itmRange)
	
	return itmRange
}

func (char *Character) recalcCharacterWeight() {
	weight := 0

	weight += char.handSlots[0].weight
	weight += char.handSlots[1].weight

	for k := 0; k < len(char.armorSlots); k++ {
		weight += char.armorSlots[k].weight
	}

	for k := 0; k < len(character.inventory); k++ {
		weight += char.inventory[k].weight
	}

	char.weight = weight
}

func (char *Character) setClearInventory() {

	char.handSlots[0] = getEmptyItem()
	char.handSlots[1] = getEmptyItem()

	for k := 0; k < len(char.armorSlots); k++ {
		char.armorSlots[k] = getEmptyItem()
	}

}

func (char *Character) giveCharacterItem(item Item) bool {
	equipped := false

	char.recalcCharacterWeight()

	// do an encumberance check
	if char.weight+item.weight > char.maxweight {
		return false
	}

	if item.equip != EQUIP_NONE {
		if item.equip < EQUIP_HAND {
			if item.typeCode == ITEM_TYPE_ARMOR {
				item.durability -= 1
			}
		
			if char.armorSlots[item.equip].id == -1 {
				char.armorSlots[item.equip] = item
				equipped = true
			}
		} else if item.equip == EQUIP_HAND {
			if item.hands == 1 {
				if char.handSlots[LEFT].id == -1 {
					char.handSlots[LEFT] = item
					equipped = true
				} else if char.handSlots[RIGHT].id == -1 {
					char.handSlots[RIGHT] = item
					equipped = true
				}
			} else if item.hands == 2 {
				if char.handSlots[LEFT].id == -1 && char.handSlots[RIGHT].id == -1 {
					char.handSlots[RIGHT] = item
					char.handSlots[LEFT] = item
					equipped = true
				}
			}
		}
	}

	if !equipped {
		char.inventory = append(char.inventory, item)
		equipped = true
	}

	char.recalcCharacterWeight()

	return equipped
}

func getName() string {

	clearConsole()
	var flag bool = true
	rsp := ""

	for flag {
		fmt.Println("--- Choose a Character Name ---")
		fmt.Println("A name is nothing more than a tool. Don't forget that.")
		fmt.Println("")
		fmt.Println("Enter a name: ")

		fmt.Scanln(&rsp)

		if len(strings.Trim(rsp, " ")) > 0 {
			rsp2 := ""
			fmt.Println("")
			fmt.Println("(Y/N) Do you wish to use " + rsp + "?")
			fmt.Scanln(&rsp2)

			if rsp2 == "y" || rsp2 == "Y" {
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
		fmt.Printf("5. Charm:      %v  (Bartering, apprentice building, soul)\n", c.cha)
		fmt.Printf("6. Guile:      %v  (Experience, skillcraft, soul)\n", c.gui)
		fmt.Println("")
		fmt.Println("7. Minutiae (Help)")
		fmt.Println("8. Reset")
		fmt.Println("9. Finished")
		fmt.Println("--------------------")
		fmt.Printf("Points remaining: %v \n", points)
		fmt.Println("Choose an attribute to add a point: ")
		rsp := ""
		fmt.Scanln(&rsp)

		if rsp == "9" {
			flag = false
		} else if rsp == "8" {
			c.per = 3
			c.str = 3
			c.intl = 3
			c.agi = 3
			c.cha = 3
			c.gui = 3
			points = 8
		} else if rsp == "7" {
			showAttributesMinutiae()

		} else {
			if points < 1 {
				rsp2 := ""
				fmt.Println("No points remain. Press enter to continue.")
				fmt.Scanln(&rsp2)
			} else {

				switch rsp {
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

func (c *Character) getTotalStats() int {
	return c.str + c.agi + c.intl + c.gui + c.cha + c.per
}

func createCharacter() Character {
	var character Character

	character.setClearInventory()

	character.name = getName()
	character.str = 3
	character.agi = 3
	character.intl = 3
	character.gui = 3
	character.cha = 3
	character.per = 3

	character.purchaseStats()

	character.lvl = 1

	character.gold = 32

	character.hp = character.str
	character.maxhp = character.hp

	character.soul = character.gui + character.cha
	character.maxsoul = character.soul

	character.maxweight = character.str * 10
	character.weight = 0

	character.handSlots[0] = getEmptyItem()
	character.handSlots[1] = getEmptyItem()

	character.exp = 0
	character.villageIndex = 0

	character.save()
	
	return character
}

func (char *Character) showStatus() {
	clearConsole()

	fmt.Println("[Health Status]             ")
	fmt.Println("            ####      ▲       ")
	fmt.Println("            ####     ◄ ►        ")
	fmt.Println("             ##       ▼       ")
	fmt.Println("         ##########         ")
	fmt.Println("        ############         ")
	fmt.Println("        ## ###### ##         ")
	fmt.Println("        ## ###### ##        ")
	fmt.Println("           ######         ")
	fmt.Println("           ##  ##            ")
	fmt.Println("           ##  ##            ")
	fmt.Println("           ##  ##            ")
	fmt.Println("          ###  ###           ")

	rsp := ""
	fmt.Println("\nPress enter to continue.")
	fmt.Scanln(&rsp)
}

func (character *Character) printCharacter(pause int) {

	clearConsole()

	fmt.Printf("Name: %s    ", character.name)
	fmt.Printf("Level: %v    ", character.lvl)
	fmt.Printf("Exp: %v    ", character.exp)

	fmt.Println()
	fmt.Printf("Hp: %v / %v  ", character.hp, character.maxhp)
	fmt.Printf("Soul: %v / %v  ", character.soul, character.maxsoul)
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

	if pause > 0 {
		rsp := "n"
		fmt.Println("\nPress enter to continue.")
		fmt.Scanln(&rsp)
	}
}

func (character *Character) showInventory() {
	clearConsole()

	seg1 := ""
	seg2 := ""

	fmt.Printf("Encumb: %v / %v  (stone) \n", character.weight, character.maxweight)

	fmt.Println("")
	fmt.Println("--Hands--")

	seg1 = packSpaceString("Left Hand: ", 14) + character.handSlots[LEFT].getInvDisplayString()
	seg2 = packSpaceString("Right Hand: ", 14) + character.handSlots[RIGHT].getInvDisplayString()
	fmt.Println(seg1)
	fmt.Println(seg2)
	
	fmt.Println("")
	fmt.Println("--Armor--")
	seg1 = packSpaceString("Head: ", 14) + character.armorSlots[EQUIP_HEAD].getInvDisplayString()
	seg2 = packSpaceString("Neck: ", 14) + character.armorSlots[EQUIP_NECK].getInvDisplayString()
	fmt.Println(seg1)
	fmt.Println(seg2)
	seg1 = packSpaceString("Chest: ", 14) + character.armorSlots[EQUIP_CHEST].getInvDisplayString()
	seg2 = packSpaceString("Arms: ", 14) + character.armorSlots[EQUIP_ARMS].getInvDisplayString()
	fmt.Println(seg1)
	fmt.Println(seg2)
	seg1 = packSpaceString("Legs: ", 14) + character.armorSlots[EQUIP_LEG].getInvDisplayString()
	seg2 = packSpaceString("Feet: ", 14) + character.armorSlots[EQUIP_FEET].getInvDisplayString()
	fmt.Println(seg1)
	fmt.Println(seg2)
	seg1 = packSpaceString("Cloak: ", 14) + character.armorSlots[EQUIP_CLOAK].getInvDisplayString()
	seg2 = packSpaceString("Ring: ", 14) + character.armorSlots[EQUIP_RING].getInvDisplayString()
	fmt.Println(seg1)
	fmt.Println(seg2)

	fmt.Println("")
	fmt.Println("--Bags--")
	count := 0
	for k := 0; k < len(character.inventory); k++ {
		fmt.Printf("%s", packSpaceString(character.inventory[k].name, 23))
		count++
		if count == 3 {
			count = 0
			fmt.Printf("\n")
		}
	}

	fmt.Println("\nPress enter to continue.")
	rsp := ""
	fmt.Scanln(&rsp)
}
