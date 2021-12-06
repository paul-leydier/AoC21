package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const dataPath = "day_2/data.txt"

type instruction struct {
	Direction string
	Units     int
}

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

type coordinates struct {
	x   int // horizontal position
	y   int // depth
	aim int
}

func readInstructions(path string) ([]instruction, error) {
	var instructions []instruction
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if len(words) != 2 {
			return nil, fmt.Errorf("invalid length instruction line: %s", line)
		}
		dist, err := strconv.ParseInt(words[1], 10, 64)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction{Direction: words[0], Units: int(dist)})
	}

	return instructions, nil
}

func move(pos *coordinates, instr instruction) error {
	switch instr.Direction {
	case forward:
		pos.x += instr.Units
		pos.y += pos.aim * instr.Units
	case up:
		pos.aim -= instr.Units
	case down:
		pos.aim += instr.Units
	default:
		return fmt.Errorf("invalid direction: %s", instr.Direction)
	}
	return nil
}

func followInstructions(pos *coordinates, instructions []instruction) error {
	for _, instr := range instructions {
		err := move(pos, instr)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	pos := coordinates{0, 0, 0}
	instructions, err := readInstructions(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	err = followInstructions(&pos, instructions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Position after instructions: %d, %d\n", pos.x, pos.y)
	fmt.Printf("Multiplication: %d\n", pos.x*pos.y)
}
