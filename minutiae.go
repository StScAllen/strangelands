// minutiae.go

package main

import "fmt"

func showSkillsMinutiae() {

	clearConsole()
	var flag bool = true
	rsp := ""

	for flag {
		fmt.Println("--- Minutiae: Skills ---")
		fmt.Println("XXXXX")
		fmt.Println("")
		fmt.Println("")
		
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)
		
		flag = false
	}
}

func showTravelMinutiae() {

	clearConsole()
	var flag bool = true
	rsp := ""

	for flag {
		clearConsole()
		fmt.Println("--- Minutiae: Travel ---")
		fmt.Println("XXXXX")
		fmt.Println("")
		fmt.Println("")
		
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)
		
		flag = false
	}
}


func showAttributesMinutiae() {

	clearConsole()
	var flag bool = true
	rsp := ""

	for flag {
		fmt.Println("--- Minutiae: Attributes ---")
		fmt.Println("Fail forwards if you can, backwards is, afterall, just out of sight...")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=Perception=")
		fmt.Println("Perception primarily governs view distance on the battle map. It also provides")
		fmt.Println("a bonus to ranged combat aim.")
		fmt.Println("")
		fmt.Println("Score	Modifiers")
		fmt.Println(" 1 	View Range: 1	Aim: -2")
		fmt.Println(" 2 	View Range: 1 	Aim: -1")
		fmt.Println(" 3 	View Range: 2 	Aim: -1")
		fmt.Println(" 4 	View Range: 2 	Aim:  0")
		fmt.Println(" 5 	View Range: 3 	Aim:  0")
		fmt.Println(" 6 	View Range: 3 	Aim: +1")
		fmt.Println(" 7 	View Range: 4 	Aim: +1")
		fmt.Println(" 8 	View Range: 4 	Aim: +2")
		fmt.Println(" 9 	View Range: 5 	Aim: +2")
		fmt.Println(" 10 	View Range: 6 	Aim: +3")
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)

		clearConsole()
		fmt.Println("--- Minutiae: Attributes ---")
		fmt.Println("Ultimately, we all carry our own burdens.")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=Strength=")
		fmt.Println("Strength governs health and encumberance. It is also a prerequisite for certain")
		fmt.Println("weapons and armor.")
		fmt.Println("")
		fmt.Println("Score	Modifiers")
		fmt.Println(" 1 	Health: 1	Carry: 1 stone 	")
		fmt.Println(" 2 	Health: 2 	Carry: 2 stone	")
		fmt.Println(" 3 	Health: 3 	Carry: 3 stone	")
		fmt.Println(" 4 	Health: 4 	Carry: 4 stone	")
		fmt.Println(" 5 	Health: 5 	Carry: 5 stone	")
		fmt.Println(" 6 	Health: 6 	Carry: 6 stone	")
		fmt.Println(" 7 	Health: 7 	Carry: 7 stone	")
		fmt.Println(" 8 	Health: 8 	Carry: 8 stone	")
		fmt.Println(" 9 	Health: 9 	Carry: 9 stone 	")
		fmt.Println(" 10 	Health: 12 	Carry: 10 stone	")
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)

		clearConsole()
		fmt.Println("--- Minutiae: Attributes ---")
		fmt.Println("Quick of feet, lithe, agile? Thats for circus folk and cutthroats.")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=Agility=")
		fmt.Println("Agility governs the number of actions you may take on the battle map per turn.")
		fmt.Println("It provides base defense adjustment and is a prerequisite for certain weapons and armor.")
		fmt.Println("")
		fmt.Println("Score	Modifiers")
		fmt.Println(" 1 	Actions: 4		Defense: -2")
		fmt.Println(" 2 	Actions: 5 		Defense: -1")
		fmt.Println(" 3 	Actions: 6 		Defense: -1")
		fmt.Println(" 4 	Actions: 7 		Defense: 0")
		fmt.Println(" 5 	Actions: 8 		Defense: 0")
		fmt.Println(" 6 	Actions: 9 		Defense: +1")
		fmt.Println(" 7 	Actions: 10		Defense: +1")
		fmt.Println(" 8 	Actions: 11 	Defense: +2")
		fmt.Println(" 9 	Actions: 12 	Defense: +2")
		fmt.Println(" 10 	Actions: 13 	Defense: +3")
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)

		clearConsole()
		fmt.Println("--- Minutiae: Attributes ---")
		fmt.Println("A strong mind encourages truth and vitatility, a weak one encourages fantasy and whimsy.")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=Intellect=")
		fmt.Println("Intellect governs the number of spells you can prepare as well as max skill")
		fmt.Println(" level.")
		fmt.Println("")
		fmt.Println("Score	Modifiers")
		fmt.Println(" 1 	Spells: 0		Max Skill: 1")
		fmt.Println(" 2 	Spells: 1 		Max Skill: 1")
		fmt.Println(" 3 	Spells: 2 		Max Skill: 2")
		fmt.Println(" 4 	Spells: 2 		Max Skill: 3")
		fmt.Println(" 5 	Spells: 3 		Max Skill: 3")
		fmt.Println(" 6 	Spells: 4 		Max Skill: 4")
		fmt.Println(" 7 	Spells: 5 		Max Skill: 4")
		fmt.Println(" 8 	Spells: 6 		Max Skill: 5")
		fmt.Println(" 9 	Spells: 7 		Max Skill: 5")
		fmt.Println(" 10 	Spells: 8 		Max Skill: 6")
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)

		clearConsole()
		fmt.Println("--- Minutiae: Attributes ---")
		fmt.Println("A silver tongue has made more men rich than wisdom or hard work.")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=Charm=")
		fmt.Println("Charm primarily impacts how many apprentices you can have, it also provides ")
		fmt.Println("a cap to politics and is a factor in determing soul. It also impacts shop prices.")
		fmt.Println("")
		fmt.Println("Score	Modifiers")
		fmt.Println(" 1 	Apprentices: 0		Max Political Level: 1		Soul: 0")
		fmt.Println(" 2 	Apprentices: 0 		Max Political Level: 2		Soul: 1")
		fmt.Println(" 3 	Apprentices: 0 		Max Political Level: 3		Soul: 2")
		fmt.Println(" 4 	Apprentices: 1 		Max Political Level: 4		Soul: 3")
		fmt.Println(" 5 	Apprentices: 1 		Max Political Level: 5		Soul: 4")
		fmt.Println(" 6 	Apprentices: 2 		Max Political Level: 6		Soul: 5")
		fmt.Println(" 7 	Apprentices: 3 		Max Political Level: 7		Soul: 6")
		fmt.Println(" 8 	Apprentices: 4 		Max Political Level: 8		Soul: 7")
		fmt.Println(" 9 	Apprentices: 5 		Max Political Level: 9		Soul: 8")
		fmt.Println(" 10 	Apprentices: 6 		Max Political Level: 10		Soul: 10")
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)

		clearConsole()
		fmt.Println("--- Minutiae: Attributes ---")
		fmt.Println("Cunning, clever, willy, sharp. Con-men share these qualities too.")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("=Guile=")
		fmt.Println("Guile provides a modifier for all skill rolls and is a factor in determining ")
		fmt.Println("soul and Magical Defense")
		fmt.Println("")
		fmt.Println("Score	Modifiers")
		fmt.Println(" 1 	Skill Bonus: -2		Soul: 1		Magic Def: -3")
		fmt.Println(" 2 	Skill Bonus: -1 	Soul: 2		Magic Def: -3")
		fmt.Println(" 3 	Skill Bonus: -1 	Soul: 3		Magic Def: -2")
		fmt.Println(" 4 	Skill Bonus: 0 		Soul: 4		Magic Def: -1")
		fmt.Println(" 5 	Skill Bonus: 0 		Soul: 5		Magic Def: 0")
		fmt.Println(" 6 	Skill Bonus: 1 		Soul: 6		Magic Def: +1")
		fmt.Println(" 7 	Skill Bonus: 2 		Soul: 7		Magic Def: +2")
		fmt.Println(" 8 	Skill Bonus: 3 		Soul: 8		Magic Def: +3")
		fmt.Println(" 9 	Skill Bonus: 4 		Soul: 9		Magic Def: +4")
		fmt.Println(" 10 	Skill Bonus: 5 		Soul: 10	Magic Def: +5")
		fmt.Println("")
		fmt.Println("[ENTER] to continue.")
		fmt.Scanln(&rsp)

		flag = false
	}
}

func openTownMinutiae() {
	rsp := ""

	clearConsole()
	fmt.Println("")
	fmt.Println("--- Minutiae: Town ---")
	fmt.Println("Dusty, corrupt, toiling. The only remaining beacons in the vast dark.")
	fmt.Println("╓")
	fmt.Println("║1. Character Minutiae")
	fmt.Println("║2. City Minutiae")
	fmt.Println("║3. Keep Minutiae")
	fmt.Println("║4. Mission Minutiae")
	fmt.Println("║5. Misc Minutiae")
	fmt.Println("╙")

	fmt.Scanln(&rsp)

}

// main minutiae entry point
func openMinutiae() {
	rsp := ""

	clearConsole()
	fmt.Println("")
	fmt.Println("--- Minutiae: Main ---")
	fmt.Println("Regardless of purported significance, its all just minutiae.")
	fmt.Println("╓")
	fmt.Println("║1. Character Minutiae")
	fmt.Println("║2. City Minutiae")
	fmt.Println("║3. Keep Minutiae")
	fmt.Println("║4. Mission Minutiae")
	fmt.Println("║5. Misc Minutiae")
	fmt.Println("╙")

	fmt.Scanln(&rsp)
}
