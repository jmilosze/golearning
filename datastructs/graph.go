package datastructs

import "fmt"

type Graph map[int]map[int]int

func (graph *Graph) AddNode(value int) (*Graph, error) {
	if _, ok := (*graph)[value]; ok == true {
		return graph, fmt.Errorf("node %v already exists", value)
	}
	(*graph)[value] = make(map[int]int)
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
	(*graph)[valueFrom][valueTo] = weight

	return graph, nil
}
