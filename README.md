# merge_lists

Takes a list of intervals and merges every overlapping interval and prints the resulting list.

## Complexitiy
Sorting runs with a complexity of O(n\*log(n)).  
Merging runs with a complexity of O(n).  
Therefore the program runs with a complexity of O(n\*log(n)).  

## Memory Usage
Every interval gets saved in a struct containing two uint64 variables.  
The same structure is used for the result list.

## Reliability
Since the length of the interval input list is determined during runtime the needed memory is allocated on the heap, which is expensive.
It should be noted that the golang-GC does a new cycle everytime the heap usage doubles.  
So in case this program does something else after parsing a huge input list, the memory for the intervals might wont get freed for a long time. 

## Usage
./merge_lists -input="[1,5] [2,3]"

## Build
go build

## Test
go test

