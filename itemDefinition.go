// itemDefinition.go

package main

import "strings"
import "fmt"
import "strconv"

var qualities = []string{"Crude", "Stand", "Crafts", "Master"}
var materials = []string{"Oak", "Bone", "Stone", "Iron", "Steel", "Silver"}
var equipStrings = []string{"Head", "Neck", "Arms", "Chest", "Leg", "Feet", "Ring", "Cloak", "Hand", "Any", "None"}
var containers = []string{"Chest", "Bag", "Satchel", "Skeleton", "Debris", "Corpse", "Barrel", "Crate"}

// TYPE constants
const ITEM_TYPE_UNSET = 0 //item struct will zeroize on init to this
const ITEM_TYPE_WEAPON = 1
const ITEM_TYPE_ARMOR = 2
const ITEM_TYPE_UNCTURE = 3
const ITEM_TYPE_INGREDIENT = 4
const ITEM_TYPE_EQUIPMENT = 5
const ITEM_TYPE_SPECIAL = 9

const RARITY_LOW = 0
const RARITY_MED = 1
const RARITY_HIGH = 2

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

const CRITICAL = 200

var HUMAN_TARGETS = []int{EQUIP_FEET, EQUIP_LEG, EQUIP_LEG, EQUIP_ARMS, EQUIP_ARMS, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_HEAD, CRITICAL}
var ORB_TARGETS = []int{EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, EQUIP_CHEST, CRITICAL}
var QUAD_TARGETS = []int{EQUIP_FEET, EQUIP_FEET, EQUIP_LEG, EQUIP_LEG, EQUIP_LEG, EQUIP_LEG, EQUIP_CHEST, EQUIP_CHEST, EQUIP_HEAD, CRITICAL}

var HUMAN_STRING = []string{"Foot", "Leg", "Leg", "Arm", "Arm", "Chest", "Chest", "Chest", "Head", "Critical"}
var ORB_STRING = []string{"Chest", "Chest", "Chest", "Chest", "Chest", "Chest", "Chest", "Chest", "Chest", "Critical"}
var QUAD_STRING = []string{"Foot", "Foot", "Leg", "Leg", "Leg", "Leg", "Chest", "Chest", "Head", "Critical"}

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
	//dmg, acc, def, wgt, armdef, shields, durabMultiplier, costMultip, atkTurnsMod, resistMod
	{-1, 0, 0, 0, 0, 0, 1, 1, 1, -1}, // crude
	{0, 0, 0, 0, 0, 1, 1, 2, 0, 0},   //standard
	{1, 1, 0, -1, 1, 2, 2, 3, -1, 1}, // craftsman
	{2, 2, 1, -2, 1, 3, 3, 4, -1, 2}, // master
}

var weapons = []Weapon{ //name, hands, dmgmin, dmgmax, acc, def, weight, durab, value, range, atkTurns, noMaterial flag, vsPad, vsLeath, vsChain, rarity
	{"Club", 1, 1, 4, 0, 0, 9, 12, 5, 1, 3, 0, 0, 1, 1, RARITY_LOW},
	{"Knife", 1, 2, 4, 0, 0, 4, 30, 7, 1, 2, 0, 1, 0, -1, RARITY_LOW},
	{"Hatchet", 1, 2, 5, 0, 0, 7, 25, 8, 1, 3, 0, 0, 0, 0, RARITY_LOW},
	{"Dagger", 1, 3, 4, 0, 0, 5, 36, 8, 1, 2, 0, 1, 0, -1, RARITY_LOW},
	{"Short Sword", 1, 3, 5, 0, 0, 7, 32, 9, 1, 3, 0, 1, 0, -1, RARITY_MED},
	{"Light Mace", 1, 3, 6, 0, 0, 8, 38, 15, 1, 3, 0, -1, 1, 1, RARITY_MED},
	{"Lt Crossbow", 2, 1, 3, 0, -1, 9, 26, 12, 3, 4, 1, 0, 0, 0, RARITY_MED},
}

var armors = []Armor{ // name, shields, defense, resistance, weight, value, slot, rarity (0-2)
	{"Cloth Shirt", 1, 0, 4, 0, 0, EQUIP_CHEST, RARITY_LOW},
	{"Thick Cloth Coat", 1, 0, 6, 4, 5, EQUIP_CHEST, RARITY_LOW},
	{"Padded Jerkin", 2, 1, 8, 7, 12, EQUIP_CHEST, RARITY_MED},
	{"Soft Leather Jerkin", 2, 1, 9, 10, 25, EQUIP_CHEST, RARITY_MED},
	{"Hard Leather Jerkin", 3, 1, 10, 14, 36, EQUIP_CHEST, RARITY_MED},
	{"Studded Leather Jerkin", 4, 2, 11, 16, 48, EQUIP_CHEST, RARITY_HIGH},
	{"Chain Shirt", 5, 2, 14, 24, 60, EQUIP_CHEST, RARITY_HIGH},

	{"Padded Sleeves", 1, 1, 8, 1, 2, EQUIP_ARMS, RARITY_LOW},
	{"Leather Sleeves", 2, 1, 9, 2, 2, EQUIP_ARMS, RARITY_MED},
	{"Chain Sleeves", 4, 2, 14, 4, 2, EQUIP_ARMS, RARITY_HIGH},

	{"Padded Coif", 1, 0, 8, 1, 2, EQUIP_HEAD, RARITY_LOW},
	{"Leather Coif", 2, 1, 9, 2, 2, EQUIP_HEAD, RARITY_MED},
	{"Chain Coif", 3, 1, 14, 4, 2, EQUIP_HEAD, RARITY_HIGH},

	{"Cloth Pants", 1, 0, 4, 0, 0, EQUIP_LEG, RARITY_LOW},
	{"Padded Greeves", 1, 1, 8, 1, 2, EQUIP_LEG, RARITY_LOW},
	{"Leather Greeves", 2, 1, 9, 1, 2, EQUIP_LEG, RARITY_MED},
	{"Chain Greeves", 4, 2, 14, 7, 2, EQUIP_LEG, RARITY_HIGH},

	{"Leather Boots", 2, 0, 9, 2, 2, EQUIP_FEET, RARITY_LOW},
	{"Hard Leather Boots", 3, 0, 10, 3, 7, EQUIP_FEET, RARITY_MED},
	{"Chain Boots", 4, 1, 14, 3, 7, EQUIP_FEET, RARITY_HIGH},

	{"Light Cape", 1, 1, 0, 1, 1, EQUIP_CLOAK, RARITY_MED},
	{"Wood Shield", 3, 3, 10, 2, 2, EQUIP_HAND, RARITY_MED},
}

var common = []Item{
	{0, "Torch", ITEM_TYPE_EQUIPMENT, 1, 1, 3, 3, EQUIP_HAND, 1, 2, "", "", 5, 0, 0, 0, 1, -2, 3, 0, 0, 0, 0, 1, 2, -1, 1},
	{0, "Cobbler Weed", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 1, "", "", 4, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Finger Bone", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 1, "", "", 6, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Iron Bar", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 5, "", "", 9, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Rope", ITEM_TYPE_EQUIPMENT, 1, 1, 1, 1, EQUIP_NONE, 1, 1, "", "", 7, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Bandage", ITEM_TYPE_EQUIPMENT, 1, 1, 1, 1, EQUIP_NONE, 1, 1, "", "", 8, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},	
	{0, "Junk", ITEM_TYPE_EQUIPMENT, 1, 1, 1, 1, EQUIP_NONE, 1, 2, "", "", 0, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},	
}

var uncommon = []Item{
	{0, "Hollow Rose", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 1, "", "", 18, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Finger Bone", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 1, "", "", 16, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Copper Bar", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 5, "", "", 24, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Salve", ITEM_TYPE_UNCTURE, 1, 1, 1, 1, EQUIP_NONE, 1, 5, "", "", 20, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},	
}

var rare = []Item{
	{0, "Lantern", ITEM_TYPE_EQUIPMENT, 99, 99, 9, 9, EQUIP_HAND, 1, 7, "", "", 56, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
	{0, "Silver Bar", ITEM_TYPE_INGREDIENT, 1, 1, 1, 1, EQUIP_NONE, 1, 5, "", "", 48, 0, 0, 0, 0, -4, 4, 0, 0, 0, 0, 1, -3, -3, -3},
}

type Loot struct {
	crowns     int
	items      []Item
	locX, locY int
	seen       bool
	empty      bool
	container  string
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
	dmgMin, dmgMax int
	wRange         int
	accuracy       int
	atkTurns       int
	// armor stuff
	defense                         int
	shields                         int
	maxShields                      int
	resistance                      int
	noMaterialFlag                  int // used for crossbows or any other item that shouldn't have a material
	paddedMod, leatherMod, chainMod int // vs vars for weapons
}

//name, hands, dmgmin, dmgmax, acc, def, weight, durab, value, range, atkTurns
type Weapon struct {
	name                            string
	hands                           int
	dmgMin                          int
	dmgMax                          int
	accuracy                        int
	defense                         int
	weight                          int
	durab                           int
	value                           int
	wRange                          int
	atkTurns                        int
	noMaterialFlag                  int
	paddedMod, leatherMod, chainMod int
	rarity 							int
}

// name, shields, defense, weight, value, equip
type Armor struct {
	name       string
	shields    int
	defense    int
	resistance int
	weight     int
	value      int
	equip      int
	rarity	   int
}

func genGameWeapon(weapon Weapon, qual string, mat string) Item {
	var item Item

	item.id = game.itemInstanceId
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
	item.maxShields = 0
	item.resistance = 0
	item.noMaterialFlag = weapon.noMaterialFlag
	item.paddedMod = weapon.paddedMod
	item.leatherMod = weapon.leatherMod
	item.chainMod = weapon.chainMod
	item.atkTurns = weapon.atkTurns

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
	if mIdx > -1 && weapon.noMaterialFlag == 0 {
		// dmgMod, accMod, defMod, weightMod, durabMultiplier, costMultip
		item.dmgMax += materialBonuses[mIdx][0]
		item.accuracy += materialBonuses[mIdx][1]
		item.defense += materialBonuses[mIdx][2]
		item.weight += materialBonuses[mIdx][3]
		item.durability *= materialBonuses[mIdx][4]
		item.maxDurability *= materialBonuses[mIdx][4]
		item.value *= materialBonuses[mIdx][5]
	}

	if item.weight < 1 {
		item.weight = 1
	}

	game.itemInstanceId += 1
	return item
}

func genGameArmor(armor Armor, qual string) Item {
	var item Item

	item.id = game.itemInstanceId
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
	item.shields = armor.shields
	item.maxShields = armor.shields
	item.noMaterialFlag = 1 // armor never has a conventional material, its all dessicated flesh and steel
	item.resistance = armor.resistance

	// apply quality modifiers
	qIdx := getQualityIndex(qual)
	if qIdx > -1 {
		// name, shields, defense, weight, value, equip		(armor)
		//dmg, acc, def, wgt, armdef, shields, durabMultiplier, costMultip, atkTurnsMod (qual)
		item.weight += qualBonuses[qIdx][3]
		item.defense += qualBonuses[qIdx][4]
		item.shields += qualBonuses[qIdx][5]
		item.maxShields += qualBonuses[qIdx][5]
		item.durability *= qualBonuses[qIdx][6]
		item.maxDurability *= qualBonuses[qIdx][6]
		item.value *= qualBonuses[qIdx][7]
		item.resistance += qualBonuses[qIdx][9]
	}

	if item.weight < 1 {
		item.weight = 1
	}

	game.itemInstanceId += 1
	return item
}

func getRandomWeapon(weight int) Item {
	var die Die

	var weapon = weapons[die.rollxdx(0, len(weapons)-1)]
	var mat = materials[die.rollxdx(0, len(materials)-1)]
	var qual = qualities[die.rollxdx(0, len(qualities)-1)]

	return genGameWeapon(weapon, qual, mat)
}

func getRandomArmor(weight int) Item {
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

func restoreSavedItem(line string) (Item, int) {
	var itm Item

	bits := strings.Split(line, ",")

	if bits[0] != "ITEM" {
		log.addError("Expected Item row - not found.")
		fmt.Println("Item row not found!")
		return itm, -1
	}

	itm.id, _ = strconv.Atoi(bits[1])
	itm.name = bits[2]
	itm.typeCode, _ = strconv.Atoi(bits[3])

	itm.uses, _ = strconv.Atoi(bits[4])
	itm.maxuses, _ = strconv.Atoi(bits[5])
	itm.durability, _ = strconv.Atoi(bits[6])
	itm.maxDurability, _ = strconv.Atoi(bits[7])
	itm.equip, _ = strconv.Atoi(bits[8])
	itm.hands, _ = strconv.Atoi(bits[9])
	itm.weight, _ = strconv.Atoi(bits[10])

	itm.material = bits[11]
	itm.quality = bits[12]

	itm.value, _ = strconv.Atoi(bits[13])
	itm.magical, _ = strconv.Atoi(bits[14])
	itm.dmgMin, _ = strconv.Atoi(bits[15])
	itm.dmgMax, _ = strconv.Atoi(bits[16])
	itm.wRange, _ = strconv.Atoi(bits[17])
	itm.accuracy, _ = strconv.Atoi(bits[18])
	itm.atkTurns, _ = strconv.Atoi(bits[19])

	itm.defense, _ = strconv.Atoi(bits[20])
	itm.shields, _ = strconv.Atoi(bits[21])
	itm.maxShields, _ = strconv.Atoi(bits[22])
	itm.resistance, _ = strconv.Atoi(bits[23])
	itm.noMaterialFlag, _ = strconv.Atoi(bits[24])
	itm.paddedMod, _ = strconv.Atoi(bits[25])
	itm.leatherMod, _ = strconv.Atoi(bits[26])
	itm.chainMod, _ = strconv.Atoi(bits[27])

	return itm, 1
}

func (itm *Item) getSaveString() string {
	saveString := "ITEM,"

	saveString += fmt.Sprintf("%v,", itm.id)
	saveString += itm.name + ","
	saveString += fmt.Sprintf("%v,", itm.typeCode)
	saveString += fmt.Sprintf("%v,", itm.uses)
	saveString += fmt.Sprintf("%v,", itm.maxuses)
	saveString += fmt.Sprintf("%v,", itm.durability)
	saveString += fmt.Sprintf("%v,", itm.maxDurability)
	saveString += fmt.Sprintf("%v,", itm.equip)
	saveString += fmt.Sprintf("%v,", itm.hands)
	saveString += fmt.Sprintf("%v,", itm.weight)

	saveString += itm.material + ","
	saveString += itm.quality + ","

	saveString += fmt.Sprintf("%v,", itm.value)
	saveString += fmt.Sprintf("%v,", itm.magical)
	saveString += fmt.Sprintf("%v,", itm.dmgMin)
	saveString += fmt.Sprintf("%v,", itm.dmgMax)

	saveString += fmt.Sprintf("%v,", itm.wRange)
	saveString += fmt.Sprintf("%v,", itm.accuracy)
	saveString += fmt.Sprintf("%v,", itm.atkTurns)

	saveString += fmt.Sprintf("%v,", itm.defense)
	saveString += fmt.Sprintf("%v,", itm.shields)
	saveString += fmt.Sprintf("%v,", itm.maxShields)
	saveString += fmt.Sprintf("%v,", itm.resistance)
	saveString += fmt.Sprintf("%v,", itm.noMaterialFlag)
	saveString += fmt.Sprintf("%v,", itm.paddedMod)
	saveString += fmt.Sprintf("%v,", itm.leatherMod)
	saveString += fmt.Sprintf("%v,", itm.chainMod)

	saveString += "◄"

	return saveString

}

func getEmptyItem() Item {
	var item Item

	item.id = -1
	item.name = "empty"
	item.equip = EQUIP_ANY
	item.weight = 0

	return item
}

func (item *Item) getInvDisplayString() string {

	durab := strings.Repeat("▲", item.shields)
	miss := ""
	miss = strings.Repeat("•", (item.maxShields - item.shields))

	if len(durab) < 1 && item.maxShields > 0 {
		durab = "X"
	}

	disp := packSpaceString(fmt.Sprintf("%s", item.name), 25)
	if item.typeCode == ITEM_TYPE_ARMOR {
		disp += fmt.Sprintf("[%s", durab)
		disp += miss
		disp += "]"
	}

	return disp
}

func (item *Item) getStatusDisplayStringArmor() string {

	durab := strings.Repeat("♦", item.shields)
	miss := ""
	miss = strings.Repeat("•", (item.maxShields - item.shields))

	if len(durab) < 1 && item.maxShields > 0 {
		durab = "X"
	}

	disp := packSpaceString(item.name, 24)
	if item.typeCode == ITEM_TYPE_ARMOR {
		disp += fmt.Sprintf("[%s", durab)
		disp += miss
		disp += "]"
	}

	return disp
}

func makeLootItem(itm Item) Item {
	itm.id = game.itemInstanceId
	game.itemInstanceId++
	return itm
}

func createRandomLoot() Loot {
	var loot Loot
	var die Die

	loot.seen = false
	loot.crowns = die.rollxdx(0, game.gameDay+4)

	roll := die.rollxdx(1, 100)
	timeMod := -2

	for k := game.gameDay; k > 20; k -= 20 {
		timeMod++
	}

	roll += timeMod

	if roll >= 98 {
		if die.rollxdx(1, 6) > 3 {
			loot.items = append(loot.items, getRandomWeapon(0))
		} else {
			loot.items = append(loot.items, getRandomArmor(0))		
		}
	} else if roll > 96 {	
		loot.items = append(loot.items, makeLootItem(rare[die.rollxdx(1, len(rare))-1]))
	} else if roll > 88 {
		loot.items = append(loot.items, makeLootItem(uncommon[die.rollxdx(1, len(uncommon))-1]))
	} else {
		loot.items = append(loot.items, makeLootItem(common[die.rollxdx(1, len(common))-1]))
	}

	if die.rollxdx(1, 100) > 88 {
		loot.items = append(loot.items, common[die.rollxdx(1, len(common))-1])
	}

	roll = die.rollxdx(1, len(containers))

	loot.container = containers[roll-1]

	return loot
}

func (w *Weapon) getWeaponStatLine() string {
	return ""
}
