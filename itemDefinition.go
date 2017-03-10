// itemDefinition.go

package main

var qualities = []string{"Crude", "Stand", "Crafts", "Master"}
var materials = []string{"Oak", "Bone", "Stone", "Iron", "Steel", "Silver"}

var itemInstanceId int = 0

// TYPE constants
const ITEM_TYPE_UNSET = 0		//item struct will zeroize on init to this
const ITEM_TYPE_WEAPON = 1
const ITEM_TYPE_ARMOR = 2
const ITEM_TYPE_UNCTURE = 3
const ITEM_TYPE_SPECIAL = 4

const ITEM_TYPE_OTHER = 9

// Equip Constants
const EQUIP_NONE = 0
const EQUIP_HEAD = 1
const EQUIP_NECK = 2
const EQUIP_HAND = 3
const EQUIP_ARMS = 4
const EQUIP_CHEST = 5
const EQUIP_LEG = 6
const EQUIP_FEET = 7
const EQUIP_RING = 8
const EQUIP_CLOAK = 9

var materialBonuses = [][]int{
	// dmgMod, accMod, defMod, weightMod, durabMultiplier, costMultip
	{-1, -1, 0, -1, 1, 1},  // oak
	{0, -1, 0, -1, 2, 2},    // bone
	{0, -1, 0, 0, 3, 3},    // stone
	{1, 0, 0, 1, 4, 4},     // iron
	{1, 1, 0, 1, 5, 5},     // steel
	{0, 1, 0, 0, 2, 5},     // silver
}
var qualBonuses = [][]int{
	//dmg, acc, def, wgt, armdef, shields, durabMultiplier, costMultip, atkTurnsMod
	{-1, 0, 0, 0, 0, 1, 1, 1, 1},   // crude
	{0, 0, 0, 0, 0, 2, 1, 2, 0},    //standard
	{1, 1, 0, -1, 1, 4, 2, 3, -1},   // craftsman
	{2, 2, 1, -2, 2, 6, 3, 4, -1},   // master
}

var weapons = []Weapon{ //name, hands, dmgmin, dmgmax, acc, def, weight, durab, value, range, atkTurns
	{"Club", 1, 1, 4, 0, 0, 1, 12, 5, 1, 3},
	{"Knife", 1, 2, 4, 0, 0, 1, 30, 7, 1, 2},
	{"Hatchet", 1, 2, 5, 0, 0, 2, 25, 8, 1, 3},
	{"Dagger", 1, 3, 4, 0, 0, 1, 36, 8, 1, 2},
	{"Short Sword", 1, 3, 5, 0, 0, 2, 32, 9, 1, 3},
}

var armors = []Armor{	// name, shields, defense, weight, value, slot
	{"Clothes", 1, 1, 0, 0, EQUIP_CHEST},
	{"Thick Cloth Coat", 1, 2, 1, 2, EQUIP_CHEST},
	{"Padded Jerkin", 2, 3, 2, 4, EQUIP_CHEST},
	{"Soft Leather Jerkin", 2, 4, 3, 8, EQUIP_CHEST},
	{"Hard Leather Jerkin", 3, 5, 4, 11, EQUIP_CHEST},
	{"Studded Leather Jerkin", 3, 6, 5, 11, EQUIP_CHEST},
}

type Item struct { // regular items
	id            				int		// instance code
	name						string	// name of item
	typeCode       			  	int		// ITEM_TYPE_* constants
	uses, maxuses 				int		// uses is for unctures/items
	durability, maxDurability 	int		// durability is for weapon/armor
	equip 						int 	// body part code or 0 EQUIP_* constants
	hands						int 	// 1 or 2 handed, for items equippable in EQUIP_HAND	
	weight						int 	// weight
	material					string 	// material
	quality						string 	// QUALITY_* constants
	value						int 	// value, in currency
	magical						int 	// flag 0, 1
	// weapon stuff
	dmgMin, dmgMax 				int 
	wRange						int
	accuracy					int
	atkTurns 					int
	// armor stuff
	defense						int
	shields						int
}

type Weapon struct {
	name 		string
	hands 		int
	dmgMin   	int
	dmgMax   	int
	accuracy 	int
	defense  	int
	durab 		int
	value    	int
	weight 		int
	wRange		int
	atkTurns	int
}

type Armor struct {
	name		string
	shields  	int
	defense  	int
	weight 		int
	value 		int
	equip		int
}

func genGameWeapon(weapon Weapon, qual string, mat string) (Item){
	var item Item
	
	item.id = itemInstanceId
	item.name = weapon.name
	item.typeCode = ITEM_TYPE_WEAPON
	item.uses = 1
	item.maxuses = 1
	item.durability = weapon.durab
	item.maxDurability = weapon.durab
	item.equip = EQUIP_HAND
	item.hands = weapon.hands
	item.weight = weapon.weight
	item.material = mat
	item.quality = qual
	item.value = weapon.value
	item.magical = 0
	item.dmgMin = weapon.dmgMin
	item.dmgMax = weapon.dmgMax
	item.wRange = weapon.wRange
	item.accuracy = weapon.accuracy
	item.defense = weapon.defense
	item.shields = 0
	
	// apply quality modifiers
	qIdx := getQualityIndex(qual)
	if qIdx > -1 {
		//dmg, acc, def, wgt, armdef, shields, durabMultiplier, costMultiplier, atkTurnsMod
		item.dmgMax += qualBonuses[qIdx][0]
		item.accuracy += qualBonuses[qIdx][1]
		item.defense += qualBonuses[qIdx][2]
		item.weight += qualBonuses[qIdx][3]
		// skip armDef, is weapon
		// skip shields, is weapon
		item.durability *= qualBonuses[qIdx][6]
		item.maxDurability *= qualBonuses[qIdx][6]		
		item.value *= qualBonuses[qIdx][7]
		item.atkTurns += qualBonuses[qIdx][8]
	}
	
	mIdx := getMaterialIndex(mat)
	if mIdx > -1 {
		// dmgMod, accMod, defMod, weightMod, durabMultiplier, costMultip	
		item.dmgMax += materialBonuses[mIdx][0]
		item.accuracy += materialBonuses[mIdx][1]	
		item.defense += materialBonuses[mIdx][2]	
		item.weight += materialBonuses[mIdx][3]
		item.durability *= materialBonuses[mIdx][4]
		item.maxDurability *= materialBonuses[mIdx][4]		
		item.value *= materialBonuses[mIdx][5]
	}

	itemInstanceId += 1
	return item
}

func getRandomWeapon() Item {
	var die Die
	
	var weapon = weapons[die.rollxdx(0, len(weapons)-1)]
	var mat = materials[die.rollxdx(0, len(materials)-1)]
	var qual = qualities[die.rollxdx(0, len(qualities)-1)]
	
	return genGameWeapon(weapon, qual, mat)
}

func getQualityIndex(quality string) int {

	if quality != "" {
		for i := 0; i < len(qualities); i++ {
			if qualities[i] == quality {
				return i
			}
		}
	}

	return -1
}

func getMaterialIndex(mat string) int {

	if mat != "" {
		for i := 0; i < len(materials); i++ {
			if materials[i] == mat {
				return i
			}
		}
	}

	return -1
}

func (w *Weapon) getWeaponStatLine() string {

	return ""
}
