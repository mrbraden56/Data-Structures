package main

import "DataStructures/trie"
import "fmt"

func main() {
	root_node := trie.Node{
		Value: "root",
		Word:  false,
		Edges: make(map[string]*trie.Node),
	}

	root_node.Insert("app")
	root_node.Insert("asx")
	root_node.Insert("asp")
	root_node.Insert("brad")
	root_node.Insert("brady")
	root_node.Insert("brap")
	trie.PrintNodes(&root_node)
	root_node.Delete("asp")
	trie.PrintNodes(&root_node)
	fmt.Println(root_node.Search("asp"))
}
