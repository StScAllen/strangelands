// itemDefinition.go

package main

import "strings"
import "fmt"

var qualities = []string{"Crude", "Stand", "Crafts", "Master"}
var materials = []string{"Oak", "Bone", "Stone", "Iron", "Steel", "Silver"}

var itemInstanceId int = 0

// TYPE constants
const ITEM_TYPE_UNSET = 0 //item struct will zeroize on init to this
const ITEM_TYPE_WEAPON = 1
const ITEM_TYPE_ARMOR = 2
const ITEM_TYPE_UNCTURE = 3
const ITEM_TYPE_SPECIAL = 4

const ITEM_TYPE_OTHER = 9

// Equip Constants
const EQUIP_HEAD = 0
const EQUIP_NECK = 1
const EQUIP_ARMS = 2
const EQUIP_CHEST = 3
const EQUIP_LEG = 4
const EQUIP_FEET = 5
const EQUIP_RING = 6
const EQUIP_CLOAK = 7
const EQUIP_HAND = 8
const EQUIP_ANY = 9 // only used for blank items
const EQUIP_NONE = 100

var materialBonuses = [][]int{
	// dmgMod, accMod, defMod, weightMod, durabMultiplier, costMultip
	{-1, -1, 0, -1, 1, 1}, // oak
	{0, -1, 0, -1, 2, 2},  // bone
	{0, -1, 0, 1, 3, 3},   // stone
	{1, 0, 0, 1, 4, 4},    // iron
	{1, 1, 0, 1, 5, 5},    // steel
	{0, 1, 0, 0, 2, 5},    // silver
}
var qualBonuses = [][]int{
	//dmg, acc, def, wgt, armdef, shields, durabMultiplier, costMultip, atkTurnsMod
	{-1, 0, 0, 0, 0, 1, 1, 1, 1},  // crude
	{0, 0, 0, 0, 0, 2, 1, 2, 0},   //standard
	{1, 1, 0, -1, 1, 4, 2, 3, -1}, // craftsman
	{2, 2, 1, -2, 2, 6, 3, 4, -1}, // master
}

var weapons = []Weapon{ //name, hands, dmgmin, dmgmax, acc, def, weight, durab, value, range, atkTurns, noMaterial flag
	{"Club", 1, 1, 4, 0, 0, 9, 12, 5, 1, 3, 0},
	{"Knife", 1, 2, 4, 0, 0, 4, 30, 7, 1, 2, 0},
	{"Hatchet", 1, 2, 5, 0, 0, 7, 25, 8, 1, 3, 0},
	{"Dagger", 1, 3, 4, 0, 0, 5, 36, 8, 1, 2, 0},
	{"Short Sword", 1, 3, 5, 0, 0, 7, 32, 9, 1, 3, 0},
	{"Lt Crossbow", 2, 1, 3, 0, -1, 9, 26, 12, 3, 4, 1},	
}

var armors = []Armor{ // name, shields, defense, weight, value, slot
	{"Cloth Shirt", 1, 1, 0, 0, EQUIP_CHEST},
	{"Thick Cloth Coat", 1, 2, 1, 2, EQUIP_CHEST},
	{"Padded Jerkin", 2, 3, 2, 4, EQUIP_CHEST},
	{"Soft Leather Jerkin", 2, 4, 3, 8, EQUIP_CHEST},
	{"Hard Leather Jerkin", 3, 5, 4, 11, EQUIP_CHEST},
	{"Studded Leather Jerkin", 3, 6, 5, 11, EQUIP_CHEST},
	{"Padded Sleeves", 1, 1, 1, 2, EQUIP_ARMS},
	{"Leather Sleeves", 2, 1, 1, 2, EQUIP_ARMS},
	{"Chain Sleeves", 4, 2, 1, 2, EQUIP_ARMS},
	{"Padded Coif", 1, 1, 1, 2, EQUIP_HEAD},
	{"Leather Coif", 2, 1, 1, 2, EQUIP_HEAD},
	{"Chain Coif", 4, 2, 1, 2, EQUIP_HEAD},
	{"Cloth Pants", 1, 1, 0, 0, EQUIP_LEG},	
	{"Padded Greeves", 1, 1, 1, 2, EQUIP_LEG},
	{"Leather Greeves", 2, 1, 1, 2, EQUIP_LEG},
	{"Chain Greeves", 4, 2, 1, 2, EQUIP_LEG},
	{"Light Cape", 1, 1, 1, 1, EQUIP_CLOAK},
	{"Wood Shield", 3, 1, 2, 2, EQUIP_HAND},
}

type Item struct { // regular items
	id                        int    // instance code
	name                      string // name of item
	typeCode                  int    // ITEM_TYPE_* constants
	uses, maxuses             int    // uses is for unctures/items
	durability, maxDurability int    // durability is for weapon/armor
	equip                     int    // body part code or 0 EQUIP_* constants
	hands                     int    // 1 or 2 handed, for items equippable in EQUIP_HAND
	weight                    int    // weight
	material                  string // material
	quality                   string // QUALITY_* constants
	value                     int    // value, in currency
	magical                   int    // flag 0, 1
	// weapon stuff
	dmgMin, dmgMax 				int
	wRange         				int
	accuracy       				int
	atkTurns      				int
	// armor stuff
	defense 					int
	shields 					int
	noMaterialFlag 				int		// used for crossbows or any other item that shouldn't have a material
}

//name, hands, dmgmin, dmgmax, acc, def, weight, durab, value, range, atkTurns
type Weapon struct {
	name     	string
	hands    	int
	dmgMin   	int
	dmgMax   	int
	accuracy 	int
	defense  	int
	weight   	int
	durab    	int
	value    	int
	wRange   	int
	atkTurns 	int
	noMaterialFlag int
}

// name, shields, defense, weight, value, equip
type Armor struct {
	name    string
	shields int
	defense int
	weight  int
	value   int
	equip   int
}

func genGameWeapon(weapon Weapon, qual string, mat string) Item {
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
	item.noMaterialFlag = weapon.noMaterialFlag
	
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
	if mIdx > -1 && weapon.noMaterialFlag == 0{
		// dmgMod, accMod, defMod, weightMod, durabMultiplier, costMultip
		item.dmgMax += materialBonuses[mIdx][0]
		item.accuracy += materialBonuses[mIdx][1]
		item.defense += materialBonuses[mIdx][2]
		item.weight += materialBonuses[mIdx][3]
		item.durability *= materialBonuses[mIdx][4]
		item.maxDurability *= materialBonuses[mIdx][4]
		item.value *= materialBonuses[mIdx][5]
	}

	if (item.weight < 1){
		item.weight = 1
	}
	
	itemInstanceId += 1
	return item
}

func genGameArmor(armor Armor, qual string) Item {
	var item Item

	item.id = itemInstanceId
	item.name = armor.name
	item.typeCode = ITEM_TYPE_ARMOR
	item.uses = 1
	item.maxuses = 1
	item.durability = armor.shields
	item.maxDurability = armor.shields
	item.equip = armor.equip
	item.hands = 1
	item.weight = armor.weight
	item.material = ""
	item.quality = qual
	item.value = armor.value
	item.magical = 0
	item.dmgMin = 1
	item.dmgMax = 1
	item.wRange = 0
	item.accuracy = 0
	item.defense = armor.defense
	item.shields = 0
	item.noMaterialFlag = 1	// armor never has a conventional material, its all dessicated flesh and steel
	
	// apply quality modifiers
	qIdx := getQualityIndex(qual)
	if qIdx > -1 {
		// name, shields, defense, weight, value, equip		(armor)
		//dmg, acc, def, wgt, armdef, shields, durabMultiplier, costMultip, atkTurnsMod (qual)
		item.weight += qualBonuses[qIdx][3]
		item.defense += qualBonuses[qIdx][4]
		item.shields += qualBonuses[qIdx][5]
		item.durability *= qualBonuses[qIdx][6]
		item.maxDurability *= qualBonuses[qIdx][6]
		item.value *= qualBonuses[qIdx][7]
	}
	
	if (item.weight < 1){
		item.weight = 1
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

func getRandomArmor() Item {
	var die Die

	var armor = armors[die.rollxdx(0, len(armors)-1)]
	var qual = qualities[die.rollxdx(0, len(qualities)-1)]

	return genGameArmor(armor, qual)
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

func getEmptyItem() Item {
	var item Item

	item.id = -1
	item.name = "empty"
	item.equip = EQUIP_ANY
	item.weight = 0

	return item
}

func (item * Item) getInvDisplayString() (string){
	
	durab := strings.Repeat("▲", item.durability)
	miss := ""
	miss = strings.Repeat("•", (item.maxDurability - item.durability))
	
	if (len(durab) < 1 && item.maxDurability > 0){
		durab = "X"
	}
	
	disp := packSpaceString(fmt.Sprintf("%s", item.name), 25)  
	if (item.typeCode == ITEM_TYPE_ARMOR){
		disp += fmt.Sprintf("[%s", durab)
		disp += miss
		disp += "]"
	}

	return disp;
}

func (w *Weapon) getWeaponStatLine() string {

	return ""
}
