// shop.go

package main

import "fmt"
import "strconv"

func genWeaponsOfWeek() []Item{
	var die Die

	shopWeapons := make([]Item, 0)

	for k := 0; k < die.rollxdx(5, 12); k++ {
		shopWeapons = append(shopWeapons, getRandomWeapon())
	}
	
	return shopWeapons

}

func genArmorOfWeek() []Item {
	var die Die

	shopArmor := make([]Item, 0)

	for k := 0; k < die.rollxdx(5, 12); k++ {
		shopArmor = append(shopArmor, getRandomArmor())
	}
	
	return shopArmor
}

func updateShops() {
	for i := range villages {
		villages[i].shopWeapons = genWeaponsOfWeek()
		villages[i].shopArmor = genArmorOfWeek()
	}
}

func (village * Village) buyWeaponScreen() {
	clearConsole()
	
	shopWeapons := village.shopWeapons
	
	charString := fmt.Sprintf("%v  Encumb: %v / %v", character.gold, character.weight, character.maxweight)

	fmt.Println("Weapon Shop      Gold:  " + charString)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("")

	fmt.Println("   Item         Dmg   \tAcc  Def  Wgt\tCost\tQuality\t\tMaterial")

	for i := 0; i < len(shopWeapons); i++ {
		fmt.Println(fmt.Sprintf("%v. %s \t %v-%v   \t%s %s %s\t%v \t%s\t\t%s",
			i, shopWeapons[i].name, shopWeapons[i].dmgMin, shopWeapons[i].dmgMax, packSpace(shopWeapons[i].accuracy, 4), packSpace(shopWeapons[i].defense, 4), packSpace(shopWeapons[i].weight, 4), shopWeapons[i].value, shopWeapons[i].quality, shopWeapons[i].material))
	}

	fmt.Println("")
	fmt.Println("[x. Back]   [n. More]")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	if len(rsp) > 0 && rsp != "x" && rsp != "n" {
		num, err := strconv.Atoi(rsp)
		fmt.Println(fmt.Sprintf("Buy %s? %v", shopWeapons[num].name, err))
		fmt.Scanln(&rsp)

		if rsp == "y" {
			if character.gold < shopWeapons[num].value {
				showPause("Not enough gold!")
				village.buyWeaponScreen()
			} else {
				item := shopWeapons[num]
				if character.giveCharacterItem(item) {
					character.gold -= item.value
					shopWeapons = append(shopWeapons[:num], shopWeapons[num+1:]...)
					village.shopWeapons = shopWeapons
				} else {
					showPause("Character weight exceeded! Purchase not made!")
				}
				village.buyWeaponScreen()
			}
		}
	}
	
	village.shopWeapons = shopWeapons
}

func (village * Village) buyArmorScreen() {
	clearConsole()

	shopArmor := village.shopArmor
	
	charString := fmt.Sprintf("%v  Encumb: %v / %v", character.gold, character.weight, character.maxweight)

	fmt.Println("Armor Shop      Gold:  " + charString)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("")

	fmt.Println("   Item                    Def\tShields\tWgt \tCost\tQuality")

	for i := 0; i < len(shopArmor); i++ {
		fmt.Printf("%v. %s %s%s\t%s \t%s\t%s \n", i, packSpaceString(shopArmor[i].name, 24), packSpace(shopArmor[i].defense, 4), packSpace(shopArmor[i].durability, 4), packSpace(shopArmor[i].weight, 4), packSpace(shopArmor[i].value, 4), shopArmor[i].quality)
	}

	fmt.Println("")
	fmt.Println("[x. Back]   [n. More]")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	if len(rsp) > 0 && rsp != "x" && rsp != "n" {
		num, err := strconv.Atoi(rsp)
		fmt.Println(fmt.Sprintf("Buy %s? %v", shopArmor[num].name, err))
		fmt.Scanln(&rsp)

		if rsp == "y" {
			if character.gold < shopArmor[num].value {
				showPause("Not enough gold!")
				village.buyArmorScreen()
			} else {
				item := shopArmor[num]
				if character.giveCharacterItem(item) {
					character.gold -= item.value
					shopArmor = append(shopArmor[:num], shopArmor[num+1:]...)
					village.shopArmor = shopArmor
				} else {
					showPause("Character weight exceeded! Purchase not made!")
				}
				village.buyArmorScreen()
			}
		}
	}
	
	village.shopArmor = shopArmor
}

func (village * Village) buySuppliesScreen() {

}

func (village * Village) buyAnimalsScreen() {

}


