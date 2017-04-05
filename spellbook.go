// spellbook.go
package main

//import "fmt"
//import "strings"

type Spellbook struct {
	knownCount    int
	known         [20]int // max twenty known spells
	preparedCount int
	prepared      [10]int
}

// Spell Ideas (Lamentations)
// ==================================
// Cauterize - mutes a bleed hit taking it from the death total, still needs time to "heal", hurts like a B^#$&
// Fuze
// Divinate - Removes fog of war for the current grid?
// Protection circle - creates defensive perimeter of salt
// Invisibility - Remain undetected
// Corpse Candle - Acts as a floating torch for the duration of the mission
// Caustic Blood - A spray attack that can damage spirits, but costs one health
// Lucky Nickel - Shows ? in squares where items/gold are hidden
// Soul Inhale - Sucks soul points from Monster

// These need some atmosphere if I am going to use them:
// Blood Seek - Attack bonus
// Leather skin - Armor bonus?
