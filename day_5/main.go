package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const dataPath = "day_5/data.txt"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type mapDiagram [][]int

func (m *mapDiagram) Text() string {
	var s strings.Builder
	for i := 0; i < m.Width(); i++ {
		for j := 0; j < m.Height(); j++ {
			switch (*m)[i][j] {
			case 0:
				s.WriteString(".\t")
			default:
				s.WriteString(strconv.Itoa((*m)[i][j]) + "\t")
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

func (m *mapDiagram) AddLine(l line) {
	m.FitLine(l)      // Make sure the line fits within the map
	if l.x1 == l.x2 { // Vertical line
		for i := Min(l.y1, l.y2); i <= Max(l.y1, l.y2); i++ {
			(*m)[i][l.x1]++
		}
	}
	if l.y1 == l.y2 { // Horizontal line
		for i := Min(l.x1, l.x2); i <= Max(l.x1, l.x2); i++ {
			(*m)[l.y1][i]++
		}
	}
	// TODO: diagonals?
}

func (m *mapDiagram) Width() int {
	if len(*m) == 0 {
		return 0
	}
	return len((*m)[0])
}

func (m *mapDiagram) Height() int {
	return len(*m)
}

func (m *mapDiagram) IncreaseWidth(size int) {
	padding := make([]int, size-m.Width())
	for i := 0; i < len(*m); i++ {
		(*m)[i] = append((*m)[i], padding...)
	}
}

func (m *mapDiagram) IncreaseHeight(size int) {
	paddingRow := make([]int, m.Width())
	padding := make([][]int, size-len(*m))
	for i := 0; i < len(padding); i++ {
		padding[i] = paddingRow
	}
	*m = append(*m, padding...)
}

func (m *mapDiagram) FitLine(l line) {
	height := m.Height()
	if l.y1 >= height || l.y2 >= height {
		m.IncreaseHeight(Max(l.y1, l.y2) + 1)
	}
	width := m.Width()
	if l.x1 >= width || l.x2 >= width {
		m.IncreaseWidth(Max(l.x1, l.x2) + 1)
	}
}

func (m *mapDiagram) CountOverlaps(n int) int {
	var s int
	for i := 0; i < m.Width(); i++ {
		for j := 0; j < m.Height(); j++ {
			if (*m)[i][j] >= n {
				s++
			}
		}
	}
	return s
}

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func parseLine(input string) (line, error) {
	points := strings.Split(input, " -> ")
	if len(points) != 2 {
		return line{}, fmt.Errorf("incorrect line: %s", input)
	}
	coordinatesString := make([][]string, 2)
	for i, point := range points {
		coordinatesString[i] = strings.Split(point, ",")
		if len(coordinatesString[i]) != 2 {
			return line{}, fmt.Errorf("incorrect point %d: %s", i, input)
		}
	}
	coordinates := make([][]int, 2)
	for i, point := range coordinatesString {
		for _, coor := range point {
			n, err := strconv.ParseInt(coor, 10, 64)
			if err != nil {
				return line{}, err
			}
			coordinates[i] = append(coordinates[i], int(n))
		}
	}
	return line{coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1]}, nil
}

func readData(path string) ([]line, error) {
	var result []line
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		l, err := parseLine(row)
		if err != nil {
			return nil, err
		}
		result = append(result, l)
	}
	return result, nil
}

func main() {
	lines, err := readData(dataPath)
	if err != nil {
		log.Fatalln(err)
	}
	m := mapDiagram{}
	for _, l := range lines {
		m.AddLine(l)
	}
	fmt.Print(m.Text())
	fmt.Printf("%d overlaps!", m.CountOverlaps(2))
}
