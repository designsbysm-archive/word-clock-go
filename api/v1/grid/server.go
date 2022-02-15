package grid

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Row []string
type Grid []Row

func random(c *gin.Context) {
	empty := generate(8, 13)
	grid := fillRandom(empty)

	c.JSON(http.StatusOK, grid)
}

func time(c *gin.Context) {
	empty := generate(8, 13)
	grid := fillTime(empty)

	c.JSON(http.StatusOK, grid)
}

func fillRandom(grid Grid) Grid {
	var last string

	for r := range grid {
		for c := range grid[r] {
			pick := randomLetter()

			for last == pick {
				pick = randomLetter()
			}

			grid[r][c] = pick
			last = pick
		}
	}

	return grid
}

func fillTime(grid Grid) Grid {
	return grid
}

func generate(rows int, cols int) Grid {
	grid := make(Grid, rows)

	for r := 0; r < rows; r++ {
		grid[r] = make(Row, cols)
	}

	return grid
}

func randomLetter() string {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	index := rand.Intn(len(letters))

	return letters[index]
}
