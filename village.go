// village.go
// Village

package main

import "fmt"
import "time"

type Village struct {
	name            string
	distanceToKeep  int
	size            int
	politicalFavor  int
	shopWeapons     []Item
	shopArmor       []Item
	shopProvisions  []Item
	shopApothecary  []Item
	shopCuriosities []Item
	mapX, mapY      int
}

func buildVillages() {
	villages = make([]Village, 8, 8)

	var crowley Village
	crowley.name = "Crowley"
	crowley.distanceToKeep = 1
	crowley.size = 1
	crowley.mapX, crowley.mapY = 20, 10
	villages[0] = crowley

	var bristal Village
	bristal.name = "Bristal"
	bristal.distanceToKeep = 2
	bristal.size = 3
	bristal.mapX, bristal.mapY = 17, 3
	villages[1] = bristal

	var faust Village
	faust.name = "Faust"
	faust.distanceToKeep = 2
	faust.size = 3
	faust.mapX, faust.mapY = 37, 2
	villages[2] = faust

	var gould Village
	gould.name = "Gould"
	gould.distanceToKeep = 2
	gould.size = 2
	gould.mapX, gould.mapY = 35, 15
	villages[3] = gould

	var elise Village
	elise.name = "Elise"
	elise.distanceToKeep = 3
	elise.size = 2
	elise.mapX, elise.mapY = 56, 15
	villages[4] = elise

	var autumn Village
	autumn.name = "Autumn"
	autumn.distanceToKeep = 4
	autumn.size = 4
	autumn.mapX, autumn.mapY = 57, 3
	villages[5] = autumn

	var hollow Village
	hollow.name = "Hollow"
	hollow.distanceToKeep = 13
	hollow.size = 2
	hollow.mapX, hollow.mapY = 2, 13
	villages[6] = hollow

	var caustus Village
	caustus.name = "Caustus"
	caustus.distanceToKeep = 1
	caustus.size = 6
	caustus.mapX, caustus.mapY = 35, 11
	villages[7] = caustus
}

func (village *Village) research() {

}

func (village *Village) visitTavern() {
	// rest		 (end day, gain hp)
	// get gossip (apprentice tips, politickal gains, )
	// gamble  (randomly gain or lose money)
	// drink  (raise soul by 1 - cost)

	endDay()
	save()
}

func (village *Village) visitChirurgeon() {
	// set wounds
	// sutures
	// increase healing time
	// costs

}

func (village *Village) politicks() {

}

func (village *Village) goMissions() {
	chooseAdventure()
	result := adventure()
	
	if result == DIED {
		// end the game and go back to main menu
	}
}

func (village *Village) visitVillage() string {
	clearConsole()

	fmt.Println("+++ Village of " + village.name + " +++")
	fmt.Println("------------")
	fmt.Println("1. Shop")
	fmt.Println("2. Research Quest")
	fmt.Println("3. Visit Tavern")
	fmt.Println("4. Visit Chirurgeon")
	fmt.Println("5. Politicks - Curry Favor / Influence")
	fmt.Println("6. Missions")
	fmt.Println("7. Travel")

	fmt.Println("q. Quit")
	fmt.Println("")
	fmt.Println("Select an Option:  ")

	rsp := ""
	fmt.Scanln(&rsp)

	if rsp == "1" {
		village.shopMenu()
	} else if rsp == "2" {
		village.research()
	} else if rsp == "3" {
		village.visitTavern()
	} else if rsp == "4" {
		village.visitChirurgeon()
	} else if rsp == "5" {
		village.politicks()
	} else if rsp == "6" {
		village.goMissions()
	} else if rsp == "7" {
		showTravelMenu()
	} else if rsp == "m" {
		openTownMinutiae()
	} else if rsp != "q" {
		village.visitVillage()
	}

	return rsp
}

func showTravelMenu() string {

	showVillages()

	validSelection := false
	dist := 0
	destination := ""
	rsp := ""
	prevIndex := character.villageIndex

	for validSelection == false {
		clearConsole()

		fmt.Println("Travel Menu")
		fmt.Printf("Day: %v \n", game.gameDay)
		fmt.Println("------------")
		dist = getVillageDistance(0)
		fmt.Println(packSpaceString("1. Crowley", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(1)
		fmt.Println(packSpaceString("2. Bristal", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(2)
		fmt.Println(packSpaceString("3. Faust", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(3)
		fmt.Println(packSpaceString("4. Gould", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(4)
		fmt.Println(packSpaceString("5. Elise", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(5)
		fmt.Println(packSpaceString("6. Autumn", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(6)
		fmt.Println(packSpaceString("7. Hollow", 20) + fmt.Sprintf("(%v days travel)", dist))
		dist = getVillageDistance(7)
		fmt.Println(packSpaceString("8. Caustus", 20) + fmt.Sprintf("(%v days travel)", dist))
		fmt.Println("")
		dist = getVillageDistance(99)
		fmt.Println(packSpaceString("k. Return to Keep", 20) + fmt.Sprintf("(%v days travel)", dist))
		fmt.Println("m. World Map")
		fmt.Println("h. Minutiae")
		fmt.Println("x. Back")
		fmt.Println("    ----    ")
		fmt.Printf("Where do you wish to travel? ")

		fmt.Scanln(&rsp)

		switch rsp {
		case "1":
			dist = getVillageDistance(0)
			character.villageIndex = 0
			destination = villages[0].name
			validSelection = true
		case "2":
			dist = getVillageDistance(1)
			character.villageIndex = 1
			destination = villages[1].name
			validSelection = true
		case "3":
			dist = getVillageDistance(2)
			character.villageIndex = 2
			destination = villages[2].name
			validSelection = true
		case "4":
			dist = getVillageDistance(3)
			character.villageIndex = 3
			destination = villages[3].name
			validSelection = true
		case "5":
			dist = getVillageDistance(4)
			character.villageIndex = 4
			destination = villages[4].name
			validSelection = true
		case "6":
			dist = getVillageDistance(5)
			character.villageIndex = 5
			destination = villages[5].name
			validSelection = true
		case "7":
			dist = getVillageDistance(6)
			character.villageIndex = 6
			destination = villages[6].name
			validSelection = true
		case "8":
			dist = getVillageDistance(7)
			character.villageIndex = 7
			destination = villages[7].name
			validSelection = true
		case "k":
			dist = getVillageDistance(99)
			character.villageIndex = 99
			destination = "Keep"
			validSelection = true
		case "m":
			drawWorldMap()
			validSelection = false
		case "x":
			validSelection = true
			dist = 0
		}
	}

	if dist > 0 {

		confirm := ""
		fmt.Printf("\nAre you sure you wish to travel to " + destination + "? ")
		fmt.Scanln(&confirm)

		if confirm == "y" {
			showTimePassageScreen(dist)
			fmt.Println("Arrived in "+destination+". After ", dist, " days of travel.")
			fmt.Println("Press any key to continue.")
			tgt := ""
			fmt.Scanln(&tgt)
		} else {
			character.villageIndex = prevIndex
		}

	}

	return rsp
}

func showTimePassageScreen(lapse int) {
	start := 0
	tick := "["

	for start < lapse {
		clearConsole()
		fmt.Println("Day: ", game.gameDay)
		fmt.Println("Time Passes...")
		endDay()
		tick += " █ █ █"
		if start == lapse-1 {
			tick += " ]"
		}
		fmt.Println(tick)
		time.Sleep(1000 * time.Millisecond)
		start++
	}
}

func (village *Village) goShop() {

	var result string = ""

	for result != "x" {
		result = village.visitVillage()

		if result == "1" {
			village.buyWeaponScreen()
		} else if result == "2" {
			village.buyArmorScreen()
		} else if result == "3" {
			village.buyProvisions()
		} else if result == "4" {
			village.buyCuriosities()
		}
	}

}
