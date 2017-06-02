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

func showWeapon(weapon Item) {
	clearConsole()

	fmt.Println(packSpaceString(weapon.name, 30) + "Value: " + packSpace(weapon.value, 6))
	fmt.Println("-------------")
	row := ""
	row = packSpaceString("Material: "+weapon.material, 30)
	row += "Quality: " + weapon.quality
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Durability: %v / %v", weapon.durability, weapon.maxDurability), 30)
	row += packSpaceString(fmt.Sprintf("Weight: %v ", weapon.weight), 20)
	row += packSpaceString(fmt.Sprintf("Hands: %v ", weapon.hands), 12)
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Attack Turns: %v ", weapon.atkTurns), 30)
	row += fmt.Sprintf("Attack Range: %v ", weapon.wRange)
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Accuracy: %v ", weapon.accuracy), 30)
	row += fmt.Sprintf("Defense: %v ", weapon.defense)
	fmt.Println(row)
	fmt.Println("")

	row = "" //paddedMod, leatherMod, chainMod
	txt := "Penetration:\n [vs Padded: %v]    [vs Leather: %v]    [vs Chain: %v]"
	row = fmt.Sprintf(txt, getSigned(weapon.paddedMod), getSigned(weapon.leatherMod), getSigned(weapon.chainMod))
	fmt.Println(row)
	fmt.Println("")
	fmt.Println("")

}

func showArmor(armor Item) {
	clearConsole()

	fmt.Println(packSpaceString(armor.name, 30) + "Value: " + packSpace(armor.value, 6))
	fmt.Println("-------------")
	row := packSpaceString("Equips: "+equipStrings[armor.equip], 28)
	row += "Quality: " + armor.quality
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Shields: %v / %v", armor.durability, armor.maxDurability), 30)
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Weight: %v ", armor.weight), 20)
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Defense: %v ", armor.defense), 30)
	fmt.Println(row)
	fmt.Println("")

	row = ""
	row = packSpaceString(fmt.Sprintf("Resistance: %v ", armor.resistance), 30)
	fmt.Println(row)
	fmt.Println("")
	fmt.Println("")

}

func giveToWho() int {
	if apprentice.instanceId < 1 {
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
			apprentice.showInventory()
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
		fmt.Println("[x. Back]   [n. More]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		rsp := ""
		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" && rsp != "i" {
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
			if apprentice.instanceId > 0 {
				apprentice.showInventory()
			}
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
		fmt.Println("[x. Back]   [n. More]   [i. inventory]")
		fmt.Println("")
		fmt.Println("Select an Option:  ")

		rsp := ""
		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" && rsp != "i" {
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
			if apprentice.instanceId > 0 {
				apprentice.showInventory()
			}
		}
	}

	village.shopArmor = shopArmor
}

func (village *Village) buyProvisions() {

	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Provisions      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Nothing available")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" {

		} else if rsp == "x" {
			exitFlag = true
		}
	}
}

func (village *Village) buyApothecary() {
	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Curiosities      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Nothing available")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" {

		} else if rsp == "x" {
			exitFlag = true
		}
	}
}

func (village *Village) buyCuriosities() {
	exitFlag := false

	for !exitFlag {
		clearConsole()
		rsp := ""

		charString := fmt.Sprintf("%v    Encumb: %v : %v", character.crowns, convertPoundsToStone(character.weight), convertPoundsToStone(character.maxweight))

		fmt.Println("Curiosities      Crowns:  " + charString)
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("Nothing available")

		fmt.Scanln(&rsp)

		if len(rsp) > 0 && rsp != "x" && rsp != "n" {

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
