package datastructs

import "fmt"

type GNode struct {
	connections map[int]int
	parent      int
}

type Graph map[int]GNode

func (graph *Graph) AddNode(value int) (*Graph, error) {
	if value < 1 {
		return graph, fmt.Errorf("node value %v is lower than minimum value of 1", value)
	}

	if _, ok := (*graph)[value]; ok == true {
		return graph, fmt.Errorf("node %v already exists", value)
	}
	(*graph)[value] = GNode{make(map[int]int), 0}
	return graph, nil
}

func (graph *Graph) AddConnection(valueFrom int, valueTo int, weight int) (*Graph, error) {
	if valueFrom == valueTo {
		return graph, fmt.Errorf("cannot connect node %v to itself", valueFrom)
	}
	if weight < 0 {
		return graph, fmt.Errorf("negative weights (%v) are not allowed", weight)
	}
	if _, ok := (*graph)[valueFrom]; ok != true {
		return graph, fmt.Errorf("no node with value %v", valueFrom)
	}
	if _, ok := (*graph)[valueTo]; ok != true {
		return graph, fmt.Errorf("no node with value %v", valueFrom)
	}
	(*graph)[valueFrom].connections[valueTo] = weight

	return graph, nil
}

//func BreadthFirst(graph Graph, valueFrom int, valueTo int) ([]int, error) {
//	if _, ok := graph[valueFrom]; ok != true {
//		return make([]int, 0), fmt.Errorf("no node with value %v", valueFrom)
//	}
//	if _, ok := graph[valueTo]; ok != true {
//		return make([]int, 0), fmt.Errorf("no node with value %v", valueFrom)
//	}
//
//	toDo := make([]int, 100)[0:0]
//	toDo[0] = valueFrom
//
//	for len(toDo) > 0 {
//		nextNode := toDo[0]
//		toDo = toDo[1:]
//
//		if nextNode == valueTo {
//			return
//		}
//	}
//}
