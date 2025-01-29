//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func Roll(faces int) int {
	return rand.Intn(faces) + 1
}

type Round struct {
	sides int
	rolls int
	dices int
}

type Result struct {
	count int
	sync.Mutex
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	round := Round{}

	fmt.Printf("sides: ")
	if scanner.Scan() {
		round.sides, _ = strconv.Atoi(scanner.Text())
	}

	fmt.Printf("rolls: ")
	if scanner.Scan() {
		round.rolls, _ = strconv.Atoi(scanner.Text())
	}

	fmt.Printf("dices: ")
	if scanner.Scan() {
		round.dices, _ = strconv.Atoi(scanner.Text())
	}

	var wg sync.WaitGroup

	result := Result{}
	for r := 0; r < round.rolls; r++ {
		fmt.Printf("Roll number: %v\n", r+1)

		for d := 0; d < round.dices; d++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				roll_result := Roll(round.sides)
				result.Lock()
				defer result.Unlock()
				result.count += roll_result
				fmt.Println(roll_result)
			}()
		}
		wg.Wait()
	}

	fmt.Println("result: ", result.count)
	if result.count == round.dices {
		fmt.Println("Snake eyes")
	}
	if result.count == 7 {
		fmt.Println("Lucky 7")
	}
	if result.count%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
