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

func TestPartitionDescendingOrder(t *testing.T) {
	participants := []models.Participant{
		{Name: "Alice", Score: 50},
		{Name: "Charlie", Score: 60},
		{Name: "Bob", Score: 70},
		{Name: "David", Score: 80},
	}

	expectedPivotIndex := 0
	expectedParticipants := []models.Participant{
		{Name: "David", Score: 80},
		{Name: "Charlie", Score: 60},
		{Name: "Bob", Score: 70},
		{Name: "Alice", Score: 50},
	}

	actualPivotIndex := partition(participants, 0, len(participants)-1)

	if actualPivotIndex != expectedPivotIndex {
		t.Errorf("Expected pivot index %d, but got %d", expectedPivotIndex, actualPivotIndex)
	}

	for i, participant := range participants {
		if participant != expectedParticipants[i] {
			t.Errorf("Expected participant %v at index %d, but got %v", expectedParticipants[i], i, participant)
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

func TestCheckIfParticipantExists(t *testing.T) {
	participants := []models.Participant{
		{Name:"Alice", Score: 85},
		{Name:"Bob", Score: 70},
	}

	// Test existence
	if !CheckIfParticipantExists(participants, "Alice") {
		t.Errorf("Expected Alice to exist, but got false")
	}

	if CheckIfParticipantExists(participants, "Charlie") {
		t.Errorf("Expected Charlie not to exist, but got true")
	}
}
