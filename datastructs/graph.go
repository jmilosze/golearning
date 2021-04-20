package datastructs

import (
	"fmt"
	"math/bits"
)

type GNode struct {
	connections map[int]int
}

type Graph map[int]*GNode

func (graph *Graph) AddNode(value int) (*Graph, error) {
	if value < 1 {
		return graph, fmt.Errorf("node value %v is lower than minimum value of 1", value)
	}

	if _, ok := (*graph)[value]; ok == true {
		return graph, fmt.Errorf("node %v already exists", value)
	}
	(*graph)[value] = &GNode{make(map[int]int)}
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

func extractBreadthFirstPath(parents map[int]int, valueTo int) []int {
	path := make([]int, 100)[0:0]
	path = append(path, valueTo)
	parentValue, ok := parents[valueTo]
	for ok {
		path = append(path, parentValue)
		parentValue, ok = parents[parentValue]
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (graph *Graph) BreadthFirst(valueFrom int, valueTo int) ([]int, error) {
	if _, ok := (*graph)[valueFrom]; ok != true {
		return make([]int, 0), fmt.Errorf("no node with value %v", valueFrom)
	}
	if _, ok := (*graph)[valueTo]; ok != true {
		return make([]int, 0), fmt.Errorf("no node with value %v", valueFrom)
	}

	parents := make(map[int]int, 100)
	checked := make(map[int]bool, 100)
	toCheck := make([]int, 100)[0:0]
	toCheck = append(toCheck, valueFrom)

	for len(toCheck) > 0 {
		currentNode := toCheck[0]
		toCheck = toCheck[1:]

		checked[currentNode] = true

		if currentNode == valueTo {
			return extractBreadthFirstPath(parents, valueTo), nil
		}

		for neighbourNode := range (*graph)[currentNode].connections {
			if _, ok := checked[neighbourNode]; ok {
				continue
			}

			if _, ok := parents[neighbourNode]; !ok {
				parents[neighbourNode] = currentNode
			}
			toCheck = append(toCheck, neighbourNode)
		}
	}
	return make([]int, 0), fmt.Errorf("no path from node %v to node %v", valueFrom, valueTo)
}

const maxInt = 1<<(bits.UintSize-1) - 1

type nodeInfo struct {
	processed bool
	cost      int
	parent    int
}

func nextNode(nodes map[int]*nodeInfo) int {
	nextNodeVal := 0
	nextNodeCost := maxInt

	for node, nodeData := range nodes {
		if nodeData.cost < nextNodeCost && !nodeData.processed {
			nextNodeVal = node
		}
	}
	return nextNodeVal
}

func extractDijkstraPath(nodes map[int]*nodeInfo, valueTo int) []int {
	path := make([]int, 100)[0:0]
	path = append(path, valueTo)

	for parentValue := nodes[valueTo].parent; parentValue > 0; parentValue = nodes[parentValue].parent {
		path = append(path, parentValue)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func (graph *Graph) Dijkstra(valueFrom int, valueTo int) ([]int, error) {
	if valueFrom == valueFrom {
		return []int{valueFrom}, nil
	}

	nodes := make(map[int]*nodeInfo, 100)

	for node := range *graph {
		if node != valueFrom {
			nodes[node] = &nodeInfo{false, maxInt, 0}
		}
	}

	for neighbourNode, weight := range (*graph)[valueFrom].connections {
		nodes[neighbourNode].parent = valueFrom
		nodes[neighbourNode].cost = weight
	}

	for nodeValue := nextNode(nodes); nodeValue > 0; nodeValue = nextNode(nodes) {
		for neighbourNode, neighbourNodeWeight := range (*graph)[valueFrom].connections {
			newCost := nodes[nodeValue].cost + neighbourNodeWeight
			if newCost > nodes[neighbourNode].cost {
				nodes[neighbourNode].cost = newCost
				nodes[neighbourNode].parent = nodeValue
			}
		}
		nodes[nodeValue].processed = true
	}

	if nodes[valueTo].cost == maxInt {
		return make([]int, 0), fmt.Errorf("no path from nodeValue %v to nodeValue %v", valueFrom, valueTo)
	}

	return extractDijkstraPath(nodes, valueTo), nil
}
