// gameio.go

package main

import "io/ioutil"
import "strings"
import "strconv"
import "fmt"
import "os"
import "bytes"
import "encoding/gob"

const BLOCK_GAME = "[GAME]"
const BLOCK_CHAR = "[CHAR]"
const BLOCK_APP = "[APP]"
const BLOCK_VILLAGE = "[VILLAGE]"
const BLOCK_KEEP = "[KEEP]"
const BLOCK_MISSION = "[MISSION]"

func getGameSaveBlock() string {
	gameBlock := BLOCK_GAME + ","

	gameBlock += fmt.Sprintf("%s,", VERSION)
	gameBlock += fmt.Sprintf("%v,", game.gameDay)
	gameBlock += fmt.Sprintf("%v,", game.itemInstanceId)
	gameBlock += fmt.Sprintf("%v,", game.dayCounter)
	gameBlock += fmt.Sprintf("%v,", game.weekCounter)
	gameBlock += fmt.Sprintf("%v,", game.monthCounter)
	gameBlock += fmt.Sprintf("%v,", game.missionInstanceId)
	gameBlock += fmt.Sprintf("%v,", game.charInstanceId)
	gameBlock += fmt.Sprintf("%v,", game.darkness)
	
	actorCount := len(actors) + len(keep.apprentices) + len(orphanage)
	gameBlock += fmt.Sprintf("%v,", actorCount)
	
	gameBlock += "■"

	return gameBlock
}

func unpackGameBlock(block string) int {
	// only 1 line for game block, no need to split lines
	// just do bits
	bits := strings.Split(block, ",")

	if bits[0] == "[GAME]" {
		fmt.Println("Loading Game Block...")
	} else {
		log.addError("Cant find game block.")
		fmt.Println("Game Block not found!")
		return -1
	}

	ver := bits[1]

	if ver != VERSION {
		log.addError("Save version is incorrect. Current version is " + VERSION + " Save version is " + ver)
		return -1
	}

	game.gameDay, _ = strconv.Atoi(bits[2])
	game.itemInstanceId, _ = strconv.Atoi(bits[3])
	game.dayCounter, _ = strconv.Atoi(bits[4])
	game.weekCounter, _ = strconv.Atoi(bits[5])
	game.monthCounter, _ = strconv.Atoi(bits[6])
	game.missionInstanceId, _ = strconv.Atoi(bits[7])
	game.charInstanceId, _ = strconv.Atoi(bits[8])
	game.darkness, _ = strconv.Atoi(bits[9])
	actorCount, _ := strconv.Atoi(bits[10])
	
	fmt.Println("            ...done!")

	return actorCount
}

func (c *Character) getCharSaveBlock() string {
	saveString := BLOCK_CHAR + ","

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

	saveString += fmt.Sprintf("%v,", c.soul)
	saveString += fmt.Sprintf("%v,", c.maxsoul)

	saveString += fmt.Sprintf("%v,", c.weight)
	saveString += fmt.Sprintf("%v,", c.maxweight)

	saveString += fmt.Sprintf("%v,", c.crowns)
	saveString += fmt.Sprintf("%v,", c.villageIndex)
	
	saveString += fmt.Sprintf("%v,", c.exp)
	saveString += fmt.Sprintf("%v,", c.lvl)

	saveString += fmt.Sprintf("%v,", c.instanceId)
	
	for k := 0; k < len(c.skills); k++ {
		saveString += fmt.Sprintf("%v,", c.skills[k])
	}

	saveString += fmt.Sprintf("%v,", c.subLoc)
	saveString += fmt.Sprintf("%v,", c.gender)
	saveString += fmt.Sprintf("%v,", c.trainingPoints)
	saveString += fmt.Sprintf("%v,", c.task)

	saveString += fmt.Sprintf("%v,", len(c.inventory)) // save count of backpack items

	saveString += "◄" // end line so we can do equipment, spells, etc on their own line

	// save inventory!
	saveString += c.handSlots[0].getSaveString()
	saveString += c.handSlots[1].getSaveString()

	for k := 0; k < len(c.armorSlots); k++ {
		saveString += c.armorSlots[k].getSaveString()
	}

	for k := 0; k < len(c.inventory); k++ {
		saveString += c.inventory[k].getSaveString()
	}

	// TODO: save wounds!

	saveString += "■"

	return saveString
}

func unpackCharacterBlock(block string) (int, Character) {
	var char Character

	lines := strings.Split(block, "◄")

	// first bit is major character stuff, attributes, etc.
	bits := strings.Split(lines[0], ",")

	if bits[0] == BLOCK_CHAR {
		fmt.Println("Loading " + BLOCK_CHAR + "...")
	} else {
		log.addError("Cant find CHAR block.")
		fmt.Println("CHAR Block not found!")
		return -1, char
	}

	lineCounter := 0

	char.name = bits[1]
	char.lvl, _ = strconv.Atoi(bits[2])

	char.str, _ = strconv.Atoi(bits[3])
	char.per, _ = strconv.Atoi(bits[4])
	char.agi, _ = strconv.Atoi(bits[5])
	char.intl, _ = strconv.Atoi(bits[6])
	char.cha, _ = strconv.Atoi(bits[7])
	char.gui, _ = strconv.Atoi(bits[8])

	char.hp, _ = strconv.Atoi(bits[9])
	char.maxhp, _ = strconv.Atoi(bits[10])

	char.soul, _ = strconv.Atoi(bits[11])
	char.maxsoul, _ = strconv.Atoi(bits[12])

	char.weight, _ = strconv.Atoi(bits[13])
	char.maxweight, _ = strconv.Atoi(bits[14])

	char.crowns, _ = strconv.Atoi(bits[15])
	char.villageIndex, _ = strconv.Atoi(bits[16])

	char.exp, _ = strconv.Atoi(bits[17])	
	char.lvl, _ = strconv.Atoi(bits[18])
	
	char.instanceId, _ = strconv.Atoi(bits[19])
	
	char.skills[0], _ = strconv.Atoi(bits[20]) 
	char.skills[1], _ = strconv.Atoi(bits[21]) 
	char.skills[2], _ = strconv.Atoi(bits[22]) 
	char.skills[3], _ = strconv.Atoi(bits[23]) 
	char.skills[4], _ = strconv.Atoi(bits[24]) 
	char.skills[5], _ = strconv.Atoi(bits[25]) 
	char.skills[6], _ = strconv.Atoi(bits[26]) 	
	char.skills[7], _ = strconv.Atoi(bits[27]) 
	char.skills[8], _ = strconv.Atoi(bits[28]) 
	char.skills[9], _ = strconv.Atoi(bits[29]) 
	char.skills[10], _ = strconv.Atoi(bits[30]) 
	char.skills[11], _ = strconv.Atoi(bits[31]) 
	
	char.subLoc, _ = strconv.Atoi(bits[32]) // count of items in backpack
	char.gender, _ = strconv.Atoi(bits[33]) // count of items in backpack
	char.trainingPoints, _ = strconv.Atoi(bits[34]) // count of items in backpack
	char.task, _ = strconv.Atoi(bits[35]) // count of items in backpack

	inventoryCount, _ := strconv.Atoi(bits[36]) // count of items in backpack
	
	// load inventory!
	char.handSlots[0], _ = restoreSavedItem(lines[1])
	char.handSlots[1], _ = restoreSavedItem(lines[2])

	lineCounter = 3
	for k := 0; k < len(char.armorSlots); k++ {
		char.armorSlots[k], _ = restoreSavedItem(lines[lineCounter])
		lineCounter++
	}

	for k := 0; k < inventoryCount; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		char.inventory = append(char.inventory, itm)
		lineCounter++
	}

	// load wounds!
	char.wounds = make([]Wound, 0, 0)

	return 1, char
}

/////////////////////////////////////////////////////////////////////////////////////////////
func save() {
	filename := "save.txt"
	filename2 := "saveenc.txt"

//	var chars []Character
//	var missions []Mission
	
//	chars = make([]Character, 0, 0)
//	missions = make([]Mission, 0, 0)
	
	file, err := os.Create(filename)
	file2, err := os.Create(filename2)

	var saveString string

	saveString = getGameSaveBlock()
	saveString += character.getCharSaveBlock()
	saveString += apprentice.getCharSaveBlock()
	saveString += getKeepSaveBlock()
	saveString += mission.getSaveString()

	for k := 0; k < len(villages); k++ {
		saveString += villages[k].getSaveString()
	}
	
		
	// add all additional actors at END (keep apprentices, other important npcs)

	for k := 0; k < len(keep.apprentices); k++ {
		keep.apprentices[k].villageIndex = 99
		saveString += keep.apprentices[k].getCharSaveBlock()	
	}
	
	for k := 0; k < len(actors); k++ {
		saveString += actors[k].getCharSaveBlock()	
	}	
	
	for k := 0; k < len(orphanage); k++ {
		orphanage[k].villageIndex = 5
		orphanage[k].subLoc = 99		
		saveString += orphanage[k].getCharSaveBlock()	
	}		
	
	if err == nil {
		defer file.Close()
		defer file2.Close()

		file.WriteString(saveString)
		
		var network bytes.Buffer        
		enc := gob.NewEncoder(&network) 
		//dec := gob.NewDecoder(&network) 
		
		_ = enc.Encode(saveString)
		
		file2.WriteString(network.String())
		
		fmt.Println("Game Saved!")
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////
func loadGame() int {
	character.setClearInventory()

	data, err := ioutil.ReadFile("save.txt")
	if err == nil {
		charData := fmt.Sprintf("%s", data)

		if len(charData) > 0 {
			blocks := strings.Split(charData, "■")
			actorCount := unpackGameBlock(blocks[0])

			if actorCount < 0 {
				log.addError("Failed to load save game.")
				return -1
			}
			
			fmt.Println(fmt.Sprintf("%v total actor blocks", actorCount))
			fmt.Println(fmt.Sprintf("Found %v total blocks", len(blocks)))

			_, character = unpackCharacterBlock(blocks[1])
			// current travel apprentice
			_, apprentice = unpackCharacterBlock(blocks[2])
			// keep
			_, keep = unpackKeepBlock(blocks[3])

			_, mission = unpackMissionBlock(blocks[4])
			
			place := 4
			for k := 0; k < len(villages); k++ {
				place += 1
				villIndx := k + 5
				
				_, villages[k] = unpackVillageBlock(k, blocks[villIndx])
			} 
			
			orphanage = make([]Character, 0, 0)	// build an empty orphanage
			
			for k := 0; k < actorCount; k++ {
				place += 1
				_, act := unpackCharacterBlock(blocks[place])
				if (act.villageIndex == 5 && act.subLoc == 99) {
					orphanage = append(orphanage, act)
				} else if (act.villageIndex == 99) {
					keep.addNewApprenticeToKeep(act)
					fmt.Println(act.name + " added to keep")
				} else {
					actors = append(actors, act)
				}
			}
			
			// blocks are broken with ■
			// blocks are character, keep, village, game
			// lines are broken with ◄
			// core line, equipment lines
			// bits are broken with,
			// individual values
		}

		fmt.Println("-----")

		fmt.Printf("\n%s", data)
		fmt.Println("Game Loaded! ")
		log.addInfo("Game loaded.\n")
		showPause("Game Loaded!")
	}

	return 1
}

func testRandom(){
	var filename string
	var die Die
	var nums []int = make([]int, 20, 20)
	
	filename = "random.txt"

	file, err := os.Create(filename)

	saveString := ""
	counter := 0
	for k:= 0; k < 10000; k++ {
		val := die.rollxdx(1, 20)-1
		nums[val]++
		saveString += fmt.Sprintf("%v, ", val)
		counter++
		
		if counter >= 10 {
			counter = 0
			saveString += "\n"
		}
	}
	
	saveString += "\n\n"
	
	for k:= 0; k < 20; k++ {
		saveString += fmt.Sprintf("%v : %v \n", k, nums[k])
	}
	
	if err == nil {
		defer file.Close()

		file.WriteString(saveString)
	}	
}

func (bg *BattleGrid) writeGridsToFile() {
	var filename string

	filename = "grids.txt"

	file, err := os.Create(filename)

	saveString := ""
	
	for k:= 0; k < len(bg.allGrids); k++ {
		if bg.allGrids[k].used {
			grid := bg.getEntityGrid(bg.allGrids[k].id)
			saveString += fmt.Sprintf("***** [%v] %s ***** \n", grid.id, grid.gridName)
			for i := 0; i < len(grid.grid); i++ {
				row := ""
				for t := 0; t < len(grid.grid[i]); t++ {
					row += grid.grid[i][t]
				}
				saveString += row + "\n"
			}
		}
	}

	if err == nil {
		defer file.Close()

		file.WriteString(saveString)
		fmt.Println("Grids Saved!")
	}
}
