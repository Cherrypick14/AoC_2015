package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

file, _ := os.Open("data.txt")

defer file.Close()

filename := bufio.NewReader(file)

scanFile := bufio.NewScanner(filename)

var finalDelivery int

for scanFile.Scan() {
	line := scanFile.Text()

	finalDelivery = Delivery(line)
}

fmt.Println(finalDelivery)
}


func Delivery(s string) int {

   //keep track of the visited houses
   count:= 0

   checkHouse := make(map[[2]int]bool)

   //both Santa and Rob start delivering presents in the first house

   checkHouse[[2]int{0,0}] = true
   count++
   
   //initialize positions for both Santa and Rob
   var santaX, santaY, robSantaX, robSantaY int

   for i := 0; i < len(s); i++ {
	  if (i%2) == 0 {
		switch string(s[i]) {
		case "^":
			santaY++
		case "v":
			santaY--
		case ">":
			santaX++
		case "<":
			santaX--
		}

		if !checkHouse[[2]int{santaX, santaY}] {
			checkHouse[[2]int{santaX,santaY}] = true
			count++
		}

	  } else {
		 switch string(s[i]){
		 case "^":
			robSantaY++
		 case "v":
			robSantaY--
		 case ">":
			robSantaX++
		 case "<":
			robSantaX--
		 }

		 if !checkHouse[[2]int{robSantaX, robSantaY}] {
			 checkHouse[[2]int{robSantaX,robSantaY}] = true
			 count++
		 }
	  }
   }
   fmt.Println(checkHouse)
   return count
}