package main

import (
	"fmt"
	h "hack/functions"
	"hack/models"
)

func main() {
	// Initializing an empty slice to store participants
	var participants []models.Participant

	// Function to add participants and check for duplicates
	addParticipant := func(name string, score int) {
		// Check if participant with the same name already exists
		for _, p := range participants {
			if p.Name == name {
				fmt.Printf("Participant with name %s already exists.\n", name)
				return
			}
		}
		// If no duplicate, add the participant
		participants = append(participants, models.Participant{Name: name, Score: score})
	}

	// Loop to accept participant input
	var numParticipants int
	fmt.Println("Enter the number of participants:")
	fmt.Scanln(&numParticipants)

	for i := 0; i < numParticipants; i++ {
		var name string
		var score int

		// Read participant name
		fmt.Printf("Enter name for participant %d: ", i+1)
		fmt.Scanln(&name)

		// Read participant score
		fmt.Printf("Enter score for participant %s: ", name)
		fmt.Scanln(&score)

		// Add participant to the list
		addParticipant(name, score)
	}

	// Display initial leaderboard
	h.DisplayLeaderboard(participants)

	// Ask for score updates
	var updateName string
	var newScore int

	for {
		fmt.Println("\nDo you want to update a score? (yes/no): ")
		var response string
		fmt.Scanln(&response)
		if response == "no" {
			break
		}

		// Read participant name for score update
		fmt.Println("Enter the participant's name to update score: ")
		fmt.Scanln(&updateName)

		// Read the new score
		fmt.Println("Enter the new score: ")
		fmt.Scanln(&newScore)

		// Update the score
		h.UpdateScore(participants, updateName, newScore)

		// Display updated leaderboard
		h.DisplayLeaderboard(participants)
	}

	// Sort the leaderboard
	h.QuickSort(participants, 0, len(participants)-1)

	// Display final leaderboard
	h.DisplayLeaderboard(participants)
}
