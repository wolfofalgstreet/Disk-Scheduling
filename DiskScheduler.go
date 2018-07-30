// Created by Isaias Perez Vega
// Disk - Scheduling Algorithms
// =================================================
// Program reads an input file with cylinder requests
// File specifies scheduling algorithm, program outputs
// The optimal schedule based on algorithm.


package main

import (
  "bufio"
  "fmt"
  "strings"
  "os"
)

// --------------------------- //
// Streamlining error checking //
func checkErr(e error) {
  if e != nil {
    fmt.Println("Error ocurred when trying to open file.\n")
    panic(e)

  }
}

// ---------------------------- //
// Search index of word in list //
func find(item string, list []string) (index int) {
  count := 0
  index = -1
  for _,  word:= range list {
    if word == item {
      index = count
    }
    count = count + 1
  }

  // Before returning check if the word was found
  if index == -1 {
    //fmt.Println("Did not find ", item, " in memory..")
  }
  return index
}


// ------------------------------------------------ //
// Looks for a word in array and converts it to int //
func lookConvert(word string, words []string) (int) {
  index := find(word, words)
  num, err := strconv.Atoi(words[index + 1])

  // Type conversion error handling
  if err != nil {
    //fmt.Println("Conversion of var ", word, " failed!\n")
  }
  return num
}


// ----------------------------------- //
// Read Input File parameter in arg[1] //
func readInputFile() {

  // Fetch name and read file
  fileName := os.Args[1]
  input, err := os.Open(fileName)
  checkErr(err)
  defer input.Close()

  // Scanning word by Word
  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanLines)

  // Words will be scanned into words[]
  var words []string
  for scanner.Scan() {
    words = append(words, scanner.Text())
  }

  // Get scheduler setup info
  algorithm := find("use", words[0])
  lowerCYL := find()

  //fmt.Println("name: ", fileName)

}


// ---------------- //
// Initiate Program //
func main() {

  readInputFile()

}
