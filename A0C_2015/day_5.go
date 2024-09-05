package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Santa needs help figuring out which strings in his text file are naughty or nice.

// Criteria for nice strings:
// > Contains at least 3 vowels (aeiou)
// > Contains at least one letter that appears twice in a row (xx, abcdde, aabbccdd)
// > Should not contain the strings (ab,cd,pq,xy)

func isNice(str string) int {
	 nonstrings := []string{"ab","cd","pq","xy"}
     

	 for _, forbidden := range nonstrings {
		if strings.Contains(str, forbidden) {
			  return 0

		}
	 }

	vowelCount := 0
	hasDouble := false

	for i := 0; i < len(str); i++ {

		if isVowel(rune(str[i])) {
			vowelCount++	
		}

		if  i > 0 && str[i] == str[i-1] {
			 hasDouble = true
		}
	}
   
	if vowelCount >= 3 && hasDouble {
		return 1

	}
	return 0
}


func isVowel(char rune) bool {
	return char == 'a'|| char == 'e'|| char == 'i'|| char == 'o'|| char == 'u'
}


func main(){

	file, _ := os.Open("data.txt")

	defer file.Close()

	filename := bufio.NewReader(file)

	newfile := bufio.NewScanner(filename)
    
	var output int

	totalNiceCount := 0

	for newfile.Scan(){
		
		line := newfile.Text()

		output = isNice(line)
        
		totalNiceCount += output

	}

	fmt.Println("The total number of nice strings is", totalNiceCount)

}