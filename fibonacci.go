package main
import (
	"fmt"
)

func main(){
    var f0,f1,fn int
	fmt.Println("fibonacci series b/w 1 to 50")
	f0 = 0
	f1 = 1
	fn = 0
	//looping upto sum of f0 and f1 less than 50 
	for ; f0 + f1 <= 50; {
		if fn == 0 {
			fmt.Printf("%v\n",f1)
		}
		fn = f0 + f1
		fmt.Printf("%v\n",fn)
		//assigning current f1,fn values to f0,f1 respectively
		f0 = f1
		f1 = fn
		
	}
}