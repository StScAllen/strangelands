// CharacterDefine.go
package main

import "fmt"
import "strings"
import "strconv"

var skills = []string{"Puzzles", "Politicking", "Investigation", "Alchemy", "Craft", "Spellcraft", "Chirurgery"}
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
	crowns int
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
	crowns                        int
	lvl                           int
	exp                           int
	turns                         int
	turnDefense 				  int  // how many turns were used as defense
	skills						  [7]int
	alive						  bool
	handSlots                     [2]Item
	armorSlots                    [9]Item
	wounds                        []Wound
	inventory                     []Item
	spellbook                     Spellbook
	villageIndex                  int
}

func (char * Character) getPowerBalance() float32 {
	var balance float32
	balance = 0.0
	
	if char.hp < 1 {
		return balance
	}
	
	balance += (float32)(char.hp * 1.0)
	
	balance += (float32)(char.getTotalStats() / 6)
	
	return balance
}

// can have special items to increase moves
func (char *Character) getCharacterMoves() int {
	totalMoves := char.agi
	
	// TODO: add equipment, other buffs
	
	return totalMoves
}

func (char *Character) giveSoul(amt int) {
	if (char.soul+amt) <= char.maxsoul {
		char.soul += amt
	} else {
		char.soul = char.maxsoul
	}
}

func (char *Character) getTotalAttackAdjustment(handSlot int) int {
	adj := 0

	adj += char.handSlots[handSlot].accuracy

	return adj
}

func (char *Character) getTotalDefenseAdjustment() int {
	adj := 0

	// TODO: this needs to calc total defense from all sources

	return adj
}

func (char *Character) getResistanceAt(charBodyIndex int) int {
	equipIndex := HUMAN_TARGETS[charBodyIndex]

	if char.armorSlots[equipIndex].id != -1 {
		return char.armorSlots[equipIndex].resistance
	}

	return 2
}

// hits - number of hits to assess, charBodyIndex - equip constant
func (char *Character) soakHits(hits, charBodyIndex int) {
	char.armorSlots[charBodyIndex].shields -= hits

	if char.armorSlots[charBodyIndex].shields < 1 {
		// destroy armor
		showPause(char.armorSlots[charBodyIndex].name + " is destroyed!")
		char.armorSlots[charBodyIndex] = getEmptyItem()
	}
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

	char.inventory = make([]Item, 0, 0)
}

// finds and remove an item from the characters inventory
func (char *Character) removeItemFromCharacter(item Item) {
	if item.id < 1 {
		return
	}

	if item.equip == EQUIP_HAND {
		showPause("Removing from hand...")
		if char.handSlots[LEFT].id == item.id {
			char.handSlots[LEFT] = getEmptyItem()
		} else if char.handSlots[RIGHT].id == item.id {
			char.handSlots[RIGHT] = getEmptyItem()	
		} else {
			for k:= 0; k < len(char.inventory); k++ {
				if char.inventory[k].id == item.id {
					if len(char.inventory) > 1 {
						char.inventory = append(char.inventory[:k], char.inventory[k+1:]...)
					} else {
						char.inventory = make([]Item, 0, 0)
					}
					break
				}
			}
		}
	} else {
		if char.armorSlots[item.equip].id == item.id {
			char.armorSlots[item.equip] = getEmptyItem()
		} else {
			for k:= 0; k < len(char.inventory); k++ {
				if char.inventory[k].id == item.id {
					if len(char.inventory) > 1 {
						char.inventory = append(char.inventory[:k], char.inventory[k+1:]...)
					} else {
						char.inventory = make([]Item, 0, 0)
					}
					break
				}
			}		
		}
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

func (c *Character) getAllAvailableItemsForSlot(slot int) ([]Item){
	availItems := make([]Item, 0, 0,)
	
	for k := range c.inventory {
		if c.inventory[k].equip == slot || c.inventory[k].equip == EQUIP_ANY {
			availItems = append(availItems, c.inventory[k])
		}
	}
	
	if (slot == EQUIP_HAND) {
		if c.handSlots[0].id > 0 {
			availItems = append(availItems, c.handSlots[0])
		}
		if c.handSlots[1].id > 0 {
			availItems = append(availItems, c.handSlots[1])
		}		
	} else {
		if c.armorSlots[slot].id > 0 {
			availItems = append(availItems, c.armorSlots[slot])
		} 	
	}

	return availItems
}

func (c *Character) chooseSkills() {
	var flag bool = true
	var points int = 3

	for j := 0; j < len(skills); j++ {
		c.skills[j] = 1
	}
	
	for flag {
		clearConsole()
		fmt.Println("--- Purchase Skills ---")
		fmt.Println("XXX A saying here is something Steve must write!")
		fmt.Println("")
		for k := 0; k < len(skills); k++ {
			bit := packSpaceString(fmt.Sprintf("%v. %s: ", k+1, skills[k]), 24)	
			bit2 := fmt.Sprintf("%v", c.skills[k])
			val := bit + bit2
			fmt.Println(val)
		}

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
			for j := 0; j < len(skills); j++ {
				c.skills[j] = 1
			}
			points = 3
		} else if rsp == "7" {
			showSkillsMinutiae()
		} else {
			if points < 1 {
				rsp2 := ""
				fmt.Println("No points remain. Press enter to return.")
				fmt.Scanln(&rsp2)
			} else {

				switch rsp {
				case "1":
					c.skills[0]++ 
					points -= 1
				case "2":
					c.skills[1]++ 
					points -= 1
				case "3":
					c.skills[2]++ 
					points -= 1
				case "4":
					c.skills[3]++ 
					points -= 1
				case "5":
					c.skills[4]++ 
					points -= 1
				case "6":
					c.skills[5]++ 
					points -= 1
				case "7":
					c.skills[6]++ 
					points -= 1
				}

			}
		}
	}
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
				fmt.Println("No points remain. Press enter to return.")
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
	character.chooseSkills()

	character.lvl = 1

	character.crowns = 32

	character.hp = character.str
	character.maxhp = character.hp

	character.soul = character.gui + character.cha
	character.maxsoul = character.soul

	character.maxweight = character.str * 10
	character.weight = 0

	character.handSlots[0] = getEmptyItem()
	character.handSlots[1] = getEmptyItem()

	character.wounds = make([]Wound, 0, 0)
	character.alive = true
	
	character.exp = 0
	character.villageIndex = 0

	return character
}

func (char *Character) getArmorStringForSlot(slot int) string {
	slotString := ""

	if char.armorSlots[slot].name != "empty" {
		slotString = char.armorSlots[slot].getStatusDisplayStringArmor()
	} else {
		slotString = packSpaceString("none", 24) + "[_]"
	}

	return slotString
}

func (char *Character) getWoundsString(slot int) string {
	woundString := ""

	if len(char.wounds) > 0 {
		for k := 0; k < len(char.wounds); k++ {
			if char.wounds[k].location == slot {
				woundString += "(" + char.wounds[k].name + ")"
			}
		}
	} else {
		woundString = "None"
	}

	return woundString
}

func (char *Character) showStatus() {

	rsp := ""
	clearConsole()

	fmt.Println("              [Armor] ")
	fmt.Println(" ")
	fmt.Println("      ┌─────── Head:  ")
	fmt.Println("     ##         " + char.getArmorStringForSlot(EQUIP_HEAD))
	fmt.Println("    ####")
	fmt.Println("    #### ┌──── Chest: ")
	fmt.Println("     ││         " + char.getArmorStringForSlot(EQUIP_CHEST))
	fmt.Println(" ##########  ")
	fmt.Println("############ ")
	fmt.Println("## ###### ## ─ Arms:  ")
	fmt.Println("## ###### ##    " + char.getArmorStringForSlot(EQUIP_ARMS))
	fmt.Println("└┘ ###### └┘   ")
	fmt.Println("   ######    ")
	fmt.Println("   ##  ## ──── Legs:  ")
	fmt.Println("   ##  ##       " + char.getArmorStringForSlot(EQUIP_LEG))
	fmt.Println("   ││  ││      ")
	fmt.Println("   ##  ##      ")
	fmt.Println("  ###  ### ─── Feet:  ")
	fmt.Println("                " + char.getArmorStringForSlot(EQUIP_FEET))

	fmt.Println("\nPress enter to continue.")
	fmt.Scanln(&rsp)

	clearConsole()

	fmt.Println("              [Wounds] ")
	fmt.Println(" ")
	fmt.Println("      ┌─────── Head:  ")
	fmt.Println("     ##         " + char.getWoundsString(EQUIP_HEAD))
	fmt.Println("    ####")
	fmt.Println("    #### ┌──── Chest: ")
	fmt.Println("     ││         " + char.getWoundsString(EQUIP_CHEST))
	fmt.Println(" ##########  ")
	fmt.Println("############ ")
	fmt.Println("## ###### ## ─ Arms:  ")
	fmt.Println("## ###### ##    " + char.getWoundsString(EQUIP_ARMS))
	fmt.Println("└┘ ###### └┘   ")
	fmt.Println("   ######    ")
	fmt.Println("   ##  ## ──── Legs:  ")
	fmt.Println("   ##  ##       " + char.getWoundsString(EQUIP_LEG))
	fmt.Println("   ││  ││      ")
	fmt.Println("   ##  ##      ")
	fmt.Println("  ###  ### ─── Feet:  ")
	fmt.Println("                " + char.getWoundsString(EQUIP_FEET))

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
	fmt.Println()
	fmt.Println(" -Attributes-")
	fmt.Printf(" Per: %v \n", character.per)
	fmt.Printf(" Str: %v \n", character.str)
	fmt.Printf(" Agi: %v \n", character.agi)
	fmt.Printf(" Int: %v \n", character.intl)
	fmt.Printf(" Cha: %v \n", character.cha)
	fmt.Printf(" Gui: %v \n", character.gui)
	fmt.Println() 
	fmt.Println(" -Skills-")
	pack1 := packSpaceString(fmt.Sprintf(" %s: %v", skills[0], character.skills[0]), 28)
	pack2 := packSpaceString(fmt.Sprintf("%s: %v", skills[1], character.skills[1]), 28)	
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[2], character.skills[2]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[3], character.skills[3]), 28)	
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[4], character.skills[4]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[5], character.skills[5]), 28)	
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[6], character.skills[6]), 28)	
	fmt.Println(pack1)
	
	fmt.Println()

	fmt.Printf("\nCrowns: %v", character.crowns)

	if pause > 0 {
		rsp := "n"
		fmt.Println("\nPress enter to continue.")
		fmt.Scanln(&rsp)
	}
}

func (char *Character) chooseItemForSlot(slot string) {
	sl,_ := strconv.Atoi(slot)
	
	if (sl >= 3){
		sl -= 3
	} else if (sl == 0){
		sl = 7
	} else {
		sl = EQUIP_HAND
	}
	
	itemsAvail := char.getAllAvailableItemsForSlot(sl)
	
	if len(itemsAvail) < 1 {
		showPause("Character does not possess any items that can be equipped to this slot.")
		return
	}
	
	cont := true
	
	for cont {
		clearConsole()

		itmCount := 0;
		fmt.Println("Available Items for Slot")
		fmt.Println("--------------------------")

		for k := 0; k < len(itemsAvail); k += 2 {
			row := ""
			row = packSpaceString(fmt.Sprintf("%v. %s", itmCount, itemsAvail[itmCount].name), 24)
			itmCount++
			
			if itmCount < len(itemsAvail) {
				row += packSpaceString(fmt.Sprintf("%v. %s", itmCount, itemsAvail[itmCount].name), 24)
				itmCount++			
			}
			
			fmt.Println(row)
		}
		
		fmt.Println("")		
		fmt.Println("n. nothing")		
		fmt.Println("e. Exit")		
		fmt.Println("--------------------------")	
		fmt.Println("Choose item number to equip: ")
		
		rsp := ""
		fmt.Scanln(&rsp)
		
		if rsp == "e" {
			cont = false
		} else if rsp == "n" {
			item := getEmptyItem()
			if slot == "1" {
				oldItem := char.handSlots[LEFT]
				char.removeItemFromCharacter(oldItem)
				char.handSlots[LEFT] = item
				if (oldItem.id > 0){
					char.inventory = append(char.inventory, oldItem)
				}					
			} else if slot == "2" {
				oldItem := char.handSlots[RIGHT]
				char.removeItemFromCharacter(oldItem)
				char.handSlots[RIGHT] = item
				if (oldItem.id > 0){
					char.inventory = append(char.inventory, oldItem)
				}					
			} else {
				if (item.equip < 9) {
					oldItem := char.armorSlots[sl]
					if char.armorSlots[sl].id > 0 {
						char.removeItemFromCharacter(oldItem)
						char.armorSlots[sl] = item
						if (oldItem.id > 0){
							char.inventory = append(char.inventory, oldItem)
						}
					}
				}
			}
			cont = false			
		} else {
			indx,exr := strconv.Atoi(rsp)
		
			if exr == nil {
				if indx < len(itemsAvail) {
					item := itemsAvail[indx]
					if slot == "1" {
						oldItem := char.handSlots[LEFT]
						char.removeItemFromCharacter(oldItem)
						char.removeItemFromCharacter(item)
						char.handSlots[LEFT] = item
						if (oldItem.id > 0){
							char.inventory = append(char.inventory, oldItem)
						}					
					} else if slot == "2" {
						oldItem := char.handSlots[RIGHT]
						char.removeItemFromCharacter(oldItem)
						char.removeItemFromCharacter(item)					
						char.handSlots[RIGHT] = item
						if (oldItem.id > 0){
							char.inventory = append(char.inventory, oldItem)
						}					
					} else {
						if (item.equip < 9) {
							oldItem := char.armorSlots[sl]
							if char.armorSlots[sl].id > 0 {
								char.removeItemFromCharacter(oldItem)
								char.removeItemFromCharacter(item)							
								char.armorSlots[sl] = item
								if (oldItem.id > 0){
									char.inventory = append(char.inventory, oldItem)
								}
							}
						}
					}
					
					cont = false
				}
			}
		}
	}	
}

func (char *Character) equipScreen() {
	cont := true
	
	for cont {
		clearConsole()
		
		fmt.Println("Equip Screen")
		fmt.Println("")
		fmt.Println("1. Left Hand")
		fmt.Println("2. Right Hand")
		fmt.Println("")
		fmt.Println("3. Head")
		fmt.Println("4. Neck")
		fmt.Println("5. Arms")
		fmt.Println("6. Chest")
		fmt.Println("7. Legs")
		fmt.Println("8. Feet")
		fmt.Println("9. Ring")
		fmt.Println("0. Cloak")
		fmt.Println("")		
		fmt.Println("e. Exit")		
		fmt.Println("")
		fmt.Println("Choose a slot number to equip: ")
		
		rsp := ""
		fmt.Scanln(&rsp)
		
		if rsp == "e" {
			cont = false
		} else if rsp == "1" || rsp == "2" || rsp == "3" || rsp == "4" || rsp == "5" || rsp == "6" || rsp == "7" || rsp == "8" || rsp == "9" || rsp == "0"{
			char.chooseItemForSlot(rsp)	
		}
	}

}

func (character *Character) showInventory() {
	cont := true
	
	for cont {
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

		fmt.Println("")
		fmt.Println("(eq. equip) (r. remove) (ex. exit)")
		fmt.Println("")
		fmt.Printf("Choose an option: ")
		rsp := ""
		fmt.Scanln(&rsp)	
		
		if rsp == "eq" {
			character.equipScreen()
		} else if rsp == "ex" {
			cont = false
		}
	}

}
