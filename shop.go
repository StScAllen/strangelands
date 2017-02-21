// shop.go

package main

import "fmt"


func buyWeaponScreen(){
	clearConsole()
	
	fmt.Println("Weapon Shop")
	fmt.Println("----Gold:  " + fmt.Sprintf("%v", character.gold) + "  ------")

	fmt.Println("Title            Dmg   \tAcc\tDef\tCost\tQuality\tMaterial")
	for i := 0; i < len(weapons); i++ {	
		fmt.Println(fmt.Sprintf("%v. %s    \t %v-%v   \t%v\t%v   \t%v \t%s\t%s", 
								  i, weapons[i].name, weapons[i].dmgMin, weapons[i].dmgMax, weapons[i].accuracy, weapons[i].defense, 
								  weapons[i].value, weapons[i].quality, weapons[i].material))
	}
	
	fmt.Println("")	
	fmt.Println("X. Back")
	fmt.Println("")
	fmt.Println("Select an Option:  ")
	
	rsp := ""
	fmt.Scanln(&rsp)

}

func buyArmorScreen(){



}

func buySuppliesScreen(){



}

func buyAnimalsScreen(){



}

func showShopMenu() (string){

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

		if (result == "1"){
			buyWeaponScreen()
		} else if (result == "2"){
			buyArmorScreen()
		} else if (result == "3"){
			buySuppliesScreen()
		} else if (result == "4"){
			buyAnimalsScreen()
		}
	}

}