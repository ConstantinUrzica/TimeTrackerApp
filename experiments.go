/* This app allows you to keep track of the time invested in different activities. The data is kept for analysis and improvements.
* Start Date: June 5th 2025
* Author: @ConstantinUrzica
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	//Testing some changesgi
	fmt.Println("Someday we're gonna have a time tracker on our hands!")
	fmt.Println("Current time object cotains this crap: ", time)
	fmt.Println("Let's break it down...")
	hours, minutes, seconds := time.Clock()
	fmt.Printf("DEBUG: This should be the result of Clock(): %d:%d:%d\n", hours, minutes, seconds)
	year, month, day := time.Date()
	fmt.Printf("Debug: THis should be the result of Date(): %d %s %dth\n", year, month, day)
}
