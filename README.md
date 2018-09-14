data struct and algorithm in golang

======>>>>>>>>
branch rand-qsort: qsort + insert sort, use random number generation to test

======>>>>>>>>
branch test-qsort: use testing package test our qsort

======>>>>>>>>
branch go-tree: api for avl tree:

root@ubuntu01:tree# go test

create a slice: [1 16 27 25 9 12 24 21 0 8 17 15 19 29 23 5 4 22 18 26 7 10 20 28 6 2 13 3 11 14]

tree created

midl read:
 0  1  2  3  4  5  6  7  8  9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29 
 
leve read:
16 5 21 1 9 18 25 0 3 7 12 17 19 23 27 2 4 6 8 10 14 20 22 24 26 29 11 13 15 28 


draw:
                                16 
                5                                 21 
    1                 9         18                 25 
  0     3     7         12   17     19     23         27 
  2   4   6   8     10     14   20   22   24   26     29 
  11   13   15   28 

min: 0, max:29
PASS

