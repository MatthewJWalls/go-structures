package main

type PathList [][]*TrieNode

// TrieNode

type TrieNode struct {
	char string
	children []*TrieNode
	doStop bool
}

func (this *TrieNode) isLeaf() bool {
	return len(this.children) == 0
}

func (this *TrieNode) isStop() bool {
	return this.doStop
}

func (this *TrieNode) getChild(char string) *TrieNode {

	for _, n := range(this.children) {
		if n.char == char {
			return n
		}
	}

	return nil
}

func (this *TrieNode) size() int {

	sum := 1

	for _, c := range(this.children){
		sum += c.size()
	}

	return sum

}

// Trie

type Trie struct {
	root *TrieNode
}

func (this *Trie) Size() int {
	return this.root.size() - 1
}

func (this *Trie) Insert(word string) {

	marker := this.root

	for _, char := range(word) {

		child := marker.getChild(string(char))
		
		if child == nil {
			child = NewTrieNode(string(char))
			marker.children = append(marker.children, child)
		}

		marker = child

	}

	// signify that this was a word end
	marker.doStop = true

}

func (this *Trie) Complete(prefix string) []string {

	var collector PathList
	var startPath []*TrieNode

	// step 1. find the appropriate start node

	var startNode = this.root

	for _, char := range(prefix) {
		child := startNode.getChild(string(char))
		if child == nil {
			return make([]string, 0, 0)
		} else {
			startNode = child
		}
	}

	// step 2. get all branches from this node

	this.branches(&collector, startPath, startNode)

	var out []string

	for _, path := range(collector) {
		word := ""
		for _, node := range(path) {
			word += node.char
		}
		out = append(out, prefix+word)
	}

	return out

}

func (this *Trie) branches(branches *PathList, path []*TrieNode, n *TrieNode) {

	path = append(path, n)

	if n.doStop {
		// this was a word end, collect it into branches
		*branches = append(*branches, path)
	}

	if ! n.isLeaf() {
		// recurse down the children
		for _, c := range(n.children) {
			this.branches(branches, path, c)
		}
	}

}

func NewTrie() *Trie {
	return &Trie{NewTrieNode("")}
}

func NewTrieNode(char string) *TrieNode {
	return &TrieNode{char, make([]*TrieNode, 0, 0), false}
}
