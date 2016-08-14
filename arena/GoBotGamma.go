package main
import( "fmt"
		"time"
		"math")
func returnMove1(board [][]int) []int {
	if testPassGamma(board){
		fmt.Println("BOT passed")
		return []int{-1, -1}
	}
	chains := botFindChainsGamma(board)
	libertySlice := make([][][]int, 0)
	curLib := make([][]int, 0)
	var tempSlice [][][]int
	for n := range chains[whoAmI %2] {
		curLib = findChainLibertiesGamma(chains[whoAmI % 2][n], board)
		for m:= 0; m <= len(libertySlice); m++ {
			if m == len(libertySlice) {
				libertySlice = append(libertySlice, curLib)
				break
			} else {
				if len(libertySlice[m]) >= len(curLib) {
					tempSlice = make([][][]int, len(libertySlice[m:]))
					copy(tempSlice, libertySlice[m:])
					libertySlice = append(libertySlice[:m], curLib)
					for k := range tempSlice {
						libertySlice = append(libertySlice, tempSlice[k])
					}
					break
				}
			}
		}
	}
	var tempTesting []int
	coords := make([]int, 2)
	testing := make([]int, 0)
	for n := range libertySlice {
		testing = make([]int, 0)
		for m:= range libertySlice[n] {
			testing = append(testing, m)
		}
		for len(testing) > 0 {
			index := testing[randomGamma(len(testing))]
			tempTesting = make([]int, len(testing))
			copy(tempTesting, testing)
			testing = make([]int, 0)
			for m := range tempTesting {
				if tempTesting[m] != index {
					testing = append(testing, tempTesting[m])
				}
			}
			copy(coords, libertySlice[n][index])
			if !suicideGamma(chains, board, coords) {
				fmt.Println("Bot Moved:", coords[1], coords[0])
				return coords
			}
		}
	}
	y, x := 9, 9
	var( yState int
	xState int)
	if board[9][9] != 0 {
		y, x, yState, xState = 8, 9, 0, -1
		for board[x][y] != 0 {
			x += xState
			y += yState
			if y == x || (x + y == 18 && y > x) || (x + y == 17 && yState == -1) {
				if xState == 0 {
					xState = int(yState)
					yState = 0
				} else {
					if yState == 0 {
						yState = int(-xState)
						xState = 0
					}
				}
			}
		}
	}
	fmt.Println("Bot Moved:", x, y)
	return []int{x, y}
/*
		y := 0
		x := 0
		for x < 19{
			y = 0
			for y < 19{
				if board[x][y] == 0{
					fmt.Println("BOT moved:", y, x)
					return []int{x, y}
				}
				y++
			}
			x++
		}
*/
	fmt.Println("BOT passed")
	return []int{-1, -1}
}

func findChainAttachedGamma(chains [][][]int, coords []int) int {
	for n := range chains {
		for m := range chains[n] {
			if chains[n][m][0] == coords[0] && chains[n][m][1] == coords[1] {
				return n
			}
		}
	}
	return -1
}

func botFindChainsGamma(board [][]int) [][][][]int {
	var( chainLayers [][][]int
	currentChain [][]int
	x int
	y int
	m int)
	chains := make([][][][]int, 2)
	n := 0
	for n < 2{
		chainLayers = findChains(board, n + 1)
		m = 0
		for m < len(chainLayers){
			currentChain = make([][]int, 0)
			x = 0
			for x < 19 {
				y = 0
				for y < 19 {
					if chainLayers[m][x][y] != 0 {
						currentChain = append(currentChain, []int{x, y})
					}
					y++
				}
				x++
			}
			chains[n] = append(chains[n], currentChain)
			m++
		}	
		n++
	}
	return chains
}

func findChainLibertiesGamma(chain [][]int, board [][]int) [][]int {
	tested := make([][]bool, 19)
	for x := range tested {
		tested[x] = make([]bool, 19)
		for y := range tested[x]{
			tested[x][y] = false
		}
	}
	liberties := make([][]int, 0)
	for n := range chain {
		if chain[n][0] != 0 {
			if board[chain[n][0] - 1][chain[n][1]] == 0 && !tested[chain[n][0] - 1][chain[n][1]] {
				liberties = append(liberties, []int{chain[n][0] - 1, chain[n][1]})
			}
			tested[chain[n][0] - 1][chain[n][1]] = true
		}
		if chain[n][1] != 0 {
			if board[chain[n][0]][chain[n][1] - 1] == 0 && !tested[chain[n][0]][chain[n][1] - 1]  {
				liberties = append(liberties, []int{chain[n][0], chain[n][1] - 1})
			}
			tested[chain[n][0]][chain[n][1] - 1] = true
		}
		if chain[n][1] != 18 {
			if board[chain[n][0]][chain[n][1] + 1] == 0 && !tested[chain[n][0]][chain[n][1] + 1] {
				liberties = append(liberties, []int{chain[n][0], chain[n][1] + 1})
			}
			tested[chain[n][0]][chain[n][1] + 1] = true
		}
		if chain[n][0] != 18 {
			if board[chain[n][0] + 1][chain[n][1]] == 0 && !tested[chain[n][0] + 1][chain[n][1]] {
				liberties = append(liberties, []int{chain[n][0] + 1, chain[n][1]})
			}
			tested[chain[n][0] + 1][chain[n][1]] = true
		}
	}
	return liberties
}

func findEmptyBordersGamma(board [][]int, emptyChain [][]int) [][][]int {
	borders := make([][][]int, 2)
	var( x, y int
	tested bool)
	for n := range emptyChain {
		x, y = emptyChain[n][0], emptyChain[n][1]
		if x != 0 {
			if board[x - 1][y] != 0 {
				tested = true
				for m:= range borders[board[x - 1][y] - 1] {
					if borders[board[x - 1][y]][m][0] == x - 1 && borders[board[x - 1][y]][m][1] == y {
						tested = false
					}
				}
				if tested {
					borders[board[x-1][y]] = append(borders[board[x-1][y]], []int{x - 1, y})
				}
			}
		}
	}
	return borders
}

func findEmptyBordersTestGamma(board [][]int, emptyChain [][]int, x int, y int, borders [][][]int) [][][]int {
	if board[x][y] != 0 {
	tested := true
		for m:= range borders[board[x][y] - 1] {
			if borders[board[x][y]][m][0] == x && borders[board[x][y]][m][1] == y {
				tested = false
			}
		}
		if tested {
			borders[board[x][y]] = append(borders[board[x][y]], []int{x, y})
		}
	}
	return borders
}

func findEmptyGamma(board [][]int) [][][]int {
	tested := make([][]bool, 19)
	emptyChains := make([][][]int, 0)
	for n := range tested {
		tested[n] = make([]bool, 19)
		for m := range tested[n] {
			tested[n][m] = false
		}
	}
	var chain [][]int
	for x := range board {
		for y := range board[x] {
			if board[x][y] == 0 && !tested[x][y] {
				chain = [][]int{{x, y}}
				tested[x][y] = true
				for n := 0; n < len(chain); n++ {
					if chain[n][0] != 0 {
						if board[chain[n][0] - 1][chain[n][1]] == 0 && !tested[chain[n][0] - 1][chain[n][1]] {
							chain = append(chain, []int{chain[n][0] - 1, chain[n][1]})
						}
						tested[chain[n][0] - 1][chain[n][1]] = true
					}
					if chain[n][0] != 18 {
						if board[chain[n][0] + 1][chain[n][1]] == 0 && !tested[chain[n][0] + 1][chain[n][1]] {
							chain = append(chain, []int{chain[n][0] + 1, chain[n][1]})
						}
						tested[chain[n][0] + 1][chain[n][1]] = true
					}
					if chain[n][1] != 0 {
						if board[chain[n][0]][chain[n][1] - 1] == 0 && !tested[chain[n][0]][chain[n][1] - 1] {
							chain = append(chain, []int{chain[n][0], chain[n][1] - 1})
						}
						tested[chain[n][0]][chain[n][1] - 1] = true
					}
					if chain[n][1] != 18 {
						if board[chain[n][0]][chain[n][1] + 1] == 0 && !tested[chain[n][0]][chain[n][1] + 1] {
							chain = append(chain, []int{chain[n][0], chain[n][1] + 1})
						}
						tested[chain[n][0]][chain[n][1] + 1] = true
					}
				}
				emptyChains = append(emptyChains, chain)
			}
		}
	}
	return emptyChains
}

func suicideGamma(chains [][][][]int, board [][]int, coords []int) bool {
	var x, y int
	if coords[0] != 0{
		x, y = coords[0] - 1, coords[1]
		if !suicideTestGamma(chains, board, x, y) {
			return false
		}
	}
	if coords[0] != 18{
		x, y = coords[0] + 1, coords[1]
		if !suicideTestGamma(chains, board, x, y) {
			return false
		}
	}
	if coords[1] != 0{
		x, y = coords[0], coords[1] - 1
		if !suicideTestGamma(chains, board, x, y) {
			return false
		}
	}
	if coords[1] != 18{
		x, y = coords[0], coords[1] + 1
		if !suicideTestGamma(chains, board, x, y) {
			return false
		}
	}
	return true
}

func suicideTestGamma(chains [][][][]int, board[][]int, x int, y int) bool{
	if board[x][y] == 0 {
			return false
	} else {
		if board[x][y] == whoAmI {
			if len(findChainLibertiesGamma(chains[(whoAmI - 1) % 2][findChainAttachedGamma(chains[(whoAmI - 1) % 2], []int{x, y})], board)) >= 2 {
				return false
			}
		} else {
			if len(findChainLibertiesGamma(chains[whoAmI % 2][findChainAttachedGamma(chains[whoAmI % 2], []int{x, y})], board)) == 1 {
				return false
			}
		}
	}
	return true
}

func testPassGamma(board [][]int) bool {
	x := 0
	y := 0
	stones := 0
	for x < 19{
		y = 0
		for y < 19{
			if board[x][y] != 0{
				stones++
			}
			y++
		}
		x++
	}
	if stones > 200 {
		return true
	}
	return false
}

func randomGamma(n int) int {
	return int(math.Abs(float64(int(float64(time.Since(time.Date(1, 1, 1, 0, 0, 0, 1, time.FixedZone("UTC", 0))))*math.Pi) % n)))
}