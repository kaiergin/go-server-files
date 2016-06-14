package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

var mainChain [][]int
var scoreOne int = 0
var scoreTwo int = 0
var move int = 1

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
	var n []int
	var tempScore1 int
	var tempScore2 int
	var tempMove int
	
	// OPTIONS
	var debug bool = true
	var timing bool = false
	var printScore bool = true
	var showCoords bool = false
	var moveCount bool = false
	
	reader := bufio.NewReader(os.Stdin)
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
		printBoard(sliceMain,showCoords)
		if debug && revert2{
			fmt.Print("-> ")
			text, _ := reader.ReadString('\n')
			if text == "stop\n" {
				break
			} else if text == "-1 -1\n" || text == "-1 -1\r\n" {
				if pass {
					break
				} else {
					pass = true
				}
			} else {
				n = runParse(text, sliceMain)
				if checkIllegal(n) {
					sliceMain = makeMove(n[0], n[1], sliceMain, 1)
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
					}
				} else {
					revert1 = false
					strike1++
					fmt.Println("Strike:",strike1)
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
			
			t0 = time.Now()
			nextMove = returnMove(sliceMain)
			t1 = time.Now()
			
			if t1.Sub(t0) > 60000000000 && timing || tempScore1 != scoreOne || tempScore2 != scoreTwo || tempMove != move{
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
				if checkIllegal(nextMove) {
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
						}
					}
				} else {
					revert2 = false
					strike2++
					fmt.Println("Strike:",strike2)
					if strike2 == 3 {
						whoWon = 1
						break
					}
				}
			}
		}
		revert1 = true
		
		if printScore {
			fmt.Println("Player One Score:", scoreOne, "\nPlayer BOT Score:", scoreTwo)
		}
	}
	if whoWon == 1 {
		fmt.Println("Player 1 won by fault in player 2")
	} else if whoWon == 2 {
		fmt.Println("Player 2 won by fault in player 1")
	} else {
		endGame(sliceMain)
		fmt.Println("\nEnd Game Score\nPlayer One Score:", scoreOne, "\nPlayer BOT Score:", scoreTwo)
	}
	
	/*
	sliceMain = makeMove(4,4,sliceMain,1)
	sliceMain = makeMove(4,5,sliceMain,1)
	sliceMain = makeMove(7,7,sliceMain,1)
	sliceMain = makeMove(6,7,sliceMain,2)
	sliceMain = makeMove(8,7,sliceMain,2)
	sliceMain = makeMove(7,6,sliceMain,2)
	sliceMain = makeMove(7,8,sliceMain,2)
	chains := findChains(sliceMain,1)
	
	for x := range chains{
		printBoard(chains[x])
		fmt.Println("-----")
	}
	
	runDelete(sliceMain, chains, 1)
	printBoard(sliceMain)
	*/
}

func checkIllegal(check []int) bool {
	if check[0] < -1 || check[1] < -1 || check[0] > 18 || check[1] > 18 {
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

func runParse(text string, sliceMain [][]int) []int {
	if isWindows() {
		text = text[:len(text)-2]
	} else {
		text = text[:len(text)-1]
	}
	s := strings.Split(text, " ")
	n := make([]int, 2)
	a, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil || -1 > a || a > 18 {
		fmt.Println(err)
		n[0], n[1] = 20, 20
		return n
	}
	b, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil || -1 > b || b > 18 {
		fmt.Println(err)
		n[0], n[1] = 20, 20
		return n
	}
	n[0], n[1] = int(b), int(a)
	return n
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

func printBoard(sliceMain [][]int, display bool) {
	if display {
		fmt.Println("   0 1 2 3 4 5 6 7 8 9 ...")
	}
	for x := range sliceMain {
		if display {
			if x<10{
				fmt.Println(x,sliceMain[x])
			} else {
				fmt.Println(" ",sliceMain[x])
			}
		} else {
			fmt.Println(sliceMain[x])
		}
	}
	fmt.Println("---------------------------------------")
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

func isWindows() bool {
    return os.PathSeparator == '\\' && os.PathListSeparator == ';'
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
