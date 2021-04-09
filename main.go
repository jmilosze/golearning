package main

import ds "golearning/datastructs"

func main() {
	newTree := ds.MakeTree(75)
	newTree.Insert(40)
	newTree.Insert(80)
	newTree.Insert(28)
	newTree.Insert(60)
	newTree.Insert(48)
	newTree.Insert(65)
	newTree.Insert(45)
	newTree.Insert(51)
	newTree.Insert(46)

	ds.PrintInOrder(&newTree.RootNode)

	newTree.Delete(40)

	ds.PrintInOrder(&newTree.RootNode)
}
