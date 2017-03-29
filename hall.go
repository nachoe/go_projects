package main

import (
		"fmt"
		"math/rand"
		"strconv"
		"time"
		)
 
func makeSelections(reselect bool, totalAttempts int) (int) {
	totalCorrect := 0
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < totalAttempts; i++ {
		initialArray := []string{"car", "goat", "goat"}
		restArray, contestantSelection := random(initialArray);
		for {
			j := rand.Intn(len(restArray))
			//remove hall's selection
			if (restArray[j] == "goat") {
				//remove slice (magic?)
				restArray = append(restArray[:j], restArray[j+1:]...)
				break;
			}
		}
		if (reselect == true) {
			for k := 0; k < len(restArray); k++ {
				//switch selections
				if (restArray[k] != contestantSelection) {
					contestantSelection = restArray[k]
					break
				}
			}
		}
		if (contestantSelection == "car") {
			totalCorrect++
		}
	}

	return totalCorrect
}

func random(array [] string) ([] string, string) {
	j := rand.Intn(len(array))
	return array, array[j]
}

func main() {
	totalAttempts := 5000
	totalCorrect := makeSelections(false, totalAttempts)
	fmt.Println("finished with no reselect. totalCorrect: " + strconv.Itoa(totalCorrect) + ". totalAttempts: " + strconv.Itoa(totalAttempts) )
	totalCorrect = makeSelections(true, totalAttempts)
	fmt.Println("finished with reselect. totalCorrect: " + strconv.Itoa(totalCorrect) + ". totalAttempts: " + strconv.Itoa(totalAttempts) )
}