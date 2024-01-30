package cow_trie

import "fmt"

/*
For the COW Trie we will have a list of Nodes that represents the Roots.
The key to these roots is represented by Latest which will be a counter. Each time we insert a new value(meaning we insert a new root)

	-we will update the Latest+=1. This way we can always have O(1) access to most recent root.
*/
type CowTrie struct {
	Roots  map[int]*Node
	Latest int
}

type Node struct {
	Value     any
	Word      bool
	Edges     map[string]*Node
	CanDelete bool
}

func (node *Node) Insert(key string, value any) {
	if len(key) == 0 {
		return
	}
	char := string(key[0])
	if nextNode, ok := node.Edges[string(char)]; ok {
		nextNode.Insert(key[1:], value)
	} else {
		newNode := &Node{
			Value: value,
			Word:  false,
			Edges: make(map[string]*Node),
		}
		node.Edges[char] = newNode
		newNode.Insert(key[1:], value)
	}

	if len(key) == 1 {
		node.Edges[char].Word = true
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

func (node *Node) Replace(key string, value any) {
	if len(key) == 0 {
		return
	}

	if nextNode, ok := node.Edges[string(key[0])]; ok {
		nextNode.Replace(key[1:], value)
		if len(key) == 1 {
			nextNode.Value = value
		}
	}
}

func (cowtrie *CowTrie) Initialize(key string, value any, latest_root int) {
	cowtrie.Roots = make(map[int]*Node) // Initialize the map
	root_node := Node{
		Value: "root",
		Word:  false,
		Edges: make(map[string]*Node),
	}
	cowtrie.Latest = latest_root
	cowtrie.Roots[cowtrie.Latest] = &root_node

}

func (cowtrie *CowTrie) option2(key string, value any) {
	//option 2
	cowtrie.Roots[cowtrie.Latest].Insert(key, value)
	for key, node := range cowtrie.Roots[cowtrie.Latest-1].Edges {
		cowtrie.Roots[cowtrie.Latest].Edges[key] = node
	}
}

func (cowtrie *CowTrie) option4(key string, value any) {
}

func (cowtrie *CowTrie) Insert(key string, value any) {
	if len(cowtrie.Roots) == 0 {
		//option 1(Drawn in notebook)
		fmt.Println("No Roots")
		cowtrie.Initialize(key, value, 0)
		cowtrie.Roots[cowtrie.Latest].Insert(key, value)
	} else {
		fmt.Println("Roots")
		if cowtrie.Roots[cowtrie.Latest].Search(key) {
			//option 3
			cowtrie.Roots[cowtrie.Latest].Replace(key, value)
		} else {
			cowtrie.Initialize(key, value, cowtrie.Latest+1)
			//option 2 or option 4
			prev_root := cowtrie.Roots[cowtrie.Latest-1]
			_, ok := prev_root.Edges[string(key[0])]
			if ok {
				cowtrie.option4(key, value)
			} else {
				cowtrie.option2(key, value)
			}

		}

	}

}
