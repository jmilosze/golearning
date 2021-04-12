package main

import (
	"fmt"
	ds "golearning/datastructs"
)

func main() {
	newTree := ds.MakeBSTree(75)
	newTree.Insert(40)
	newTree.Insert(80)

	fmt.Println(newTree.ReturnInOrder())
}
