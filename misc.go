// misc.go

package main

var allWounds = []Wound{
	{
		0, EQUIP_HEAD, "Concussion", "", 5,
		12, false, false, 0,
		0, 0, 0, 0,
		-2, 0, 0,
	},
	{
		1, EQUIP_HEAD, "Eye Wound", "", 5,
		8, false, false, 0,
		0, 0, 0, 0,
		0, -1, 0,
	},
	{
		2, EQUIP_ARMS, "Broken Arm", "", 5,
		24, false, false, 0,
		0, -2, -2, 0,
		0, 0, 0,
	},
	{
		3, EQUIP_ARMS, "Deep Laceration", "", 5,
		8, false, true, 12,
		0, 0, 0, 0,
		0, 0, 0,
	},
	{
		4, EQUIP_CHEST, "Deep Laceration", "", 5,
		8, false, true, 12,
		0, 0, 0, 0,
		0, 0, 0,
	},
	{
		5, EQUIP_CHEST, "Strained Back", "", 5,
		10, false, false, 0,
		-1, 0, 0, -10,
		0, 0, 1,
	},
	{
		6, EQUIP_LEG, "Broken Leg", "", 5,
		24, false, false, 0,
		-1, -1, -1, 0,
		0, 0, 1,
	},
	{
		7, EQUIP_LEG, "Deep Laceration", "", 5,
		8, false, true, 12,
		0, 0, 0, 0,
		0, 0, 0,
	},
}

type Wound struct {
	id            int
	location      int // use EQUIP_X constants
	name          string
	description   string
	treatmentCost int

	healTime  int // days to heal
	treated   bool
	bleeding  bool
	bleedTime int // amount of turns -untreated- before a hit is assessed

	moveMod       int
	attackMod     int
	defMod        int
	encumbPenalty int

	spellcraftPenalty int
	viewRangePenalty  int
	travelPenalty     int
}

func genNewWound(loc int) Wound {
	var die Die
	locWounds := make([]Wound, 0, 0)

	for k := 0; k < len(allWounds); k++ {
		if allWounds[k].location == loc {
			locWounds = append(locWounds, allWounds[k])
		}
	}

	newWound := locWounds[die.rollxdx(1, len(locWounds))-1]

	newWound.healTime = (die.rollxdx(1, 4) + (newWound.healTime - 2))

	return newWound
}
