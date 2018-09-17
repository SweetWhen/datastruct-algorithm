1) an avl tree api--------------------------------------------------------------------------:

//==>> operator tree api:

//insert an x to AvlTree, invoker format: t = tree.Insert(t, x)

func Insert(tree *AvlNode, x ElementType) *AvlNode 

//del an x in AvlTree, invoker format: t = tree.Delete(t, x)

func Delete(tree *AvlNode, x ElementType) *AvlNode 	

//make a tree to be empty

func MakeEmpty( tree *AvlNode) *AvlNode 

//==>> print tree api:

//draw a tree

func DrawTree1(tree *AvlNode) 

//leve print a tree

func LevePrintTree(tree *AvlNode) 

//middle print a tree

func MidPrintTree(tree *AvlNode) 

//prev print a tree

func PrevPrintTree(tree *AvlNode)

//return the height of the node

func (node *AvlNode) NodeHeight() int 

//==>> find   node  api:

//find Element x position in the avl tree

func (tree *AvlNode) Find(x ElementType) *AvlNode

//find min elem in the tree

func (tree *AvlNode) FinMin() *AvlNode 

//find max elem in the tree

func (tree *AvlNode) FinMax() *AvlNode 

root@ubuntu01:tree# go test

create a slice: [19 13 0 9 4 29 3 14 21 23 1 12 2 11 24 26 28 20 7 17 25 15 18 6 16 5 8 27 10 22]

tree created

midl read:
 0  1  2  3  4  5  6  7  8  9  10  11  12  13  14  15  16  17  18  19  20  21  22  23  24  25  26  27  28  29 

leve read:
13 4 23 1 7 19 26 0 3 6 11 15 21 24 28 2 5 9 12 14 17 20 22 25 27 29 8 10 16 18 

draw:
                                13 
                4                                 23 
    1                 7                 19         26 
  0     3     6         11         15     21     24     28 
  2   5     9   12   14     17   20   22   25   27   29 
  8   10   16   18 

min: 0, max:29
PASS
2) heap api: --------------------------------------------------------------------------
An mult-goroutines safe package in struct-algorit/myheap/priorityqueue
init heap api:
(pq *PriorityQueue) Init(maxSize int)

Push/update api:
(pq *PriorityQueue) PushItem(key, value interface{},priority int) (bPushed bool) {

Pop api:
(pq *PriorityQueue) PopItem() interface {}

get all item api:
(pq PriorityQueue) GetQueueItems() []*Item

//testing:
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



