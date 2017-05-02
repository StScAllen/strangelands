// missions.go

package main
import "strings"
import "fmt"
import "strconv"

var blankNPC Character

const TITLE = 0
const DESC = 1

const PHASE_FIGHT = 0
const PHASE_RESEARCH = 1
const PHASE_PUZZLE = 2

var missionDescrips = [][]string	{
										{"None", "None"},
										{"Corpse Candle Haunts the Bog.", "A corpse candle draws our sheep into the bog to drown."},
										{"Hung the wrong man.", "Now his corpse lingers in the cemetary!"},
									}

var phaseDescrips = []string 	{
									"Its time to confront the darkness.",
									"A witness holds the key to the beast's lair.",
									"The secret to the beasts power lies somewhere in this tome.",
								}									
									
var BLANK_MISSION = Mission{-1, 0, 0, 0, 0, "", "", 0, 0, 0, []Phase{}, 0, 0, 0, 0, "", 0, 0, 0, 0, 0, 0, 0}

var missions = []Mission 	{
								{0, 0, 1, 1, 1, "", "", 2, 1, 2, []Phase{}, 0, 50, 0, 0, "", 60, 90, 15, 15, 5, 0, 1},
								{1, 0, 1, 2, 2, "", "", 2, 1, 2, []Phase{}, 0, 60, 0, 0, "", 60, 90, 15, 15, 0, 1, 1},
							}

type Mission struct {
	typeId					int
	instanceId				int
	complexity				int
	monsterType				int
	txtIndex 				int
	title					string	
	description				string
	phasesTotal				int
	currentPhase			int
	minimumPhases			int		// how many phases must be completed before monster can be faced
	phases					[]Phase
	missionBaseLocation		int
	crownReward				int
	apprenticeReward		int		// 1 true
	apprenticeRewardVariant	int		// id that determines type, (child, adult, girl, boy, etc.)
	appenticeRewardName		string 	// string name for apprentice reward (character name)
	startDays				int		// how many days mission is available to accept
	completeDays			int		// how many days until mission must be completed
	impactDays 				int 	// how many days until the village receives an impact from the quest being unsolved
	impactDaysLeft			int		// remaining days until next impact
	financialImpact			int
	livesImpact				int
	politicalImpact			int
}

type Phase struct {
	id						int
	locationIndex			int
	itemRequiredId			int
	puzzlePips				int
	researchPips			int
	descIndex				int
	description				string
	rewardId				int
	rewardItemId			int
	rewardExperience		int
	rewardCrowns			int
	complete 				int			// 0 - incomplete, 1 - complete
}

func genNewMission() (Mission) {
	var die Die

	numMissions := len(missions)
	missionIdx := (die.rollxdx(1, numMissions) - 1)
	
	tMission := missions[missionIdx]
	
	tMission.title = missionDescrips[tMission.txtIndex][TITLE]
	tMission.description = missionDescrips[tMission.txtIndex][DESC]	
	
	game.missionInstanceId++	
	tMission.instanceId = game.missionInstanceId
	
	var phase Phase
	
	phase.id = PHASE_PUZZLE
	phase.locationIndex = die.rollxdx(1, 7)
	phase.itemRequiredId = -1
	phase.puzzlePips = die.rollxdx(1, 8) + 5
	phase.researchPips = 0
	phase.descIndex = 2
	phase.description = phaseDescrips[2]
	phase.rewardId = 1
	phase.rewardItemId = -1
	phase.rewardExperience = 10
	phase.rewardCrowns	= 25
	phase.complete = 0
	
	tMission.phases = make([]Phase, 0, 0)
	tMission.phases = append(tMission.phases, phase)
	
	var phase2 Phase
	
	phase2.id = PHASE_FIGHT
	phase2.locationIndex = 0
	phase2.itemRequiredId = -1
	phase2.puzzlePips = 0
	phase2.researchPips = 0
	phase2.descIndex = 0
	phase2.description = phaseDescrips[0]
	phase2.rewardId = 1
	phase2.rewardItemId = -1
	phase2.rewardExperience = 0
	phase2.rewardCrowns	= 0
	phase2.complete = 0	
	
	tMission.phases = append(tMission.phases, phase2)
	
	return tMission
}

func unpackMissionBlock(block string) (int, Mission) {
	var miss Mission
	
	lines := strings.Split(block, "◄")
	bits := strings.Split(lines[0], ",")

	if bits[0] != BLOCK_MISSION {
		log.addError("Expected Mission - row not found.")
		fmt.Println("Mission row not found!")
		return -1, miss
	}

	miss.typeId, _ = strconv.Atoi(bits[1])
	miss.instanceId, _ = strconv.Atoi(bits[2])
	miss.monsterType, _ = strconv.Atoi(bits[3])
	miss.txtIndex, _ = strconv.Atoi(bits[4])	
	miss.title = ""
	miss.description = ""
	miss.phasesTotal, _ = strconv.Atoi(bits[5])
	miss.currentPhase, _ = strconv.Atoi(bits[6])
	miss.minimumPhases, _ = strconv.Atoi(bits[7])
	miss.missionBaseLocation, _ = strconv.Atoi(bits[8])
	miss.crownReward, _ = strconv.Atoi(bits[9])
	miss.apprenticeReward, _ = strconv.Atoi(bits[10])
	miss.apprenticeRewardVariant, _ = strconv.Atoi(bits[11])
	miss.appenticeRewardName = bits[12]
	
	miss.startDays, _ = strconv.Atoi(bits[13])
	miss.completeDays, _ = strconv.Atoi(bits[14])
	miss.impactDays, _ = strconv.Atoi(bits[15])
	miss.impactDaysLeft, _ = strconv.Atoi(bits[16])
	miss.financialImpact, _ = strconv.Atoi(bits[17])	
	miss.livesImpact, _ = strconv.Atoi(bits[18])	
	miss.politicalImpact, _ = strconv.Atoi(bits[19])	

	miss.phases = make([]Phase, 0, 0)
	
	// reload phases...
	for k := 1; k < len(lines)-1; k++ {
		pbits := strings.Split(lines[k], ",")
		var tphase Phase
		
		if (len(lines[k]) < 5) {
			continue
		}
		
		// pbits[0] should be PHASE
		tphase.id, _ = strconv.Atoi(pbits[1])
		tphase.locationIndex, _ = strconv.Atoi(pbits[2])
		
		tphase.itemRequiredId, _ = strconv.Atoi(pbits[3])
		tphase.puzzlePips, _ = strconv.Atoi(pbits[4])
		tphase.researchPips, _ = strconv.Atoi(pbits[5])
		tphase.descIndex, _ = strconv.Atoi(pbits[6])
		tphase.description = ""
		
		tphase.rewardId, _ = strconv.Atoi(pbits[7])
		tphase.rewardItemId, _ = strconv.Atoi(pbits[8])
		tphase.rewardExperience, _ = strconv.Atoi(pbits[9])	
		tphase.rewardCrowns, _ = strconv.Atoi(pbits[10])
		tphase.complete, _ = strconv.Atoi(pbits[11])		
		
		miss.phases = append(miss.phases, tphase)
	}

	return 1, miss
}

func (miss *Mission) getSaveString() string {
	saveString := BLOCK_MISSION + ","

	saveString += fmt.Sprintf("%v,", miss.typeId)
	saveString += fmt.Sprintf("%v,", miss.instanceId)
	saveString += fmt.Sprintf("%v,", miss.complexity)
	saveString += fmt.Sprintf("%v,", miss.monsterType)
	saveString += fmt.Sprintf("%v,", miss.txtIndex)
	saveString += fmt.Sprintf("%v,", miss.phasesTotal)
	saveString += fmt.Sprintf("%v,", miss.currentPhase)
	saveString += fmt.Sprintf("%v,", miss.minimumPhases)
	saveString += fmt.Sprintf("%v,", miss.missionBaseLocation)
	saveString += fmt.Sprintf("%v,", miss.crownReward)
	saveString += fmt.Sprintf("%v,", miss.apprenticeReward)
	saveString += fmt.Sprintf("%v,", miss.apprenticeRewardVariant)
	saveString += fmt.Sprintf("%s,", miss.appenticeRewardName)
	saveString += fmt.Sprintf("%v,", miss.startDays)
	saveString += fmt.Sprintf("%v,", miss.completeDays)
	saveString += fmt.Sprintf("%v,", miss.impactDays)
	saveString += fmt.Sprintf("%v,", miss.impactDaysLeft)
	saveString += fmt.Sprintf("%v,", miss.financialImpact)
	saveString += fmt.Sprintf("%v,", miss.livesImpact)
	saveString += fmt.Sprintf("%v,", miss.politicalImpact)	

	saveString += "◄" // end line
	
	for k := 0; k < miss.phasesTotal; k++ {
		cphase := miss.phases[k]
		row := "PHASE,"

		row += fmt.Sprintf("%v,", cphase.id)
		row += fmt.Sprintf("%v,", cphase.locationIndex)
		row += fmt.Sprintf("%v,", cphase.puzzlePips)
		row += fmt.Sprintf("%v,", cphase.researchPips)
		row += fmt.Sprintf("%v,", cphase.descIndex)
		row += fmt.Sprintf("%v,", cphase.rewardId)
		row += fmt.Sprintf("%v,", cphase.rewardItemId)
		row += fmt.Sprintf("%v,", cphase.rewardExperience)
		row += fmt.Sprintf("%v,", cphase.rewardCrowns)
		row += fmt.Sprintf("%v,", cphase.complete)		
		
		row += "◄"
		
		saveString += row
	}

	saveString += "■"

	return saveString
}