//ok yeah I know this is completely ridiculous but I just wanted to do this
//and its a pretty fun practice at how to run goroutines and select
package main

import (
	"fmt"
	"time"
)

func blueHorse(ch chan string) {
	ch <- "Blue Horse"
}

func orangeHorse(ch chan string) {
	ch <- "Orange Horse"
}

func redHorse(ch chan string) {
	ch <- "Red Horse"
}

func yellowHorse(ch chan string) {
	ch <- "Yellow Horse"
}

func greenHorse(ch chan string) {
	ch <- "Green Horse"
}

func main() {
	//making the channel for each of the goroutines
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)
	ch3 := make(chan string, 3)
	ch4 := make(chan string, 3)
	ch5 := make(chan string, 3)

	for i := 0; i < 3; i++ {
		go blueHorse(ch1)
		go orangeHorse(ch2)
		go redHorse(ch3)
		go yellowHorse(ch4)
		go greenHorse(ch5)
		time.Sleep(3 * time.Second)

		if i == 2 {
			fmt.Println("They are heading down the stretch!")
			time.Sleep(2 * time.Second)

			select {
			case s1 := <-ch1:
				fmt.Println(s1, " has won!")
			case s2 := <-ch2:
				fmt.Println(s2, " has won!")
			case s3 := <-ch3:
				fmt.Println(s3, " has won!")
			case s4 := <-ch4:
				fmt.Println(s4, " has won!")
			case s5 := <-ch5:
				fmt.Println(s5, " has won!")
			} //end of the select
		} else {
			select {
			case s1 := <-ch1:
				fmt.Println("Now ", s1, " is in the lead!")
			case s2 := <-ch2:
				fmt.Println("Now ", s2, " is in the lead!")
			case s3 := <-ch3:
				fmt.Println("Now ", s3, " is in the lead!")
			case s4 := <-ch4:
				fmt.Println("Now ", s4, " is in the lead!")
			case s5 := <-ch5:
				fmt.Println("Now ", s5, " is in the lead!")
			} //end of the select
		} //end of else

	} //end of for statement
} //end of main
