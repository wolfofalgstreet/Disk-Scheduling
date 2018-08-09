# Disk-Scheduling


## Objective

The goal of this program is to implement various disk scheduling algorithms, the program reads an input file containing several
configuration parameters to simulate first-come first-served (FCFS), shortest seek time first (SSTF), SCAN, C-SCAN, LOOK, C-LOOK  disk scheduling algorithms using Go. The output will reflect the results of the configured disk scheduling
algorithm along with its traversed cylinders and traversal count.

## Input

The program will read a file specified in the first comman line parameter, `ARGV[1]`, which will be formatted as follows. The program ignores anything that comes afer "#" mark for every line and additional spaces in input file. 

There may be any number of cylinder requests but there may only be n cylinder requests for any algorithm where each request is unique, where n corresponds to the bound size. In this case the program is bounded to lowerCYL and upperCYL, thus the number of total cylinders that can be serviced only once is upperCYL - lowerCYL. 

In the case that a cylinder request is not bounded by both bounds, an error message is generated and the rest of the requests are serviced. 


### Sample Input
```
Use	sstf  #	fcfs,	sstf,	scan,	c-scan,	look,	or c-look
lowerCYL  00000 # valid	lower	cylinder number
upperCYL 00000 # valid upper cylinder number (> lower cylinder)
initCYL 00000 #	initial	cylinder position (0 < initCYL < upperCYL)
cylreq  00000 # a single cylinder request, where the lowerCYL < cylreq < upperCYL
cylreq	00000 # a single cylinder request
cylreq	00000 # a single cylinder request up to 20 requests
end

```

## Test Files

The following input files test each golrithm ussing a different number of cylinder requests, corresponding output files have a .base extension. 


| File Name     | Description                                                    | Output File |
| ------------- |:--------------------------------------------------------------:|:------------:|
| fcfs01.txt    | 6 cylinder requests, scheduled by *First-Come First-Served*    |fcfs01.base|
| fcfsPA.txt    | 12 cylinder requests, scheduled by *First-Come First-Served*   |fcfsPA.base|
| fcfs20.txt    | 20 cylinder requests, scheduled by *First-Come First-Served*   |fcfs20.base|
| sstf01.txt    | 6 cylinder requests, scheduled by *Shortest-Seek First*        |sstf01.base|
| sstfPA.txt    | 12 cylinder requests, scheduled by *Shortest-Seek First*       |sstfPA.base|
| sstf20.txt    | 20 cylinder requests, scheduled by *Shortest-Seek First*       |sstf20.base|
| scan01.txt    | 6 cylinder requests, scheduled by *Scan*                       |scan01.base|
| scanPA.txt    | 12 cylinder requests, scheduled by *Scan*                      |scanPA.base|
| scan20.txt    | 20 cylinder requests, scheduled by *Scan*                      |scan20.base|
| c-scan01.txt    | 6 cylinder requests, scheduled by *C-Scan*                   |c-scan01.base|
| c-scanPA.txt    | 12 cylinder requests, scheduled by *C-Scan*                  |c-scanPA.base|
| c-scan20.txt    | 20 cylinder requests, scheduled by *C-Scan*                  |c-scan20.base|
| look01.txt    | 6 cylinder requests, scheduled by *Look*                       |look01.base|
| lookPA.txt    | 12 cylinder requests, scheduled by *Look*                      |lookPA.base|
| look20.txt    | 20 cylinder requests, scheduled by *Look*                      |look20.base|
| c-look01.txt    | 6 cylinder requests, scheduled by *C-Look*                   |c-look01.base|
| c-lookPA.txt    | 12 cylinder requests, scheduled by *C-Look*                  |c-lookPA.base|
| c-look20.txt    | 20 cylinder requests, scheduled by *C-Look*                  |c-look20.base|

## Compiling and Running 

To compile run the following command:

```
go build DiskScheduler.go
```

Before running the program there should be an input file at the same location as the source .go file. Specify the name of the formatted input file as the first command line parameter. 

```
DiskScheduler.exe fcfs01.txt
```






