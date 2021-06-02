package main
import (
	"fmt"
)

func main() {
	fmt.Println("even numbers from 1 to 50:")
	//looping from 1 to 50
	for i := 1; i<= 50; i++ {
		//finding reminder zero by mod division of each i by 2 and printing even no.s
	    if i % 2 == 0 {
	       fmt.Printf("%v \n",i)
	    }
	}
}
