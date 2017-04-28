// missions.go

package main

//import "fmt"

var npc Character

var missions = []Mission 	{
								{0, 0, 1, 1, "Corpse Candle Haunts the Bog.", "A corpse candle draws our sheep into the bog to drown.", 2, 1, 2, []Phase{}, 0, 50, npc, 60, 90, 15, 15, 5, 0, 1},
								{1, 0, 1, 2, "Hung the wrong man.", "Now his corpse lingers in the cemetary!", 2, 1, 2, []Phase{}, 0, 60, npc, 60, 90, 15, 15, 0, 1, 1},
							}

type Mission struct {
	typeId					int
	instanceId				int
	complexity				int
	monsterType				int
	title					string	
	description				string
	phasesTotal				int
	currentPhase			int
	minimumPhases			int		// how many phases must be completed before monster can be faced
	phases					[]Phase
	missionBaseLocation		int
	crownReward				int
	apprenticeReward		Character
	startDays				int		// how many days mission is available to accept
	completeDays			int		// how many days until mission must be completed
	impactDays 				int 	// how many days until the village receives an impact from the quest being unsolved
	impactDaysLeft			int		// remaining days until next impact
	financialImpact			int
	livesImpact				int
	politicalImpact			int
}

type Phase	struct {
	id						int
	locationIndex			int
	itemRequired			Item
	puzzlePips				int
	researchPips			int
	description				string
	rewardId				int
	rewardItem				Item
	rewardExperience		int
	rewardCrowns			int
}

func genNewMission() (Mission) {
	var mission Mission

	game.missionInstanceId++	
	mission.instanceId = game.missionInstanceId

	
	return mission
}