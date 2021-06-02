package main
import (
	"fmt"
)

func main() {
	fmt.Println("odd numbers from 1 to 50:")
	//looping from 1 to 50
	for i := 1; i<= 50; i++ {
		//finding reminder not equal to zero by mod division of each i by 2 and print odd no.s
	    if i % 2 != 0 {
	       fmt.Printf("%v\n",i)
	    }
	}
}