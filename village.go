// village.go
// Village

package main

import "fmt"

type Village struct {
	name    					string
	distanceToKeep 				int
	size 						int
}

func buildVillages() {
	villages = make(Village, 9, 9)
	
	var crowley Village	
	crowley.name = "Crowley"
	crowley.distanceToKeep = 1
	crowley.size = 1
	villages[0] = crowley	
		
	var pritchard Village	
	pritchard.name = "Pritchard"
	pritchard.distanceToKeep = 1
	pritchard.size = 6
	villages[1] = pritchard
	
	var maline Village	
	maline.name = "Maline"
	maline.distanceToKeep = 2
	maline.size = 3
	villages[2] = maline 
	
	var alistaire Village	
	alistaire.name = "Alistair"
	alistaire.distanceToKeep = 2
	alistaire.size = 3
	villages[3] = alistaire 
	
	var dauntun Village	
	dauntun.name = "Dauntun"
	dauntun.distanceToKeep = 2
	dauntun.size = 2
	villages[4] = dauntun 	
	
	var elice Village	
	elice.name = "Elice"
	elice.distanceToKeep = 3
	elice.size = 2
	villages[5] = elice 	
	
	var hollow Village	
	hollow.name = "Hollow"
	hollow.distanceToKeep = 13
	hollow.size = 2
	villages[4] = hollow 	
	
	var hastur Village	
	hastur.name = "Hastur"
	hastur.distanceToKeep = 4
	hastur.size = 4
	villages[4] = hastur 	
}

func (village * Village) visitVillage(){
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

	fmt.Println("x. Back")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	return rsp
}

func (village * Village) goShop() {

	var result string = ""

	for result != "9" {
		result = village.visitVillage()

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