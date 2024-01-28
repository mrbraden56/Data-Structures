package trie

import "fmt"
import "container/list"

type Node struct {
	Value     string
	Word      bool
	Edges     map[string]*Node
	CanDelete bool
}

func PrintNodes(rootNode *Node) {
	queue := list.New()
	queue.PushBack(rootNode)
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			element := queue.Front()
			queue.Remove(element)
			node, ok := element.Value.(*Node)
			if !ok {
				continue
			}

			fmt.Print(node.Value + " ")
			for _, val := range node.Edges {
				queue.PushBack(val)
			}
		}
		fmt.Println("")
	}
}

func (node *Node) Insert(value string) {
	if len(value) == 0 {
		return
	}
	char := string(value[0])
	if nextNode, ok := node.Edges[string(char)]; ok {
		nextNode.Insert(value[1:])
	} else {
		newNode := &Node{
			Value: char,
			Word:  false,
			Edges: make(map[string]*Node),
		}
		node.Edges[char] = newNode
		newNode.Insert(value[1:])
	}

	if len(value) == 1 {
		node.Edges[char].Word = true
	}

}

/*
Performe using DFS
 1. Navigate to end of the word...
    - a If node has no children, delete
    - b If node has children, do not delete but set Word as false
 2. After dealing with child node, recurse back up...
    - a If node does not have any children, delete, else continue

When at a node, check above conditions
*/
func (node *Node) Delete(value string) {
	if len(value) == 0 {
		return
	}
	char := string(value[0])
	var nextNode *Node
	var ok bool

	if nextNode, ok = node.Edges[string(char)]; ok {
		nextNode.Delete(value[1:])
		if len(value) == 1 {
			nextNode.Word = false
		}
	} else {
		//Throw error
		fmt.Println("Node does not exist")
	}

	//This is where we handle deletion
	if len(nextNode.Edges) == 0 {
		nextNode.CanDelete = true
	} else {
		child := string(value[1])
		if nodeChild, ok := nextNode.Edges[child]; ok {
			if nodeChild.CanDelete {
				delete(nextNode.Edges, child)
			}
			if len(nextNode.Edges) == 0 {
				nextNode.CanDelete = true
			}
		}
	}

}

func (node *Node) Search(value string) bool {
	if len(value) == 0 {
		return false
	}

	var endOfWord bool
	var isFound bool
	if nextNode, ok := node.Edges[string(value[0])]; ok {
		isFound = nextNode.Search(value[1:])
		if len(value) == 1 {
			endOfWord = nextNode.Word
		}

	}

	return isFound || endOfWord

}

func (node *Node) Content() {
	fmt.Println(node.Value)
}
