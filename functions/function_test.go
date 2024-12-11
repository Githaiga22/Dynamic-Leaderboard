package functions

import (
	"testing"

	"hack/models"
)

func TestQuickSort(t *testing.T) {
	participants := []models.Participant{
		{Name: "Alice", Score: 50},
		{Name: "Bob", Score: 70},
		{Name: "Charlie", Score: 60},
		{Name: "David", Score: 80},
	}

	expected := []models.Participant{
		{Name: "David", Score: 80},
		{Name: "Bob", Score: 70},
		{Name: "Charlie", Score: 60},
		{Name: "Alice", Score: 50},
	}

	QuickSort(participants, 0, len(participants)-1)

	for i, participant := range participants {
		if participant != expected[i] {
			t.Errorf("Expected participant %v at index %d, but got %v", expected[i], i, participant)
		}
	}
}

func TestPartition(t *testing.T) {
	participants := []models.Participant{
		{Name: "Alice", Score: 85},
		{Name: "Bob", Score: 70},
		{Name: "Charlie", Score: 90},
		{Name: "Diana", Score: 75},
	}

	pivotIndex := partition(participants, 0, len(participants)-1)
	pivotScore := participants[pivotIndex].Score

	// Validate partitioning logic
	for i := 0; i < pivotIndex; i++ {
		if participants[i].Score < pivotScore {
			t.Errorf("Element %v before pivot is less than pivot score %d", participants[i], pivotScore)
		}
	}
	for i := pivotIndex + 1; i < len(participants); i++ {
		if participants[i].Score >= pivotScore {
			t.Errorf("Element %v after pivot is greater than or equal to pivot score %d", participants[i], pivotScore)
		}
	}
}

func TestUpdateScore_ExistingParticipant(t *testing.T) {
	participants := []models.Participant{
		{Name: "Alice", Score: 50},
		{Name: "Bob", Score: 70},
		{Name: "Charlie", Score: 60},
	}

	UpdateScore(participants, "Bob", 75)

	expected := []models.Participant{
		{Name: "Alice", Score: 50},
		{Name: "Bob", Score: 75},
		{Name: "Charlie", Score: 60},
	}

	for i, participant := range participants {
		if participant != expected[i] {
			t.Errorf("Expected participant %v at index %d, but got %v", expected[i], i, participant)
		}
	}
}

// func TestCheckIfParticipantExists(t *testing.T) {
// 	participants := []models.Participant{
// 		{Name: "Alice", Score: 85},
// 		{Name: "Bob", Score: 70},
// 	}

// 	// Test existence
// 	if !CheckIfParticipantExists(participants, "Alice") {
// 		t.Errorf("Expected Alice to exist, but got false")
// 	}

// 	if CheckIfParticipantExists(participants, "Charlie") {
// 		t.Errorf("Expected Charlie not to exist, but got true")
// 	}
// }
