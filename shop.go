// shop.go

package main

import "fmt"
import "strconv"

var shopWeapons []Item

func genWeaponsOfWeek() {
	var die Die

	shopWeapons = make([]Item, 0)

	for k := 0; k < die.rollxdx(5, 12); k++ {
		shopWeapons = append(shopWeapons, getRandomWeapon())
	}
}

func updateShops() {
	genWeaponsOfWeek()
}

func packSpace(num int, digits int) string {
	ret := fmt.Sprintf("%v", num)

	for len(ret) < digits {
		ret += " "
	}

	return ret
}

func packSpaceString(str string, digits int) string {
	for len(str) < digits {
		str += " "
	}

	return str
}

func buyWeaponScreen() {

	clearConsole()

	charString := fmt.Sprintf("%v  Encumb: %v / %v", character.gold, character.weight, character.maxweight)

	fmt.Println("Weapon Shop      Gold:  " + charString)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("")

	fmt.Println("Title            Dmg   \tAcc  Def  Wgt\tCost\tQuality\t\tMaterial")

	for i := 0; i < len(shopWeapons); i++ {
		fmt.Println(fmt.Sprintf("%v. %s \t %v-%v   \t%s %s %s  \t%v \t%s\t\t%s",
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
				buyWeaponScreen()
			} else {
				item := shopWeapons[num]
				if character.giveCharacterItem(item) {
					character.gold -= item.value
					shopWeapons = append(shopWeapons[:num], shopWeapons[num+1:]...)
				} else {
					showPause("Character weight exceeded! Purchase not made!")
				}
				buyWeaponScreen()
			}
		}
	}
}

func buyArmorScreen() {

}

func buySuppliesScreen() {

}

func buyAnimalsScreen() {

}

func showShopMenu() string {

	clearConsole()

	fmt.Println("Town Menu")
	fmt.Println("------------")
	fmt.Println("1. Shop Weapons")
	fmt.Println("2. Shop Armor")
	fmt.Println("3. Shop Equipment")
	fmt.Println("4. Shop Curiosities (Recipes & Spells)")
	fmt.Println("5. Research Quest")
	fmt.Println("6. Visit Tavern")
	fmt.Println("7. Politicks - Curry Favor / Influence")
	fmt.Println("8. View/Accept Missions")
	fmt.Println("")

	fmt.Println("9. Back")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	return rsp
}

func goShop() {

	var result string = ""

	for result != "9" {
		result = showShopMenu()

		if result == "1" {
			buyWeaponScreen()
		} else if result == "2" {
			buyArmorScreen()
		} else if result == "3" {
			buySuppliesScreen()
		} else if result == "4" {
			buyAnimalsScreen()
		}
	}

}
