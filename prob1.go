package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	move := "R4,R4,L1,R3,L5,R2,R5,R1,L4,R3,L5,R2,L3,L4,L3,R1,R5,R1,L3,L1,R3,L1,R2,R2,L2,R5,L3,L4,R4,R4,R2,L4,L1,R5,L1,L4,R4,L1,R1,L2,R5,L2,L3,R2,R1,L194,R2,L4,R49,R1,R3,L5,L4,L1,R4,R2,R1,L5,R3,L5,L4,R4,R4,L2,L3,R78,L5,R4,R191,R4,R3,R1,L2,R1,R3,L1,R3,R4,R2,L2,R1,R4,L5,R2,L2,L4,L2,R1,R2,L3,R5,R2,L3,L3,R3,L1,L1,R5,L4,L4,L2,R5,R1,R4,L3,L5,L4,R5,L4,R5,R4,L3,L2,L5,R4,R3,L3,R1,L5,R5,R1,L3,R2,L5,R5,L3,R1,R4,L5,R4,R2,R3,L4,L5,R3,R4,L5,L5,R4,L4,L4,R1,R5,R3,L1,L4,L3,L4,R1,L5,L1,R2,R2,R4,R4,L5,R4,R1,L1,L1,L3,L5,L2,R4,L3,L5,L4,L1,R3"
	moves := strings.Split(move, ",")
	fmt.Println("Len: ", len(moves))
	curr_x := 0
	curr_y := 0

	prev := "N"
	for _, curr := range moves {
		//fmt.Printf("%d", i)
		runes := []rune(curr)
		dir := string(runes[0:1])
		fmt.Println("dir: ", dir)
		s1 := string(runes[1:len(runes)])
		fmt.Println("num: ", s1)
		num, err := strconv.Atoi(s1)
		if err != nil {
			log.Fatal("Input is wrong!")
		}
		switch prev {
		case "N":
			switch dir {
			case "R":
				curr_x = curr_x + num
				prev = "E"
			case "L":
				curr_x = curr_x - num
				prev = "W"
			}
		case "E":
			switch dir {
			case "R":
				curr_y = curr_y - num
				prev = "S"
			case "L":
				curr_y = curr_y + num
				prev = "N"
			}
		case "S":
			switch dir {
			case "R":
				curr_x = curr_x - num
				prev = "W"
			case "L":
				curr_x = curr_x + num
				prev = "E"
			}
		case "W":
			switch dir {
			case "R":
				curr_y = curr_y + num
				prev = "N"
			case "L":
				curr_y = curr_y - num
				prev = "S"
			}
		}
	}

	fmt.Printf("%d , %d", curr_x, curr_y)
}
