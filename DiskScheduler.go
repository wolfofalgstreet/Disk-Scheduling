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
  "math"
  "strconv"
)

// ------------ //
// Struct       //
type cylinder struct {
  location int
  cost     int
}



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
  fmt.Println(" Init cylinder: ", initCYL)
  fmt.Println(" Cylinder requests:")
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


//-------------------------------- //
// Perform an Inversed Bubble sort //
func sort(cylReqs []cylinder, sortBy string)([]cylinder) {


  for x := 0; x < len(cylReqs); x = x + 1 {
    for i := 0; i < len(cylReqs) - 1; i = i + 1 {

      // Sort the structs by cost
      if sortBy == "cost" {
        if cylReqs[i].cost < cylReqs[i + 1].cost {
          cylReqs[i], cylReqs[i + 1] = cylReqs[i + 1], cylReqs[i]
        }
      } else {
        // Sort the structs by location
        if cylReqs[i].location < cylReqs[i + 1].location {
          cylReqs[i], cylReqs[i + 1] = cylReqs[i + 1], cylReqs[i]
        }
      }
    }
  }

  return cylReqs
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


// ------------------------------------------------------------------------- //
// Execute the Shortest seek time first disk scheduling algorithm will print //
// cylinder number as it processes them and calculate total seek distance    //
func runSSTF(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {

    // Create array of struct cylinders with respective locations
    // and initial costs
    var requests = make([]cylinder, len(cylReqs))
    var cyl cylinder
    current := initCYL
    for i := 0; i < len(cylReqs); i = i + 1 {
      cyl.location = cylReqs[i]
      cyl.cost = int(math.Abs(float64(cylReqs[i] - initCYL)))
      requests[i] = cyl
    }

    // Traversal Distance
    seekDistance := 0

    // For every requests, calculate costs and service the
    // one with the shortest seek time
    for i := 0; i < len(cylReqs); i = i + 1 {
      requests = sort(requests, "cost")

      fmt.Println("Servicing   ", requests[len(requests)-1].location)

      // Recalculate costs without current cylinder
      // and update traversal distance
      if requests[len(requests)-1].location < upperCYL && requests[len(requests)-1].location > lowerCYL {
        seekDistance = seekDistance + int(math.Abs(float64(requests[len(requests)-1].location - current)))
        current = requests[len(requests)-1].location
      } else {
        fmt.Println("Cylinder is out of bounds")
      }

      for i := 0; i < len(requests); i = i + 1 {
        //fmt.Println("cyl.loc: ", requests[i].location, " cyl.cost: ", requests[i].cost)
        requests[i].cost = int(math.Abs(float64(requests[i].location - current)))

      }

      requests = requests[:len(requests)-1]

    }
    fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)
}


// -------------------------------------------------------------------- //
// Execute the SCAN scheduling algorithm will print cylinder number as  //
// it processes them and calculate total seek distance                  //
func runSCAN(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {

  var requests = make([]cylinder, len(cylReqs))
  var cyl cylinder
  current := initCYL
  startIndex := 0

  // Create array of struct cylinders with respective locations
  // and initial costs
  for i := 0; i < len(cylReqs); i = i + 1 {
    cyl.location = cylReqs[i]
    cyl.cost = int(math.Abs(float64(cylReqs[i] - initCYL)))
    requests[i] = cyl
  }

  seekDistance := 0

  // Start by sorting request by location
  requests = sort(requests, "location")

  // Find First cylinder location to visit
  for i := 0; i < len(requests); i = i + 1 {
    if initCYL <= requests[i].location {
      startIndex = i
    }
  }

  // Service cylinders to the right(up) of the disk
  for i := startIndex; i >= 0; i = i - 1 {
    fmt.Println("Servicing   ", requests[i].location)

    // Recalculate costs without current cylinder
    // and update traversal distance
    if requests[i].location < upperCYL && requests[i].location > lowerCYL {
      seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
      current = requests[i].location
      fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance, " i: ", i)
    } else {
      fmt.Println("Cylinder is out of bounds")
    }
  }

  fmt.Println("finished right") //

  // Service cylinders to the left(down) of the disk
  for i := startIndex + 1; i < len(requests); i = i + 1 {
    fmt.Println("i: ", i)
    fmt.Println("Servicing   ", requests[i].location, " with curr: ", current)

    // Recalculate costs without current cylinder
    // and update traversal distance
    if requests[i].location < upperCYL && requests[i].location > lowerCYL {
      seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
      current = requests[i].location
      fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance, " i: ", i)
    } else {
      fmt.Println("Cylinder is out of bounds")
    }
  }

  fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)

}


// ---------------------------------------------------------------------- //
// Execute the C-SCAN scheduling algorithm will print cylinder number as  //
// it processes them and calculate total seek distance                    //
func runCSCAN(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {
  var requests = make([]cylinder, len(cylReqs))
  var cyl cylinder
  current := initCYL
  startIndex := 0
  seekDistance := 0
  turnHead := false

  // Create array of struct cylinders with respective locations
  // and initial costs
  for i := 0; i < len(cylReqs); i = i + 1 {
    cyl.location = cylReqs[i]
    cyl.cost = int(math.Abs(float64(cylReqs[i] - initCYL)))
    requests[i] = cyl
  }

  // Start by sorting request by location
  requests = sort(requests, "location")

  // Find First cylinder location to visit
  for i := 0; i < len(requests); i = i + 1 {
    if initCYL <= requests[i].location {
      startIndex = i
    }

    // Check if head must turn
    if requests[i].location < initCYL {
      turnHead = true
    }
  }

  // Service cylinders to the right(up) of the disk
  for i := startIndex; i >= 0; i = i - 1 {
    fmt.Println("Servicing   ", requests[i].location)

    // Recalculate costs without current cylinder
    // and update traversal distance
    if requests[i].location < upperCYL && requests[i].location > lowerCYL {
      seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
      current = requests[i].location

    } else {
      fmt.Println("Cylinder is out of bounds")
    }
  }


  // Update distance, account for head to reach end of disk
  // and traverse back to otherside.
  if turnHead {
    seekDistance = seekDistance + (upperCYL - current) + (upperCYL - lowerCYL)
    current = 0
  }

  // Service cylinders to the left(down) of the disk
  for i := len(requests) - 1; i > startIndex; i = i - 1 {
    fmt.Println("Servicing   ", requests[i].location)

    // Recalculate costs without current cylinder
    // and update traversal distance
    if requests[i].location < upperCYL && requests[i].location > lowerCYL {
      seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
      current = requests[i].location
    } else {
      fmt.Println("Cylinder is out of bounds")
    }
  }

  fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)
}


// -------------------------------------------------------------------- //
// Execute the LOOK scheduling algorithm will print cylinder number as  //
// it processes them and calculate total seek distance                  //
func runLOOK(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {
  var requests = make([]cylinder, len(cylReqs))
  var cyl cylinder
  current := initCYL
  startIndex := 0
  seekDistance := 0
  turnHead := false

  // Create array of struct cylinders with respective locations
  // and initial costs
  for i := 0; i < len(cylReqs); i = i + 1 {
    cyl.location = cylReqs[i]
    cyl.cost = int(math.Abs(float64(cylReqs[i] - initCYL)))
    requests[i] = cyl
  }

  // Start by sorting request by location
  requests = sort(requests, "location")

  // Find First cylinder location to visit
  for i := 0; i < len(requests); i = i + 1 {
    if initCYL <= requests[i].location {
      startIndex = i
    }

    // Check if head must turn
    if requests[i].location < initCYL {
      turnHead = true
    }
  }

  // Service cylinders to the right(up) of the disk
  for i := startIndex; i >= 0; i = i - 1 {

    fmt.Println("Servicing   ", requests[i].location)

    // Recalculate costs without current cylinder
    // and update traversal distance
    if requests[i].location < upperCYL && requests[i].location > lowerCYL {
      seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
      current = requests[i].location

    } else {
      fmt.Println("Cylinder is out of bounds")
    }
  }


  // Update distance, account for head to reach end of disk
  // and traverse back to otherside, only if it has to
  if turnHead {
    lastCYL := current
    current = requests[startIndex + 1].location
    seekDistance  = seekDistance + (lastCYL - current)

    // Service cylinders to the left(down) of the disk
    for i := startIndex + 1; i < len(requests); i = i + 1 {
      fmt.Println("Servicing   ", requests[i].location)

      // Recalculate costs without current cylinder
      // and update traversal distance
      if requests[i].location < upperCYL && requests[i].location > lowerCYL {
        seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
        current = requests[i].location
      } else {
        fmt.Println("Cylinder is out of bounds")
      }
    }
  }

  fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)
}


// -------------------------------------------------------------------- //
// Execute the LOOK scheduling algorithm will print cylinder number as  //
// it processes them and calculate total seek distance                  //
func runCLOOK(algorithm string, lowerCYL, upperCYL, initCYL int, cylReqs []int)() {
  var requests = make([]cylinder, len(cylReqs))
  var cyl cylinder
  current := initCYL
  startIndex := 0
  seekDistance := 0
  turnHead := false

  // Create array of struct cylinders with respective locations
  // and initial costs
  for i := 0; i < len(cylReqs); i = i + 1 {
    cyl.location = cylReqs[i]
    cyl.cost = int(math.Abs(float64(cylReqs[i] - initCYL)))
    requests[i] = cyl
  }

  // Start by sorting request by location
  requests = sort(requests, "location")

  // Find First cylinder location to visit
  for i := 0; i < len(requests); i = i + 1 {
    if initCYL <= requests[i].location {
      startIndex = i
    }

    // Check if head must turn
    if requests[i].location < initCYL {
      turnHead = true
    }
  }

  // Service cylinders to the right(up) of the disk
  for i := startIndex; i >= 0; i = i - 1 {

    fmt.Println("Servicing   ", requests[i].location, " i: ", )

    // Recalculate costs without current cylinder
    // and update traversal distance
    if requests[i].location < upperCYL && requests[i].location > lowerCYL {
      seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
      current = requests[i].location

    } else {
      fmt.Println("Cylinder is out of bounds")
    }
  }


  // Update distance, account for head to reach last cylinder
  // and traverse back to otherside from smallest cylinder, only if it has to
  if turnHead {
    lastCYL := current
    current = requests[len(requests) - 1].location
    seekDistance  = seekDistance + (lastCYL - current)

    // Service cylinders to the left(down) of the disk
    for i := len(requests) - 1; i > startIndex; i = i - 1 {
      fmt.Println("Servicing   ", requests[i].location)

      // Recalculate costs without current cylinder
      // and update traversal distance
      if requests[i].location < upperCYL && requests[i].location > lowerCYL {
        seekDistance = seekDistance + int(math.Abs(float64(requests[i].location - current)))
        current = requests[i].location
      } else {
        fmt.Println("Cylinder is out of bounds")
      }
    }
  }

  fmt.Println(strings.ToUpper(algorithm), " traversal count = ", seekDistance)
}

// ---------------- //
// Initiate Program //
func main() {

  // Read input file and structure data
  algorithm, lowerCYL, upperCYL, initCYL, cylReqs := readInputFile()

  // Print to STDOUT structured data
  printReadInput(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)

  // Execute chosen scheduling algorithm
  switch algorithm {

    case "fcfs":
      runFCFS(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)
    case "sstf":
      runSSTF(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)
    case "scan":
      runSCAN(algorithm, lowerCYL, upperCYL, initCYL, cylReqs) // Debug comming down
    case "c-scan":
      runCSCAN(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)
    case "look":
      runLOOK(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)
    case "c-look":
      runCLOOK(algorithm, lowerCYL, upperCYL, initCYL, cylReqs)

  }

  algorithm, lowerCYL, upperCYL, initCYL, cylReqs = algorithm, lowerCYL, upperCYL, initCYL, cylReqs //

}
