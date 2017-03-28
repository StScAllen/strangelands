// village.go
// Village

package main

import "fmt"

type Village struct {
	name    					string
	distanceToKeep 				int
	size 						int
	shopWeapons 				[]Item
	shopArmor 					[]Item
	mapX, mapY					int
}

func buildVillages() {
	villages = make([]Village, 8, 8)
	
	var crowley Village	
	crowley.name = "Crowley"
	crowley.distanceToKeep = 1
	crowley.size = 1
	crowley.mapX, crowley.mapY = 20, 10
	villages[0] = crowley	
		
	var pritchard Village	
	pritchard.name = "Pritchard"
	pritchard.distanceToKeep = 1
	pritchard.size = 6
	pritchard.mapX, pritchard.mapY = 25, 10
	villages[1] = pritchard
	
	var maline Village	
	maline.name = "Maline"
	maline.distanceToKeep = 2
	maline.size = 3
	maline.mapX, maline.mapY = 17, 3
	villages[2] = maline 
	
	var faust Village	
	faust.name = "Faust"
	faust.distanceToKeep = 2
	faust.size = 3
	faust.mapX, faust.mapY = 38, 3
	villages[3] = faust 
	
	var dauntun Village	
	dauntun.name = "Dauntun"
	dauntun.distanceToKeep = 2
	dauntun.size = 2
	dauntun.mapX, dauntun.mapY = 25, 16
	villages[4] = dauntun 	
	
	var elice Village	
	elice.name = "Elice"
	elice.distanceToKeep = 3
	elice.size = 2
	elice.mapX, elice.mapY = 46, 16
	villages[5] = elice 	
	
	var hollow Village	
	hollow.name = "Hollow"
	hollow.distanceToKeep = 13
	hollow.size = 2
	hollow.mapX, hollow.mapY = 2, 13
	villages[6] = hollow 	
	
	var hastur Village	
	hastur.name = "Hastur"
	hastur.distanceToKeep = 4
	hastur.size = 4
	hastur.mapX, hastur.mapY = 46, 2
	villages[7] = hastur 	

}

func (village * Village) visitVillage() (string) {
	clearConsole()

	fmt.Println("+ Village of" + village.name + "+")
	fmt.Println("------------")
	fmt.Println("1. Shop Weapons")
	fmt.Println("2. Shop Armor")
	fmt.Println("3. Shop Equipment")
	fmt.Println("4. Shop Curiosities (Recipes & Spells)")
	fmt.Println("5. Research Quest")
	fmt.Println("6. Visit Tavern")
	fmt.Println("7. Politicks - Curry Favor / Influence")
	fmt.Println("8. View/Accept Missions")
	fmt.Println("9. Travel")

	fmt.Println("q. Quit")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	return rsp
}

func (village * Village) goShop() {

	var result string = ""

	for result != "x" {
		result = village.visitVillage()

		if result == "1" {
			village.buyWeaponScreen()
		} else if result == "2" {
			village.buyArmorScreen()
		} else if result == "3" {
			village.buySuppliesScreen()
		} else if result == "4" {
			village.buyAnimalsScreen()
		}
	}

}