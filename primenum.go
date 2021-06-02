package main
import(
	"fmt"
)

func main() {
	fmt.Println("prime numbers between 1 to 50:")
	
	//looping through each no. upto 50
	for i := 2; i < 50; i++ {
		flag := 0

		//looping through each no.half of i
        for j := 2; j <= i/2; j++ {
          //dividing current no. i with no.s upto half of i to check whether i is divisible
		   if i % j == 0 {
				flag = 1
			}
		}
		if flag == 0 {
		   //printing prime no.s
            fmt.Printf("%v\n",i)   
		}

	}
}

	
	
	