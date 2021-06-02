package main
import(
	"fmt"
)

func main() {
	m := 5
	//loop for rows
	for i := 1; i <= 5; i++ {
	    //loop for printing stars
		for j := 1; j <= m; j++ {
			fmt.Print("* ")
		}
		
		fmt.Print("\n")
		//loop for spacing in each row
	    for k := 1; i >= k; k++ {
	        fmt.Print(" ")
	    }
	    m = m-1        
	}
}