data struct and algorithm in golang

======>>>>>>>>
branch rand-qsort: qsort + insert sort, use random number generation to test

======>>>>>>>>
branch test-qsort: use testing package test our qsort

======>>>>>>>>
branch go-tree: api for avl tree, and a package for muti-goroutines safe api
Here are the test of go-tree branch:
1) avl tree:
root@ubuntu01:tree# go test
draw the avl tree:
                                16 
                5                                 21 
    1                 9         18                 25 
  0     3     7         12   17     19     23         27 
  2   4   6   8     10     14   20   22   24   26     29 
  11   13   15   28 

min: 0, max:29
PASS

2) muti-goroutines safe heap:
root@ubuntu01:priorityqueue# go test
Before update Apple data:
Banana, UintPrice:2.4, Cnt:10
Orange, UintPrice:2.4, Cnt:10
Apple, UintPrice:2.4, Cnt:10
Updated Apple data and PopItem:
Apple, UintPrice:69, Cnt:3
After PopItem:
Banana, UintPrice:2.4, Cnt:10
Orange, UintPrice:2.4, Cnt:10
PASS
ok  	_/home/agan/data-struct/struct-algorit/myheap/priorityqueue	0.010s


