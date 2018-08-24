// keep.go

package main

import "fmt"
import "strings"
import "strconv"

var keepDescriptions = []string{
	"It's cold and dark here. Shadows from my waning fire dance across the\n vacant expanse.",
	"It's empty and barren, but it's mine.",
}

type Keep struct {
	name             string
	acres, usedacres int
	descriptionId    int
	maxStorage 		 int	// max number of items storage can hold
	apprentices      []Character
	storage			 []Item		
	mapX, mapY       int
}

// Politicks
// Menu:  Village Status, Curry Favor, Offer Assistance, Donate Crowns, Spend Political Currency
// Village Status is not available until it is purchased with political currency.
// Curry Favor - Spend time to gain favor.
// Offer Assistance - Sometimes the mayor needs assistance, new mission
// Donate Crowns - Simple monetary exchange for favor, crowns donated it this way will improve village metrics.
// Spend Political Currency - Purchase acres for keep, open village status, approve apprentice(?), options vary by village
// uses political favors to gain acres - use land to build useful structures
// View village status
// Approve Apprentices
// Request Assistance - bonus to skill check for mission arch
// ONE VILLAGE HAS AN ORPHANAGE - Apprentices can be purchased but they are not always available.)


// building ideas - each has various levels, land requirements and upgrade prices
// farmland - "bleeding hands scrape weeds from the earth in hopes of a satiated belly - but funds must be made to pay the taxman"
// Greenhouse - produces small amounts (1 or 2 at a time max) of alchemical plants / spell components
// walls - defense?
// hut, shack, house, tower, keep, stronghold - various comfort, healing benefits, maximum apprentice quarters
// Blacksmith - produce weapons
// Chicanery - produce magical stuff
// Study - Prepare lamentations (spells)
// Unctuary - Place to prepare unctures.

// assign Apprentices to work various structures, their skill will affect performance
// thought: an apprentice can be assigned to repair equipment, will pick repair objects randomly unless assigned
// thought: an apprentice assigned to a green house will produce random ingredients

// apprentice should have strongly typed strengths - they should grow towards them
// apprentice new skills will be hard to acquire, will require training that takes time slots for both character and apprentice
// Rarely, a "blank slate" apprentice will be available (Tabula Rasa) - the character can shape them however they see fit.

// MISSIONS are posted in the villages.
// ONLY 1 mission can be worked at a time.
// Missions have ARCHS
// They start with a series of puzzles, each must be solved with a skill roll. CHARM/INVESTIGATE/PUZZLES
// Player can spend a day researching to gain a clue token, this provides a +1 to whatever skill is needed for that arch
// Archs can force players to travel between villages for the next puzzle
// Once all archs are complete, player travels to battlegrid to face the beast.
//
// Incomplete missions can have different consequences. Death toll, financial, political
// the longer the mission is unresolved the larger the impact, death reduces village size, financial
// reduces what is available in the stores, political reduces favor.
// Each descriptively describes the encounter:
// Ex. A corpse candle draws our sheep into the bog to drown.
// Ex. We hung the wrong man, his corpse is back for revenge!
// Ex. A grip (goblin) absconds with our cattle after dark!
// Ex. Valuable property tormented by wily ghast.
// Ex. Something is eating our children.

// POTIONS are called UNCTURES

// HAND OF GLORY - make some creatures flee (5 fingers, 5 uses)
// Cross of St Martin - Randomly move monster to random grid/loc
// Splinter of the True Cross - ???
// Stone Cast at the Lord during the Crucificion - ??
// ??? - bonus moves
// ??? - bonus attack
// ??? - bonus healing

// SPELLS are called LAMENTATIONS
// "Only the majesty of grief and sorrow separate us from the dark.  The dead never grieve, and the dark knows little of sorrow."
// Can only prepare spells at keep.  Require ingredients & preparation time. Maximum number prepared at one time is a figment of intl
// very difficult to learn new spells and the effects are usually muted

// Wander Action - create a random set of grids with ingredients/objects/npcs - potential apprentices, maybe a mugger

// Combat round:
// Player makes a contested attack against opponent.
//		Player attack rating + d20 vs player defense rating + d20
// 			- Player Attack rating is comprised of skill bonus + weapon quality/material bonuses
//			- Player Defense rating is comprised of agi bonus + shield bonus + defense posture bonus
//
//		On HIT
//			Target roll is a d10 roll + (atk roll) bonus that determine what body location will be targeted
// 				For every 5 points over the attack roll - defense roll character can add +1 to target roll
//				Different body locations will provide different wound potentials, and measure armor performance (existance of, etc.)
// 				A target roll of 10+ is considered a critical hit and will result in an additional hit being scored
//					If a target roll is awarded a critical the target roll is remade to assess location.
//					Multiple criticals can be stacked in this way.
//
//			Penetration roll is a d20 (+bonuses) vs the Penetration Rating of the armor.
//				For every 2 points over (attack roll - defense roll) attacker receives a +1 bonus
//				Each weapon has performance criteria vs various armors (either bonus or penalty)
//				If the penetration roll fails, a hit is deducted from the armor durability.
//				If the penetration roll succeeds a hit is assessed against the defender.

// EXAMPLE ATTACK:
// Character swings at Monster's HEAD with MACE - hits!  Contested attack roll: 13 vs 9
// Penetration roll is: 13 Mace (+2) vs Leather Coif (12) - Mace Penetrates!
// Monster takes 1 hit!
//
// Character swings at Monster's HEAD with MACE - hits!
// Penetration roll is: 7 Mace (+2) vs Leather Coif (12) - Leather Coif Protects!
// Leather Coif takes 1 hit!
// Leather Coif is destroyed!!!

func getKeepSaveBlock() string {
	keepBlock := BLOCK_KEEP + ","

	keepBlock += keep.name + ","
	keepBlock += fmt.Sprintf("%v,", keep.acres)
	keepBlock += fmt.Sprintf("%v,", keep.usedacres)
	keepBlock += fmt.Sprintf("%v,", keep.descriptionId)
	keepBlock += fmt.Sprintf("%v,", keep.mapX)
	keepBlock += fmt.Sprintf("%v,", keep.mapY)
	keepBlock += fmt.Sprintf("%v,", keep.maxStorage)
	keepBlock += fmt.Sprintf("%v,", len(keep.storage))
	
	// add new lines for each storage item
	keepBlock += "◄"
	for k := 0; k < len(keep.storage); k++ {
		keepBlock += keep.storage[k].getSaveString()
	}
	
	keepBlock += "■"

	return keepBlock
}

func unpackKeepBlock(block string) (int, Keep) {
	var keep Keep

	lines := strings.Split(block, "◄")
	bits := strings.Split(lines[0], ",")
	lineCounter := 0
	storageCount := 0
	
	if bits[0] == BLOCK_KEEP {
		fmt.Println("Loading Keep Block...")
	} else {
		log.addError("Cant find Keep block.")
		fmt.Println("Keep Block not found!")
		return -1, keep
	}

	keep.name = bits[1]
	keep.acres, _ = strconv.Atoi(bits[2])

	keep.usedacres, _ = strconv.Atoi(bits[3])
	keep.descriptionId, _ = strconv.Atoi(bits[4])
	keep.mapX, _ = strconv.Atoi(bits[5])
	keep.mapY, _ = strconv.Atoi(bits[6])
	keep.maxStorage, _ = strconv.Atoi(bits[7])

	storageCount, _ = strconv.Atoi(bits[8])
	lineCounter = 1
	for k :=0 ; k < storageCount; k++ {
		itm, _ := restoreSavedItem(lines[lineCounter])
		keep.storage = append(keep.storage, itm)
		lineCounter++
	}

	fmt.Println("            ...done!")

	return 1, keep
}

func (keep *Keep) addApprenticeToKeep() {
	keep.apprentices = append(keep.apprentices, apprentice)
	apprentice.villageIndex = 99
}
func (keep *Keep) addNewApprenticeToKeep(app Character) {
	app.villageIndex = 99
	app.subLoc = 0
	keep.apprentices = append(keep.apprentices, app)
}

func (keep *Keep) getApprenticeList() ([]Character) {
	
	var apprentices = make([]Character, 0, 0)
	
	if apprentice.exists() {
		apprentices = append(apprentices, apprentice)
	}
	
	for k := 0; k < len(keep.apprentices); k++ {
		apprentices = append(apprentices, keep.apprentices[k])
	} 
	
	return apprentices
}

func (keep *Keep) rolodexViewApprentices() {
	rsp := ""
	
	apprentices := keep.getApprenticeList()
	
	if len(apprentices) < 1 {
		showPause("No Apprentices Available")
		return
	}

	index := 0	
	for rsp != "x" {
		apprentices[index].printCharacter(0)
		fmt.Println("")
		fmt.Println("(s. status) (i. inventory) (n. next) (x. exit)")
		fmt.Scanln(&rsp)
		
		if rsp == "n" && index + 1 < len(apprentices) {
			index++
		} else if rsp == "n" {
			index = 0
		} else if rsp == "i" {
			apprentices[index].showInventoryChar(false)
		} else if rsp == "s" {
			apprentices[index].showStatus()
		}
	}	
		
}

func (keep *Keep) manageApprentices() {
	rsp := ""
	
	for rsp != "x" {
		clearConsole()
		fmt.Println("╔ Manage Apprentices ╗")
		fmt.Println("")	
		fmt.Println("1. Select Companion")	
		fmt.Println("2. Assign Position")
		fmt.Println("3. Rolodex")	
		fmt.Println("")
		fmt.Println("x. Exit")
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)
		showPause(fmt.Sprintf("App instance id is %v ",  apprentice.instanceId))

		if rsp == "1" {
			ret := keep.selectApprentice()		
			if ret > -1 {
				if ret != 99 {	// 99 signifies current companion
					if apprentice.exists() {
						keep.addApprenticeToKeep()
						apprentice = keep.apprentices[ret]
						if len(keep.apprentices) > 1 {
							keep.apprentices = append(keep.apprentices[:ret], keep.apprentices[ret+1:]...)
						} else {
							keep.apprentices = make([]Character, 0, 0)
						}
					} else {	// no current companion, just assign
						apprentice = keep.apprentices[ret]
						if len(keep.apprentices) > 1 {
							keep.apprentices = append(keep.apprentices[:ret], keep.apprentices[ret+1:]...)
						} else {
							keep.apprentices = make([]Character, 0, 0)
						}
					}
				}
			}
		} else if rsp == "2" {
			// TODO: manage work in keep	
		} else if rsp == "3" {
			keep.rolodexViewApprentices()
		}	
	}
}
		
func (keep *Keep) selectApprentice() int {
	rsp := ""
	result := -1
	count := 0
	companion := false
	
	for rsp != "x" {
		clearConsole()
		fmt.Println("╔ Choose Apprentice ╗")
		fmt.Println("")		
		if apprentice.exists() {
			count++
			status := taskCodes[apprentice.task]
			row := packSpaceString(fmt.Sprintf("%v. %s  ", count, apprentice.name), 45) + status
			fmt.Println(row)
			companion = true
		}
		for k := 0; k < len(keep.apprentices); k++ {
			count++
			status := taskCodes[keep.apprentices[k].task]
			row := packSpaceString(fmt.Sprintf("%v. %s  ", count, keep.apprentices[k].name), 45) + status			
			fmt.Println(row)
		}
		
		if count == 0 {
			fmt.Println("No Apprentices Available")				
		}
		
		fmt.Println("")
		fmt.Println("x. Exit")
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)	
		
		idx := 0
		
		if rsp == "1" {
			if companion && apprentice.exists() {
				if apprentice.trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return 99				
				}
			} else if companion && !apprentice.exists() {
				result = -1
			} else if !companion {
				if len(keep.apprentices) > 0 {
					if keep.apprentices[0].trainingTime > 0 {
						showPause("Character is in training and not currently available.")
					} else {
						return 0						
					}
				} else {
					result = -1					
				}
			} else {
				showPause("Not a valid option.")				
			}
		} else if rsp == "2" {
			if companion {
				idx = 0
			} else {
				idx = 1
			}
			if len(keep.apprentices) > 1 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		} else if rsp == "3" {
			if companion {
				idx = 1
			} else {
				idx = 2
			}
			if len(keep.apprentices) > 2 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		} else if rsp == "4" {
			if companion {
				idx = 2
			} else {
				idx = 3
			}			
			if len(keep.apprentices) > 3 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		} else if rsp == "5" {
			if companion {
				idx = 3
			} else {
				idx = 4
			}			
			if len(keep.apprentices) > 4 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}

		} else if rsp == "6" {
			if companion {
				idx = 4
			} else {
				idx = 5
			}			
			if len(keep.apprentices) > 5 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		} else if rsp == "7" {
			if companion {
				idx = 5
			} else {
				idx = 6
			}			
			if len(keep.apprentices) > 6 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		} else if rsp == "8" {
			if companion {
				idx = 6
			} else {
				idx = 7
			}			
			if len(keep.apprentices) > 7 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		} else if rsp == "9" {
			if companion {
				idx = 7
			} else {
				idx = 8
			}			
			if len(keep.apprentices) > 8 {
				if keep.apprentices[idx].trainingTime > 0 {
					showPause("Character is in training and not currently available.")
				} else {
					return idx					
				}
			} else {
				showPause("Not a valid option.")		
			}
		}
	}
	
	return result
}

func (keep *Keep) train() {
	rsp := ""

	for rsp != "x" {
		clearConsole()
		fmt.Println("╔ Training ╗")	
		fmt.Println(fmt.Sprintf("1. Train Character  [%v pts avail]", character.trainingPoints))
		fmt.Println("2. Train Apprentice")
		fmt.Println("")
		fmt.Println("x. Exit")
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)
		
		if rsp == "1" {
			character.train()
			character.trainingTime = 0
		} else if rsp == "2" {
			appCode := keep.selectApprentice()
			
			if appCode == 99 {
				apprentice.train()
				
				if apprentice.trainingTime > 0 {
					keep.addApprenticeToKeep()
					var blankApp Character
					blankApp.instanceId = 0
					blankApp.name = ""
	
					apprentice = blankApp					
				}				

			} else if appCode > -1 {
				keep.apprentices[appCode].train()
			}
			
		} else if rsp != "x" {
			showPause("Invalid selection!")
		}
	}

}

func (keep *Keep) getFilteredStorageList(filters []bool) ([]Item){
	filteredList := make([]Item, 0, 0)
	
	for k := 0; k < len(keep.storage); k++ {
		itm := keep.storage[k]
		
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

func (keep *Keep) getKeepStorageIndex(itm Item) (int) {
	for k := 0; k < len(keep.storage); k++ {
		if itm.id == keep.storage[k].id {
			return k
		}
	}
	return -1	
}

func getDisplayFilters(filters []bool) (string){
	
	dispFilt := ""
	
	if filters[0] {
		dispFilt += "☼ "
	} else {
		dispFilt += "  "
	}
	
	if filters[1] {
		dispFilt += "⌂ "
	} else {
		dispFilt += "  "
	}
	
	if filters[2] {
		dispFilt += "♥ "
	} else {
		dispFilt += "  "
	}	
	
	if filters[3] {
		dispFilt += "♣ "
	} else {
		dispFilt += "  "
	}

	if filters[4] {
		dispFilt += "♦ "
	} else {
		dispFilt += "  "
	}	
	
	if filters[5] {
		dispFilt += "∞ "
	} else {
		dispFilt += "  "
	}
		
	return dispFilt
}

func setFilters(filters []bool) ([]bool) {
	
	rsp := ""
	
	for rsp != "x" {
		clearConsole()
		
		fmt.Println("  ─────────────────  Set Filters  ─────────────────")
		onOff := getOnOffString(filters[0])
		itmTxt := packSpaceString("  1.  Weapons", 28)
		itmTxt += filterIcons[0] +  "  "
		fmt.Println(itmTxt + "[" + onOff + "]")
		
		onOff = getOnOffString(filters[1])
		itmTxt = packSpaceString("  2.  Armor", 28)
		itmTxt += filterIcons[1] +  "  "
		fmt.Println(itmTxt + "[" + onOff + "]")	
		
		onOff = getOnOffString(filters[2])
		itmTxt = packSpaceString("  3.  Unctures", 28)
		itmTxt += filterIcons[2] +  "  "
		fmt.Println(itmTxt + "[" + onOff + "]")	
		
		onOff = getOnOffString(filters[3])
		itmTxt = packSpaceString("  4.  Ingredients", 28)
		itmTxt += filterIcons[3] +  "  "		
		fmt.Println(itmTxt + "[" + onOff + "]")			
		
		onOff = getOnOffString(filters[4])
		itmTxt = packSpaceString("  5.  Equipment", 28)
		itmTxt += filterIcons[4] +  "  "		
		fmt.Println(itmTxt + "[" + onOff + "]")			
		
		onOff = getOnOffString(filters[5])
		itmTxt = packSpaceString("  6.  Special", 28)
		itmTxt += filterIcons[5] +  "  "		
		fmt.Println(itmTxt + "[" + onOff + "]")			
		
		fmt.Println("")		
		choices := "[#. Toggle Item]  [x. Exit]"
		fmt.Println(choices)
		fmt.Println("")		
		fmt.Printf("Choose an option: ")

		fmt.Scanln(&rsp)	
	
		if rsp != "x" {
			num, err := strconv.Atoi(rsp)

			if err == nil && num < 7 && num > 0 {
				num -= 1  // index is 0 based but options start at 1
				
				if filters[num] {
					filters[num] = false
				} else {
					filters[num] = true				
				}
			}
		}
	}
	
	return filters
}

func (keep *Keep) showStorage() {
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
		
		filteredList := keep.getFilteredStorageList(filters)
		dispFilters = getDisplayFilters(filters)
		
		clearConsole()
		fmt.Println("╔═══════════════════════ Storage ═══════════════════════╗")		
		
		pages = 1
		if len(keep.storage) > ITEMS_PER_PAGE {
			for j := len(filteredList); j > ITEMS_PER_PAGE; j -= ITEMS_PER_PAGE {
				pages++
			} 
		}
		
		dispStr := packSpaceStringCenter(fmt.Sprintf("  %v / %v ", len(keep.storage), keep.maxStorage), 24)
		dispStr += "           Filters: " + dispFilters
		
		fmt.Println(dispStr)
		fmt.Println("  ─────────────────────           ─────────────────────")
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
		commands += "  [f. filters]  [#. Take]  [x. Exit]"
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
				
				tgt := giveToWho()
				
				if tgt == 0 {
					if character.giveCharacterItem(storeItem) {
						showPause(fmt.Sprintf("%s given to %s!", storeItem.name, character.name))	
						idx := keep.getKeepStorageIndex(storeItem)						
						keep.storage = append(keep.storage[:idx], keep.storage[idx+1:]...)
					} else {
						showPause(character.name + " cannot hold this item (over-encumbered).")
					}
				} else {
					if apprentice.giveCharacterItem(storeItem) {
						showPause(fmt.Sprintf("%s given to %s!", storeItem.name, apprentice.name))		
						idx := keep.getKeepStorageIndex(storeItem)					
						keep.storage = append(keep.storage[:idx], keep.storage[idx+1:]...)						
					} else {
						showPause(apprentice.name + " cannot hold this item (over-encumbered).")					
					}					
				}				
			} else {
				showPause("Invalid selection.")
			}
		}				
	}
}

func (keep *Keep) endDay() (string) {
	var die Die
	updates := ""
	
	for k := 0; k < len(keep.apprentices); k++ {
		if die.rollxdx(1, 4) <= 2 {
			if keep.apprentices[k].hp < keep.apprentices[k].maxhp {
				keep.apprentices[k].hp++
			}
		}
		if keep.apprentices[k].trainingTime > 0 {
			keep.apprentices[k].trainingTime -= 1
			
			if keep.apprentices[k].trainingTime == 0 {
				keep.apprentices[k].task = STATUS_REST
				updates += keep.apprentices[k].name + " has completed training. \n"
			}			
		} else if keep.apprentices[k].task == STATUS_TRAINING && keep.apprentices[k].trainingTime == 0 {
			keep.apprentices[k].task = STATUS_REST
		}
	}
	
	return updates
}

func (keep *Keep) visitKeep() string {
	rsp := ""

	apprentice.villageIndex = 99

	for rsp != "q" {
		clearConsole()
		fmt.Println("╔ Keep ╗")
		fmt.Println(makeDialogString(keepDescriptions[keep.descriptionId]))
		fmt.Println("")
		fmt.Printf("Day: %v \n", game.gameDay)
		fmt.Printf("Acres: %v / %v \n", keep.usedacres, keep.acres)
		fmt.Println("------------")

		fmt.Println("1. Structures")
		fmt.Println("2. Apprentices")
		fmt.Println("3. Storage")
		fmt.Println("4. Train")
		fmt.Println("")
		fmt.Println("r. Rest (End Day)")		
		fmt.Println("t. Travel")
		fmt.Println("q. Exit")
		fmt.Println("")
		fmt.Println(BASE_ACTIONS)
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)

		//const BASE_ACTIONS = "[s. status   i. inventory   m. mission   w. world map   h. minutiae]"
		
		if rsp == "r" {
			endDay(2, true)
			save()
		} else if rsp == "s" {
			character.printCharacter(1)
			character.showStatus()
			if apprentice.exists() {
				apprentice.printCharacter(1)
				apprentice.showStatus()
			}
		} else if rsp == "m" {	
			mission.viewMissionStatus()
		} else if rsp == "i" {	
			character.showInventory()
		} else if rsp == "w" {	
			drawWorldMap()	
		} else if rsp == "h" {	
			// show keep minutiae		
		} else if rsp == "t" {
			travel := showTravelMenu()
			return travel
		} else if rsp == "1" {	
			// TODO
		} else if rsp == "2" {	
			keep.manageApprentices()
		} else if rsp == "3" {	
			keep.showStorage() 		
		} else if rsp == "4" {	
			keep.train()			
		} 
	}

	return rsp
}

func createKeep() Keep {
	var keep Keep

	keep.acres = 0
	keep.usedacres = 0
	keep.name = "Campground"
	keep.descriptionId = 0
	keep.mapX, keep.mapY = 23, 12
	keep.maxStorage = 50
	
	return keep
}
