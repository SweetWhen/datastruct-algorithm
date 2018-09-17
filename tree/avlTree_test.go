package tree

import (
	"testing"
	"math/rand"
	"fmt"
	"time"	
)

func TestCreateTree(t *testing.T) {
	r := rand.New(rand.NewSource(10))	
	r.Seed(time.Now().UnixNano())
	a := r.Perm(30)	
	fmt.Println("create a slice:", a)
	var tree *AvlNode
	for _, v := range a {
		tree = Insert(tree, ElementType(v)) 
	}
		
	fmt.Println("tree created\nmidl read:")
	MidPrintTree(tree)
	fmt.Println("\nleve read:")
	LevePrintTree(tree)
	fmt.Println("\ndraw:")
	DrawTree1(tree);
	fmt.Printf("\nmin: %v, max:%v\n", tree.FinMin().Elem, tree.FinMax().Elem)		
}





