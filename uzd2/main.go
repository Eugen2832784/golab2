//Jevgenijs Kovalonoks 251RDB125 7.grupa

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main(){
	synonyms:= make(map[string] []string)
	reader := bufio.NewReader(os.Stdin)
	for{
		teikums:=sanemtTeikumu(reader)
		sliceVardi:= strings.Split(teikums, " ")
	if len(sliceVardi) == 0{
		return
	}
	if teikums == "done" {
			break
		}
	switch sliceVardi[0]{
	case "add":
		add(synonyms,sliceVardi[1], sliceVardi[2] )
	case "count":
		count(synonyms, sliceVardi[1])
	case "check":
		check(synonyms,sliceVardi[1], sliceVardi[2])
	default: 
		fmt.Println("error")
	}	
	}
}

func sanemtTeikumu(reader *bufio.Reader) string {
	teikums, _ := reader.ReadString('\n')
	teikums = strings.TrimSpace(teikums)
	teikums = strings.ToLower(teikums)
	return teikums
}

func add(synonyms map[string][]string, word1, word2 string){
	if !contains(synonyms[word1], word2){
		synonyms[word1] = append(synonyms[word1],word2)
	}
	if !contains(synonyms[word2], word1){
		synonyms[word2] = append(synonyms[word2],word1)
	}
	
}

func contains(slice []string, vards string)bool{
	for _, val:= range slice{
		if val == vards{
			return true
		} 
	}
	return false
}

func count(synonyms map[string][]string, word string){
	if list, exists := synonyms[word]; exists {
		fmt.Println(len(list))
	} else {
		fmt.Println(0)
	}
}

func check(synonyms map[string][]string, word1, word2 string){
	if list, exists := synonyms[word1]; exists {
		if contains(list, word2) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No") 

}
