package graph

import "fmt"

func NewGraph() *Graph {
	return &Graph{
		RootNodes: make([]string, 0),
		Nodes:     make(map[string]*DependencyNode),
		Adjacency: make(map[string][]string),
	}
}

func (g *Graph) AddNode(node *DependencyNode) error {
	if _, exists := g.Nodes[node.ID]; exists {
		return nil
	}
	g.Nodes[node.ID] = node

	if node.Direct {
		g.RootNodes = append(g.RootNodes, node.ID)
	}
	return nil
}

func (g *Graph) AddEdge(parentID, childID string) error {
	if _, exists := g.Nodes[parentID]; !exists {
		return fmt.Errorf("parent node %s not found", parentID)
	}
	if _, exists := g.Nodes[childID]; !exists {
		return fmt.Errorf("child node %s not found", childID)
	}

	for _, existingChild := range g.Adjacency[parentID] {
		if existingChild == childID {
			return nil
		}
	}

	g.Adjacency[parentID] = append(g.Adjacency[parentID], childID)
	return nil
}
