// itemDefinition.go

package main

var qualities = []string {"Crude", "Simple", "Standard", "Improved", "Craftsman", "Expert", "Master"}
var materials = []string {"Pine", "Oak", "Stone", "Copper", "Iron", "Steel", "Black Steel", "Liberine"}

const ITEM_TYPE_WEAPON = 1
const ITEM_TYPE_ARMOR = 2
const ITEM_TYPE_OTHER = 3

var materialBonuses = [][]int {
								// dmgMod, accMod, defMod, weightMod
								{-1, -1, -1, -2},	// pine
								{-1, -1, 0, -1},	// oak
								{0, -1, 0, 0},		// stone
								{0, 0, 0, 0},		// copper
								{1, 0, 0, 1},		// iron
								{1, 1, 0, 1},		// steel
								{1, 1, 1, 1},		// black steel
								{2, 2, 1, 0},		// liberine						
							  }
var qualBonuses = [][]int {
							//dmg, acc, def, wgt, armdef, shields
							{-1, -1, -1, 1, 0, 0}, // crude
							{-1, 0, 0, 0, 0, 1},	// simple
							{0, 0, 0, 0, 0, 2},	//standard
							{1, 0, 0, 0, 1, 3},	// Improved
							{1, 1, 0, -1, 1, 4},	// craftsman
							{2, 1, 1, -1, 2, 5},	// expert
							{2, 2, 1, -2, 2, 6},	// master							
						  }		
					
var weapons = []Weapon {	//name, hands, dmgmin, dmgmax, acc, def, material, weight, value
							{"Stick", 1, 1, 4, 0, 0, "", "", 1, 5},
							{"Knife", 1, 2, 4, 0, 0, "", "", 1, 7},
							{"Hatchet", 1, 2, 5, 0, 0, "", "", 2, 8},
							{"Dagger", 1, 3, 4, 0, 0, "", "", 1, 8},
						}

var armors = []Armor {
						{"Clothes", 1, 2, "", "", 0, 0},
						{"Thick Coat", 1, 3, "", "", 1, 2},
						{"Padded Jerkin", 2, 3, "", "", 2, 4},
						{"Full Padded Suit", 3, 5, "", "", 3, 9},
						{"Soft Leather Jerkin", 2, 5, "", "", 3, 8},
						{"Soft Leather Suit", 3, 8, "", "", 4, 12},
						{"Hard Leather Jerkin", 3, 7, "", "", 4, 11},
						{"Hard Leather Suit", 4, 10, "", "", 5, 16},		
						{"Studded Leather Jerkin", 3, 9, "", "", 5, 11},
						{"Studded Leather Suit", 5, 12, "", "", 7, 16},						
					 }	
					 

type Item struct {		// regular items
	id int
	food int
	special int
	uses, maxuses int
}

type Weapon struct {
	name string
	hands int
	dmgMin int
	dmgMax int
	accuracy int
	defense int
	material string
	quality string
	weight int
	value int
}

type Armor struct {
	name string
	defense int
	shields int
	material string
	quality string
	weight int
	value int
}					

func getQualityIndex(quality string) (int){

	if quality != "" {	
		for i := 0; i < len(qualities); i++ {		
			if (qualities[i] == quality){
				return i
			}	
		}	
	}
	
	return -1
}

func getMaterialIndex(mat string) (int){

	if mat != "" {	
		for i := 0; i < len(materials); i++ {		
			if (materials[i] == mat){
				return i
			}	
		}	
	}
	
	return -1
} 

func (w *Weapon) getWeaponStatLine() (string){

	
	return ""
}