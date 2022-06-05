package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	sumPlayer := ""
	sumPlayerInt := 0
	sumDice := ""
	sumDiceInt := 0
	step := 0

	for time := 0; time < 100; time++ {
		if step == 0 {
			step = 1
			reader := bufio.NewScanner(os.Stdin)
			fmt.Print("Pemain	: ")
			reader.Scan()
			sumPlayer = reader.Text()
			sumPlayer = strings.TrimSuffix(sumPlayer, "\r\n")
			sumPlayerInt, _ = strconv.Atoi(sumPlayer)
		}

		if step == 1 {
			step = 2
			reader := bufio.NewScanner(os.Stdin)
			fmt.Print("Dadu	: ")
			reader.Scan()
			sumDice = reader.Text()
			sumDice = strings.TrimSuffix(sumDice, "\r\n")
			sumDiceInt, _ = strconv.Atoi(sumDice)
		}

		if step == 2 {
			step = 3
			permainan_dadu(sumDiceInt, sumPlayerInt)
		}
	}
}

func permainan_dadu(sumDice int, sumPlayer int) {
	var pointPlayer []int
	var dicePlayer []int
	var addDice []int

	fmt.Println("=================================================================")

	for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
		dicePlayer = append(dicePlayer, sumDice)
		pointPlayer = append(pointPlayer, 0)
		addDice = append(addDice, 0)
	}

	for time := 0; time < 100; time++ {

		playerWinner := 0
		velueWinner := 0
		endGame := 0
		idPlayer := 0
		a := ""
		var Temp []string
		fmt.Println("Giliran", time+1, "lempar dadu:")
		fmt.Println()
		for idPlayer = 1; idPlayer <= sumPlayer; idPlayer++ {
			if dicePlayer[idPlayer-1] > 0 {
				allVelueDice := ""
				pemain(idPlayer, pointPlayer)
				a = dadu(dicePlayer, idPlayer, allVelueDice, addDice, pointPlayer, sumPlayer)
				fmt.Println(a)
				Temp = append(Temp, a)
			}
			if idPlayer == sumPlayer {
				fmt.Println("=================================================================")
				fmt.Println("Setelah evaluasi:")
				fmt.Println()
				for idPlayer = 1; idPlayer <= sumPlayer; idPlayer++ {
					if dicePlayer[idPlayer-1] > 0 {
						pemain(idPlayer, pointPlayer)
						for k, i := range Temp {
							k = k + 1
							if len(Temp) == sumPlayer {
								if k == idPlayer {
									replace(i, k, idPlayer)
								}
							} else {
								if k+1 == idPlayer {
									replace(i, k, idPlayer)
								}
							}
						}
					}
				}
			}
		}

		endGame = 0
		for idPlayer := 1; idPlayer <= sumPlayer; idPlayer++ {
			dicePlayer[idPlayer-1] = dicePlayer[idPlayer-1] + addDice[idPlayer-1]
			addDice[idPlayer-1] = 0
			if velueWinner < pointPlayer[idPlayer-1] {
				playerWinner = idPlayer
				velueWinner = pointPlayer[idPlayer-1]
			}
			if dicePlayer[idPlayer-1] == 0 {
				endGame = endGame + 1
			}
		}

		if endGame == sumPlayer-1 {
			fmt.Println("=================================================================")
			fmt.Print("Game dimenangkan oleh pemain #")
			fmt.Print(playerWinner)
			fmt.Print(" karena memiliki poin lebih banyak dari pemain lainnya dan menyelesaikan game lebih awal.")
			fmt.Println()
			return
		}
		fmt.Println("=================================================================")

	}
}

func dadu(dicePlayer []int, idPlayer int, allVelueDice string, addDice []int, pointPlayer []int, sumPlayer int) string {
	for rollDice := 0; rollDice < dicePlayer[idPlayer-1]; rollDice++ {
		velueDice := (rand.Intn(6) + 1)
		if allVelueDice == "" {
			allVelueDice = strconv.Itoa(velueDice)
		} else {
			allVelueDice = allVelueDice + ", " + strconv.Itoa(velueDice)
		}
		if velueDice == 6 {
			addDice[idPlayer-1] = addDice[idPlayer-1] - 1
			pointPlayer[idPlayer-1] = pointPlayer[idPlayer-1] + 1
		}
		if velueDice == 1 {
			plusDice := 1
			addDice[idPlayer-1] = addDice[idPlayer-1] - 1
			for a := idPlayer; a < sumPlayer; a++ {
				if dicePlayer[a] > 0 {
					addDice[a] = addDice[a] + plusDice
					plusDice = 0
				}
			}
			for b := 0; b < idPlayer-1; b++ {
				if dicePlayer[b] > 0 {
					addDice[b] = addDice[b] + plusDice
					plusDice = 0
				}
			}
		}
	}
	return allVelueDice
}

func replace(i string, k int, idPlayer int) {
	if strings.Contains(i, "6") || strings.Contains(i, "6,") {
		i = strings.ReplaceAll(i, "6,", "")
		i = strings.ReplaceAll(i, "6", "")
	}
	if strings.Contains(i, "1") || strings.Contains(i, "1,") {
		for a := k; a-1 < idPlayer; a++ {
			i = strings.ReplaceAll(i, "1,", "")
			i = strings.ReplaceAll(i, "1", "")
		}
	}
	fmt.Println(i)
}

func pemain(idPlayer int, pointPlayer []int) {
	fmt.Print("Pemain #")
	fmt.Print(idPlayer, " ", "(", pointPlayer[idPlayer-1], ")", " ")
}
