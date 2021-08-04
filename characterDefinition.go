// CharacterDefine.go
package main

import "fmt"
import "strings"
import "strconv"

// *** be pragmatic and keep first and last names < 11 characters each ***
var femaleNames = []string{"Sarah", "Donna", "Kathryn", "Sheila", "Clarissa", "Coral", "Elizabeth", "Ailla", "Elaine", "Halley"}
var maleNames = []string{"Sam", "Richard", "Mason", "Hunter", "Conner", "Bentley", "Garriot", "Tanner", "Norris", "Robert"}
var lastNames = []string{"Snow", "Smith", "Unknown", "Peters", "Matthew", "Vague", "Mason", "Haston", "Carpathia", "Lennox"}

var skills = []string{"Puzzles", "Politicking", "Investigation", "Alchemy", "Craft", "Spellcraft", "Chirurgery", "unused", "Blades", "Crossbow", "Polearms", "Blunt"}
var weaponSkills = []string{"Blades", "Crossbow", "Polearms", "Blunt"}

var experienceReqs = []int{	0, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 		// 60 levels
							85, 90, 95, 100, 105, 110, 115, 120, 125, 130, 135,
							140, 145, 150, 155, 160, 165, 170, 175, 180, 185,
							190, 195, 200, 205, 210, 215, 220, 225, 230, 235,
							240, 245, 250, 255, 260, 265, 270, 275, 280, 285,
							290, 295, 300, 305, 310}

const NUM_SKILLS = 12
const LEFT = 0
const RIGHT = 1

const STATUS_REST = 0
const STATUS_COMPANION = 1
const STATUS_TRAINING = 2

var taskCodes = []string{"RESTING", "COMPANION", "TRAINING"}

// TODO: add status codes for various Keep jobs (grow, tend crops, craft, etc.)

// new char attributes - Perception (view distance, searching)
//					   - Agility (actions per turn)
//					   - Strength
// 					   - Intellect
//					   - Charm
//					   - Guile

// Experience and leveling
// Upon attaining enough experience to advance to the  next level the character/apprentice is given
// 2 enhancement points.  Points can be spent as follows:  1pt +1 Skill  2pt +1 Att
// Once points are spent the character/apprentice must spend so much time training (they are unavailable)
// Time spent equals  (new level  * 1 day) * 2 (att)
// The character trains much faster.

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
	instanceId                    int
	agi, str, per, intl, cha, gui int
	name                          string // max name length for FORMATTING reasons is 23 characters!
	hp, maxhp                     int
	soul, maxsoul                 int // soul is both a spiritual hp and a tool to craft/power artefacts
	weight, maxweight             int
	crowns                        int
	lvl                           int
	exp                           int
	turns                         int
	turnDefense                   int // how many turns were used as defense
	skills                        [12]int
	alive                         bool
	handSlots                     [2]Item
	armorSlots                    [9]Item
	wounds                        []Wound
	inventory                     []Item
	spellbook                     Spellbook
	villageIndex                  int
	subLoc                        int // better defines where in a village/mission the character can be found (apprentices/npcs)
	gender                        int // 1 - boy, 2 - girl
	trainingPoints                int
	trainingTime 				  int
	task						  int	// a status code for apprentice jobs/training/rest
}

func getNewBlankCharacter(name string) Character {
	var char Character

	game.charInstanceId++

	char.setClearInventory()

	char.name = name
	char.instanceId = game.charInstanceId

	char.str = 2
	char.agi = 2
	char.intl = 2
	char.gui = 2
	char.cha = 2
	char.per = 2

	for k := 0; k < len(char.skills); k++ {
		char.skills[k] = 0
	}

	char.lvl = 1

	char.crowns = 0

	char.hp = char.str
	char.maxhp = char.hp

	char.soul = char.gui + char.cha
	char.maxsoul = char.soul

	char.maxweight = char.str * 10
	char.weight = 0

	char.handSlots[0] = getEmptyItem()
	char.handSlots[1] = getEmptyItem()

	char.wounds = make([]Wound, 0, 0)
	char.alive = true

	char.exp = 0

	return char
}

func (char *Character) getPowerBalance() float32 {
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

	totalMoves += 3

	// TODO: add equipment, other buffs

	return totalMoves
}

func (char *Character) giveSoul(amt int) {
	if (char.soul + amt) <= char.maxsoul {
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

	adj += char.turnDefense
	adj += char.handSlots[LEFT].defense
	adj += char.handSlots[RIGHT].defense

	return adj
}

func (char *Character) getResistanceAt(charBodyIndex int) int {
	equipIndex := HUMAN_TARGETS[charBodyIndex]

	if char.armorSlots[equipIndex].id != -1 {
		return char.armorSlots[equipIndex].resistance
	}

	return 2
}

// check to make sure actor is alive
func (char *Character) isAlive() bool {
	return char.hp > 0
}

// check to make sure actor exists
func (char *Character) exists() bool {
	if char.instanceId > 0 {
		return true
	}
	return false
}

// actor exists & is still living, an actionable actor
// status effects could impact this later
func (char *Character) isMotile() bool {
	if char.exists() && char.isAlive() {
		return true
	}
	return false
}

func (char *Character) getHealthString() string {

	healthString := ""

	for k := 0; k < char.maxhp; k++ {
		if k < char.hp {
			healthString += "♥"
		} else {
			healthString += "-"
		}
	}

	return healthString
}

func getRandomName(gender int) string {
	var die Die
	name := ""

	if gender == 1 { // boy
		roll1 := die.rollxdx(1, len(maleNames)) - 1
		roll2 := die.rollxdx(1, len(lastNames)) - 1

		name = maleNames[roll1] + " " + lastNames[roll2]
	} else { // girl
		roll1 := die.rollxdx(1, len(femaleNames)) - 1
		roll2 := die.rollxdx(1, len(lastNames)) - 1
		name = femaleNames[roll1] + " " + lastNames[roll2]
	}

	return name
}

// a slightly sexist random character generator
func getRandomApprentice(genderbias int) Character {
	var die Die
	female := false
	randomApprentice := getNewBlankCharacter("")

	if genderbias == 0 { // dont care
		if die.rollxdx(1, 4) > 2 {
			female = true
		}
	} else {
		if genderbias == 1 {
			female = false
		} else {
			female = true
		}
	}
	if female {
		roll1 := die.rollxdx(1, len(femaleNames)) - 1
		roll2 := die.rollxdx(1, len(lastNames)) - 1

		randomApprentice.name = femaleNames[roll1] + " " + lastNames[roll2]
		randomApprentice.gender = 2
	} else {
		roll1 := die.rollxdx(1, len(maleNames)) - 1
		roll2 := die.rollxdx(1, len(lastNames)) - 1

		randomApprentice.name = maleNames[roll1] + " " + lastNames[roll2]
		randomApprentice.gender = 1
	}

	// give them a random skill and attribute pt based on gender (sexist, huh?)
	if female {
		roll := die.rollxdx(1, 3)

		if roll == 1 {
			randomApprentice.intl++
		} else if roll == 2 {
			randomApprentice.cha++
		} else if roll == 3 {
			randomApprentice.gui++
		}

	} else {
		roll := die.rollxdx(1, 3)

		if roll == 1 {
			randomApprentice.agi++
		} else if roll == 2 {
			randomApprentice.str++
		} else if roll == 3 {
			randomApprentice.per++
		}
	}

	roll := die.rollxdx(1, 7) - 1

	randomApprentice.skills[roll]++

	// TODO: Add a cool random interesting thing to each character (maybe minus an att, or give them an item, 'er something)

	randomApprentice.hp = randomApprentice.str
	randomApprentice.maxhp = randomApprentice.hp

	randomApprentice.soul = randomApprentice.gui + randomApprentice.cha
	randomApprentice.maxsoul = randomApprentice.soul

	randomApprentice.maxweight = randomApprentice.str * 10
	randomApprentice.weight = 0

	return randomApprentice
}

func (c *Character) trainSkills() {
	var flag bool = true

	for flag {
		clearConsole()
		counter := 0
		fmt.Println("--- Train Skills ---")
		fmt.Println("XXX A saying here is something Steve must write!")
		fmt.Println(packSpaceString("  Skill", 24) + packSpaceString("Curr", 6) + packSpaceString("Cost", 6) + "Training Time")
		fmt.Println("")

		for k := 0; k < len(skills); k++ {
			counter++
			bit := packSpaceString(fmt.Sprintf("%v. %s ", counter, skills[k]), 24)
			bit2 := packSpaceString(fmt.Sprintf("%v", c.skills[k]), 6)
			bit3 := packSpaceString(fmt.Sprintf("%v", 1), 6)
			bit4 := fmt.Sprintf("%v", c.skills[k]+1)
			val := bit + bit2 + bit3 + bit4
			fmt.Println(val)
		}

		fmt.Println("")
		fmt.Println("m. Minutiae (Help)")
		fmt.Println("f. Finished")
		fmt.Println("--------------------")
		fmt.Printf("Points remaining: %v \n", c.trainingPoints)
		fmt.Println("Choose an attribute to add a point: ")
		rsp := ""
		fmt.Scanln(&rsp)

		if rsp == "f" {
			flag = false
		} else if rsp == "m" {
			showSkillsMinutiae()
		} else {
			if c.trainingPoints < 1 {
				rsp2 := ""
				fmt.Println("No points remain. Press enter to return.")
				fmt.Scanln(&rsp2)
				flag = false;
			} else {

				switch rsp {
				case "1":
					c.skills[0]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[0]
					c.task = STATUS_TRAINING

				case "2":
					c.skills[1]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[1]
					c.task = STATUS_TRAINING
					
				case "3":
					c.skills[2]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[2]
					c.task = STATUS_TRAINING

				case "4":
					c.skills[3]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[3]	
					c.task = STATUS_TRAINING
									
				case "5":
					c.skills[4]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[4]	
					c.task = STATUS_TRAINING
									
				case "6":
					c.skills[5]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[5]
					c.task = STATUS_TRAINING
										
				case "7":
					c.skills[6]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[6]
					c.task = STATUS_TRAINING
					
				case "8":
					c.skills[7]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[7]
					c.task = STATUS_TRAINING
										
				case "9":
					c.skills[8]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[8]
					c.task = STATUS_TRAINING
										
				case "10":
					c.skills[9]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[9]
					c.task = STATUS_TRAINING
										
				case "11":
					c.skills[10]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[10]
					c.task = STATUS_TRAINING
										
				case "12":
					c.skills[11]++
					c.trainingPoints -= 1
					c.trainingTime += c.skills[11]
					c.task = STATUS_TRAINING				
				}
			}
		}
	}
}

func (c *Character) trainAttribute() {
	var flag bool = true

	for flag {
		clearConsole()
		fmt.Println("--- Choose Attribute to Train---")
		fmt.Println("Its not what you have done, but what you will do.")
		fmt.Println("  Attribute   Curr     Cost      Training Time")
		fmt.Printf("1. Perception: %v        %v        %v\n", c.per, 2, c.per*2)
		fmt.Printf("2. Strength:   %v        %v        %v\n", c.str, 2, c.str*2)
		fmt.Printf("3. Agility:    %v        %v        %v\n", c.agi, 2, c.agi*2)
		fmt.Printf("4. Intellect:  %v        %v        %v\n", c.intl, 2, c.intl*2)
		fmt.Printf("5. Charm:      %v        %v        %v\n", c.cha, 2, c.cha*2)
		fmt.Printf("6. Guile:      %v        %v        %v\n", c.gui, 2, c.gui*2)
		fmt.Println("")
		fmt.Println("7. Minutiae (Help)")
		fmt.Println("x. Finished")
		fmt.Println("--------------------")
		fmt.Printf("Points remaining: %v \n", c.trainingPoints)
		fmt.Println("Choose an attribute to add a point: ")
		rsp := ""
		fmt.Scanln(&rsp)

		if rsp == "x" {
			flag = false
		} else if rsp == "7" {
			showAttributesMinutiae()
		} else {
			if c.trainingPoints < 2 {
				rsp2 := ""
				fmt.Println("Not enough training points to train attribute. Press enter to return.")
				fmt.Scanln(&rsp2)
				flag = false
			} else {

				switch rsp {
				case "1":
					c.per++
					c.trainingPoints -= 2
					c.trainingTime += c.per*2
					c.task = STATUS_TRAINING
				case "2":
					c.str++
					c.trainingPoints -= 2
					c.trainingTime += c.str*2
					c.task = STATUS_TRAINING

				case "3":
					c.agi++
					c.trainingPoints -= 2
					c.trainingTime += c.agi*2
					c.task = STATUS_TRAINING

				case "4":
					c.intl++
					c.trainingPoints -= 2
					c.trainingTime += c.intl*2		
					c.task = STATUS_TRAINING
								
				case "5":
					c.cha++
					c.trainingPoints -= 2
					c.trainingTime += c.cha*2
					c.task = STATUS_TRAINING
					
				case "6":
					c.gui++
					c.trainingPoints -= 2
					c.trainingTime += c.gui*2
					c.task = STATUS_TRAINING
										
				}
			}
		}
	}
}

func (char *Character) train() {
	rsp := ""

	for rsp != "x" {
		clearConsole()
		fmt.Println("╔ Training ╗")	
		fmt.Println(fmt.Sprintf("Training: %s    Training Pts Avail: %v", char.name, char.trainingPoints))
		fmt.Println("")
		fmt.Println(packSpaceString("Activity ", 24) + "Training Point Cost")
		fmt.Println(packSpaceString("1. Raise Attribute ", 24) + "2")
		fmt.Println(packSpaceString("2. Raise Skill ", 24) + "1")	
		fmt.Println("")
		fmt.Println("x. Finish/Cancel")
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)
		
		if rsp == "1" {
			char.trainAttribute()
		} else if rsp == "2" {
			char.trainSkills()
		} else if rsp != "x" {
			showPause("Invalid selection!")
		}
	}
	
	if (char.instanceId == character.instanceId){
		char.task = STATUS_REST
	}
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

	for k := 0; k < len(char.inventory); k++ {
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
		
		if item.hands == 1 {
			if char.handSlots[LEFT].id == item.id {
				char.handSlots[LEFT] = getEmptyItem()
			} else if char.handSlots[RIGHT].id == item.id {
				char.handSlots[RIGHT] = getEmptyItem()
			} else {
				for k := 0; k < len(char.inventory); k++ {
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
		} else if item.hands == 2 {
			if char.handSlots[LEFT].id == item.id {
				char.handSlots[LEFT] = getEmptyItem()
				char.handSlots[RIGHT] = getEmptyItem()
			} else {
				for k := 0; k < len(char.inventory); k++ {
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

	} else if item.equip == EQUIP_NONE {
		idx := 0
		for k := 0; k < len(char.inventory); k++ {
			if char.inventory[k].id == item.id {
				idx = k
				break
			}
		}

		char.inventory = append(char.inventory[:idx], char.inventory[idx+1:]...)

	} else { // armor
		if char.armorSlots[item.equip].id == item.id {
			char.armorSlots[item.equip] = getEmptyItem()
		} else {
			for k := 0; k < len(char.inventory); k++ {
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

func (char *Character) giveCharacterExperience(exp int) {
	
	char.exp += exp
	
	for experienceReqs[char.lvl] <= char.exp {
		showPause(char.name + " has gone up a level!")
		char.exp -= experienceReqs[char.lvl]
		char.lvl++
		char.trainingPoints += 2
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
	name1 := ""
	name2 := ""
	rsp := ""

	for flag {
		fmt.Println("--- Choose a Character Name ---")
		fmt.Println("A name is nothing more than a tool. Don't forget that.")
		fmt.Println("")
		fmt.Println("Enter a name: ")

		fmt.Scanln(&name1, &name2)
		rsp = name1 + " " + name2

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

func getGender() int {

	clearConsole()
	var flag bool = true
	rsp := ""

	for flag {
		fmt.Println("--- Choose Gender  ---")
		fmt.Println("Meaningless in the long run, but everyone needs something to derive a name.")
		fmt.Println("")
		fmt.Println("1. Male")
		fmt.Println("2. Female")

		fmt.Scanln(&rsp)

		if rsp == "1" {
			flag = false
			return 1
		} else if rsp == "2" {
			flag = false
			return 2
		}
	}

	return -1
}

func (c *Character) isWearing(itm Item) bool {
	if c.handSlots[0].id > 0 {
		if itm.id == c.handSlots[0].id{
			return true
		} 
	}
	if c.handSlots[1].id > 0 {
		if itm.id == c.handSlots[1].id{
			return true
		} 	
	}

	if c.armorSlots[0].id > 0 {
		if itm.id == c.armorSlots[0].id{
			return true
		} 	
	}
	if c.armorSlots[1].id > 0 {
		if itm.id == c.armorSlots[1].id{
			return true
		} 	
	}
	
	if c.armorSlots[2].id > 0 {
		if itm.id == c.armorSlots[2].id{
			return true
		} 	
	}
	if c.armorSlots[3].id > 0 {
		if itm.id == c.armorSlots[3].id{
			return true
		} 	
	}	
	if c.armorSlots[4].id > 0 {
		if itm.id == c.armorSlots[4].id{
			return true
		} 	
	}
	if c.armorSlots[5].id > 0 {
		if itm.id == c.armorSlots[5].id{
			return true
		} 	
	}
	if c.armorSlots[6].id > 0 {
		if itm.id == c.armorSlots[6].id{
			return true
		} 	
	}
	if c.armorSlots[7].id > 0 {
		if itm.id == c.armorSlots[7].id{
			return true
		} 	
	}		
		
	for k := range c.inventory {
		if c.inventory[k].id > 0 {
			if itm.id == c.inventory[k].id {
				return true	
			}
		}
	}
	
	return false
}

func (c *Character) getListOfPossessions() []Item {
	allPossessions := make([]Item, 0)

	if c.handSlots[0].id > 0 {
		allPossessions = append(allPossessions, c.handSlots[0])
	}
	if c.handSlots[1].id > 0 {
		allPossessions = append(allPossessions, c.handSlots[1])
	}

	if c.armorSlots[0].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[0])
	}
	if c.armorSlots[1].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[1])
	}
	if c.armorSlots[2].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[2])
	}
	if c.armorSlots[3].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[3])
	}
	if c.armorSlots[4].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[4])
	}
	if c.armorSlots[5].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[5])
	}
	if c.armorSlots[6].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[6])
	}
	if c.armorSlots[7].id > 0 {
		allPossessions = append(allPossessions, c.armorSlots[7])
	}

	for k := range c.inventory {
		if c.inventory[k].id > 0 {
			allPossessions = append(allPossessions, c.inventory[k])
		}
	}

	return allPossessions
}

func (c *Character) getAllAvailableItemsForSlot(slot int) []Item {
	availItems := make([]Item, 0, 0)

	for k := range c.inventory {
		if c.inventory[k].equip == slot || c.inventory[k].equip == EQUIP_ANY {
			availItems = append(availItems, c.inventory[k])
		}
	}

	if slot == EQUIP_HAND {
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
		counter := 0
		fmt.Println("--- Purchase Skills ---")
		fmt.Println("XXX A saying here is something Steve must write!")
		fmt.Println(packSpaceString("  Skill", 24) + "  Level")
		fmt.Println("")

		for k := 0; k < len(skills); k++ {
			counter++
			bit := packSpaceString(fmt.Sprintf("%v. %s ", counter, skills[k]), 24)
			bit2 := fmt.Sprintf("%v", c.skills[k])
			val := bit + bit2
			fmt.Println(val)
		}

		fmt.Println("")
		fmt.Println("m. Minutiae (Help)")
		fmt.Println("r. Reset")
		fmt.Println("f. Finished")
		fmt.Println("--------------------")
		fmt.Printf("Points remaining: %v \n", points)
		fmt.Println("Choose an attribute to add a point: ")
		rsp := ""
		fmt.Scanln(&rsp)

		if rsp == "f" {
			flag = false
		} else if rsp == "r" {
			for j := 0; j < len(skills); j++ {
				c.skills[j] = 1
			}
			points = 3
		} else if rsp == "m" {
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
				case "8":
					c.skills[7]++
					points -= 1
				case "9":
					c.skills[8]++
					points -= 1
				case "10":
					c.skills[9]++
					points -= 1
				case "11":
					c.skills[10]++
					points -= 1
				case "12":
					c.skills[11]++
					points -= 1
				}

			}
		}
	}
}

func (c *Character) purchaseStats() {
	var flag bool = true
	var points int = 6

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

func createPlayerCharacter() Character {
	var char Character

	char.setClearInventory()

	char.instanceId = 1
	char.name = getName()
	char.gender = getGender()

	char.str = 3
	char.agi = 3
	char.intl = 3
	char.gui = 3
	char.cha = 3
	char.per = 3

	char.purchaseStats()
	char.chooseSkills()

	char.lvl = 3

	char.crowns = 32

	char.hp = char.str
	char.maxhp = char.hp

	char.soul = char.gui + char.cha
	char.maxsoul = char.soul

	char.maxweight = char.str * 10
	char.weight = 0

	char.handSlots[0] = getEmptyItem()
	char.handSlots[1] = getEmptyItem()

	char.wounds = make([]Wound, 0, 0)
	char.alive = true

	char.exp = 0
	char.villageIndex = 0

	return char
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

func (char *Character) getStatus(task int) (string){
	taskDescrip := taskCodes[task]
	if task == STATUS_TRAINING {
		taskDescrip += " " + fmt.Sprintf("%v", char.trainingTime)
	}
	
	str := packSpaceString(taskDescrip, 18)
	str = "* " + str + " *" 
	return str
}


func (char *Character) printCharacter(pause int) {

	clearConsole()

	fmt.Printf("Name: %s    ", char.name)
	fmt.Printf("Level: %v    ", char.lvl)
	fmt.Printf("Exp: %v  /  %v  ", char.exp, experienceReqs[char.lvl])

	fmt.Println()
	fmt.Printf("Hp: %v / %v  ", char.hp, char.maxhp)
	fmt.Printf("Soul: %v / %v  ", char.soul, char.maxsoul)
	fmt.Printf("Encumb: %v / %v  st", char.weight, char.maxweight)
	fmt.Println()
	fmt.Println()
	fmt.Println(" -Attributes-")

	tStr := packSpaceString(fmt.Sprintf(" Per: %v", char.per), 15) + "*************** "
	fmt.Println( tStr)
	tStr = packSpaceString(fmt.Sprintf(" Str: %v ", char.str), 15) + char.getStatus(char.task)
	fmt.Println(tStr)
	tStr = packSpaceString(fmt.Sprintf(" Agi: %v ", char.agi), 15) + "*************** "		
	fmt.Println(tStr)	

	fmt.Printf(" Int: %v \n", char.intl)
	fmt.Printf(" Cha: %v \n", char.cha)
	fmt.Printf(" Gui: %v \n", char.gui)
	fmt.Println()
	fmt.Println(" -Skills-")
	pack1 := packSpaceString(fmt.Sprintf(" %s: %v", skills[0], char.skills[0]), 28)
	pack2 := packSpaceString(fmt.Sprintf("%s: %v", skills[1], char.skills[1]), 28)
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[2], char.skills[2]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[3], char.skills[3]), 28)
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[4], char.skills[4]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[5], char.skills[5]), 28)
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[6], char.skills[6]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[7], char.skills[7]), 28)
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[8], char.skills[8]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[9], char.skills[9]), 28)
	fmt.Println(pack1 + pack2)
	pack1 = packSpaceString(fmt.Sprintf(" %s: %v", skills[10], char.skills[10]), 28)
	pack2 = packSpaceString(fmt.Sprintf("%s: %v", skills[11], char.skills[11]), 28)
	fmt.Println(pack1 + pack2)

	fmt.Println()

	fmt.Printf("\nCrowns: %v", char.crowns)

	if pause > 0 {
		rsp := "n"
		fmt.Println("\nPress enter to continue.")
		fmt.Scanln(&rsp)
	}
}

func (char *Character) chooseItemForSlot(slot string) {
	sl, _ := strconv.Atoi(slot)

	if sl >= 3 {
		sl -= 3
	} else if sl == 0 {
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

		itmCount := 0
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
				if oldItem.id > 0 {
					char.inventory = append(char.inventory, oldItem)
				}
			} else if slot == "2" {
				oldItem := char.handSlots[RIGHT]
				char.removeItemFromCharacter(oldItem)
				char.handSlots[RIGHT] = item
				if oldItem.id > 0 {
					char.inventory = append(char.inventory, oldItem)
				}
			} else {
				if item.equip < 9 {
					oldItem := char.armorSlots[sl]
					if char.armorSlots[sl].id > 0 {
						char.removeItemFromCharacter(oldItem)
						char.armorSlots[sl] = item
						if oldItem.id > 0 {
							char.inventory = append(char.inventory, oldItem)
						}
					}
				}
			}
			cont = false
		} else {
			indx, exr := strconv.Atoi(rsp)

			if exr == nil {
				if indx < len(itemsAvail) {
					item := itemsAvail[indx]
					if slot == "1" {						
						if item.hands == 1 {
							oldItem := char.handSlots[LEFT]
	
							char.removeItemFromCharacter(oldItem)
							char.removeItemFromCharacter(item)
	
							char.handSlots[LEFT] = item			
	
							if oldItem.id > 0 {
								char.inventory = append(char.inventory, oldItem)
							}							
						} else if item.hands == 2 {
							oldItem := char.handSlots[LEFT]
							oldItem2 := char.handSlots[RIGHT]
	
							char.removeItemFromCharacter(oldItem)
							char.removeItemFromCharacter(oldItem)

							char.removeItemFromCharacter(item)
	
							char.handSlots[LEFT] = item			
							char.handSlots[RIGHT] = item			
	
							if oldItem.id > 0 {
								char.inventory = append(char.inventory, oldItem)
							}	
							if oldItem2.id > 0 {
								char.inventory = append(char.inventory, oldItem2)
							}														
						}					

					} else if slot == "2" {
						if item.hands == 1 {
							oldItem := char.handSlots[RIGHT]
	
							char.removeItemFromCharacter(oldItem)
							char.removeItemFromCharacter(item)
	
							char.handSlots[RIGHT] = item			
	
							if oldItem.id > 0 {
								char.inventory = append(char.inventory, oldItem)
							}							
						} else if item.hands == 2 {
							oldItem := char.handSlots[LEFT]
							oldItem2 := char.handSlots[RIGHT]
	
							char.removeItemFromCharacter(oldItem)
							char.removeItemFromCharacter(oldItem)

							char.removeItemFromCharacter(item)
	
							char.handSlots[LEFT] = item			
							char.handSlots[RIGHT] = item			
	
							if oldItem.id > 0 {
								char.inventory = append(char.inventory, oldItem)
							}	
							if oldItem2.id > 0 {
								char.inventory = append(char.inventory, oldItem2)
							}														
						}	
					} else {
						if item.equip < 9 {
							oldItem := char.armorSlots[sl]
							if char.armorSlots[sl].id > 0 {
								char.removeItemFromCharacter(oldItem)
								char.removeItemFromCharacter(item)
								char.armorSlots[sl] = item
								if oldItem.id > 0 {
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
		} else if rsp == "1" || rsp == "2" || rsp == "3" || rsp == "4" || rsp == "5" || rsp == "6" || rsp == "7" || rsp == "8" || rsp == "9" || rsp == "0" {
			char.chooseItemForSlot(rsp)
		}
	}
}

func tradeItems(direction int) {

	allPossessions := make([]Item, 0)
	charString := ""

	if direction == 0 {
		allPossessions = character.getListOfPossessions()
		charString = apprentice.name
	} else {
		allPossessions = apprentice.getListOfPossessions()
		charString = character.name
	}

	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		fmt.Println("What item do you wish to give  " + charString + "?")
		fmt.Println("-----------------------------------------------------------------")

		if len(allPossessions) < 1 {
			fmt.Println("Nothing available")
		} else {
			counter := 0
			for k := range allPossessions {
				pack1 := packSpaceString(fmt.Sprintf("%v.  %s", counter, allPossessions[k].name), 36)
				fmt.Println(pack1)
				counter++
			}
		}

		fmt.Println("")
		fmt.Println("e. Exit")
		fmt.Println("--------------------------")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "e" {
			num, err := strconv.Atoi(rsp)

			if err == nil {
				if num < len(allPossessions) {
					item := allPossessions[num]

					if direction == 0 {
						giveOK := apprentice.giveCharacterItem(item)

						if giveOK {
							character.removeItemFromCharacter(item)
							allPossessions = character.getListOfPossessions()
							showPause(fmt.Sprintf("%s given to %s!", item.name, charString))
						}

					} else if direction == 1 {
						giveOK := character.giveCharacterItem(item)

						if giveOK {
							apprentice.removeItemFromCharacter(item)
							allPossessions = apprentice.getListOfPossessions()
							showPause(fmt.Sprintf("%s given to %s!", item.name, charString))
						}
					}
				}
			}

		} else if rsp == "e" {
			exitFlag = true
		}
	}
}

func (char *Character) dropItems() {
	exitFlag := false

	for !exitFlag {
		
		clearConsole()
	
		fmt.Println("Drop which item?")
		fmt.Println("--------------------------")
	
		for k := range char.inventory {
			fmt.Println(fmt.Sprintf("%v. %s ", k, char.inventory[k].name))
		}
		
		fmt.Println("")
		fmt.Println("[x. Exit]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")
	
		rsp := ""
		fmt.Scanln(&rsp)
		
		if len(rsp) > 0 && rsp != "x" {
			num, _ := strconv.Atoi(rsp)
			item := char.inventory[num]
			char.removeItemFromCharacter(item)
			if game.disposition == 1 {
				game.battleGrid.dropItem(item)
			}	
		} else if rsp == "x" {	
			exitFlag = true
		}
	}
}

func (char *Character) sellItems(typ int) {
	
	exitFlag := false
	rsp := ""
	for !exitFlag {
		applicableItems := make([]Item, 0, 0)
		allItems := char.getListOfPossessions()

		for k := range allItems {
			if allItems[k].typeCode == typ {
				applicableItems = append(applicableItems, allItems[k])
			}
		}
		
		clearConsole()
		fmt.Println(fmt.Sprintf("Sell Items      Crowns:  %v", character.crowns))
		fmt.Println("-----------------------------------------------------------------")

		fmt.Println("   Item                                 \tWgt \tCost")

		if len(applicableItems) < 1 {
			fmt.Println("")
			fmt.Println("  (No Items of this category to sell.)")	
		}

		for i := 0; i < len(applicableItems); i++ {
			s:= ""
			if char.isWearing(applicableItems[i]) {
				s = "  (Worn)"
			}
			fmt.Printf("%v. %s %s       \t%v \t%v \n", i, packSpaceString(applicableItems[i].name, 24), s, applicableItems[i].weight, getItemSellPrice(applicableItems[i]))
		}

		fmt.Println("")
		fmt.Println("[x. Back]   [n. More]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		fmt.Scanln(&rsp)
		
		if len(rsp) > 0 && rsp != "x" {
			idx, _ := strconv.Atoi(rsp)			
			
			fmt.Println(fmt.Sprintf("Sell %s for %v?", applicableItems[idx].name, getItemSellPrice(applicableItems[idx])))
			fmt.Scanln(&rsp)
			
			if rsp == "y" {
				character.crowns += getItemSellPrice(applicableItems[idx])
				char.removeItemFromCharacter(applicableItems[idx])
				char.recalcCharacterWeight()
				
				showPause(applicableItems[idx].name + " sold!")			
			}
			
		} else if rsp == "x" {	
			exitFlag = true		
		} 
	}	
	
	
}


func (char *Character) storeItems() {
	if len(keep.storage) >= keep.maxStorage {
		showPause("Keep storage is maxed out. Build larger storage facilities or clean the dump up! Ya fucking pack rat.")
		return
	}

	const ITEMS_PER_PAGE = 16

	cont := true
	rsp := ""

	range1 := 0
	range2 := ITEMS_PER_PAGE
	pages := 0
	page := 0

	for cont {
		clearConsole()

		items := char.getListOfPossessions()

		pages = 1
		if len(items) > ITEMS_PER_PAGE {
			for j := len(items); j > ITEMS_PER_PAGE; j -= ITEMS_PER_PAGE {
				pages++
			}
		}

		fmt.Println("-- Keep Storage: " + fmt.Sprintf("%v of %v used.", len(keep.storage), keep.maxStorage))
		fmt.Println("")

		fmt.Println(fmt.Sprintf("-- Character Inventory --  [Page %v : %v]", page+1, pages))

		range1 = page * ITEMS_PER_PAGE
		range2 = range1 + ITEMS_PER_PAGE

		if range2 > len(items) {
			range2 = len(items)
		}

		for k := range1; k < range2; k++ {
			fmt.Println(fmt.Sprintf("%v. %s ", k, items[k].name))
		}

		fmt.Println("")
		if pages > 1 {
			fmt.Println("[n. next page]")
		} else {
			fmt.Println("")
		}

		fmt.Println("--------------------")
		choices := "(#. Store Item) (a. Store All) (u. Store Unequiped) (x. Exit)"
		fmt.Println(choices)
		fmt.Println("")
		fmt.Printf("Choose an option: ")

		fmt.Scanln(&rsp)

		if rsp == "x" {
			cont = false
		} else if rsp == "a" {

		} else if rsp == "u" {

		} else if rsp == "n" && pages > 1 {
			page++
			if page >= pages {
				page = 0
			}
		} else {
			num, err := strconv.Atoi(rsp)

			if err == nil {
				selection := (page * 12) + num
				storeItem := items[selection]

				char.removeItemFromCharacter(storeItem)
				keep.storage = append(keep.storage, storeItem)

				showPause(storeItem.name + " stored in Keep!")

			} else {
				showPause("Invalid selection.")
			}
		}
	}
}

func (char *Character) getFilteredInventoryList(filters []bool) []Item {
	filteredList := make([]Item, 0, 0)
	allList := char.getListOfPossessions()

	for k := 0; k < len(allList); k++ {
		itm := allList[k]

		if itm.typeCode == ITEM_TYPE_WEAPON && filters[0] {
			filteredList = append(filteredList, itm)
		} else if itm.typeCode == ITEM_TYPE_ARMOR && filters[1] {
			filteredList = append(filteredList, itm)
		} else if itm.typeCode == ITEM_TYPE_UNCTURE && filters[2] {
			filteredList = append(filteredList, itm)
		} else if itm.typeCode == ITEM_TYPE_INGREDIENT && filters[3] {
			filteredList = append(filteredList, itm)
		} else if itm.typeCode == ITEM_TYPE_EQUIPMENT && filters[4] {
			filteredList = append(filteredList, itm)
		} else if itm.typeCode == ITEM_TYPE_SPECIAL && filters[5] {
			filteredList = append(filteredList, itm)
		}
	}

	return filteredList
}

func (char *Character) showInventoryFilteredList() {
	const ITEMS_PER_PAGE = 18
	rsp := ""
	range1 := 0
	range2 := ITEMS_PER_PAGE
	pages := 0
	page := 0

	filters := make([]bool, 6, 6)
	for k := 0; k < 6; k++ {
		filters[k] = true
	}

	dispFilters := "☼ ⌂ ♥ ♣ ♦ ∞"

	for rsp != "x" {

		filteredList := char.getFilteredInventoryList(filters)
		dispFilters = getDisplayFilters(filters)

		clearConsole()
		fmt.Println("╔═══════════════════════ Inventory ═══════════════════════╗")

		pages = 1
		if len(keep.storage) > ITEMS_PER_PAGE {
			for j := len(filteredList); j > ITEMS_PER_PAGE; j -= ITEMS_PER_PAGE {
				pages++
			}
		}

		dispStr := packSpaceStringCenter(fmt.Sprintf("  %v :: %v ", char.weight, char.maxweight), 24)
		dispStr += "            Filters: " + dispFilters

		fmt.Println(dispStr)
		fmt.Println("  ─────────────────────             ─────────────────────")
		fmt.Println("")

		range1 = page * ITEMS_PER_PAGE
		range2 = range1 + ITEMS_PER_PAGE

		if range2 > len(filteredList) {
			range2 = len(filteredList)
		}

		for k := range1; k < range2; k++ {
			numBit := fmt.Sprintf("   %v.", k)
			numBit = packSpaceString(numBit, 8)
			fmt.Println(numBit + filteredList[k].name)
		}

		fmt.Println("")
		fmt.Println("  ─────────────────────────────────────────────────────")

		commands := ""
		if pages > 1 {
			commands += "  [n. next]"
		}
		commands += "  [f. filters]  [#. View]  [x. Exit]"
		commands = packSpaceString(commands, 46)

		fmt.Println(commands)
		fmt.Println("╚═══════════════════════════════════════════════════════╝")
		fmt.Println("")
		fmt.Printf("Choose an option: ")

		fmt.Scanln(&rsp)

		if rsp == "f" {
			setFilters(filters)
		} else if rsp == "n" && pages > 1 {
			page++
			if page >= pages {
				page = 0
			}
		} else if rsp != "x" {
			num, err := strconv.Atoi(rsp)

			if err == nil {
				selection := (page * ITEMS_PER_PAGE) + num
				storeItem := filteredList[selection]
				show(storeItem)
				showPause("Press Enter to continue.")
				// TODO: View the item
			} else {
				showPause("Invalid selection.")
			}
		}
	}
}

func (char *Character) showInventory() {

	rsp, _ := char.showInventoryChar(true)
	id := char.instanceId

	for rsp != "x" {
		if rsp == "n" {
			if id == character.instanceId && apprentice.exists() {
				rsp, id = apprentice.showInventoryChar(true)
			} else {
				rsp, id = character.showInventoryChar(true)
			}
		} else {
			rsp = "x"
		}
	}
}

func (char *Character) showInventoryChar(canTransfer bool) (string, int) {
	cont := true

	for cont {
		clearConsole()

		seg1 := ""
		seg2 := ""

		weightStr := fmt.Sprintf("Encumb: %v / %v  (stone)", char.weight, char.maxweight)
		fmt.Println(packSpaceString(char.name, 22) + "  " + weightStr)

		fmt.Println("")
		fmt.Println("--Hands--")

		seg1 = packSpaceString("Left Hand: ", 14) + char.handSlots[LEFT].getInvDisplayString()
		seg2 = packSpaceString("Right Hand: ", 14) + char.handSlots[RIGHT].getInvDisplayString()
		fmt.Println(seg1)
		fmt.Println(seg2)

		fmt.Println("")
		fmt.Println("--Armor--")
		seg1 = packSpaceString("Head: ", 14) + char.armorSlots[EQUIP_HEAD].getInvDisplayString()
		seg2 = packSpaceString("Neck: ", 14) + char.armorSlots[EQUIP_NECK].getInvDisplayString()
		fmt.Println(seg1)
		fmt.Println(seg2)
		seg1 = packSpaceString("Chest: ", 14) + char.armorSlots[EQUIP_CHEST].getInvDisplayString()
		seg2 = packSpaceString("Arms: ", 14) + char.armorSlots[EQUIP_ARMS].getInvDisplayString()
		fmt.Println(seg1)
		fmt.Println(seg2)
		seg1 = packSpaceString("Legs: ", 14) + char.armorSlots[EQUIP_LEG].getInvDisplayString()
		seg2 = packSpaceString("Feet: ", 14) + char.armorSlots[EQUIP_FEET].getInvDisplayString()
		fmt.Println(seg1)
		fmt.Println(seg2)
		seg1 = packSpaceString("Cloak: ", 14) + char.armorSlots[EQUIP_CLOAK].getInvDisplayString()
		seg2 = packSpaceString("Ring: ", 14) + char.armorSlots[EQUIP_RING].getInvDisplayString()
		fmt.Println(seg1)
		fmt.Println(seg2)

		fmt.Println("")
		fmt.Println("--Bags--")
		count := 0
		for k := 0; k < len(char.inventory); k++ {
			fmt.Printf("%s", packSpaceString(char.inventory[k].name, 23))
			count++
			if count == 3 {
				count = 0
				fmt.Printf("\n")
			}
		}

		fmt.Println("")
		choices := "(e. equip) (r. remove) "

		if canTransfer && char.instanceId == character.instanceId && apprentice.exists() {
			choices += "(g. give) (n. apprentice) "
		} else if canTransfer && apprentice.exists() && apprentice.instanceId == char.instanceId {
			choices += "(g. give) (n. character) "
		}
		
		if game.disposition == 1 { // in battlegrid
			choices += "(d. drop) "
		}

		if char.villageIndex == 99 {	// in keep
			choices += "(s. store) "
		}

		choices += "(l. list) "

		choices += "(x. exit)"

		fmt.Println(choices)
		fmt.Println("")
		fmt.Printf("Choose an option: ")
		rsp := ""
		fmt.Scanln(&rsp)

		if rsp == "e" {
			char.equipScreen()
		} else if rsp == "l" {
			char.showInventoryFilteredList()
		} else if rsp == "x" {
			cont = false
		} else if canTransfer && rsp == "n" {
			cont = false
			return rsp, char.instanceId
		} else if canTransfer && rsp == "g" {
			if apprentice.exists() {
				if char.instanceId == character.instanceId {
					tradeItems(0) // character to apprentice
				} else {
					tradeItems(1) // apprentice to character
				}
			} else {
				showPause("No apprentice to trade with.")
			}
		} else if char.villageIndex == 99 && rsp == "s" {
			char.storeItems() // if in keep, we can store inventory items
		} else if game.disposition == 1 && rsp == "d" {
			char.dropItems() 			
		}
	}

	return "x", -1
}
