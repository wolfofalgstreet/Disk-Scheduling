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
  "math"
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

  // Temporary storage
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

  return algorithm, lowerCYL, upperCYL, initCYL, cylReqs
}


// ----------------------------------------------------- //
// Print formatted setup info and list cylinder requests //
func printReadInput(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {

  fmt.Println("Seek algorithm: ", strings.ToUpper(algorithm))
  fmt.Println(" Lower cylinder: ", lowerCYL)
  fmt.Println(" Upper cylinder: ", upperCYL)
  fmt.Println(" Init cylinder requests: ")
  for i := 0; i < len(cylReqs); i = i + 1 {
    fmt.Println("  Cylinder   ", cylReqs[i])
  }
}


// --------------------------------------- //
// Check if cylinder is within disk bounds //
func checkBounds(lowerCYL, upperCYL, cylinder int)(bool) {
  inBounds := false
  if cylinder > lowerCYL && cylinder < upperCYL {
      inBounds = true
  }
  return inBounds
}


// ------------------------------------------------------------------------ //
// Execute the First-come First-Served disk scheduling algorithm will print //
// cylinder number as it processes them and calculate total seek distance   //
func runFCFS(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {
  nextDistance := float64(cylReqs[0] - initCYL)
  seekDistance := int(nextDistance)
  fmt.Println("Servicing   ", cylReqs[0])

  for i := 1; i < len(cylReqs); i = i + 1 {
    fmt.Println("Servicing   ", cylReqs[i])
    if checkBounds(lowerCYL, upperCYL, cylReqs[i]) {
      nextDistance = math.Abs(float64(cylReqs[i] - cylReqs[i - 1]))
      seekDistance = seekDistance + int(nextDistance)
    } else {
      fmt.Println("Cylinder is out of bounds")
    }

  }
  fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)
}


//---------------------- //
// Perform a Bubble sort //
func sort(cylReqs []int)([]int) {

  // Bubble sort the structs by arrival time
  for x := 0; x < len(cylReqs); x = x + 1 {
    for i := 0; i < len(cylReqs) - 1; i = i + 1 {
      if cylReqs[i] > cylReqs[i + 1] {
        cylReqs[i], cylReqs[i + 1] = cylReqs[i + 1], cylReqs[i]
      }
    }
  }

  return cylReqs
}


// ------------------------------------------------------------------------ //
// Execute the First-come First-Served disk scheduling algorithm will print //
// cylinder number as it processes them and calculate total seek distance   //
func runSSTF2(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {
    var usedIndex []int     // Stores order of used cylinders
    var seekSorted []int    // Stores sorted cylinders
    seekSorted = cylReqs
    var nextCylinder int
    //seekDistance := 0
    nextDistance := 0;
    seekDistance := int(nextDistance)

    for i := 0; i < len(cylReqs); i = i + 1 {
      seekSorted = sort(seekSorted)
      nextCylinder = seekSorted[0]

      fmt.Println("Servicing   ", cylReqs[0])
      nextDistance = float64(nextCylinder - initCYL)

      for i := 1; i < len(cylReqs); i = i + 1 {
        fmt.Println("Servicing   ", cylReqs[i])
        if checkBounds(lowerCYL, upperCYL, cylReqs[i]) {
          nextDistance = math.Abs(float64(cylReqs[i] - cylReqs[i - 1]))
          seekDistance = seekDistance + int(nextDistance)
        } else {
          fmt.Println("Cylinder is out of bounds")
        }

      }
      seekSorted = seekSorted[1:]
      fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)

    }

    usedIndex, seekSorted, nextCylinder = usedIndex, seekSorted, nextCylinder

}


func runSSTF(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {
    var usedIndex []int     // Stores order of used cylinders
    var seekSorted []int    // Stores sorted cylinders
    seekSorted = cylReqs
    var nextCylinder int
    var min int

    min = upperCYL
    //seekDistance := 0
    nextDistance := 0;
    seekDistance := int(nextDistance)

    for i := 0; i < len(cylReqs); i = i + 1 {

    }

    for i := 0; i < len(cylReqs); i = i + 1 {
      seekSorted = sort(seekSorted)
      nextCylinder = seekSorted[0]

      fmt.Println("Servicing   ", cylReqs[0])
      nextDistance = float64(nextCylinder - initCYL)

      for i := 1; i < len(cylReqs); i = i + 1 {
        fmt.Println("Servicing   ", cylReqs[i])
        if checkBounds(lowerCYL, upperCYL, cylReqs[i]) {
          nextDistance = math.Abs(float64(cylReqs[i] - cylReqs[i - 1]))
          seekDistance = seekDistance + int(nextDistance)
        } else {
          fmt.Println("Cylinder is out of bounds")
        }

      }
      seekSorted = seekSorted[1:]
      fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)

    }

    usedIndex, seekSorted, nextCylinder = usedIndex, seekSorted, nextCylinder

}

// ---------------- //
// Initiate Program //
func main() {

  // Read input file and structure data
  algorithm, lowerCYL, upperCYL, initCYL, cylReqs := readInputFile()

  //
  printReadInput(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)

  // Execute chosen scheduling algorithm
  switch algorithm {

    case "fcfs":
      runFCFS(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)
    case "sstf":
      runSSTF(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)
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
