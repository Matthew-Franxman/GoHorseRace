//ok yeah I know this is completely ridiculous but I just wanted to do this
//and its a pretty fun practice at how to run goroutines and select
package main

import (
	"fmt"
	"time"
)

//im gonna try and help y'all learn how all this shit works
//so below each one of these functions are called goroutines
//basically what that means is instead of running in order, each one of these runs simultaneously with the main function
//but that creates a problem at times cause sometimes the main will terminate before the goroutines are done running
//but we'll get into that more later

//so this is our first of five goroutines blueHorse
//it requires a channel to run
//basically a channel is a pipe that connects all of the routines together
//for thisprogram all of the pipes just go straight through the main
//I arbitrarily just named all of their channels ch but they are all different in the main
//so each of these five functions are running simultaneously with the main
func blueHorse(ch chan string) {
	ch <- "Blue Horse" //this means that we are throwing 'Blue Horse' into the channel
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
	//each channel has a data type associated with it (ie string, int, whatever)
	//and each has a maximum capacity of 3, but you don't really have to have one of those
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)
	ch3 := make(chan string, 3)
	ch4 := make(chan string, 3)
	ch5 := make(chan string, 3)

	//I then made a for loop to loop through the goroutines 3 times each, meaning there are 16 routines running
	//in this file alltogether, a maximum of 6 at a time
	//the 'go' functions tell the routine to start, and each routine will pass data to their respective channel
	for i := 0; i < 3; i++ {
		go blueHorse(ch1)
		go orangeHorse(ch2)
		go redHorse(ch3)
		go yellowHorse(ch4)
		go greenHorse(ch5)
		time.Sleep(3 * time.Second) //sleep means the computer waits for 3 seconds

		//the reason the sleeps are in means that all of the pipes are guaranteed to have data by the time the main wakes back up
		//so that way the computer is forced to pick at random which channel to pull from
		//when the main pulls from a channel, it takes it completely out of there and the channel is now empty
		//select works as a switch for channels, picking 1 to choose from and returning something
		//in this case that something is a Println with the specific horse in the line
		if i == 2 {
			fmt.Println("They are heading down the stretch!")
			time.Sleep(2 * time.Second)

			select {
			case s1 := <-ch1: //the <-ch1 syntax is for taking data out of a channel
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

			//I fucking hate this syntax for else statements
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
