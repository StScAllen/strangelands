package main

import "fmt"
import "os"
import "os/exec"			

var character Character
var apprentice Character
var gameDay int

func clearConsole(){
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func showGameMenu() (string){

	clearConsole()
	
	fmt.Println("Main Menu")
	fmt.Printf("Day: %v \n", gameDay)
	fmt.Println("------------")
	fmt.Println("1. Visit Town")
	fmt.Println("2. Rest at Keep")
	fmt.Println("3. Status")
	fmt.Println("4. Missions")
	fmt.Println("")
	fmt.Println("5. Bestiary")
	fmt.Println("6. Quit")
	fmt.Println("")
	fmt.Println("Select an Option:  ")
	
	rsp := ""
	fmt.Scanln(&rsp)
	
	return rsp
}

func showTopMenu() (string){

	clearConsole()
	
	fmt.Println("Blather   (v" + VERSION + ")")
	fmt.Println("------------")
	fmt.Println("1. New Game")
	fmt.Println("2. Load Game")
	fmt.Println("3. Quit")
	fmt.Println("")
	fmt.Println("Select an Option:  ")
	
	rsp := ""
	fmt.Scanln(&rsp)
	
	return rsp
}

func init(){
	gameDay = 1

}

func main() {
	rsp := ""
	rsp = showTopMenu()
	
	if (rsp == "1") {	// new game, make a character
		rsp = "n"
		for rsp != "y" && rsp != "Y" {
			character = createCharacter()
			character.printCharacter(0)
		
			fmt.Print("\nUse this character? ")
			fmt.Scanln(&rsp)
			clearConsole()
		}	
	} else if (rsp == "2"){
		character = loadGame()
		character.printCharacter(1)
	} else if (rsp == "3"){
		return
	}

	clearConsole()

	rsp = ""

	for rsp != "6" {
		rsp = showGameMenu()
		
		if (rsp == "1") {
			goShop()
		} else if (rsp == "2") {
			character.hp = character.maxhp
			character.save()
		} else if (rsp == "3") {
			chooseAdventure()
			adventure()
		}
	}
}

