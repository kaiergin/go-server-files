package main

// Competition Engine

import (
	"time"
	"os"
	"strconv"
)

var mainChain [][]int
var scoreOne int = 0
var scoreTwo int = 3
var move int = 1
var whoAmI int = 0
var turns int = 1

func main() {
	var strike1 int = 0
	var strike2 int = 0
	ko1 := make([][]int, 19)
	for x := range ko1 {
		ko1[x] = make([]int, 19)
		for y := range ko1[x]{
			ko1[x][y] = 0
		}
	}
	ko2 := make([][]int, 19)
	for x := range ko2 {
		ko2[x] = make([]int, 19)
		for y := range ko2[x]{
			ko2[x][y] = 0
		}
	}
	var chains [][][]int
	var whoWon int
	var t0 time.Time
	var t1 time.Time
	var tempScore1 int
	var tempScore2 int
	var tempMove int
	var tempTurn int
	
	var timing bool = true
	var moveCount bool = true
	
	sliceMain := make([][]int, 19)
	for x := range sliceMain {
		sliceMain[x] = make([]int, 19)
		for y := range sliceMain[x]{
			sliceMain[x][y] = 0
		}
	}
	
	revert1 := true
	revert2 := true
	
	pass := false
	mainLoop := true
	var nextMove []int
	
	for mainLoop {
		if revert2 {
			tempScore1 = scoreOne
			tempScore2 = scoreTwo
			tempMove = move
			whoAmI = 1
			tempTurn = turns
			
			t0 = time.Now()
			nextMove = returnMove1(sliceMain)
			t1 = time.Now()
			
			if (t1.Sub(t0) > 60000000000 && timing) || tempScore1 != scoreOne || tempScore2 != scoreTwo || tempMove != move || turns != tempTurn {
				whoWon = 2
				break
			}
			
			if nextMove[0] == -1 && nextMove[1] == -1 {
				if pass {
					break
				} else {
					pass = true
				}
			} else {
				if checkIllegal(nextMove, sliceMain) {
					sliceMain = makeMove(nextMove[0], nextMove[1], sliceMain, 1)
					chains = findChains(sliceMain,2)
					runDelete(sliceMain, chains, 2)
					chains = findChains(sliceMain,1)
					runDelete(sliceMain, chains, 1)
					pass = false
					if testEqu(ko1,sliceMain) {
						sliceMain = copyBoard(sliceMain, ko2)
						scoreOne--
						revert1 = false
						strike1++
						if strike1 == 3 {
							whoWon = 2
							break
						}
					} else {
						ko1 = copyBoard(ko1, sliceMain)
						turns++
					}
				} else {
					revert1 = false
					strike1++
					if strike1 == 3 {
						whoWon = 2
						break
					}
				}
			}
		}
		
		revert2 = true
		
		if revert1 {
			tempScore1 = scoreOne
			tempScore2 = scoreTwo
			tempMove = move
			whoAmI = 2
			tempTurn = turns
			
			t0 = time.Now()
			nextMove = returnMove2(sliceMain)
			t1 = time.Now()
			
			if t1.Sub(t0) > 60000000000 && timing || tempScore1 != scoreOne || tempScore2 != scoreTwo || tempMove != move || turns != tempTurn {
				whoWon = 1
				break
			}
			
			if nextMove[0] == -1 && nextMove[1] == -1 {
				if pass {
					break
				} else {
					pass = true
				}
			} else {
				if checkIllegal(nextMove, sliceMain) {
					sliceMain = makeMove(nextMove[0], nextMove[1], sliceMain, 2)
					chains = findChains(sliceMain,1)
					runDelete(sliceMain, chains, 1)
					chains = findChains(sliceMain,2)
					runDelete(sliceMain, chains, 2)
					pass = false
					if testEqu(ko2,sliceMain) {
						sliceMain = copyBoard(sliceMain, ko1)
						scoreTwo--
						revert2 = false
						strike2++
						if strike2 == 3 {
							whoWon = 1
							break
						}
					} else {
						ko2 = copyBoard(ko2, sliceMain)
						if moveCount {
							move++
							turns++
						}
					}
				} else {
					revert2 = false
					strike2++
					if strike2 == 3 {
						whoWon = 1
						break
					}
				}
			}
		}
		revert1 = true
	}
	
	f, err := os.Create("complete.txt")
	check(err)
	defer f.Close()
	f.WriteString("Completed!")

}

func checkIllegal(check []int, sliceMain [][]int) bool {
	if check[0] < -1 || check[1] < -1 || check[0] > 18 || check[1] > 18 || sliceMain[check[0]][check[1]] != 0 {
		return false
	}
	return true
}

func copyBoard(a [][]int, b [][]int) [][]int {
	for x := range a {
		copy(a[x],b[x])
	}
	return a
}

func makeMove(a int, b int, sliceMain [][]int, changed int) [][]int {
	if sliceMain[a][b] == 0 {
		sliceMain[a][b] = changed
	}
	return sliceMain
}

func Delete(a int, b int, sliceMain [][]int) [][]int {
	sliceMain[a][b] = 0
	return sliceMain
}

func runDelete(sliceMain [][]int, chains [][][]int, player int) [][]int{
	for x := range chains{
		if deleteChains(sliceMain, chains[x], player){
			for a := range sliceMain{
				for b := range sliceMain[a]{
					if chains[x][a][b] == player{
						sliceMain = Delete(a,b,sliceMain)
						if player == 1 {
							scoreTwo++
						} else if player == 2 {
							scoreOne++
						}
					}
				}
			}
		}
	}
	return sliceMain
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findChains(sliceMain [][]int, player int) [][][]int{
	sliceOfChains := make([][][]int, 1)
	for i := 0; i < 1; i++ {
        sliceOfChains[i] = make([][]int, 19)
        for j := 0; j < 19; j++ {
        	sliceOfChains[i][j] = make([]int, 19)
    	}
    }
    var tempVar bool
    first := true
	for x := range sliceMain {
		for y := range sliceMain[x]{
			if sliceMain[x][y] == player {
				tempVar=true
				for a := range sliceOfChains{
					if sliceOfChains[a][x][y] == player {
						tempVar=false
						break
					}
				}
				if tempVar {
					mainChain = make([][]int, 19)
					for i := 0; i < 19; i++ {
						mainChain[i] = make([]int, 19)
					}
					mainChain[x][y] = player
					Chain(sliceMain,x,y,player)
					if first {
						sliceOfChains[0] = mainChain
						first = false
					} else{
						sliceOfChains = append(sliceOfChains,mainChain)
					}
				}
			}
		}
	}
	return sliceOfChains
}

func Chain(sliceMain [][]int, x int, y int, player int) {
	
	a, b, c, d := true, true, true, true
	
	if x == 0 {
		b = false
	}
	if y == 0 {
		d = false
	}
	if x == 18 {
		a = false
	}
	if y == 18 {
		c = false
	}
	
	if a && sliceMain[x+1][y] == player{
		if mainChain[x+1][y] != player {
			mainChain[x+1][y] = player
			Chain(sliceMain, x+1, y, player)
		}
	}
	if b && sliceMain[x-1][y] == player{
		if mainChain[x-1][y] != player {
			mainChain[x-1][y] = player
			Chain(sliceMain, x-1, y, player)
		}
	}
	if c && sliceMain[x][y+1] == player{
		if mainChain[x][y+1] != player {
			mainChain[x][y+1] = player
			Chain(sliceMain, x, y+1, player)
		}
	}
	if d && sliceMain[x][y-1] == player{
		if mainChain[x][y-1] != player {
			mainChain[x][y-1] = player
			Chain(sliceMain, x, y-1, player)
		}
	}
}

func deleteChains(sliceMain [][]int, chain [][]int, player int) bool {
	var a, b, c ,d int
	for x := range sliceMain {
		for y := range sliceMain[x] {
			if sliceMain[x][y] == player && chain[x][y] == player {
				a, b, c, d = 1, 1, 1, 1
				if x == 0 {
					b = 0
				}
				if y == 0 {
					d = 0
				}
				if x == 18 {
					a = 0
				}
				if y == 18 {
					c = 0
				}
				if sliceMain[x+a][y] == 0 || sliceMain[x-b][y] == 0 || sliceMain[x][y+c] == 0 || sliceMain[x][y-d] == 0 {
					return false
				}
			}
		}
	}
	return true
}

func testEqu(a, b [][]int) bool {

    if a == nil && b == nil { 
        return true; 
    }

    if a == nil || b == nil { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

	for i := range a {
		for j := range a[i]{
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

    return true
}

func endGame(sliceMain [][]int) {
	chains := findEndgame(sliceMain)
	runScoring(sliceMain, chains, 1)
	runScoring(sliceMain, chains, 2)
}

func whoLand(sliceMain [][]int, chain [][]int, player int) bool {
	var a, b, c, d int
	for x := range sliceMain {
		for y := range sliceMain[x] {
			if sliceMain[x][y] == 0 && chain[x][y] == 3 {
				a, b, c, d = 1, 1, 1, 1
				if x == 0 {
					b = 0
				}
				if y == 0 {
					d = 0
				}
				if x == 18 {
					a = 0
				}
				if y == 18 {
					c = 0
				}
				if sliceMain[x+a][y] == player || sliceMain[x-b][y] == player || sliceMain[x][y+c] == player || sliceMain[x][y-d] == player {
					return false
				}
			}
		}
	}
	return true
}

func findEndgame(sliceMain [][]int) [][][]int{
	sliceOfChains := make([][][]int, 1)
	for i := 0; i < 1; i++ {
        sliceOfChains[i] = make([][]int, 19)
        for j := 0; j < 19; j++ {
        	sliceOfChains[i][j] = make([]int, 19)
    	}
    }
    var tempVar bool
    first := true
	for x := range sliceMain {
		for y := range sliceMain[x]{
			if sliceMain[x][y] == 0 {
				tempVar=true
				for a := range sliceOfChains{
					if sliceOfChains[a][x][y] == 0 {
						tempVar = false
						break
					}
				}
				if tempVar {
					mainChain = make([][]int, 19)
					for i := 0; i < 19; i++ {
						mainChain[i] = make([]int, 19)
					}
					mainChain[x][y] = 3
					Chain(sliceMain,x,y,3)
					if first {
						sliceOfChains[0] = mainChain
						first = false
					} else{
						sliceOfChains = append(sliceOfChains,mainChain)
					}
				}
			}
		}
	}
	return sliceOfChains
}

func runScoring(sliceMain [][]int, chains [][][]int, player int) [][]int{
	for x := range chains{
		if whoLand(sliceMain, chains[x], player){
			for a := range sliceMain{
				for b := range sliceMain[a]{
					if chains[x][a][b] == player{
						if player == 1 {
							scoreTwo++
						} else if player == 2 {
							scoreOne++
						}
					}
				}
			}
		}
	}
	return sliceMain
}
/*
func deadStrings(sliceMain [][]int, chains [][][]int) [][]int {
	var sliceTrial [][]int
	copy(sliceTrial, sliceMain)
	checkOne := false
	var checkTwo bool
	var f int
	var g int
	var h int
	var i int
	for x := range chains {
		checkTwo = false
		for a := range chains[x]{
			if checkTwo {
				break
			}
			for b := range chains[x][a] {
				if a == 0 {
					f = 0
				}
				if a == 18 {
					g = 0
				}
				if b == 0 {
					h = 0
				}
				if b == 18 {
					i = 0
				}
				if chains[x][a][b] != 0 {
					if sliceMain[a+g][b] == 0 {
						if checkOne {
							checkTwo = true
							break
						} else {
							checkOne = true
						}
					}
					if sliceMain[a][b+i] == 0 {
						if checkOne {
							checkTwo = true
							break
						} else {
							checkOne = true
						}
					}
					if sliceMain[a-f][b] == 0 {
						if checkOne {
							checkTwo = true
							break
						} else {
							checkOne = true
						}
					}
					if sliceMain[a][b-h] == 0 {
						if checkOne {
							checkTwo = true
							break
						} else {
							checkOne = true
						}
					}
				}
			}
		}
		if !checkTwo {
			
		}
	}

}
*/
