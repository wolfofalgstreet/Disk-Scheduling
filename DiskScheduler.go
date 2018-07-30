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
  //"reflect"
  "strconv"
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
func lookConvert(word string, lines []string) (int) {
  index := find(word, lines)
  num, err := strconv.Atoi(lines[index + 1])

  // Type conversion error handling
  if err != nil {
    //fmt.Println("Conversion of var ", word, " failed!\n")
  }
  return num
}


// ----------------------------------- //
// Read Input File parameter in arg[1] //
func readInputFile()(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int) {

  // Fetch name and read file
  fileName := os.Args[1]
  input, err := os.Open(fileName)
  checkErr(err)
  defer input.Close()

  // Scanning each line of file
  scanner := bufio.NewScanner(input)
  scanner.Split(bufio.ScanLines)

  // Lines will be scanned into lines[]
  var lines []string
  var words []string

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  // Extracting info for each line
  for i := 0; i < len(lines); i = i + 1 {

    // Each line in file gets split into a string array
    words = strings.Split(lines[i], " ")

    // Get algo type, cylinder bounds, and start location info
    // Avoids processing comments '#'
    switch words[0] {
    case "use":
      algorithm = words[1]
    case "lowerCYL":
      lowerCYL = lookConvert("lowerCYL", words)
    case "upperCYL":
      upperCYL = lookConvert("upperCYL", words)
    case "initCYL":
      initCYL = lookConvert("initCYL", words)
    case "cylreq":
      cylReqs = append(cylReqs, lookConvert("cylreq", words))
    }
  }

  // Confirming read input
  fmt.Println("algorithm: ", algorithm, " lowerCYL: ", lowerCYL, " upperCYL: ", upperCYL, " initCYL: ", initCYL)
  fmt.Println("Cylinder Requests:")
  for i := 0; i < len(cylReqs); i = i + 1 {
    fmt.Println(cylReqs[i])
  }

  return algorithm, lowerCYL, upperCYL, initCYL, cylReqs
}


// ---------------- //
// Initiate Program //
func main() {

  // Read input file and structure data
  algorithm, lowerCYL, upperCYL, initCYL, cylReqs := readInputFile()

  // Execute chosen scheduling algorithm
  switch algorithm {
  case "fcfs":
    //
  case "sstf":
    //
  case "scan":
    //
  case "c-scan":
    //
  case "look":
    //
  case "c-look":
    //
  }

  algorithm, lowerCYL, upperCYL, initCYL, cylReqs = algorithm, lowerCYL, upperCYL, initCYL, cylReqs //

}
