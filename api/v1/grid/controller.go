package grid

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Row []string
type Grid []Row

type GridResponse struct {
	Empty  Grid
	Random Grid
	Words  Words
}

const ROWS = 8
const COLS = 13

type Word struct {
	Chars string
	Row   int
	Col   int
}
type List map[string]Word
type Words struct {
	Hours   List
	Minutes List
}

// private minutes: List = {
// 	5: {
// 		Chars: 'FIVE',
// 		Col: 0,
// 		Row: 2,
// 	},
// 	10: {
// 		Chars: 'TEN',
// 		Col: 6,
// 		Row: 0,
// 	},
// 	15: {
// 		Chars: 'QUARTER',
// 		Col: 0,
// 		Row: 1,
// 	},
// 	20: {
// 		Chars: 'TWENTY',
// 		Col: 7,
// 		Row: 1,
// 	},
// 	30: {
// 		Chars: 'HALF',
// 		Col: 9,
// 		Row: 0,
// 	},
// };

func grid(c *gin.Context) {
	empty := generate(ROWS, COLS)
	random := fillRandom(generate(ROWS, COLS))

	hours := make(List)

	hours["1"] = Word{
		Chars: "ONE",
		Row:   3,
		Col:   7,
	}
	hours["2"] = Word{
		Chars: "TWO",
		Row:   3,
		Col:   9,
	}
	hours["3"] = Word{
		Chars: "THREE",
		Row:   4,
		Col:   0,
	}
	hours["4"] = Word{
		Chars: "FOUR",
		Row:   4,
		Col:   5,
	}
	hours["5"] = Word{
		Chars: "FIVE",
		Row:   4,
		Col:   9,
	}
	hours["6"] = Word{
		Chars: "SIX",
		Row:   5,
		Col:   0,
	}
	hours["7"] = Word{
		Chars: "SEVEN",
		Row:   5,
		Col:   4,
	}
	hours["8"] = Word{
		Chars: "EIGHT",
		Row:   5,
		Col:   8,
	}
	hours["9"] = Word{
		Chars: "NINE",
		Row:   6,
		Col:   0,
	}
	hours["10"] = Word{
		Chars: "TEN",
		Row:   6,
		Col:   0,
	}
	hours["11"] = Word{
		Chars: "ELEVEN",
		Row:   6,
		Col:   0,
	}
	hours["12"] = Word{
		Chars: "TWELVE",
		Row:   7,
		Col:   0,
	}

	minutes := make(List)

	minutes["a"] = Word{
		Chars: "A",
		Row:   0,
		Col:   4,
	}
	minutes["its"] = Word{
		Chars: "ITS",
		Row:   0,
		Col:   0,
	}
	minutes["minutes"] = Word{
		Chars: "MINUTES",
		Row:   2,
		Col:   5,
	}
	minutes["oclock"] = Word{
		Chars: "OCLOCK",
		Row:   7,
		Col:   7,
	}
	minutes["past"] = Word{
		Chars: "PAST",
		Row:   3,
		Col:   0,
	}
	minutes["to"] = Word{
		Chars: "TO",
		Row:   3,
		Col:   4,
	}

	response := GridResponse{
		Empty:  empty,
		Random: random,
		Words: Words{
			Hours:   hours,
			Minutes: minutes,
		},
	}

	c.JSON(http.StatusOK, response)
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
