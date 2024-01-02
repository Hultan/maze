package maze_gen

import (
	"math/rand"
	"time"
)

const size = 25

const (
	North    = 1
	East     = 2
	South    = 4
	West     = 8
	notNorth = 14
	notEast  = 13
	notSouth = 11
	notWest  = 7
	all      = 15
)

type Maze [size][size]Cell

type Cell struct {
	visited bool
	Walls   uint8
}

type Location struct {
	x, y int
}

var rnd *rand.Rand

func NewMaze() *Maze {
	m := &Maze{}

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			m[y][x] = Cell{visited: false, Walls: all}
		}
	}

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	s := NewStack()
	s.Push(Location{0, 0})
	m[0][0].visited = true
	generate(m, s)

	return m
}

// https://en.wikipedia.org/wiki/Maze_generation_algorithm
func generate(m *Maze, s *Stack) {
	for s.Length() > 0 {
		loc := s.Pop()

		n := getNeighbors(m, loc)
		if len(n) == 0 {
			continue
		}

		s.Push(loc)
		nLoc := n[rnd.Intn(len(n))]

		// Remove walls
		if nLoc.x == loc.x+1 && nLoc.y == loc.y {
			m[nLoc.y][nLoc.x].Walls &= notWest
			m[loc.y][loc.x].Walls &= notEast
		}
		if nLoc.x == loc.x && nLoc.y == loc.y+1 {
			m[nLoc.y][nLoc.x].Walls &= notNorth
			m[loc.y][loc.x].Walls &= notSouth
		}
		if nLoc.x == loc.x-1 && nLoc.y == loc.y {
			m[nLoc.y][nLoc.x].Walls &= notEast
			m[loc.y][loc.x].Walls &= notWest
		}
		if nLoc.x == loc.x && nLoc.y == loc.y-1 {
			m[nLoc.y][nLoc.x].Walls &= notSouth
			m[loc.y][loc.x].Walls &= notNorth
		}

		m[nLoc.y][nLoc.x].visited = true
		s.Push(nLoc)
	}
}

func getNeighbors(m *Maze, loc Location) []Location {
	var n []Location

	if loc.x > 0 && !m[loc.y][loc.x-1].visited {
		n = append(n, Location{x: loc.x - 1, y: loc.y})
	}
	if loc.y > 0 && !m[loc.y-1][loc.x].visited {
		n = append(n, Location{x: loc.x, y: loc.y - 1})
	}
	if loc.x < size-1 && !m[loc.y][loc.x+1].visited {
		n = append(n, Location{x: loc.x + 1, y: loc.y})
	}
	if loc.y < size-1 && !m[loc.y+1][loc.x].visited {
		n = append(n, Location{x: loc.x, y: loc.y + 1})
	}

	return n
}
