package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// todo add functional tests
// todo load configuration from a config file
// todo add input XY validation
// todo add command parameters options verbosity
var l = 3
var h = 2
var RedTreesXY = []string{
	"3-1",
	"2-2",
}

var GreenTrees []Tree
var RedTrees []Tree

func main() {
	buildRedTrees()
	buildGreenTrees()

	fmt.Printf("GreenTrees : %v", GreenTrees)
	fmt.Println("")
	fmt.Printf("RedTrees : %v", RedTrees)
	fmt.Println("")

	// get a random burning tree
	oneRedRandom := pickOneTreeRandomRed()

	fmt.Println("one red random : " + oneRedRandom.XY)

	process(oneRedRandom, GreenTrees)
}

func buildGreenTrees() {
	for i := 1; i <= l; i++ {
		for j := 1; j <= h; j++ {
			concat := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			if !isRedTree(concat) {
				GreenTrees = append(GreenTrees, Tree{
					XY:      concat,
					Burning: false,
				})
			}
		}
	}
}

func buildRedTrees() {
	for _, xy := range RedTreesXY {
		RedTrees = append(RedTrees, Tree{
			XY:      xy,
			Burning: true,
		})
	}
}

func process(tree Tree, trees []Tree) {
	// get all neighbors
	greenNeighbours := neighbours(tree, trees)

	// if neighbors not burned exist -> process
	if len(greenNeighbours) > 0 {
		// get a random neighbor
		oneGreenRandom := pickOneTreeRandomGreenFromNeighbors(greenNeighbours)

		// burn the tree
		GreenTrees = burn(oneGreenRandom, GreenTrees)

		fmt.Println("-- burned : " + oneGreenRandom.XY)

		// process to the next
		process(oneGreenRandom, GreenTrees)
	}
}

type Tree struct {
	XY      string
	Burning bool
}

func isRedTree(str string) bool {
	for _, v := range RedTrees {
		if v.XY == str {
			return true
		}
	}

	return false
}

func isGreenTree(s []Tree, str string) bool {
	for _, v := range s {
		if v.XY == str && v.Burning == false {
			return true
		}
	}

	return false
}

func pickOneTreeRandomRed() Tree {
	rand.Seed(time.Now().UnixNano())

	// todo remove picked random red tree
	random := RedTrees[rand.Intn(len(RedTrees))]

	return random
}

func pickOneTreeRandomGreenFromNeighbors(neighBoors []Tree) Tree {
	rand.Seed(time.Now().UnixNano())

	random := neighBoors[rand.Intn(len(neighBoors))]

	return random
}

func burn(tree Tree, gt []Tree) []Tree {
	var s2 []Tree

	// todo do not rebuild all slice : just update the element to be burned into the slice
	for _, i := range gt {
		if i.XY == tree.XY {
			i.Burning = true
		}
		s2 = append(s2, i)
	}

	return s2
}

func neighbours(tree Tree, trees []Tree) []Tree {
	var neighbours []Tree

	s := strings.Split(tree.XY, "-")

	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])

	xm := x - 1
	xp := x + 1
	ym := y - 1
	yp := y + 1

	if xm > 0 {
		xy := strconv.Itoa(xm) + "-" + s[1]
		// todo encapsulate
		if isGreenTree(trees, xy) {
			neighbours = append(neighbours, Tree{
				XY:      xy,
				Burning: false,
			})
		}
	}

	if xp <= l {
		xy := strconv.Itoa(xp) + "-" + s[1]
		if isGreenTree(trees, xy) {
			neighbours = append(neighbours, Tree{
				XY:      xy,
				Burning: false,
			})
		}
	}

	if ym > 0 {
		xy := s[0] + "-" + strconv.Itoa(ym)
		if isGreenTree(trees, xy) {
			neighbours = append(neighbours, Tree{
				XY:      xy,
				Burning: false,
			})
		}
	}

	if yp <= h {
		xy := s[0] + "-" + strconv.Itoa(yp)
		if isGreenTree(trees, xy) {
			neighbours = append(neighbours, Tree{
				XY:      xy,
				Burning: false,
			})
		}
	}

	return neighbours
}
