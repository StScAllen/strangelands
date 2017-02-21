// dice.go

package main

import "math/rand"
import "time"

type Die struct {
	lastRoll int	
}

func (d *Die) roll(max int) (int){
	d.lastRoll = (rand.Intn(max) + 1) // 1 to max

	return d.lastRoll
}

func (d *Die) rollxdx(min int, max int) (int){
	d.lastRoll = (rand.Intn((max-min)+1) + min) // min to max

	return d.lastRoll
}

func roll3x6() (int) {
	var d1, d2, d3 Die
	d1.roll(6)
	d2.roll(6)
	d3.roll(6)
	
	return (d1.lastRoll + d2.lastRoll + d3.lastRoll)
}

func roll2x6() (int) {
	var d1, d2 Die
	d1.roll(6)
	d2.roll(6)
	
	return (d1.lastRoll + d2.lastRoll)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}