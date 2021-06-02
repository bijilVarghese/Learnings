package main
import (
	"fmt"
)

func main() {
	printOneToFifty()
	printEvenNumbersBetweenOneToFifty()
	printOddNumbersBetweenOneToFifty()
    printPrimeNummbersBetweenOneToFifty()
}

func printOneToFifty() {
	fmt.Println("numbers from 1 to 50:")

	//looping from 1 to 50 and print each no.
	for i := 1; i <= 50; i++ {
	   fmt.Printf("%v \n",i)
	}
}

func printEvenNumbersBetweenOneToFifty() {
	fmt.Println("even numbers from 1 to 50:")

	//looping from 1 to 50
	for i := 1; i<= 50; i++ {
		//finding reminder zero by mod division of each i by 2 and printing even no.s
	    if i % 2 == 0 {
	       fmt.Printf("%v \n",i)
	    }
	}  
}

func printOddNumbersBetweenOneToFifty() {
	fmt.Println("odd numbers from 1 to 50:")
	
	//looping from 1 to 50
	for i := 1; i<= 50; i++ {
		//finding reminder not equal to zero by mod division of each i by 2 and print odd no.s
	    if i % 2 != 0 {
	       fmt.Printf("%v\n",i)
	    }
	}
}

func printPrimeNummbersBetweenOneToFifty() {
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

