package mazeGen

import (
	"math/rand"
	"time"
)

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

type Maze struct {
	Cells         [][]Cell
	Unvisited     *Stack
	width, height int32

	// For solving
	SolutionStack   []Step
	SolutionVisited map[Location]bool
	SolutionParent  map[Location]*Location
	SolutionStarted bool
	SolutionDone    bool
	SolutionPath    []Location
}

type Cell struct {
	visited bool
	Walls   uint8
}

type Location struct {
	X, Y int32
}

type Step struct {
	Loc  Location
	From *Location
}

var rnd *rand.Rand

func NewMaze(width, height int32) *Maze {
	m := &Maze{}
	m.width = width
	m.height = height
	m.Cells = make([][]Cell, height)
	for i := range m.Cells {
		m.Cells[i] = make([]Cell, width)
	}
	m.clear()

	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

	m.Unvisited = NewStack()
	m.Unvisited.Push(Location{0, 0})
	m.Cells[0][0].visited = true
	return m
}

func (m *Maze) Width() int32 {
	return int32(m.width)
}

func (m *Maze) Height() int32 {
	return int32(m.height)
}

func (m *Maze) IsDone() bool {
	return m.Unvisited.Length() == 0
}

func (m *Maze) clear() {
	for x := int32(0); x < m.width; x++ {
		for y := int32(0); y < m.height; y++ {
			m.Cells[y][x] = Cell{visited: false, Walls: all}
		}
	}
}

// https://en.wikipedia.org/wiki/Maze_generation_algorithm
func (m *Maze) Generate() {
	if m.IsDone() {
		return
	}

	loc := m.Unvisited.Pop()

	n := m.getNeighbors(loc)
	if len(n) == 0 {
		return
	}

	m.Unvisited.Push(loc)
	nLoc := n[rnd.Intn(len(n))]

	// Remove walls
	if nLoc.X == loc.X+1 && nLoc.Y == loc.Y {
		m.Cells[nLoc.Y][nLoc.X].Walls &= notWest
		m.Cells[loc.Y][loc.X].Walls &= notEast
	}
	if nLoc.X == loc.X && nLoc.Y == loc.Y+1 {
		m.Cells[nLoc.Y][nLoc.X].Walls &= notNorth
		m.Cells[loc.Y][loc.X].Walls &= notSouth
	}
	if nLoc.X == loc.X-1 && nLoc.Y == loc.Y {
		m.Cells[nLoc.Y][nLoc.X].Walls &= notEast
		m.Cells[loc.Y][loc.X].Walls &= notWest
	}
	if nLoc.X == loc.X && nLoc.Y == loc.Y-1 {
		m.Cells[nLoc.Y][nLoc.X].Walls &= notSouth
		m.Cells[loc.Y][loc.X].Walls &= notNorth
	}

	m.Cells[nLoc.Y][nLoc.X].visited = true
	m.Unvisited.Push(nLoc)
}

func (m *Maze) getNeighbors(loc Location) []Location {
	var n []Location

	if loc.X > 0 && !m.Cells[loc.Y][loc.X-1].visited {
		n = append(n, Location{X: loc.X - 1, Y: loc.Y})
	}
	if loc.Y > 0 && !m.Cells[loc.Y-1][loc.X].visited {
		n = append(n, Location{X: loc.X, Y: loc.Y - 1})
	}
	if loc.X < m.width-1 && !m.Cells[loc.Y][loc.X+1].visited {
		n = append(n, Location{X: loc.X + 1, Y: loc.Y})
	}
	if loc.Y < m.height-1 && !m.Cells[loc.Y+1][loc.X].visited {
		n = append(n, Location{X: loc.X, Y: loc.Y + 1})
	}

	return n
}

func (m *Maze) StartSolving() {
	m.SolutionStack = []Step{{Loc: Location{0, 0}, From: nil}}
	m.SolutionVisited = make(map[Location]bool)
	m.SolutionParent = make(map[Location]*Location)
	m.SolutionDone = false
	m.SolutionPath = nil
}

func (m *Maze) SolveStep() {
	if m.SolutionDone || len(m.SolutionStack) == 0 {
		return
	}

	current := m.SolutionStack[len(m.SolutionStack)-1]
	m.SolutionStack = m.SolutionStack[:len(m.SolutionStack)-1]

	if m.SolutionVisited[current.Loc] {
		return
	}
	m.SolutionVisited[current.Loc] = true
	m.SolutionParent[current.Loc] = current.From

	end := Location{m.width - 1, m.height - 1}
	if current.Loc == end {
		// Reconstruct the path
		path := []Location{}
		for at := &end; at != nil; at = m.SolutionParent[*at] {
			path = append([]Location{*at}, path...)
		}
		m.SolutionPath = path
		m.SolutionDone = true
		return
	}

	directions := []struct {
		dx, dy int32
		wall   uint8
	}{
		{0, -1, North},
		{1, 0, East},
		{0, 1, South},
		{-1, 0, West},
	}

	for _, dir := range directions {
		nx, ny := current.Loc.X+dir.dx, current.Loc.Y+dir.dy
		if nx < 0 || ny < 0 || nx >= m.width || ny >= m.height {
			continue
		}
		currCell := m.Cells[current.Loc.Y][current.Loc.X]
		if currCell.Walls&dir.wall != 0 {
			continue // Wall blocks path
		}
		neighbor := Location{nx, ny}
		if !m.SolutionVisited[neighbor] {
			m.SolutionStack = append(m.SolutionStack, Step{Loc: neighbor, From: &current.Loc})
		}
	}
}
