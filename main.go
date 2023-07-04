package main

import (
	"log"
	"math"
	"strconv"
)

func coinChange(coins []string, amount float64) (int, []string) {
	// Convert coin strings to float64 values
	coinValues := make(map[string]float64)
	for _, coin := range coins {
		coinValues[coin] = stringToFloat(coin)
	}

	// Initialize dp array with maximum possible number of coins needed
	dp := make([][]int, len(coinValues))
	for i := range dp {
		dp[i] = make([]int, int(amount*100+1))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// Initialize first column of dp array to 0
	for i := range dp {
		dp[i][0] = 0
	}

	// Fill dp array
	for i := range dp {
		for j := 1; j < len(dp[i]); j++ {
			if i > 0 {
				dp[i][j] = dp[i-1][j]
			}
			if j >= int(coinValues[coins[i]]*100) {
				dp[i][j] = min(dp[i][j], dp[i][j-int(coinValues[coins[i]]*100)]+1)
			}
		}
	}

	// Get coins used to make change
	coinCount := make(map[string]int)
	i, j := len(dp)-1, len(dp[0])-1
	for i >= 0 && j > 0 {
		if i > 0 && dp[i][j] == dp[i-1][j] {
			i--
		} else {
			coinCount[coins[i]]++
			j -= int(coinValues[coins[i]] * 100)
		}
	}

	// Return -1 if change cannot be made
	if dp[len(dp)-1][len(dp[0])-1] == math.MaxInt32 {
		return -1, nil
	}

	// Convert coinCount map to slice of coins used
	usedCoins := make([]string, 0, len(coinCount))
	for coin, count := range coinCount {
		for i := 0; i < count; i++ {
			usedCoins = append(usedCoins, coin)
		}
	}

	return dp[len(dp)-1][len(dp[0])-1], usedCoins
}

// Helper function to convert string to float64
func stringToFloat(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

// Helper function to find minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Started Coin Change program")

	coins := []string{"100.00", "200.00", "50.00", "500.00", "2.00"}
	var amount float64 = 10001.00
	minCoins, coins2 := coinChange(coins, amount)
	var coinAmount float64
	for _, c := range coins2 {
		c1, _ := strconv.ParseFloat(c, 64)
		coinAmount = coinAmount + c1
	}
	log.Println("total amount of the coins received in change management:", coinAmount)
	log.Printf("Minimum number of coins required to make change for %f rupees: %d\n", amount, minCoins)
	log.Println("Coins used:", coins2)

	if amount == coinAmount {
		log.Println("Change management done successfully")
	} else {
		log.Println("Change management failed")
	}

}
