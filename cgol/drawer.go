package cgol

import "fmt"

func Draw(cm CgolManager) {
	println("     ======================================     ")
	for x := 0; x < cm.xSize; x++ {
		for y := 0; y < cm.ySize; y++ {
			if cm.currentGeneration[x][y] == 1 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	println("     ======================================     ")
}
