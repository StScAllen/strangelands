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

	// add new lines
	//	keepBlock += "◄"
	//	keepBlock += fmt.Sprintf("%v,", -1)

	keepBlock += "■"

	return keepBlock
}

func unpackKeepBlock(block string) (int, Keep) {
	var keep Keep

	lines := strings.Split(block, "◄")
	bits := strings.Split(lines[0], ",")

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

	fmt.Println("            ...done!")

	return 1, keep
}

func (keep *Keep) addApprenticeToKeep() {
	keep.apprentices = append(keep.apprentices, apprentice)
}
func (keep *Keep) addNewApprenticeToKeep(app Character) {
	keep.apprentices = append(keep.apprentices, app)
}

func (keep *Keep) manageApprentices() {
	rsp := ""
	
	for rsp != "x" {
		clearConsole()
		fmt.Println("╔ Manage Apprentices ╗")
		fmt.Println("")	
		fmt.Println("1. Select Companion")	
		fmt.Println("2. Assign Position")	
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
					if apprentice.instanceId > 0 {
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
		if apprentice.instanceId > 0 {
			count++
			row := fmt.Sprintf("%v. %s", count, apprentice.name)
			fmt.Println(row)
			companion = true
		}
		for k := 0; k < len(keep.apprentices); k++ {
			count++
			row := fmt.Sprintf("%v. %s", count, keep.apprentices[k].name)
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
		
		if rsp == "1" {
			if companion && apprentice.instanceId > 0 {
				return 99
			} else if companion && apprentice.instanceId < 1 {
				result = -1
			} else if !companion {
				if len(keep.apprentices) > 0 {
					return 0
				} else {
					result = -1					
				}
			}
			showPause("Not a valid option.")
		} else if rsp == "2" {
			if len(keep.apprentices) > 1 {
				return 1	
			}
			showPause("Not a valid option.")	
		} else if rsp == "3" {
			if len(keep.apprentices) > 2 {
				return 2	
			}
			showPause("Not a valid option.")	
		} else if rsp == "4" {
			if len(keep.apprentices) > 3 {
				return 3	
			}
			showPause("Not a valid option.")
		} else if rsp == "5" {
			if len(keep.apprentices) > 4 {
				return 4	
			}
			showPause("Not a valid option.")

		} else if rsp == "6" {
			if len(keep.apprentices) > 5 {
				return 5	
			}
			showPause("Not a valid option.")	
		} else if rsp == "7" {
			if len(keep.apprentices) > 6 {
				return 6	
			}
			showPause("Not a valid option.")	
		} else if rsp == "8" {
			if len(keep.apprentices) > 7 {
				return 7	
			}
			showPause("Not a valid option.")	
		} else if rsp == "8" {
			if len(keep.apprentices) > 7 {
				return 7	
			}
			showPause("Not a valid option.")	
		}
	}
	
	return result
}

func (keep *Keep) train() {
	rsp := ""

	for rsp != "x" {
		clearConsole()
		fmt.Println("╔ Training ╗")	
		fmt.Println("1. Train Character")
		fmt.Println("2. Train Apprentice")
		fmt.Println("")
		fmt.Println("x. Exit")
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)
		
		if rsp == "1" {
			character.train()
		} else if rsp == "2" {
			appCode := keep.selectApprentice()
			
			if appCode == 99 {
				keep.addApprenticeToKeep()
				var blankApp Character
				blankApp.instanceId = 0
				blankApp.name = ""
				apprentice = blankApp
				apprentice.train()
			} else if appCode > -1 {
				keep.apprentices[appCode].train()
			}
		} else if rsp != "x" {
			showPause("Invalid selection!")
		}
	}

}

func (keep *Keep) visitKeep() string {
	rsp := ""

	for rsp != "q" {
		clearConsole()
		fmt.Println("╔ Keep ╗")
		fmt.Println(makeDialogString(keepDescriptions[keep.descriptionId]))
		fmt.Printf("Day: %v \n", game.gameDay)
		fmt.Printf("Acres: %v / %v \n", keep.usedacres, keep.acres)
		fmt.Println("------------")

		fmt.Println("1. Structures")
		fmt.Println("2. Apprentices")
		fmt.Println("3. Keep Storage")
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
			if apprentice.instanceId > 0 {
				apprentice.printCharacter(1)
				apprentice.showStatus()
			}
		} else if rsp == "m" {	
			mission.viewMissionStatus()
		} else if rsp == "i" {	
			character.showInventory()
			if apprentice.instanceId > 0 {
				apprentice.showInventory()
			}
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
			// TODO		
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

	return keep
}
