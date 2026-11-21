package main

import "fmt"

func main() {

	// Students
	// (using a slice)
	students := []string{"Alice", "Bob", "Carol", "Dave"}
	// Scores
	// (using a  map)
	scores := map[string][]int{
		"Alice": {90, 85, 92},
		"Bob":   {70, 65, 80},
		"Carol": {100, 95, 98},
		"Dave":  {40, 55, 60},
	}
	//

	printReport(students, scores)

}

func average(scores []int) float64 {

	sumOfScore := 0
	for _, score := range scores {
		sumOfScore += score
	}

	lengthOfScores := len(scores)

	if lengthOfScores < 1 {
		return 0.0
	}
	avgScore := float64(sumOfScore) / float64(lengthOfScores)
	return avgScore
}

func letterGrade(avg float64) string {
	if avg >= 90 {
		return "A"
	} else if avg >= 80 {
		return "B"
	} else if avg >= 70 {
		return "C"
	} else if avg >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func printReport(students []string, scores map[string][]int) {

	grandTotal := 0.0
	for _, student := range students {
		score := scores[student]

		// Getting the average
		avg := average(score)

		grandTotal += avg
		// Getting the grade
		grade := letterGrade(avg)

		fmt.Printf("%s - avg: %.2f grade: %s\n", student, avg, grade)
	}

	studentCount := len(students)

	classAvg := grandTotal / float64(studentCount)
	fmt.Printf("class average: %.2f\n", classAvg)

}
