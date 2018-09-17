
package priorityqueue

import (
	"fmt"
	"testing"
)

type EatFruit struct {
	key string
	UintPrice float64
	Cnt int 
	Pri int	
}

func (eat* EatFruit)Print() {
		fmt.Printf("%s, UintPrice:%v, Cnt:%v\n",eat.key, eat.UintPrice, eat.Cnt )
}

func TestPriorityQue(t *testing.T) {
	//create a priority queue, put the items in it
	items := [] EatFruit {
		 EatFruit{
		 	key: "Banana",
			UintPrice:2.4, 
			Cnt:10, Pri:4, 				
		}, 			
		EatFruit{
			key: "Apple",
			UintPrice:2.4, 
			Cnt:10, Pri:7, 
		},

		EatFruit{
			key: "Orange",
			UintPrice:2.4, 
			Cnt:10, Pri:6,
		},		
	}
	
	pq := &PriorityQueue {
		maxSize :  32,		
	} 
	pq.Init(pq.maxSize)	
	
	for _, item := range items {
		i := item
		pq.PushItem(i.key, &i, i.Pri)
	}
	fmt.Println("Before update Apple data:")
	for _, item := range pq.GetQueueItems() {
		eat := item.Value.(*EatFruit)
		eat.Print()
	}
	fmt.Println("Updated Apple data and PopItem:")
	i := &EatFruit {
		key: "Apple",
		UintPrice:69, 
		Cnt:3, Pri:2, 
	}
	pq.PushItem(i.key, i, i.Pri)	
	item := pq.PopItem()	
	if item != nil {
		item := item.(*EatFruit)
		item.Print()
	}
	fmt.Println("After PopItem:")
	for _, item := range pq.GetQueueItems() {
		eat := item.Value.(*EatFruit)
		eat.Print()
	}
	
}


