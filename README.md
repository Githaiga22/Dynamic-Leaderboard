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

## Testing
To run tests:
```bash
go test ./functions
```
Sample output for test cases is included in the terminal logs.

## Example Demo
Below is an example of the leaderboard in action:

```bash
Enter the number of participants: 2
Enter name for participant 1: Allan
Enter score for participant Allan: 70
Enter name for participant 2: Brian
Enter score for participant Brian: 70.5

--- Hackathon Leaderboard ---
1. Allan - 70
2. Brian - 70.5

Do you want to update a score? (yes/no): yes
Enter the participant's name to update score: Brian
Enter the new score: 71

--- Hackathon Leaderboard ---
1. Brian - 71
2. Allan - 70
```

## Visual Demo
![Dynamic Leaderboard Demo](Demo/Demo.png)

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests for new features or improvements.

## License
This project is licensed under the [MIT License](LICENSE)

---
## Article
Explore the full explained codebase on [Dev.to](https://dev.to/githaiga22/mastering-algorithms-with-go-a-beginners-guide-to-sorting-small-data-sets-1nn9)
