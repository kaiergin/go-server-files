package main
import( "fmt"
		"math/rand")
func returnMove2(board [][]int) []int {
	if testPassAlpha(board){
		fmt.Println("Bot Passed")
		return []int{-1, -1}
	}
	chains := botFindChainsAlpha(board)
	minLib := make([][]int, 127)
	var curLib [][]int
	for n := range chains[whoAmI % 2] {
		curLib = findChainLibertiesAlpha(chains[whoAmI % 2][n], board)
		if len(minLib) > len(curLib) {
			minLib = curLib
		}
	}
	if len(minLib) > 0 {
		index := rand.Intn(len(minLib))
		fmt.Println("BOT moved:", minLib[index][1], minLib[index][0])
		return minLib[index]
	} else {
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
	}
	fmt.Println("Bot Passed")
	return []int{-1, -1}
}
func botFindChainsAlpha(board [][]int) [][][][]int {
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

func findChainLibertiesAlpha(chain [][]int, board [][]int) [][]int {
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

func testPassAlpha(board [][]int) bool {
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