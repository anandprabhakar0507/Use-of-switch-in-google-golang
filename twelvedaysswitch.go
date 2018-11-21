package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		fmt.Println("on the", day, "of christmas my true love sent to me")

		if day == 12 {
			fmt.Println("twelve drummers drumming")
		} else if day == 11 {
			fmt.Println("Eleven pipers piping")
		} else if day == 10 {
			fmt.Println("ten lords a leaping")
		} else if day == 9 {
			fmt.Println("nine ladies dancing")
		} else if day == 8 {
			fmt.Println("eight maids a milking")
		} else if day == 7 {
			fmt.Println("seven swans swimming")
		} else if day == 6 {
			fmt.Println("six geese a laying")
		} else if day == 5 {
			fmt.Println("five golden rings")
		} else if day == 4 {
			fmt.Println("four calling birds")
		} else if day == 3 {
			fmt.Println("three french hens")
		} else if day == 2 {
			fmt.Println("two turtle doves")
		} else if day == 1 {
			fmt.Println("a partridge in fear tree.")
		}
	}

}
