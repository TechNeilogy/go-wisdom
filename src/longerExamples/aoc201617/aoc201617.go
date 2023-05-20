package aoc201617

// Coding Challenge Solution March 2023
// Based on Advent of Code 2016, problem 17 parts A and B.
// Note: Almost everything is exported,
// to allow it to be documented.

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// Maze defines the structure of the maze.
// Since all the internal doors are computed
// based on the path, the internals of the
// maze need not be stored.
type Maze struct {
	XSize, YSize   int
	XStart, YStart int
	XGoal, YGoal   int
}

// Dir defines the direction updates
// for a cardinal point;
type Dir struct {
	Path   string
	DX, DY int
}

// Dirs defines the direction updates
// for each cardinal point;
var Dirs = []Dir{
	{"U", 0, -1},
	{"D", 0, 1},
	{"L", -1, 0},
	{"R", 1, 0},
}

// GetOpenDirs gets the open directions for path,
// based on the algorithm in the problem statement.
func GetOpenDirs(path string) (rtn []*Dir) {
	hash := md5.Sum([]byte(path))
	four := hex.EncodeToString(hash[:])[0:4]
	for i, c := range four {
		if c >= 'b' {
			rtn = append(rtn, &Dirs[i])
		}
	}
	return rtn
}

// Path is used mainly to cache X, Y values for
// breadth first search, and connect them to the Path.Path.
// This is cheaper than re-computing them from Path.Path.
type Path struct {
	Path string
	X, Y int
}

// Move adds xd and yd to x and y, and if valid,
// returns the updated (x, y) values.
// If invalid, it returns (-1, -1)
func (m *Maze) Move(x, y, xd, yd int) (int, int) {
	x += xd
	if x < 0 || x >= m.XSize {
		return -1, -1
	}
	y += yd
	if y < 0 || y >= m.YSize {
		return -1, -1
	}
	return x, y
}

// BreadthFirstSearch runs a breadth first search.
//
// paths is the list of paths (plus key) at the search front.
// Initially, this contains only the key specified in the problem.
func (m *Maze) BreadthFirstSearch(
	paths []Path,
) *Path {
	// While there is a path in the search front.
	for len(paths) > 0 {

		// The algorithm makes sure the first key
		// is the current shortest path.
		// Pop it from the list and use it
		// as the new starting point.

		h := paths[0]

		if h.X == m.XGoal && h.Y == m.YGoal {
			return &h
		}

		paths = paths[1:]

		// Append each valid new path to the current path,
		// and re-run the algorithm.

		for _, d := range GetOpenDirs(h.Path) {
			xn, yn := m.Move(h.X, h.Y, d.DX, d.DY)
			if xn >= 0 {
				paths = append(
					paths,
					Path{
						h.Path + d.Path,
						xn,
						yn,
					},
				)
			}
		}
	}

	return nil
}

// DepthFirstSearchExt runs a depth first search.
//
// pathCurrent is the path to this point,
// including the key specified in the problem.
//
// x and y are the current position.
// (Cheaper to pass than to re-compute from pathCurrent.)
//
// If findAll is true, all valid paths are returned.
// If findAll is false, only the first valid path is returned.
//
// pathAcc accumulates a list of valid paths.
func (m *Maze) DepthFirstSearchExt(
	pathCurrent string,
	x int,
	y int,
	findAll bool,
	pathAcc *[]string,
) bool {
	if x == m.XGoal && y == m.YGoal {
		*pathAcc = append(*pathAcc, pathCurrent)
		return true
	}

	for _, d := range GetOpenDirs(pathCurrent) {
		xn, yn := m.Move(x, y, d.DX, d.DY)
		if xn >= 0 {
			if m.DepthFirstSearchExt(
				pathCurrent+d.Path,
				xn,
				yn,
				findAll,
				pathAcc,
			) && !findAll {
				return true
			}
		}
	}

	return false
}

// DepthFirstSearch runs a depth first search.
//
// key is the maze key specified in the problem.
//
// If findAll is true, all valid paths are returned.
// If findAll is false, only the first valid path is returned.
func (m *Maze) DepthFirstSearch(
	key string,
	findAll bool,
) []string {
	var pathAcc []string
	m.DepthFirstSearchExt(key, m.XStart, m.YStart, findAll, &pathAcc)
	return pathAcc
}

// RunBreadthFirst runs a breadth first search
// and reports the results.
//
// key is the maze key specified in the problem.
func RunBreadthFirst(
	key string,
) {
	maze := &Maze{
		4, 4,
		0, 0,
		3, 3,
	}

	path := maze.BreadthFirstSearch(
		[]Path{
			{
				key,
				maze.XStart,
				maze.YStart,
			},
		},
	)

	fmt.Printf("Breadth First\n")
	if path != nil {
		shortest := path.Path[len(key):]
		fmt.Printf("Shortest: (%v) %s\n", len(shortest), shortest)
	} else {
		fmt.Println("No valid path.")
	}
}

// RunDepthFirst runs a depth first search
// and reports the results.//
//
// key is the maze key specified in the problem.
func RunDepthFirst(key string) {
	maze := &Maze{
		4, 4,
		0, 0,
		3, 3,
	}

	paths := maze.DepthFirstSearch(
		key,
		true,
	)

	fmt.Printf("Depth First\n")
	if len(paths) > 0 {
		shortest := paths[0]
		longest := paths[0]
		for _, p := range paths[1:] {
			if len(p) < len(shortest) {
				shortest = p
			}
			if len(p) > len(longest) {
				longest = p
			}
		}
		shortest = shortest[len(key):]
		longest = longest[len(key):]
		fmt.Printf("Shortest: (%v) %s\n", len(shortest), shortest)
		fmt.Printf("Longest: (%v) %s\n", len(longest), longest)
	} else {
		fmt.Println("No valid path.")
	}
}

func Run(key string) {
	RunBreadthFirst(key)
	RunDepthFirst(key)
}
