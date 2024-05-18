package banner

import (
	"bufio"
	//"fmt"
	"log"
	"os"
	"strings"
)

// Read banner files and store the data in a map
func ReadBannerFiles(txt string) map[int][]string {
	DATA := make(map[int][]string)
	file, err := os.Open(txt)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	read := bufio.NewScanner(file)
	ascciStart := 31
	for read.Scan() {
		if read.Text() == "" {
			ascciStart++
		} else {
			DATA[ascciStart] = append(DATA[ascciStart], read.Text())
		}
	}
	return DATA
}

// Check if all characters in the input are within printable ASCII range
func CheckIfAllCharInFile(words []string) bool {
	Temp := strings.Join(words, "")
	for _, char := range Temp {
		if char < ' ' || char > '~' {
			return false
		}
	}
	return true
}

// Generate and print the result
func Result(words []string, newLineCounter int, banner map[int][]string) string {
	counter := 1
	result := ""
	for _, word := range words {
		if word == "" && counter <= newLineCounter {
			//fmt.Println()
			result += "<br/>"
			counter++
			continue
		}
		for i := 0; i < 8; i++ {
			for j, char := range word {
				if j == len(word)-1 {
					result += banner[int(char)][i] + "<br/>"
					//fmt.Println()
					continue
				}
				//fmt.Print(banner[int(char)][i])
				result += banner[int(char)][i]

			}
		}
	}
	return result
}
