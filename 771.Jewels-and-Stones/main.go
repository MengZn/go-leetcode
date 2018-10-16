package main
import "fmt"

func main(){
	fmt.Printf("%v",numJewelsInStones("aA","aAAbbbb"))
}

func numJewelsInStones(J string, S string) int {
	index:=0
	for _, r := range S{
		c := string(r)
		for _ , j :=range J{
			c1 := string(j)
			if(c==c1){
				index++
			}
		}
	}
	return index
}

