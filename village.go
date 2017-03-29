// village.go
// Village

package main

import "fmt"
import "time"

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
		
	var maline Village	
	maline.name = "Maline"
	maline.distanceToKeep = 2
	maline.size = 3
	maline.mapX, maline.mapY = 17, 3
	villages[1] = maline 
	
	var faust Village	
	faust.name = "Faust"
	faust.distanceToKeep = 2
	faust.size = 3
	faust.mapX, faust.mapY = 38, 3
	villages[2] = faust 
	
	var dauntun Village	
	dauntun.name = "Dauntun"
	dauntun.distanceToKeep = 2
	dauntun.size = 2
	dauntun.mapX, dauntun.mapY = 25, 16
	villages[3] = dauntun 		
		
	var elice Village	
	elice.name = "Elice"
	elice.distanceToKeep = 3
	elice.size = 2
	elice.mapX, elice.mapY = 46, 16
	villages[4] = elice 	
	
	var hastur Village	
	hastur.name = "Hastur"
	hastur.distanceToKeep = 4
	hastur.size = 4
	hastur.mapX, hastur.mapY = 46, 2
	villages[5] = hastur 
	
	var hollow Village	
	hollow.name = "Hollow"
	hollow.distanceToKeep = 13
	hollow.size = 2
	hollow.mapX, hollow.mapY = 2, 13
	villages[6] = hollow 	
	
	var pritchard Village	
	pritchard.name = "Pritchard"
	pritchard.distanceToKeep = 1
	pritchard.size = 6
	pritchard.mapX, pritchard.mapY = 25, 10
	villages[7] = pritchard
}

func (village *Village) research() {


}

func (village *Village) visitTavern() {
	endDay()
	character.save()
}

func (village *Village) politicks() {


}

func (village *Village) goMissions() {


}

func (village * Village) visitVillage() (string) {
	clearConsole()

	fmt.Println("+++ Village of " + village.name + " +++")
	fmt.Println("------------")
	fmt.Println("1. Shop")
	fmt.Println("2. Research Quest")
	fmt.Println("3. Visit Tavern")
	fmt.Println("4. Politicks - Curry Favor / Influence")
	fmt.Println("5. Missions")
	fmt.Println("6. Travel")

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
		village.politicks()
	} else if rsp == "5" {		
		village.goMissions()
	} else if rsp == "6" {
		showTravelMenu()
	} else if rsp == "m" {
		openTownMinutiae()
	} 
	
	return rsp
}

func showTravelMenu() string {
	validSelection := false
	dist := 0
	destination := ""
	rsp := ""
	prevIndex := character.villageIndex

	for validSelection == false {
		clearConsole()	
	
		fmt.Println("Travel Menu")
		fmt.Printf("Day: %v \n", gameDay)
		fmt.Println("------------")
		dist = getVillageDistance(0)
		fmt.Println("1. Crowley 		(", dist, " days travel)")
		dist = getVillageDistance(1)
		fmt.Println("2. Maline			(", dist, " days travel)")
		dist = getVillageDistance(2)
		fmt.Println("3. Faust			(", dist, " days travel)")
		dist = getVillageDistance(3)		
		fmt.Println("4. Dauntun 		(", dist, " days travel)")
		dist = getVillageDistance(4)
		fmt.Println("5. Elice 			(", dist, " days travel)")
		dist = getVillageDistance(5)		
		fmt.Println("6. Hastur 			(", dist, " days travel)")
		dist = getVillageDistance(6)
		fmt.Println("7. Hollow 			(", dist, " days travel)")
		dist = getVillageDistance(7)		
		fmt.Println("8. Pritchard 		(", dist, " days travel)")	
		fmt.Println("")	
		dist = getVillageDistance(99)
		fmt.Println("k. Return to Keep 	(", dist, " days travel)")
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

	if (dist > 0){
	
		confirm := ""
		fmt.Printf("\nAre you sure you wish to travel to " + destination + "? ")
		fmt.Scanln(&confirm)
	
		if confirm == "y" {
			showTimePassageScreen(dist)
			fmt.Println("Arrived in " + destination + ". After ", dist, " days of travel.")
			fmt.Println("Press any key to continue." )
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
		fmt.Println("Day: ", gameDay)
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

func (village * Village) goShop() {

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