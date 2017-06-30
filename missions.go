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

const STATUS_ACTIVE = 1		// mission is active
const STATUS_AVAIL	= 0		// mission is available to accept
const STATUS_NOSTART = 1	// mission start timed out
const STATUS_TIMED	= 2		// mission not completed in time
const STATUS_FAILED = 3		// mission failed	(killed by monster)
const STATUS_COMPLETE = 99

var monsterNames = []string {
								"Will-O-Wisp", "Revenant Corpse", "Minstrel Piper",
							}

var missionDescrips = [][]string	{
										{"None", "None"},
										{"Corpse Candle Haunts the Bog.", "A corpse candle draws our sheep into the bog to drown."},
										{"Hung the wrong man.", "Now his corpse lingers in the cemetary!"},
										{"Missing Child", "My child has been lured into the woods by a minstrel piper."},										
									}

var phaseDescrips = []string 	{
									"Its time to confront the darkness.",
									"A witness holds the key to the fright's lair.",
									"A cunning riddle contains secrets to the fright's whereabouts.",
									"The secret to the fright's power lies somewhere in this tome.",
								}									
									
var BLANK_MISSION = Mission{-1, 0, 0, 0, 0, "No Mission", "", 0, 0, 0, []Phase{}, 0, 0, 0, 0, 0, "", 0, 0, 0, 0, 0, 0, 0, STATUS_AVAIL}

var missions = []Mission 	{
								{0, 0, 1, 1, 1, "", "", 3, 1, 2, []Phase{}, 0, 50, 10, 0, 0, "", 60, 90, 15, 15, 5, 0, 1, STATUS_AVAIL},
								{1, 0, 1, 2, 2, "", "", 3, 1, 2, []Phase{}, 0, 60, 10, 0, 0, "", 60, 90, 15, 15, 0, 1, 1, STATUS_AVAIL},
								{2, 0, 1, 3, 3, "", "", 3, 1, 2, []Phase{}, 0, 60, 10, 1, 0, "", 60, 90, 15, 15, 0, 1, 1, STATUS_AVAIL},
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
	experienceReward		int
	apprenticeReward		int		// 1 true
	apprenticeRewardVariant	int		// id that determines type, (child, adult, girl, boy, etc.) boy = 1, girl = 2
	apprenticeRewardName		string 	// string name for apprentice reward (character name)
	startDays				int		// how many days mission is available to accept
	completeDays			int		// how many days until mission must be completed
	impactDays 				int 	// how many days until the village receives an impact from the quest being unsolved
	impactDaysLeft			int		// remaining days until next impact
	financialImpact			int
	livesImpact				int
	politicalImpact			int
	status					int		
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

func (miss * Mission) getDisplayString(detail int) (string) {
	dispString := ""
	
	if detail == 0 {
		dispString = miss.title
	} else if detail == 1 {
		dispString = fmt.Sprintf("%s    Reward: %v    Challenge: %v", packSpaceString(miss.title, 34), miss.crownReward, miss.complexity)
	}
	
	
	return dispString
}

func showMissionComplete() {
	rsp := ""

	clearConsole()
	
	fmt.Println(packSpaceStringCenter("┌─────────────╡ " + packSpaceStringCenter("Mission Complete!", 30) + " ╞─────────────┐", 76))
	fmt.Println("│ " + packSpaceString(fmt.Sprintf("Gold Earned: %v", mission.crownReward), 58) + " │")
	fmt.Println("│ " + packSpaceString(fmt.Sprintf("Exp Earned: %v", mission.experienceReward), 58) + " │")

	fmt.Println("└────────────────────────────────────────────────────────────┘")		

	fmt.Println()
	fmt.Printf("\nPress enter to exit.")
	
	fmt.Scanln(&rsp)
	
	character.crowns += mission.crownReward
	character.exp += mission.experienceReward
	
	mission = getBlankMission()
}

func (miss * Mission) viewMissionStatus() {
	exitFlag := false
	rsp := ""
	
	if miss.typeId == -1 {
		return
	}
	
	for !exitFlag {
		clearConsole()
		
		fmt.Println(packSpaceStringCenter("┌────────────────────╡ " + packSpaceStringCenter(miss.title, 30) + " ╞────────────────────┐", 76))
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│ " + packSpaceStringCenter(makeDialogString(miss.description), 72) + " │")
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		pack1 := packSpaceString("  Start Village: " + villages[miss.missionBaseLocation].name, 36)
		pack2 := fmt.Sprintf("Reward: %v", miss.crownReward)
		pack3 := ""
		if miss.apprenticeReward == 1 {
			pack3 = "  [Potential Apprentice]"
		}
		fmt.Println("│ " + packSpaceString(pack1 + pack2 + pack3, 72) + " │")		
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│ " + packSpaceString(fmt.Sprintf("  Phase: %v / %v", miss.currentPhase, miss.phasesTotal), 72) + " │")
		fmt.Println("│ " + packSpaceString(fmt.Sprintf("  Days Remaining: %v ", miss.completeDays), 72) + " │")	
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")	
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")		
		fmt.Println("│  " + packSpaceStringCenter(" ┌ " + packSpaceStringCenter("- - - Current Phase - - -", 64) + " ┐ ", 72) + "  │")
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│ " + packSpaceStringCenter(makeDialogString(miss.phases[miss.currentPhase-1].description), 72) + " │")
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│ " + packSpaceString("    Phase Location: " + villages[miss.phases[miss.currentPhase-1].locationIndex].name, 72) + " │")
		pack1 = packSpaceString(fmt.Sprintf("    Research Remaining: %v", miss.phases[miss.currentPhase-1].researchPips), 30)
		pack2 = packSpaceString(fmt.Sprintf("  Puzzles Remaining: %v", miss.phases[miss.currentPhase-1].puzzlePips), 30)
		fmt.Println("│ " + packSpaceString(pack1 + pack2, 72) + " │")
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│ " + packSpaceString(" ", 72) + " │")
		fmt.Println("│  " + packSpaceStringCenter(" └ " + packSpaceStringCenter("- - -", 64) + " ┘ ", 72) + "  │")
		fmt.Println("└──────────────────────────────────────────────────────────────────────────┘")		
		fmt.Printf("\nPress enter to exit.")
		
		fmt.Scanln(&rsp)
		
		exitFlag = true
	}

}

func (miss * Mission) viewAcceptDialog() (bool) {
	accepted := false
	
	exitFlag := false
	rsp := ""
	
	for !exitFlag {
		clearConsole()
		
		fmt.Println("*** " + miss.title + " ***")
		fmt.Println()
		fmt.Println(miss.description)
		fmt.Println()
		pack1 := packSpaceString(fmt.Sprintf("Monster: %s ", monsterNames[miss.monsterType-1]), 34)
		pack2 := fmt.Sprintf("Reward: %v crowns", miss.crownReward)
		fmt.Println(pack1 + pack2)			
		fmt.Println()	
		pack1 = packSpaceString(fmt.Sprintf("Difficulty: %v ", miss.complexity), 34)
		pack2 = fmt.Sprintf("Phases: %v ", miss.phasesTotal)
		fmt.Println(pack1 + pack2)			
		fmt.Println()
		pack1 = packSpaceString(fmt.Sprintf("Days to Accept: %v", miss.startDays), 34)		
		pack2 = fmt.Sprintf("Days to Complete: %v", miss.completeDays)
		fmt.Println(pack1 + pack2)	
		fmt.Println()
		
		fmt.Printf("\nDo you wish to accept this mission? ")
		
		fmt.Scanln(&rsp)
		
		if rsp == "y" {
			if mission.typeId != -1 && mission.status == STATUS_ACTIVE {
				showPause("You already have an active mission. You must complete \nthat mission before accepting a new one.")
			} else {
				exitFlag = true
				accepted = true			
			}

		} else if rsp == "n" {
			exitFlag = true
			accepted = false
		}
	}
	
	return accepted
}

func genNewMission(villageIndex int) (Mission) {
	var die Die

	numMissions := len(missions)
	missionIdx := (die.rollxdx(1, numMissions) - 1)
	
	tMission := missions[missionIdx]
	
	tMission.title = missionDescrips[tMission.txtIndex][TITLE]
	tMission.description = missionDescrips[tMission.txtIndex][DESC]	
	tMission.missionBaseLocation = villageIndex
	
	game.missionInstanceId++	
	tMission.instanceId = game.missionInstanceId
	
	if tMission.apprenticeReward == 1 {
		gender := die.rollxdx(1, 2)
		tempChar := getRandomApprentice(gender)
		
		tMission.apprenticeRewardName = tempChar.name
		tMission.apprenticeRewardVariant = tempChar.gender
	}
	
	var phase Phase
	
	phase.id = PHASE_PUZZLE
	phase.locationIndex = die.rollxdx(1, 8) - 1
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
	
	tMission.experienceReward += (phase.researchPips + phase.puzzlePips)
	
	tMission.phases = make([]Phase, 0, 0)
	tMission.phases = append(tMission.phases, phase)
			
	var phase2 Phase
	
	phase2.id = PHASE_RESEARCH
	phase2.locationIndex = die.rollxdx(1, 7) - 1
	phase2.itemRequiredId = -1
	phase2.puzzlePips = 0
	phase2.researchPips = die.rollxdx(1, 8) + 5
	phase2.descIndex = 0
	phase2.description = phaseDescrips[3]
	phase2.rewardId = 1
	phase2.rewardItemId = -1
	phase2.rewardExperience = 0
	phase2.rewardCrowns	= 0
	phase2.complete = 0	
	
	tMission.experienceReward += (phase.researchPips + phase.puzzlePips)
	tMission.phases = append(tMission.phases, phase2)
	
	var phase3 Phase
	
	phase3.id = PHASE_FIGHT
	phase3.locationIndex = villageIndex
	phase3.itemRequiredId = -1
	phase3.puzzlePips = 0
	phase3.researchPips = 0
	phase3.descIndex = 0
	phase3.description = phaseDescrips[0]
	phase3.rewardId = 1
	phase3.rewardItemId = -1
	phase3.rewardExperience = 0
	phase3.rewardCrowns	= 0
	phase3.complete = 0	
	
	tMission.phases = append(tMission.phases, phase3)	
	
	return tMission
}

func removeExpiredMissions() {
	if mission.typeId != -1 {
		mission.completeDays -= 1		
		if mission.completeDays < 0 && mission.status != STATUS_COMPLETE {
			mission.status = STATUS_TIMED
			log.addInfo("Mission timed out!")
			
			showPause("Failed mission in " + villages[mission.missionBaseLocation].name + ". Political favor has dropped.")
			// TODO:  Need to add a journal entry that this mission failed.
			
			mission = BLANK_MISSION
		}
	}

	for k := 0; k < len(villages); k++ {
		remove := -1
		if len(villages[k].missions) > 0 {
			for m := 0; m < len(villages[k].missions); m++ {
				villages[k].missions[m].startDays -= 1
				if villages[k].missions[m].startDays < 0 {
					remove = m
					break
				}
				if remove > -1 {
					villages[k].missions[m].status = STATUS_NOSTART
					villages[k].missions = append(villages[k].missions[:remove], villages[k].missions[remove+1:]...)
					
					// TODO:  Need to add a journal entry that this mission failed.
					showPause("Failed mission in " + villages[k].name + ". Political favor has dropped.")
				}
			}
		}
	} 	
}

func maybeAddMission() {
	var die Die
	chance := 0	
	count := 0
	
	for k := 0; k < len(villages); k++ {
		if len(villages[k].missions) > 0 {
			count++;
		}
	}
	
	chance += (8 - count) * 3
	chance += game.gameDay
	
	if chance < 1 {
		chance = 1
	}
	
	if (chance > 300 && count < 8) {
		chance = 300
	}
	
	// random chance to add mission, always add if count is < 1
	if (die.rollxdx(1, 1000) < chance) || (count < 1) {
		look := true
		counter := 0
		for look {
			roll := die.rollxdx(1, 8) - 1
			counter++
			
			// look for a village without a mission and add, if we check a few times and can't find an empty one then
			// add an additional mission to a random village			
			if (len(villages[roll].missions) < 1) || (counter >= 20) {
				villages[roll].missions = append(villages[roll].missions, genNewMission(roll))
				look = false
			}
		}
	}
}

func getBlankMission() (Mission) {
	var miss Mission 
	
	miss = BLANK_MISSION
	
	miss.instanceId = 0
	miss.typeId = -1
	miss.title = "No Mission"
	
	return miss
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
	miss.complexity, _ = strconv.Atoi(bits[3])
	miss.monsterType, _ = strconv.Atoi(bits[4])
	miss.txtIndex, _ = strconv.Atoi(bits[5])	
	miss.title = missionDescrips[miss.txtIndex][TITLE]
	miss.description = missionDescrips[miss.txtIndex][DESC]
	miss.phasesTotal, _ = strconv.Atoi(bits[6])
	miss.currentPhase, _ = strconv.Atoi(bits[7])
	miss.minimumPhases, _ = strconv.Atoi(bits[8])
	miss.missionBaseLocation, _ = strconv.Atoi(bits[9])
	miss.crownReward, _ = strconv.Atoi(bits[10])
	miss.experienceReward, _ = strconv.Atoi(bits[11])
	miss.apprenticeReward, _ = strconv.Atoi(bits[12])
	miss.apprenticeRewardVariant, _ = strconv.Atoi(bits[13])
	miss.apprenticeRewardName = bits[14]
	
	miss.startDays, _ = strconv.Atoi(bits[15])
	miss.completeDays, _ = strconv.Atoi(bits[16])
	miss.impactDays, _ = strconv.Atoi(bits[17])
	miss.impactDaysLeft, _ = strconv.Atoi(bits[18])
	miss.financialImpact, _ = strconv.Atoi(bits[19])	
	miss.livesImpact, _ = strconv.Atoi(bits[20])	
	miss.politicalImpact, _ = strconv.Atoi(bits[21])	

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
		tphase.description = phaseDescrips[tphase.descIndex]
		
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
	saveString += fmt.Sprintf("%v,", miss.experienceReward)

	saveString += fmt.Sprintf("%v,", miss.apprenticeReward)
	saveString += fmt.Sprintf("%v,", miss.apprenticeRewardVariant)
	saveString += fmt.Sprintf("%s,", miss.apprenticeRewardName)
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
		row += fmt.Sprintf("%v,", cphase.itemRequiredId)
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