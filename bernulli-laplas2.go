package main

import (
	"fmt"
	"math"
)

// Calculate the factorial of a number
func factorial(n int) float64 {
	if n <= 1 {
		return 1
	}
	return float64(n) * factorial(n-1)
}

// Calculates the probability using the Bernoulli formula
func bernoulliFormula(n int, events []int, probabilities []float64) float64 {
	// Calculate the numerator (n!)
	numerator := factorial(n)

	// Calculate the denominator (n1! * n2! * ... * nk!)
	denominator := 1.0
	for _, event := range events {
		denominator *= factorial(event)
	}

	probabilityProduct := 1.0
	for i, event := range events {
		probabilityProduct *= math.Pow(probabilities[i], float64(event))
	}

	result := (numerator / denominator) * probabilityProduct
	return result
}

// Calculates the probability using the Moivre-Laplace formula
func moivreLaplaceFormula(n int, k int, p float64) float64 {
	q := 1.0 - p
	x := (float64(k) - float64(n)*p) / math.Sqrt(float64(n)*p*q)
	phi := (1.0 / math.Sqrt(2*math.Pi)) * math.Exp(-math.Pow(x, 2)/2)
	result := phi / math.Sqrt(float64(n)*p*q)
	return result
}

func main() {
	var n, k int
	fmt.Print("Введите количество экспериментов (n): ")
	fmt.Scan(&n)

	fmt.Print("Введите количество вероятностей (k): ")
	fmt.Scan(&k)

	// Initialize slices to store events and probabilities
	events := make([]int, k)
	probabilities := make([]float64, k)

	// Get the number of events for each probability
	fmt.Println("Введите количество событий для каждой вероятности:")
	sum := 0
	for i := 0; i < k; i++ {
		fmt.Printf("n%d: ", i+1)
		fmt.Scan(&events[i])
		sum += events[i]
	}

	// Validate that the sum of events equals n
	if sum != n {
		fmt.Printf("Ошибка: Сумма всех событий (%d) должна равняться количеству экспериментов (%d)\n", sum, n)
		return
	}

	// Get the probability values
	fmt.Println("Введите вероятность для каждого события (от 0 до 1):")
	sumProb := 0.0
	for i := 0; i < k; i++ {
		fmt.Printf("P%d: ", i+1)
		fmt.Scan(&probabilities[i])

		if probabilities[i] < 0 || probabilities[i] > 1 {
			fmt.Println("Ошибка: Вероятность должна быть в диапазоне от 0 до 1")
			return
		}
		sumProb += probabilities[i]
	}

	// Validate that the sum of probabilities equals 1 (within a small margin of error for floating-point)
	if math.Abs(sumProb-1.0) > 0.0001 {
		fmt.Printf("Предупреждение: Сумма вероятностей (%.4f) должна быть приблизительно равна 1\n", sumProb)
	}

	// Calculate the result using the Bernoulli formula
	bernoulliResult := bernoulliFormula(n, events, probabilities)
	fmt.Printf("\nРезультат по формуле Бернулли: %.8f\n", bernoulliResult)

	// Calculate the result using the Moivre-Laplace formula
	moivreLaplaceResult := moivreLaplaceFormula(n, events[0], probabilities[0])
	fmt.Printf("Результат по формуле Муавра-Лапласа: %.8f\n", moivreLaplaceResult)

	// Display the Bernoulli formula with values
	fmt.Print("\nРешение по формуле Бернулли:\n(")
	fmt.Print(n)
	fmt.Print("!/(")
	for i, event := range events {
		// 'if' to not display * in front of first number
		if i > 0 {
			fmt.Print("*")
		}
		fmt.Print(event, "!")
	}
	fmt.Print("))")

	for i, prob := range probabilities {
		fmt.Printf("*%.2f^%d", prob, events[i])
	}

	fmt.Printf(" = %.8f\n", bernoulliResult)

	// Display the Moivre-Laplace formula
	fmt.Print("\nРешение по формуле Муавра-Лапласа:\n")
	x := (float64(events[0]) - float64(n)*probabilities[0]) / (math.Sqrt(float64(n) * probabilities[0] * (1 - probabilities[0])))
	fmt.Printf("x = (%.0f - %d*%.2f)/sqrt(%d*%.2f*(1-%.2f)) = %.2f\n", float64(events[0]), n, probabilities[0], n, probabilities[0], 1-probabilities[0], x)
	phi := (1.0 / math.Sqrt(2*math.Pi)) * math.Exp(-math.Pow((float64(events[0])-float64(n)*probabilities[0])/math.Sqrt(float64(n)*probabilities[0]*(1-probabilities[0])), 2)/2)
	fmt.Printf("φ = (1/sqrt(2*π))*exp^(-(%.2f^2/2)) = %.2f\n", x, phi)
	fmt.Printf("Р = %.2f/sqrt(%d*%.2f*(%.2f)) = %.8f\n", phi, n, probabilities[0], 1-probabilities[0], moivreLaplaceResult)
}
