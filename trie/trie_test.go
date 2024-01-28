package trie

import "testing"
import "container/list"
import "sort"

func TestInsert(t *testing.T) {
	root_node := Node{
		Value: "root",
		Word:  false,
		Edges: make(map[string]*Node),
	}

	root_node.Insert("app")
	root_node.Insert("ass")
	root_node.Insert("aps")
	root_node.Insert("brad")
	root_node.Insert("brady")
	root_node.Insert("brap")
	root_node.Insert("bred")

	var levels [][]string
	levels = append(levels, []string{"root"})
	levels = append(levels, []string{"a", "b"})
	levels = append(levels, []string{"p", "s", "r"})
	levels = append(levels, []string{"p", "s", "s", "a", "e"})
	levels = append(levels, []string{"d", "p", "d"})
	levels = append(levels, []string{"y"})

	queue := list.New()
	queue.PushBack(&root_node)
	level := 0
	for queue.Len() != 0 { // This counts the level
		size := queue.Len()
		current_level_values := []string{}
		for i := 0; i < size; i++ {
			element := queue.Front()
			queue.Remove(element)
			node, ok := element.Value.(*Node)
			if !ok {
				continue
			}

			current_level_values = append(current_level_values, node.Value)

			for _, val := range node.Edges {
				queue.PushBack(val)
			}
		}
		level_values := levels[level]
		sort.Strings(level_values)
		sort.Strings(current_level_values)
		if len(level_values) != len(current_level_values) {
			t.Fatalf(`Level %d: Expected %d nodes, but got %d`, level, len(level_values), len(current_level_values))
		}
		for index, _ := range level_values {
			if level_values[index] != current_level_values[index] {
				t.Fatalf(`Actual Value: %s; Received Value: %s`, level_values[index], current_level_values[index])
			}
		}
		level++
	}
}
func TestDelete(t *testing.T) {
	root_node := Node{
		Value: "root",
		Word:  false,
		Edges: make(map[string]*Node),
	}

	root_node.Insert("app")
	root_node.Insert("asx")
	root_node.Insert("aps")
	root_node.Insert("brad")
	root_node.Insert("brady")
	root_node.Insert("brap")
	root_node.Insert("bred")

	root_node.Delete("asx")

	var levels [][]string
	levels = append(levels, []string{"root"})
	levels = append(levels, []string{"a", "b"})
	levels = append(levels, []string{"p", "r"})
	levels = append(levels, []string{"p", "s", "a", "e"})
	levels = append(levels, []string{"d", "p", "d"})
	levels = append(levels, []string{"y"})

	queue := list.New()
	queue.PushBack(&root_node)
	level := 0
	for queue.Len() != 0 { // This counts the level
		size := queue.Len()
		current_level_values := []string{}
		for i := 0; i < size; i++ {
			element := queue.Front()
			queue.Remove(element)
			node, ok := element.Value.(*Node)
			if !ok {
				continue
			}

			current_level_values = append(current_level_values, node.Value)

			for _, val := range node.Edges {
				queue.PushBack(val)
			}
		}
		level_values := levels[level]
		sort.Strings(level_values)
		sort.Strings(current_level_values)
		if len(level_values) != len(current_level_values) {
			t.Fatalf(`Level %d: Expected %d nodes, but got %d`, level, len(level_values), len(current_level_values))
		}
		for index, _ := range level_values {
			if level_values[index] != current_level_values[index] {
				t.Fatalf(`Actual Value: %s; Received Value: %s`, level_values[index], current_level_values[index])
			}
		}
		level++
	}
}
