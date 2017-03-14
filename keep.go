// keep.go

package main

import "fmt"

var keepDescriptions = []string{
	"It's cold and dark here. Shadows from my waning fire dance across the vacant \nexpanse.",
	"It's empty and barren, but it's mine.",
}

type Keep struct {
	name             string
	acres, usedacres int
	description      string
}

// uses political favors to gain acres
// use land to build useful structures

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

// POTIONS are called UNCTURES
// HAND OF GLORY - make some creatures flee (5 fingers, 5 uses)
// Cross of St Martin - Stun some creatures for a turn
//

// SPELLS are called LAMENTATIONS
// "Only the majesty of grief and sorrow separates us from the dark.  The dead never grieve, and the dark knows little of sorrow."
// Can only prepare spells at keep.  Require ingredients & preparation time. Maximum number prepared at one time is a figment of intl
// very difficult to learn new spells and the effects are usually muted

// Wander Action - create a random set of grids with ingredients/objects/npcs - potential apprentices, maybe a mugger

func (keep *Keep) goKeep() {
	rsp := ""

	for rsp != "7" {
		clearConsole()
		fmt.Println("╔ Keep ╗")
		fmt.Println(makeDialogString(keep.description))
		fmt.Printf("Day: %v \n", gameDay)
		fmt.Printf("Acres: %v / %v \n", keep.usedacres, keep.acres)
		fmt.Println("------------")
		fmt.Println("1. Rest (End Day)")
		fmt.Println("2. Manage Keep")
		fmt.Println("3. Apprentices")
		fmt.Println("4. Inventory")
		fmt.Println("5. Status")
		fmt.Println("")
		fmt.Println("6. Minutiae")
		fmt.Println("7. Exit")
		fmt.Println("")
		fmt.Printf("Select an Option:  ")

		fmt.Scanln(&rsp)

		if rsp == "1" {
			endDay()
		} else if rsp == "5" {
			character.showStatus()
			character.printCharacter(1)
		}
	}
}

func createKeep() Keep {
	var keep Keep

	keep.acres = 0
	keep.usedacres = 0
	keep.name = "Campground"
	keep.description = keepDescriptions[0]
	return keep
}
