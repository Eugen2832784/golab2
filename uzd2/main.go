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
		teikums:=sanemtTeikumu()
	sliceVardi:= strings.Split(teikums, " ")
	if len(sliceVardi) == 0{
		return
	}
	switch sliceVardi[0]{
	case "add":
		add(synonyms,sliceVardi[1], sliceVardi[2] )
	case "count":
	case "check":
	default: 
	fmt.Println("error")
	}	
	}
}

func sanemtTeikumu() string {
	
	teikums, _ := reader.ReadString('\n')
	teikums = strings.ToLower(teikums)
	return teikums
}

func add(synonyms map[string][]string, word1, word2 string){
	if !contains(synonyms[word1], word2){
		synonyms[word2] = append(synonyms[word1],word2)
	}
	if !contains(synonyms[word2], word1){
		synonyms[word2] = append(synonyms[word2],word1)
	}
	
}