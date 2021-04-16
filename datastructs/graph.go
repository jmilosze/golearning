package datastructs

import "fmt"

type GNode struct {
	connections map[int]int
	parent      int
}

type Graph map[int]*GNode

func (graph *Graph) AddNode(value int) (*Graph, error) {
	if value < 1 {
		return graph, fmt.Errorf("node value %v is lower than minimum value of 1", value)
	}

	if _, ok := (*graph)[value]; ok == true {
		return graph, fmt.Errorf("node %v already exists", value)
	}
	(*graph)[value] = &GNode{make(map[int]int), 0}
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

func (graph *Graph) BreadthFirst(valueFrom int, valueTo int) (bool, error) {
	if _, ok := (*graph)[valueFrom]; ok != true {
		return false, fmt.Errorf("no node with value %v", valueFrom)
	}
	if _, ok := (*graph)[valueTo]; ok != true {
		return false, fmt.Errorf("no node with value %v", valueFrom)
	}

	checked := make(map[int]bool, 100)
	toCheck := make([]int, 100)[0:0]
	toCheck[0] = valueFrom

	for len(toCheck) > 0 {
		currentNode := toCheck[0]
		toCheck = toCheck[1:]

		checked[currentNode] = true

		if currentNode == valueTo {
			return true, nil
		}

		for neighbourNode := range (*graph)[currentNode].connections {
			if _, ok := checked[neighbourNode]; ok {
				continue
			}

			if (*graph)[neighbourNode].parent == 0 {
				(*graph)[neighbourNode].parent = currentNode
			}
		}
	}
	return false, fmt.Errorf("no path from node %v to node %v", valueFrom, valueTo)
}
