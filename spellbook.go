// spellbook.go
package main

//import "fmt"		
//import "strings"

type Spellbook struct {
	knownCount int
	known [20]int 	// max twenty known spells
	preparedCount int
	prepared [10]int
}