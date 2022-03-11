package main

import "fmt"

/*
	Create a combat function that takes the player's
	current health and the amount of damage recieved,
	and returns the player's new health. Health can't
	be less than 0.
*/


func combat(health, damage float64) float64 {
	new_health := health - damage
	if new_health >= 0 {
		return new_health
	}
	return 0
}

func main() {
	solution := combat(100, 5)
	fmt.Println(solution)
}