# Dynamic Leaderboard

## Introduction
The **Dynamic Leaderboard** is a simple Go application designed to manage and sort participant rankings for events like hackathons. It allows users to input participants, update their scores in real-time, and view the leaderboard dynamically sorted by score. This project demonstrates the implementation of sorting algorithms and highlights their real-world application in small datasets.

## Features
- Add participants and their scores.
- Dynamically update scores for existing participants.
- Automatically sort the leaderboard based on scores.
- Demonstrates both sorting logic and testing in Go.

## Getting Started
### Prerequisites
- Go (1.20 or later)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Githaiga22/Dynamic-Leaderboard.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Dynamic-Leaderboard
   ```
3. Initialize the Go module:
   ```bash
   go mod tidy
   ```

### Running the Application
To run the application, execute:
```bash
go run .
```
Follow the prompts to add participants, update scores, and view the leaderboard.


## Directory Structure
```
Dynamic-Leaderboard/
├── functions/
│   ├── function.go          # Contains core logic (e.g., sorting, score updates)
│   ├── function_test.go     # Unit tests for the core logic
├── models/
│   └── model.go             # Defines participant structure
├── main.go                  # Entry point for the application
├── go.mod                   # Go module definition
└── README.md                # Project documentation
```

## Key Files and Their Purpose
1. **main.go**
   - The main entry point for the application. It handles user input and outputs the leaderboard.

2. **functions/function.go**
   - Contains the logic for sorting participants and updating scores.

3. **models/model.go**
   - Defines the `Participant` struct used throughout the application.

4. **functions/function_test.go**
   - Contains test cases to ensure the sorting and updating logic works correctly.

## How It Works
1. **Adding Participants**
   Users are prompted to input the number of participants, their names, and their scores. These details are stored in a list of `Participant` structs.

2. **Updating Scores**
   Users can update the score for a specific participant. The leaderboard is re-sorted after each update.

3. **Sorting Logic**
   The leaderboard is sorted using a simple partitioning technique, ensuring the highest scores are listed first.
