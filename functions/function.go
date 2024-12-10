package functions

import (
	"fmt"
	"hack/models"
)


func QuickSort(participants []models.Participant, low, high int) {
	if low < high {
		// partition the array
		p := partition(participants, low, high)
		QuickSort(participants, low, p-1)
		QuickSort(participants, p+1, high)
	}
}

func partition(participants []models.Participant, low, high int) int {
	pivot := participants[high].Score
	i := low - 1

	for j := low; j < high; j++ {
		if participants[j].Score >= pivot {
			i++
			participants[i], participants[j] = participants[j], participants[i]
		}
	}

	participants[i+1], participants[high] = participants[high], participants[i+1]
	return i + 1
}

// UpdateScore updates a participant's score.
func UpdateScore(participants []models.Participant, name string, newScore int) {
	for i := range participants {
		if participants[i].Name == name {
			participants[i].Score = newScore
			return
		}
	}
	fmt.Println("Participant not found!")
}

// DisplayLeaderboard prints the sorted leaderboard.
func DisplayLeaderboard(participants []models.Participant) {
	fmt.Println("\n--- Hackathon Leaderboard ---")
	for i, participant := range participants {
		fmt.Printf("%d. %s - %d\n", i+1, participant.Name, participant.Score)
	}
}

// CheckIfParticipantExists checks if a participant already exists by name
func CheckIfParticipantExists(participants []models.Participant, name string) bool {
	for _, p := range participants {
		if p.Name == name {
			return true
		}
	}
	return false
}

// AddParticipant adds a new participant if their name doesn't already exist
func AddParticipant(participants []models.Participant, name string, score int) ([]models.Participant, error) {
	if CheckIfParticipantExists(participants, name) {
		return participants, fmt.Errorf("participant with name %s already exists", name)
	}
	participants = append(participants, models.Participant{Name: name, Score: score})
	return participants, nil
}