// shop.go

package main

import "fmt"
import "strconv"

func genWeaponsOfWeek(size int) []Item {
	var die Die

	shopWeapons := make([]Item, 0)
	
	for k := 0; k < die.rollxdx(4 + size, 6 + (size*2)); k++ {
		shopWeapons = append(shopWeapons, getRandomWeapon(size))
	}

	return shopWeapons
}

func genArmorOfWeek(size int) []Item {
	var die Die

	shopArmor := make([]Item, 0)

	for k := 0; k < die.rollxdx(4 + size, 6 + (size*2)); k++ {
		shopArmor = append(shopArmor, getRandomArmor(size))
	}

	return shopArmor
}

func updateShops() {
	for i := range villages {
		villages[i].shopWeapons = genWeaponsOfWeek(villages[i].size)
		villages[i].shopArmor = genArmorOfWeek(villages[i].size)
	}
}

func giveToWho() int {
	if !apprentice.exists() {
		return 0	// character
	}

	exitFlag := false

	for !exitFlag {
		clearConsole()
		fmt.Println(fmt.Sprintf("1. %s ", character.name))
		fmt.Println(fmt.Sprintf("2. %s (Apprentice)", apprentice.name))
		fmt.Println("")		
		fmt.Println("[i. inventory]")
		fmt.Println("")	
		fmt.Println("Show should receive the item? ")	

		rsp := ""
		fmt.Scanln(&rsp)
		
		if rsp == "1" {
			return 0
		} else if rsp == "2" {
			return 1
		} else if rsp == "i" {
			character.showInventory()
		}
	}
	
	return 0
}

func (village *Village) buyWeaponScreen() {
	shopWeapons := village.shopWeapons
	exitFlag := false

	for !exitFlag {
		clearConsole()

		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Weapon Shop      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("")

		fmt.Println("   Item         Dmg   \tAcc  Def  Wgt\tCost\tQuality\t\tMaterial")

		for i := 0; i < len(shopWeapons); i++ {
			fmt.Println(fmt.Sprintf("%v. %s \t %v-%v   \t%s %s %s\t%v \t%s\t\t%s",
				i, shopWeapons[i].name, shopWeapons[i].dmgMin, shopWeapons[i].dmgMax, packSpace(shopWeapons[i].accuracy, 4), packSpace(shopWeapons[i].defense, 4), packSpace(shopWeapons[i].weight, 4), shopWeapons[i].value, shopWeapons[i].quality, shopWeapons[i].material))
		}

		fmt.Println("")
		fmt.Println("[x. Back]   [n. More]   [s. Sell]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		rsp := ""
		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" && rsp != "i" && rsp != "s"  {
			num, _ := strconv.Atoi(rsp)
			showWeapon(shopWeapons[num])
			fmt.Println(fmt.Sprintf("Buy %s? ", shopWeapons[num].name))
			rsp2 := ""
			fmt.Scanln(&rsp2)

			if rsp2 == "y" {
				if character.crowns < shopWeapons[num].value {
					showPause("Not enough crowns!")
					village.buyWeaponScreen()
				} else {
					item := shopWeapons[num]					
					ret := giveToWho()
					
					if ret == 0 {
						if character.giveCharacterItem(item) {
							character.crowns -= item.value
							shopWeapons = append(shopWeapons[:num], shopWeapons[num+1:]...)
							village.shopWeapons = shopWeapons
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					} else if ret == 1 {
						if apprentice.giveCharacterItem(item) {
							character.crowns -= item.value
							shopWeapons = append(shopWeapons[:num], shopWeapons[num+1:]...)
							village.shopWeapons = shopWeapons
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					}
				}
			}
		} else if rsp == "x" {
			exitFlag = true
		} else if rsp == "i" {
			character.showInventory()
		} else if rsp == "s" {
			village.sellItems(ITEM_TYPE_WEAPON)			
		}
	}

	village.shopWeapons = shopWeapons
}

func (village *Village) buyArmorScreen() {
	shopArmor := village.shopArmor

	exitFlag := false

	for !exitFlag {
		clearConsole()
		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Armor Shop      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("")

		fmt.Println("   Item                    Resist Def\tShields\tWgt \tCost\tQuality")

		for i := 0; i < len(shopArmor); i++ {
			fmt.Printf("%v. %s %s%s %s\t%s \t%s\t%s \n", i, packSpaceString(shopArmor[i].name, 24), packSpace(shopArmor[i].resistance, 7), packSpace(shopArmor[i].defense, 4), packSpace(shopArmor[i].durability, 4), packSpace(shopArmor[i].weight, 4), packSpace(shopArmor[i].value, 4), shopArmor[i].quality)
		}

		fmt.Println("")
		fmt.Println("[x. Back]   [n. More]   [s. sell]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		rsp := ""
		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" && rsp != "i" && rsp != "s"  {
			num, _ := strconv.Atoi(rsp)
			showArmor(shopArmor[num])
			fmt.Println(fmt.Sprintf("Buy %s?", shopArmor[num].name))
			fmt.Scanln(&rsp)

			if rsp == "y" {
				if character.crowns < shopArmor[num].value {
					showPause("Not enough crowns!")
				} else {
					item := shopArmor[num]
					ret := giveToWho()
					
					if ret == 0 {
						if character.giveCharacterItem(item) {
							character.crowns -= item.value
							shopArmor = append(shopArmor[:num], shopArmor[num+1:]...)
							village.shopArmor = shopArmor
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					} else if ret == 1 {
						if apprentice.giveCharacterItem(item) {
							character.crowns -= item.value
							shopArmor = append(shopArmor[:num], shopArmor[num+1:]...)
							village.shopArmor = shopArmor
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					}

				}
			}
		} else if rsp == "x" {
			exitFlag = true
		}  else if rsp == "i" {
			character.showInventory()
		} else if rsp == "s" {
			village.sellItems(ITEM_TYPE_ARMOR)					
		}
	}

	village.shopArmor = shopArmor
}

func (village *Village) buyProvisions() {

	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		provisions := getAllProvisions()
		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Provisions      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")

		fmt.Println("   Item                    \tWgt \tCost")

		for i := 0; i < len(provisions); i++ {
			fmt.Printf("%v. %s \t%v \t%v \n", i, packSpaceString(provisions[i].name, 24), provisions[i].weight, provisions[i].value)
		}

		fmt.Println("")
		fmt.Println("[x. Back]   [n. More]   [s. sell]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n"  && rsp != "i" && rsp != "s" {
			num, _ := strconv.Atoi(rsp)
			// showProvision()
			fmt.Println(fmt.Sprintf("Buy %s?", provisions[num].name))
			fmt.Scanln(&rsp)
			
			if rsp == "y" {
				if character.crowns < provisions[num].value {
					showPause("Not enough crowns!")
				} else {
					item := genGeneralItem(provisions[num])
					ret := giveToWho()
					
					if ret == 0 {
						if character.giveCharacterItem(item) {
							character.crowns -= item.value
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					} else if ret == 1 {
						if apprentice.giveCharacterItem(item) {
							character.crowns -= item.value
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					}
				}
			}
			
		} else if rsp == "x" {
			exitFlag = true
		}  else if rsp == "i" {
			character.showInventory()
		} else if rsp == "s" {
			village.sellItems(ITEM_TYPE_EQUIPMENT)					
		}
	}
}

func (village *Village) buyApothecary() {
	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		ingredients := getAllApothecary()
		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Apothecary      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")

		fmt.Println("   Item                    \tWgt \tCost")

		for i := 0; i < len(ingredients); i++ {
			fmt.Printf("%v. %s \t%v \t%v \n", i, packSpaceString(ingredients[i].name, 24), ingredients[i].weight, ingredients[i].value)
		}

		fmt.Println("")
		fmt.Println("[x. Back]   [n. More]   [s. sell]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n"  && rsp != "i" && rsp != "s"{
			num, _ := strconv.Atoi(rsp)
			// showIngredient()
			fmt.Println(fmt.Sprintf("Buy %s?", ingredients[num].name))
			fmt.Scanln(&rsp)
			
			if rsp == "y" {
				if character.crowns < ingredients[num].value {
					showPause("Not enough crowns!")
				} else {
					item := genGeneralItem(ingredients[num])
					ret := giveToWho()
					
					if ret == 0 {
						if character.giveCharacterItem(item) {
							character.crowns -= item.value
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					} else if ret == 1 {
						if apprentice.giveCharacterItem(item) {
							character.crowns -= item.value
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					}
				}
			}
			
		} else if rsp == "x" {
			exitFlag = true
		}  else if rsp == "i" {
			character.showInventory()
		} else if rsp == "s" {
			village.sellItems(ITEM_TYPE_INGREDIENT)					
		}
	}
}

func (village *Village) buyCuriosities() {
	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		curios := getAllCuriosities()
		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Curiosities      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")

		fmt.Println("   Item                    \t\tWgt \tCost")

		for i := 0; i < len(curios); i++ {
			fmt.Printf("%v. %s \t%v \t%v \n", i, packSpaceString(curios[i].name, 28), curios[i].weight, curios[i].value)
		}

		fmt.Println("")
		fmt.Println("[x. Back]   [n. More]   [s. sell]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n"  && rsp != "i" && rsp != "s"{
			num, _ := strconv.Atoi(rsp)
			// showCurio()
			fmt.Println(fmt.Sprintf("Buy %s?", curios[num].name))
			fmt.Scanln(&rsp)
			
			if rsp == "y" {
				if character.crowns < curios[num].value {
					showPause("Not enough crowns!")
				} else {
					item := genGeneralItem(curios[num])
					ret := giveToWho()
					
					if ret == 0 {
						if character.giveCharacterItem(item) {
							character.crowns -= item.value
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					} else if ret == 1 {
						if apprentice.giveCharacterItem(item) {
							character.crowns -= item.value
						} else {
							showPause("Character weight exceeded! Purchase not made!")
						}					
					}
				}
			}
			
		} else if rsp == "x" {
			exitFlag = true
		}  else if rsp == "i" {
			character.showInventory()
		} else if rsp == "s" {
			village.sellItems(ITEM_TYPE_SPECIAL)					
		}
	}
}

func (village *Village) sellItems(itmType int) {
	if !apprentice.exists() {
		character.sellItems(itmType)
	}

	exitFlag := false

	for !exitFlag {
		clearConsole()
		fmt.Println(fmt.Sprintf("1. %s ", character.name))
		fmt.Println(fmt.Sprintf("2. %s (Apprentice)", apprentice.name))
		fmt.Println("")	
		fmt.Println("[x. Exit]")
		fmt.Println("")
		fmt.Println("Which inventory do you wish to sell from? ")	

		rsp := ""
		fmt.Scanln(&rsp)
		
		if rsp == "1" {
			character.sellItems(itmType)
		} else if rsp == "2" {
			apprentice.sellItems(itmType)
		} else if rsp == "x" {	
			exitFlag = true		
		} 
	}
}

func (village *Village) shopMenu() {
	exitFlag := false
	rsp := ""

	for exitFlag != true {
		clearConsole()

		fmt.Println("+++ Shops of " + village.name + " +++")
		fmt.Println("------------")
		fmt.Println("1. Weapons")
		fmt.Println("2. Armor")
		fmt.Println("3. Provisions")
		fmt.Println("4. Apothecary")
		fmt.Println("5. Curiosities")

		fmt.Println("x. Exit")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		fmt.Scanln(&rsp)

		if rsp == "x" || rsp == "X" {
			exitFlag = true
		} else if rsp == "1" {
			village.buyWeaponScreen()
		} else if rsp == "2" {
			village.buyArmorScreen()
		} else if rsp == "3" {
			village.buyProvisions()
		} else if rsp == "4" {
			village.buyApothecary()
		} else if rsp == "5" {
			village.buyCuriosities()
		}
	}
}
