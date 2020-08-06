package main

func main() {

}

func checkStraightLine(coordinates [][]int) bool {
	pDeltaX := coordinates[1][0] - coordinates[0][0]
	pDeltaY := coordinates[1][1] - coordinates[0][1]
	for i := 1; i < len(coordinates)-1; i++ {
		deltaX := coordinates[i+1][0] - coordinates[i][0]
		deltaY := coordinates[i+1][1] - coordinates[i][1]
		if deltaX*pDeltaY != pDeltaX*deltaY {
			return false
		}
	}
	return true
}
